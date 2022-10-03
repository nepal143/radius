// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------
package resourcegroups

import (
	"context"
	"errors"
	"fmt"
	http "net/http"
	"strings"

	armrpc_rest "github.com/project-radius/radius/pkg/armrpc/rest"
	"github.com/project-radius/radius/pkg/middleware"
	ctrl "github.com/project-radius/radius/pkg/ucp/frontend/controller"
	"github.com/project-radius/radius/pkg/ucp/resources"
	"github.com/project-radius/radius/pkg/ucp/rest"
	"github.com/project-radius/radius/pkg/ucp/store"
	"github.com/project-radius/radius/pkg/ucp/ucplog"
)

var _ ctrl.Controller = (*GetResourceGroup)(nil)

// GetResourceGroup is the controller implementation to get the details of a UCP resource group
type GetResourceGroup struct {
	ctrl.BaseController
}

// NewGetResourceGroup creates a new GetResourceGroup.
func NewGetResourceGroup(opts ctrl.Options) (ctrl.Controller, error) {
	return &GetResourceGroup{ctrl.NewBaseController(opts)}, nil
}

func (r *GetResourceGroup) Run(ctx context.Context, w http.ResponseWriter, req *http.Request) (armrpc_rest.Response, error) {
	path := middleware.GetRelativePath(r.Options.BasePath, req.URL.Path)
	logger := ucplog.GetLogger(ctx)
	id := strings.ToLower(path)
	resourceID, err := resources.ParseScope(id)
	if err != nil {
		return armrpc_rest.NewBadRequestResponse(err.Error()), nil
	}

	logger.Info(fmt.Sprintf("Getting resource group %s from db", resourceID))
	rg := rest.ResourceGroup{}
	_, err = r.GetResource(ctx, resourceID.String(), &rg)
	if err != nil {
		if errors.Is(err, &store.ErrNotFound{}) {
			logger.Info(fmt.Sprintf("Resource group %s not found in db", resourceID))
			restResponse := armrpc_rest.NewNotFoundResponse(resourceID)
			return restResponse, nil
		}
		return nil, err
	}
	restResponse := armrpc_rest.NewOKResponse(rg)
	return restResponse, nil
}
