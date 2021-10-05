// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Azure/radius/pkg/azure/azresources"
	"github.com/Azure/radius/pkg/radlogger"
	"github.com/Azure/radius/pkg/radrp/armerrors"
	"github.com/Azure/radius/pkg/radrp/frontend/resourceprovider"
	"github.com/Azure/radius/pkg/radrp/frontend/server"
	"github.com/Azure/radius/pkg/radrp/rest"
	"github.com/Azure/radius/pkg/radrp/schema"
	"github.com/go-logr/logr"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

// These tests cover the mechanics of how the handler turns HTTP requests into
// the resource IDs and strongly typed data model of the RP.
//
// There's basically no business logic in the handler, all of that is delegated to mocks.

const providerURI = "/subscriptions/test-subscription/resourceGroups/test-resource-group/providers/Microsoft.CustomProviders/resourceProviders/radiusv3"
const baseURI = providerURI + "/Application"

type test struct {
	t         *testing.T
	ctrl      *gomock.Controller
	server    *httptest.Server
	handler   http.Handler
	rp        *resourceprovider.MockResourceProvider
	validator *FakeValidator
}

func createContext(t *testing.T) context.Context {
	logger, err := radlogger.NewTestLogger(t)
	if err != nil {
		t.Log("Unable to initialize logger")
		return context.Background()
	}
	return logr.NewContext(context.Background(), logger)
}

func start(t *testing.T) *test {
	ctrl := gomock.NewController(t)
	rp := resourceprovider.NewMockResourceProvider(ctrl)

	validator := &FakeValidator{}
	options := server.ServerOptions{
		Address:      httptest.DefaultRemoteAddr,
		Authenticate: false,
		Configure: func(router *mux.Router) {
			AddRoutes(rp, router, func(resourceType string) (schema.Validator, error) {
				if validator.RejectType {
					return nil, errors.New("unsupported type")
				}

				return validator, nil
			})
		},
	}

	s := server.NewServer(createContext(t), options)
	server := httptest.NewServer(s.Handler)
	h := injectLogger(t, server.Config.Handler)
	t.Cleanup(server.Close)

	return &test{
		t:         t,
		rp:        rp,
		ctrl:      ctrl,
		server:    server,
		handler:   h,
		validator: validator,
	}
}

