// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package resource_test

import (
	"testing"

	"github.com/Azure/radius/pkg/radrp/outputresource"
	"github.com/Azure/radius/pkg/radrp/rest"
	"github.com/Azure/radius/pkg/resourcekinds"
	"github.com/Azure/radius/test/kubernetestest"
	"github.com/Azure/radius/test/validation"
)

func Test_ContainerHttpBinding(t *testing.T) {
	template := "testdata/kubernetes-resources-container-httpbinding.bicep"
	application := "kubernetes-resources-container-httpbinding"
	test := kubernetestest.NewApplicationTest(t, application, []kubernetestest.Step{
		{
			Executor: kubernetestest.NewDeployStepExecutor(template),
			RadiusResources: &validation.ResourceSet{
				Resources: []validation.RadiusResource{
					{
						ApplicationName: application,
						ResourceName:    "frontend",
						OutputResources: map[string]validation.ExpectedOutputResource{
							outputresource.LocalIDDeployment: validation.NewOutputResource(outputresource.LocalIDDeployment, outputresource.TypeKubernetes, resourcekinds.Kubernetes, true, false, rest.OutputResourceStatus{}),
							outputresource.LocalIDService:    validation.NewOutputResource(outputresource.LocalIDService, outputresource.TypeKubernetes, resourcekinds.Kubernetes, true, false, rest.OutputResourceStatus{}),
							outputresource.LocalIDSecret:     validation.NewOutputResource(outputresource.LocalIDSecret, outputresource.TypeKubernetes, resourcekinds.Kubernetes, true, false, rest.OutputResourceStatus{}),
						},
					},
					{
						ApplicationName: application,
						ResourceName:    "backend",
						OutputResources: map[string]validation.ExpectedOutputResource{
							outputresource.LocalIDDeployment: validation.NewOutputResource(outputresource.LocalIDDeployment, outputresource.TypeKubernetes, resourcekinds.Kubernetes, true, false, rest.OutputResourceStatus{}),
							outputresource.LocalIDService:    validation.NewOutputResource(outputresource.LocalIDService, outputresource.TypeKubernetes, resourcekinds.Kubernetes, true, false, rest.OutputResourceStatus{}),
						},
					},
				},
			},
			Pods: &validation.K8sObjectSet{
				Namespaces: map[string][]validation.K8sObject{
					"default": {
						validation.NewK8sObjectForResource(application, "frontend"),
						validation.NewK8sObjectForResource(application, "backend"),
					},
				},
			},
		},
	})

	test.Test(t)
}

func Test_ContainerManualScale(t *testing.T) {
	t.Skip("Need to readd manual scale support")

	template := "testdata/kubernetes-resources-container-manualscale.bicep"
	application := "kubernetes-resources-container-manualscale"
	test := kubernetestest.NewApplicationTest(t, application, []kubernetestest.Step{
		{
			Executor: kubernetestest.NewDeployStepExecutor(template),
			RadiusResources: &validation.ResourceSet{
				Resources: []validation.RadiusResource{
					{
						ApplicationName: application,
						ResourceName:    "frontend",
						OutputResources: map[string]validation.ExpectedOutputResource{
							outputresource.LocalIDDeployment: validation.NewOutputResource(outputresource.LocalIDDeployment, outputresource.TypeKubernetes, resourcekinds.Kubernetes, true, false, rest.OutputResourceStatus{}),
							outputresource.LocalIDService:    validation.NewOutputResource(outputresource.LocalIDService, outputresource.TypeKubernetes, resourcekinds.Kubernetes, true, false, rest.OutputResourceStatus{}),
							outputresource.LocalIDSecret:     validation.NewOutputResource(outputresource.LocalIDSecret, outputresource.TypeKubernetes, resourcekinds.Kubernetes, true, false, rest.OutputResourceStatus{}),
						},
					},
					{
						ApplicationName: application,
						ResourceName:    "backend",
						OutputResources: map[string]validation.ExpectedOutputResource{
							outputresource.LocalIDDeployment: validation.NewOutputResource(outputresource.LocalIDDeployment, outputresource.TypeKubernetes, resourcekinds.Kubernetes, true, false, rest.OutputResourceStatus{}),
							outputresource.LocalIDService:    validation.NewOutputResource(outputresource.LocalIDService, outputresource.TypeKubernetes, resourcekinds.Kubernetes, true, false, rest.OutputResourceStatus{}),
						},
					},
				},
			},
			Pods: &validation.K8sObjectSet{
				Namespaces: map[string][]validation.K8sObject{
					"default": {
						validation.NewK8sObjectForResource(application, "frontend"),
						// Verify two backend pods are created.
						validation.NewK8sObjectForResource(application, "backend"),
						validation.NewK8sObjectForResource(application, "backend"),
					},
				},
			},
		},
	})

	test.Test(t)
}
