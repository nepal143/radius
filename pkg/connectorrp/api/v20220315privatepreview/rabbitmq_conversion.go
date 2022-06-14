// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package v20220315privatepreview

import (
	"github.com/project-radius/radius/pkg/armrpc/api/conv"
	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/connectorrp/datamodel"

	"github.com/Azure/go-autorest/autorest/to"
)

// ConvertTo converts from the versioned RabbitMQMessageQueue resource to version-agnostic datamodel.
func (src *RabbitMQMessageQueueResource) ConvertTo() (conv.DataModelInterface, error) {
	secrets := datamodel.RabbitMQSecrets{}
	if src.Properties.Secrets != nil {
		secrets = datamodel.RabbitMQSecrets{
			ConnectionString: to.String(src.Properties.Secrets.ConnectionString),
		}
	}

	converted := &datamodel.RabbitMQMessageQueue{
		TrackedResource: v1.TrackedResource{
			ID:       to.String(src.ID),
			Name:     to.String(src.Name),
			Type:     to.String(src.Type),
			Location: to.String(src.Location),
			Tags:     to.StringMap(src.Tags),
		},
		Properties: datamodel.RabbitMQMessageQueueProperties{
			BasicResourceProperties: v1.BasicResourceProperties{
				Status: v1.ResourceStatus{
					OutputResources: src.Properties.BasicResourceProperties.Status.OutputResources,
				},
			},
			ProvisioningState: toProvisioningStateDataModel(src.Properties.ProvisioningState),
			Environment:       to.String(src.Properties.Environment),
			Application:       to.String(src.Properties.Application),
			Queue:             to.String(src.Properties.Queue),
			Secrets:           secrets,
		},
		InternalMetadata: v1.InternalMetadata{
			UpdatedAPIVersion: Version,
		},
	}
	return converted, nil
}

// ConvertFrom converts from version-agnostic datamodel to the versioned RabbitMQMessageQueue resource.
func (dst *RabbitMQMessageQueueResource) ConvertFrom(src conv.DataModelInterface) error {
	rabbitmq, ok := src.(*datamodel.RabbitMQMessageQueue)
	if !ok {
		return conv.ErrInvalidModelConversion
	}

	dst.ID = to.StringPtr(rabbitmq.ID)
	dst.Name = to.StringPtr(rabbitmq.Name)
	dst.Type = to.StringPtr(rabbitmq.Type)
	dst.SystemData = fromSystemDataModel(rabbitmq.SystemData)
	dst.Location = to.StringPtr(rabbitmq.Location)
	dst.Tags = *to.StringMapPtr(rabbitmq.Tags)
	dst.Properties = &RabbitMQMessageQueueProperties{
		BasicResourceProperties: BasicResourceProperties{
			Status: &ResourceStatus{
				OutputResources: rabbitmq.Properties.BasicResourceProperties.Status.OutputResources,
			},
		},
		ProvisioningState: fromProvisioningStateDataModel(rabbitmq.Properties.ProvisioningState),
		Environment:       to.StringPtr(rabbitmq.Properties.Environment),
		Application:       to.StringPtr(rabbitmq.Properties.Application),
		Queue:             to.StringPtr(rabbitmq.Properties.Queue),
	}
	if (rabbitmq.Properties.Secrets != datamodel.RabbitMQSecrets{}) {
		dst.Properties.Secrets = &RabbitMQMessageQueuePropertiesSecrets{
			ConnectionString: to.StringPtr(rabbitmq.Properties.Secrets.ConnectionString),
		}
	}

	return nil
}