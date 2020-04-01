# \InfrastructureResourcesApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInfrastructureViewTree**](InfrastructureResourcesApi.md#GetInfrastructureViewTree) | **Get** /api/infrastructure-monitoring/graph/views | Get view tree
[**GetMonitoringState**](InfrastructureResourcesApi.md#GetMonitoringState) | **Get** /api/infrastructure-monitoring/monitoring-state | Monitored host count
[**GetRelatedHosts**](InfrastructureResourcesApi.md#GetRelatedHosts) | **Get** /api/infrastructure-monitoring/graph/related-hosts/{snapshotId} | Related hosts
[**SoftwareVersions**](InfrastructureResourcesApi.md#SoftwareVersions) | **Get** /api/infrastructure-monitoring/software/versions | Get installed software



## GetInfrastructureViewTree

> TreeNodeResult GetInfrastructureViewTree(ctx, )

Get view tree

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**TreeNodeResult**](TreeNodeResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitoringState

> MonitoringState GetMonitoringState(ctx, )

Monitored host count

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**MonitoringState**](MonitoringState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelatedHosts

> []string GetRelatedHosts(ctx, snapshotId)

Related hosts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string**|  | 

### Return type

**[]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoftwareVersions

> []SoftwareVersion SoftwareVersions(ctx, optional)

Get installed software

Retrieve information about the software you are running. This includes runtime and package manager information.  The `name`, `version`, `origin` and `type` parameters are optional filters that can be used to reduce the result data set.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwareVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SoftwareVersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time** | **optional.Int64**|  | 
 **origin** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **version** | **optional.String**|  | 

### Return type

[**[]SoftwareVersion**](SoftwareVersion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

