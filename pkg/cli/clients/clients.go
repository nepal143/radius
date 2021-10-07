// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package clients

import (
	"context"
	"io"
	"os"

	"github.com/Azure/radius/pkg/azure/radclient"
)

// DeploymentClient is used to deploy ARM-JSON templates (compiled Bicep output).
type DeploymentClient interface {
	Deploy(ctx context.Context, content string) error
}

// DiagnosticsClient is used to interface with diagnostics features like logs and port-forwards.
type DiagnosticsClient interface {
	Expose(ctx context.Context, options ExposeOptions) (failed chan error, stop chan struct{}, signals chan os.Signal, err error)
	Logs(ctx context.Context, options LogsOptions) ([]LogStream, error)
}

type ExposeOptions struct {
	Application string
	Resource    string
	Port        int
	RemotePort  int
	Replica     string
}

type LogsOptions struct {
	Application string
	Resource    string
	Follow      bool
	Container   string
	Replica     string
}

type LogStream struct {
	Name   string
	Stream io.ReadCloser
}

// ManagementClient is used to interface with management features like listing applications and resources.
type ManagementClient interface {
	ListApplications(ctx context.Context) (*radclient.ApplicationList, error)
	ShowApplication(ctx context.Context, applicationName string) (*radclient.ApplicationResource, error)
	DeleteApplication(ctx context.Context, applicationName string) error

	ShowResource(ctx context.Context, applicationName string, resourceType string, resourceName string) (interface{}, error)
	ListAllResourcesByApplication(ctx context.Context, applicationName string) (*radclient.RadiusResourceList, error)
}
