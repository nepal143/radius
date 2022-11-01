// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package ucp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	awsgo "github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/google/uuid"
	"github.com/project-radius/radius/pkg/ucp/api/v20220901privatepreview"
	"github.com/project-radius/radius/pkg/ucp/aws"
	"github.com/project-radius/radius/test/validation"
	"github.com/stretchr/testify/require"
)

var streamName = "my-stream" + uuid.NewString()
var resourceType = "AWS::Kinesis::Stream"

func Test_AWS_DeleteResource(t *testing.T) {
	ctx := context.Background()
	setupTestAWSResource(t, ctx)

	test := NewUCPTest(t, "Test_AWS_DeleteResource", func(t *testing.T, url string, roundTripper http.RoundTripper) {
		// Call UCP Delete AWS Resource API
		resourceID := validation.GetResourceIdentifier(t, "AWS.Kinesis/Stream", streamName)

		// Remove the stream name from the to form the post URL and add the stream name to the body
		resourceIDParts := strings.Split(resourceID, "/")
		resourceIDParts = resourceIDParts[:len(resourceIDParts)-1]
		resourceID = strings.Join(resourceIDParts, "/")
		deleteURL := fmt.Sprintf("%s%s/:delete?api-version=%s", url, resourceID, v20220901privatepreview.Version)
		deleteRequestBody := map[string]interface{}{
			"properties": map[string]interface{}{
				"Name": streamName,
			},
		}
		deleteBody, err := json.Marshal(deleteRequestBody)
		require.NoError(t, err)

		// Issue the Delete Request
		deleteRequest, err := http.NewRequest(http.MethodPost, deleteURL, bytes.NewBuffer(deleteBody))
		require.NoError(t, err)
		deleteResponse, err := roundTripper.RoundTrip(deleteRequest)
		require.NoError(t, err)
		require.Equal(t, http.StatusAccepted, deleteResponse.StatusCode)

		// Get the operation status url from the Azure-Asyncoperation header
		deleteResponseCompletionUrl := deleteResponse.Header["Azure-Asyncoperation"][0]
		getRequest, err := http.NewRequest(http.MethodGet, deleteResponseCompletionUrl, nil)
		require.NoError(t, err)
		maxRetries := 100
		deleteSucceeded := false
		for i := 0; i < maxRetries; i++ {
			getResponse, err := roundTripper.RoundTrip(getRequest)
			require.NoError(t, err)
			require.Equal(t, http.StatusOK, getResponse.StatusCode)

			// Read the request status from the body
			payload, err := io.ReadAll(getResponse.Body)
			require.NoError(t, err)
			body := map[string]interface{}{}
			err = json.Unmarshal(payload, &body)
			require.NoError(t, err)
			if body["status"] == "Succeeded" {
				deleteSucceeded = true
				break
			}
			// Give it more time
			time.Sleep(1 * time.Second)
		}
		require.True(t, deleteSucceeded)
	})
	test.Test(t)

}

func setupTestAWSResource(t *testing.T, ctx context.Context) {
	// Test setup - Create AWS resource using AWS APIs
	cfg, err := awsconfig.LoadDefaultConfig(ctx)
	require.NoError(t, err)
	var awsClient aws.AWSCloudControlClient = cloudcontrol.NewFromConfig(cfg)
	desiredState := map[string]interface{}{
		"Name":                 streamName,
		"RetentionPeriodHours": 180,
		"ShardCount":           4,
	}
	desiredStateBytes, err := json.Marshal(desiredState)
	require.NoError(t, err)

	response, err := awsClient.CreateResource(ctx, &cloudcontrol.CreateResourceInput{
		TypeName:     &resourceType,
		DesiredState: awsgo.String(string(desiredStateBytes)),
	})
	require.NoError(t, err)
	waitForSuccess(t, ctx, awsClient, response.ProgressEvent.RequestToken)

	t.Cleanup(func() {
		// Check if resource exists before issuing a delete because the AWS SDK async delete operation
		// seems to fail if the resource does not exist
		_, err := awsClient.GetResource(ctx, &cloudcontrol.GetResourceInput{
			Identifier: &streamName,
			TypeName:   &resourceType,
		})
		if aws.IsAWSResourceNotFound(err) {
			return
		}
		// Just in case delete fails
		deleteOutput, err := awsClient.DeleteResource(ctx, &cloudcontrol.DeleteResourceInput{
			Identifier: &streamName,
			TypeName:   &resourceType,
		})
		require.NoError(t, err)

		// Ignoring status of delete since AWS command fails if the resource does not already exist
		waitForSuccess(t, ctx, awsClient, deleteOutput.ProgressEvent.RequestToken)
	})
	// End of test setup
}

func waitForSuccess(t *testing.T, ctx context.Context, awsClient aws.AWSCloudControlClient, requestToken *string) {
	// Wait till the create is complete
	maxWaitTime := 300 * time.Second
	waiter := cloudcontrol.NewResourceRequestSuccessWaiter(awsClient)
	err := waiter.Wait(ctx, &cloudcontrol.GetResourceRequestStatusInput{
		RequestToken: requestToken,
	}, maxWaitTime)
	require.NoError(t, err)
}