# \ApplicationResourcesApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationResourcesEndpoints**](ApplicationResourcesApi.md#ApplicationResourcesEndpoints) | **Get** /api/application-monitoring/applications/services/endpoints | Get endpoints
[**GetApplicationServices**](ApplicationResourcesApi.md#GetApplicationServices) | **Get** /api/application-monitoring/applications/services | Get applications/services
[**GetApplications**](ApplicationResourcesApi.md#GetApplications) | **Get** /api/application-monitoring/applications | Get applications
[**GetServices**](ApplicationResourcesApi.md#GetServices) | **Get** /api/application-monitoring/services | Get services



## ApplicationResourcesEndpoints

> EndpointResult ApplicationResourcesEndpoints(ctx, optional)

Get endpoints

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationResourcesEndpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationResourcesEndpointsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameFilter** | **optional.String**|  | 
 **types** | [**optional.Interface of []string**](string.md)|  | 
 **technologies** | [**optional.Interface of []string**](string.md)|  | 
 **windowSize** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**EndpointResult**](EndpointResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationServices

> ServiceResult GetApplicationServices(ctx, optional)

Get applications/services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetApplicationServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationServicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameFilter** | **optional.String**|  | 
 **windowSize** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**ServiceResult**](ServiceResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> ApplicationResult GetApplications(ctx, optional)

Get applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameFilter** | **optional.String**|  | 
 **windowSize** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**ApplicationResult**](ApplicationResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServices

> ServiceResult GetServices(ctx, optional)

Get services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetServicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameFilter** | **optional.String**|  | 
 **windowSize** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**ServiceResult**](ServiceResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