func injectLogger(t *testing.T, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		logger, err := radlogger.NewTestLogger(t)
		if err != nil {
			t.Error(err)
			h.ServeHTTP(w, r.WithContext(context.Background()))
			return
		}
		ctx := logr.NewContext(context.Background(), logger)
		h.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func requireJSON(t *testing.T, expected interface{}, w *httptest.ResponseRecorder) {
	bytes, err := json.Marshal(expected)
	require.NoError(t, err)
	require.JSONEq(t, string(bytes), w.Body.String())
}

type FaultingResponse struct {
}

func (r *FaultingResponse) Apply(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	return fmt.Errorf("write failure!")
}

type FakeValidator struct {
	RejectType bool
	Errors     []schema.ValidationError
}

func (v *FakeValidator) ValidateJSON(body []byte) []schema.ValidationError {
	return v.Errors
}

func Test_Handler(t *testing.T) {
	testcases := []struct {
		Method      string
		Description string
		URI         string
		Expect      func(*resourceprovider.MockResourceProvider) *gomock.Call
		Body        interface{}
	}{
		{
			Method:      "GET",
			Description: "ListApplications",
			URI:         baseURI,
			Expect: func(mock *resourceprovider.MockResourceProvider) *gomock.Call {
				return mock.EXPECT().ListApplications(gomock.Any(), gomock.Any())
			},
		},
		{
			Method:      "GET",
			Description: "GetApplication",
			URI:         baseURI + "/test-application",
			Expect: func(mock *resourceprovider.MockResourceProvider) *gomock.Call {
				return mock.EXPECT().GetApplication(gomock.Any(), gomock.Any())
			},
		},
		{
			Method:      "PUT",
			Description: "UpdateApplication",
			URI:         baseURI + "/test-application",
			Expect: func(mock *resourceprovider.MockResourceProvider) *gomock.Call {
				return mock.EXPECT().UpdateApplication(gomock.Any(), gomock.Any(), gomock.Any())
			},

			// We don't need to include any significant data in the body, we're using mocks for testing
			// here. We just want to cover the code path.
			Body: map[string]interface{}{
				"tags": map[string]interface{}{
					"test-tag": "test-value",
				},
			},
		},
		{
			Method:      "DELETE",
			Description: "DeleteApplication",
			URI:         baseURI + "/test-application",
			Expect: func(mock *resourceprovider.MockResourceProvider) *gomock.Call {
				return mock.EXPECT().DeleteApplication(gomock.Any(), gomock.Any())
			},
		},
		{
			Method:      "GET",
			Description: "ListResources",
			URI:         baseURI + "/test-application/test-resource-type",
			Expect: func(mock *resourceprovider.MockResourceProvider) *gomock.Call {
				return mock.EXPECT().ListResources(gomock.Any(), gomock.Any())
			},
		},
		{
			Method:      "GET",
			Description: "GetResource",
			URI:         baseURI + "/test-application/test-resource-type/test-resource",
			Expect: func(mock *resourceprovider.MockResourceProvider) *gomock.Call {
				return mock.EXPECT().GetResource(gomock.Any(), gomock.Any())
			},
		},
		{
			Method:      "PUT",
			Description: "UpdateResource",
			URI:         baseURI + "/test-application/test-resource-type/test-resource",
			Expect: func(mock *resourceprovider.MockResourceProvider) *gomock.Call {
				return mock.EXPECT().UpdateResource(gomock.Any(), gomock.Any(), gomock.Any())
			},

			// We don't need to include any significant data in the body, we're using mocks for testing
			// here. We just want to cover the code path.
			Body: map[string]interface{}{
				"tags": map[string]interface{}{
					"test-tag": "test-value",
				},
			},
		},
		{
			Method:      "DELETE",
			Description: "DeleteResource",
			URI:         baseURI + "/test-application/test-resource-type/test-resource",
			Expect: func(mock *resourceprovider.MockResourceProvider) *gomock.Call {
				return mock.EXPECT().DeleteResource(gomock.Any(), gomock.Any())
			},
		},
		{
			Method:      "POST",
			Description: "ListSecrets",
			URI:         providerURI + "/listSecrets",
			Expect: func(mock *resourceprovider.MockResourceProvider) *gomock.Call {
				return mock.EXPECT().ListSecrets(gomock.Any(), gomock.Any())
			},

			Body: map[string]interface{}{
				"targetId": baseURI + "/test-application/test-resource-type/test-resource",
			},
		},
		{
			Method:      "GET",
			Description: "GetOperation",
			URI:         baseURI + "/test-application/test-resource-type/test-resource/OperationResults/test-operation",
			Expect: func(mock *resourceprovider.MockResourceProvider) *gomock.Call {
				return mock.EXPECT().GetOperation(gomock.Any(), gomock.Any())
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Description, func(t *testing.T) {
			t.Run("Success", func(t *testing.T) {
				test := start(t)
				if testcase.Method == "PUT" {
					testcase.Expect(test.rp).Times(1).DoAndReturn(func(ctx context.Context, id azresources.ResourceID, body []byte) (rest.Response, error) {
						return rest.NewOKResponse(map[string]interface{}{}), nil // Empty JSON
					})
				} else if testcase.Method == "POST" {
					testcase.Expect(test.rp).Times(1).DoAndReturn(func(ctx context.Context, input resourceprovider.ListSecretsInput) (rest.Response, error) {
						return rest.NewOKResponse(map[string]interface{}{}), nil // Empty JSON
					})
				} else {
					testcase.Expect(test.rp).Times(1).DoAndReturn(func(ctx context.Context, id azresources.ResourceID) (rest.Response, error) {
						return rest.NewOKResponse(map[string]interface{}{}), nil // Empty JSON
					})
				}

				var err error
				body := []byte{}
				if testcase.Body != nil {
					body, err = json.Marshal(testcase.Body)
					require.NoError(t, err)
				}

				req := httptest.NewRequest(testcase.Method, testcase.URI, bytes.NewBuffer(body))
				w := httptest.NewRecorder()

				test.handler.ServeHTTP(w, req)

				require.Equal(t, 200, w.Code)
				requireJSON(t, map[string]interface{}{}, w)
			})

			t.Run("Error", func(t *testing.T) {
				test := start(t)
				if testcase.Method == "PUT" {
					testcase.Expect(test.rp).Times(1).DoAndReturn(func(ctx context.Context, id azresources.ResourceID, body []byte) (rest.Response, error) {
						return nil, fmt.Errorf("error!")
					})
				} else if testcase.Method == "POST" {
					testcase.Expect(test.rp).Times(1).DoAndReturn(func(ctx context.Context, input resourceprovider.ListSecretsInput) (rest.Response, error) {
						return nil, fmt.Errorf("error!")
					})
				} else {
					testcase.Expect(test.rp).Times(1).DoAndReturn(func(ctx context.Context, id azresources.ResourceID) (rest.Response, error) {
						return nil, fmt.Errorf("error!")
					})
				}

				var err error
				body := []byte{}
				if testcase.Body != nil {
					body, err = json.Marshal(testcase.Body)
					require.NoError(t, err)
				}

				req := httptest.NewRequest(testcase.Method, testcase.URI, bytes.NewBuffer(body))
				w := httptest.NewRecorder()

				test.handler.ServeHTTP(w, req)

				require.Equal(t, 500, w.Code)
				requireJSON(t, &armerrors.ErrorResponse{
					Error: armerrors.ErrorDetails{
						Message: "error!",
					},
				}, w)
			})

			t.Run("Write-Failure", func(t *testing.T) {
				test := start(t)
				if testcase.Method == "PUT" {
					testcase.Expect(test.rp).Times(1).DoAndReturn(func(ctx context.Context, id azresources.ResourceID, body []byte) (rest.Response, error) {
						return &FaultingResponse{}, nil
					})
				} else if testcase.Method == "POST" {
					testcase.Expect(test.rp).Times(1).DoAndReturn(func(ctx context.Context, input resourceprovider.ListSecretsInput) (rest.Response, error) {
						return &FaultingResponse{}, nil
					})
				} else {
					testcase.Expect(test.rp).Times(1).DoAndReturn(func(ctx context.Context, id azresources.ResourceID) (rest.Response, error) {
						return &FaultingResponse{}, nil
					})
				}

				var err error
				body := []byte{}
				if testcase.Body != nil {
					body, err = json.Marshal(testcase.Body)
					require.NoError(t, err)
				}

				req := httptest.NewRequest(testcase.Method, testcase.URI, bytes.NewBuffer(body))
				w := httptest.NewRecorder()

				test.handler.ServeHTTP(w, req)

				require.Equal(t, 500, w.Code)
				requireJSON(t, &armerrors.ErrorResponse{
					Error: armerrors.ErrorDetails{
						Message: "write failure!",
					},
				}, w)
			})

			// Remaining sub-tests are for PUT methods - they deal with the request body.
			if testcase.Method != "PUT" {
				return
			}

			t.Run("Validator-Missing", func(t *testing.T) {
				test := start(t)

				// Simulate an unknown type
				test.validator.RejectType = true

				var err error
				body := []byte{}
				if testcase.Body != nil {
					body, err = json.Marshal(testcase.Body)
					require.NoError(t, err)
				}

				req := httptest.NewRequest(testcase.Method, testcase.URI, bytes.NewBuffer(body))
				w := httptest.NewRecorder()

				test.handler.ServeHTTP(w, req)

				require.Equal(t, 400, w.Code)
				requireJSON(t, &armerrors.ErrorResponse{
					Error: armerrors.ErrorDetails{
						Code:    armerrors.Invalid,
						Message: "unsupported type",
					},
				}, w)
			})

			t.Run("Validation-Failure", func(t *testing.T) {
				test := start(t)

				// Simulate a validation failure
				test.validator.Errors = []schema.ValidationError{
					{
						Position: "test-position1",
						Message:  "test-message1",
					},
					{
						Position: "test-position2",
						Message:  "test-message2",
					},
					{
						JSONError: errors.New("test-error3"),
						Message:   "test-message3",
					},
				}

				var err error
				body := []byte{}
				if testcase.Body != nil {
					body, err = json.Marshal(testcase.Body)
					require.NoError(t, err)
				}

				req := httptest.NewRequest(testcase.Method, testcase.URI, bytes.NewBuffer(body))
				w := httptest.NewRecorder()

				test.handler.ServeHTTP(w, req)

				require.Equal(t, 400, w.Code)
				requireJSON(t, &armerrors.ErrorResponse{
					Error: armerrors.ErrorDetails{
						Code:    armerrors.Invalid,
						Message: "Validation error",
						Details: []armerrors.ErrorDetails{
							{
								Message: "test-position1: test-message1",
							},
							{
								Message: "test-position2: test-message2",
							},
							{
								Message: "test-message3: test-error3",
							},
						},
					},
				}, w)
			})
		})
	}
}
