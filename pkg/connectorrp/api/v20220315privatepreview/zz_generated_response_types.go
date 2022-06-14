//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20220315privatepreview

import "net/http"

// DaprInvokeHTTPRoutesCreateOrUpdateResponse contains the response from method DaprInvokeHTTPRoutes.CreateOrUpdate.
type DaprInvokeHTTPRoutesCreateOrUpdateResponse struct {
	DaprInvokeHTTPRoutesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprInvokeHTTPRoutesCreateOrUpdateResult contains the result from method DaprInvokeHTTPRoutes.CreateOrUpdate.
type DaprInvokeHTTPRoutesCreateOrUpdateResult struct {
	DaprInvokeHTTPRouteResource
}

// DaprInvokeHTTPRoutesDeleteResponse contains the response from method DaprInvokeHTTPRoutes.Delete.
type DaprInvokeHTTPRoutesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprInvokeHTTPRoutesGetResponse contains the response from method DaprInvokeHTTPRoutes.Get.
type DaprInvokeHTTPRoutesGetResponse struct {
	DaprInvokeHTTPRoutesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprInvokeHTTPRoutesGetResult contains the result from method DaprInvokeHTTPRoutes.Get.
type DaprInvokeHTTPRoutesGetResult struct {
	DaprInvokeHTTPRouteResource
}

// DaprInvokeHTTPRoutesListByRootScopeResponse contains the response from method DaprInvokeHTTPRoutes.ListByRootScope.
type DaprInvokeHTTPRoutesListByRootScopeResponse struct {
	DaprInvokeHTTPRoutesListByRootScopeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprInvokeHTTPRoutesListByRootScopeResult contains the result from method DaprInvokeHTTPRoutes.ListByRootScope.
type DaprInvokeHTTPRoutesListByRootScopeResult struct {
	DaprInvokeHTTPRouteList
}

// DaprPubSubBrokersCreateOrUpdateResponse contains the response from method DaprPubSubBrokers.CreateOrUpdate.
type DaprPubSubBrokersCreateOrUpdateResponse struct {
	DaprPubSubBrokersCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprPubSubBrokersCreateOrUpdateResult contains the result from method DaprPubSubBrokers.CreateOrUpdate.
type DaprPubSubBrokersCreateOrUpdateResult struct {
	DaprPubSubBrokerResource
}

// DaprPubSubBrokersDeleteResponse contains the response from method DaprPubSubBrokers.Delete.
type DaprPubSubBrokersDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprPubSubBrokersGetResponse contains the response from method DaprPubSubBrokers.Get.
type DaprPubSubBrokersGetResponse struct {
	DaprPubSubBrokersGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprPubSubBrokersGetResult contains the result from method DaprPubSubBrokers.Get.
type DaprPubSubBrokersGetResult struct {
	DaprPubSubBrokerResource
}

// DaprPubSubBrokersListByRootScopeResponse contains the response from method DaprPubSubBrokers.ListByRootScope.
type DaprPubSubBrokersListByRootScopeResponse struct {
	DaprPubSubBrokersListByRootScopeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprPubSubBrokersListByRootScopeResult contains the result from method DaprPubSubBrokers.ListByRootScope.
type DaprPubSubBrokersListByRootScopeResult struct {
	DaprPubSubBrokerList
}

// DaprSecretStoresCreateOrUpdateResponse contains the response from method DaprSecretStores.CreateOrUpdate.
type DaprSecretStoresCreateOrUpdateResponse struct {
	DaprSecretStoresCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprSecretStoresCreateOrUpdateResult contains the result from method DaprSecretStores.CreateOrUpdate.
type DaprSecretStoresCreateOrUpdateResult struct {
	DaprSecretStoreResource
}

// DaprSecretStoresDeleteResponse contains the response from method DaprSecretStores.Delete.
type DaprSecretStoresDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprSecretStoresGetResponse contains the response from method DaprSecretStores.Get.
type DaprSecretStoresGetResponse struct {
	DaprSecretStoresGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprSecretStoresGetResult contains the result from method DaprSecretStores.Get.
type DaprSecretStoresGetResult struct {
	DaprSecretStoreResource
}

// DaprSecretStoresListByRootScopeResponse contains the response from method DaprSecretStores.ListByRootScope.
type DaprSecretStoresListByRootScopeResponse struct {
	DaprSecretStoresListByRootScopeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprSecretStoresListByRootScopeResult contains the result from method DaprSecretStores.ListByRootScope.
type DaprSecretStoresListByRootScopeResult struct {
	DaprSecretStoreList
}

// DaprStateStoresCreateOrUpdateResponse contains the response from method DaprStateStores.CreateOrUpdate.
type DaprStateStoresCreateOrUpdateResponse struct {
	DaprStateStoresCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprStateStoresCreateOrUpdateResult contains the result from method DaprStateStores.CreateOrUpdate.
type DaprStateStoresCreateOrUpdateResult struct {
	DaprStateStoreResource
}

// DaprStateStoresDeleteResponse contains the response from method DaprStateStores.Delete.
type DaprStateStoresDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprStateStoresGetResponse contains the response from method DaprStateStores.Get.
type DaprStateStoresGetResponse struct {
	DaprStateStoresGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprStateStoresGetResult contains the result from method DaprStateStores.Get.
type DaprStateStoresGetResult struct {
	DaprStateStoreResource
}

// DaprStateStoresListByRootScopeResponse contains the response from method DaprStateStores.ListByRootScope.
type DaprStateStoresListByRootScopeResponse struct {
	DaprStateStoresListByRootScopeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprStateStoresListByRootScopeResult contains the result from method DaprStateStores.ListByRootScope.
type DaprStateStoresListByRootScopeResult struct {
	DaprStateStoreList
}

// ExtendersCreateOrUpdateResponse contains the response from method Extenders.CreateOrUpdate.
type ExtendersCreateOrUpdateResponse struct {
	ExtendersCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtendersCreateOrUpdateResult contains the result from method Extenders.CreateOrUpdate.
type ExtendersCreateOrUpdateResult struct {
	ExtenderResource
}

// ExtendersDeleteResponse contains the response from method Extenders.Delete.
type ExtendersDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtendersGetResponse contains the response from method Extenders.Get.
type ExtendersGetResponse struct {
	ExtendersGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtendersGetResult contains the result from method Extenders.Get.
type ExtendersGetResult struct {
	ExtenderResource
}

// ExtendersListByRootScopeResponse contains the response from method Extenders.ListByRootScope.
type ExtendersListByRootScopeResponse struct {
	ExtendersListByRootScopeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtendersListByRootScopeResult contains the result from method Extenders.ListByRootScope.
type ExtendersListByRootScopeResult struct {
	ExtenderList
}

// MongoDatabasesCreateOrUpdateResponse contains the response from method MongoDatabases.CreateOrUpdate.
type MongoDatabasesCreateOrUpdateResponse struct {
	MongoDatabasesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MongoDatabasesCreateOrUpdateResult contains the result from method MongoDatabases.CreateOrUpdate.
type MongoDatabasesCreateOrUpdateResult struct {
	MongoDatabaseResource
}

// MongoDatabasesDeleteResponse contains the response from method MongoDatabases.Delete.
type MongoDatabasesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MongoDatabasesGetResponse contains the response from method MongoDatabases.Get.
type MongoDatabasesGetResponse struct {
	MongoDatabasesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MongoDatabasesGetResult contains the result from method MongoDatabases.Get.
type MongoDatabasesGetResult struct {
	MongoDatabaseResource
}

// MongoDatabasesListByRootScopeResponse contains the response from method MongoDatabases.ListByRootScope.
type MongoDatabasesListByRootScopeResponse struct {
	MongoDatabasesListByRootScopeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MongoDatabasesListByRootScopeResult contains the result from method MongoDatabases.ListByRootScope.
type MongoDatabasesListByRootScopeResult struct {
	MongoDatabaseList
}

// MongoDatabasesListSecretsResponse contains the response from method MongoDatabases.ListSecrets.
type MongoDatabasesListSecretsResponse struct {
	MongoDatabasesListSecretsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MongoDatabasesListSecretsResult contains the result from method MongoDatabases.ListSecrets.
type MongoDatabasesListSecretsResult struct {
	MongoDatabaseSecrets
}

// RabbitMQMessageQueuesCreateOrUpdateResponse contains the response from method RabbitMQMessageQueues.CreateOrUpdate.
type RabbitMQMessageQueuesCreateOrUpdateResponse struct {
	RabbitMQMessageQueuesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RabbitMQMessageQueuesCreateOrUpdateResult contains the result from method RabbitMQMessageQueues.CreateOrUpdate.
type RabbitMQMessageQueuesCreateOrUpdateResult struct {
	RabbitMQMessageQueueResource
}

// RabbitMQMessageQueuesDeleteResponse contains the response from method RabbitMQMessageQueues.Delete.
type RabbitMQMessageQueuesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RabbitMQMessageQueuesGetResponse contains the response from method RabbitMQMessageQueues.Get.
type RabbitMQMessageQueuesGetResponse struct {
	RabbitMQMessageQueuesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RabbitMQMessageQueuesGetResult contains the result from method RabbitMQMessageQueues.Get.
type RabbitMQMessageQueuesGetResult struct {
	RabbitMQMessageQueueResource
}

// RabbitMQMessageQueuesListByRootScopeResponse contains the response from method RabbitMQMessageQueues.ListByRootScope.
type RabbitMQMessageQueuesListByRootScopeResponse struct {
	RabbitMQMessageQueuesListByRootScopeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RabbitMQMessageQueuesListByRootScopeResult contains the result from method RabbitMQMessageQueues.ListByRootScope.
type RabbitMQMessageQueuesListByRootScopeResult struct {
	RabbitMQMessageQueueList
}

// RedisCachesCreateOrUpdateResponse contains the response from method RedisCaches.CreateOrUpdate.
type RedisCachesCreateOrUpdateResponse struct {
	RedisCachesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RedisCachesCreateOrUpdateResult contains the result from method RedisCaches.CreateOrUpdate.
type RedisCachesCreateOrUpdateResult struct {
	RedisCacheResource
}

// RedisCachesDeleteResponse contains the response from method RedisCaches.Delete.
type RedisCachesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RedisCachesGetResponse contains the response from method RedisCaches.Get.
type RedisCachesGetResponse struct {
	RedisCachesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RedisCachesGetResult contains the result from method RedisCaches.Get.
type RedisCachesGetResult struct {
	RedisCacheResource
}

// RedisCachesListByRootScopeResponse contains the response from method RedisCaches.ListByRootScope.
type RedisCachesListByRootScopeResponse struct {
	RedisCachesListByRootScopeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RedisCachesListByRootScopeResult contains the result from method RedisCaches.ListByRootScope.
type RedisCachesListByRootScopeResult struct {
	RedisCacheList
}

// SQLDatabasesCreateOrUpdateResponse contains the response from method SQLDatabases.CreateOrUpdate.
type SQLDatabasesCreateOrUpdateResponse struct {
	SQLDatabasesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLDatabasesCreateOrUpdateResult contains the result from method SQLDatabases.CreateOrUpdate.
type SQLDatabasesCreateOrUpdateResult struct {
	SQLDatabaseResource
}

// SQLDatabasesDeleteResponse contains the response from method SQLDatabases.Delete.
type SQLDatabasesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLDatabasesGetResponse contains the response from method SQLDatabases.Get.
type SQLDatabasesGetResponse struct {
	SQLDatabasesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLDatabasesGetResult contains the result from method SQLDatabases.Get.
type SQLDatabasesGetResult struct {
	SQLDatabaseResource
}

// SQLDatabasesListByRootScopeResponse contains the response from method SQLDatabases.ListByRootScope.
type SQLDatabasesListByRootScopeResponse struct {
	SQLDatabasesListByRootScopeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLDatabasesListByRootScopeResult contains the result from method SQLDatabases.ListByRootScope.
type SQLDatabasesListByRootScopeResult struct {
	SQLDatabaseList
}
