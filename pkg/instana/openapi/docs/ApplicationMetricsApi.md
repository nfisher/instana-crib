# \ApplicationMetricsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationMetrics**](ApplicationMetricsApi.md#GetApplicationMetrics) | **Post** /api/application-monitoring/metrics/applications | Get Application Metrics
[**GetEndpointsMetrics**](ApplicationMetricsApi.md#GetEndpointsMetrics) | **Post** /api/application-monitoring/metrics/endpoints | Get Endpoint metrics
[**GetServicesMetrics**](ApplicationMetricsApi.md#GetServicesMetrics) | **Post** /api/application-monitoring/metrics/services | Get Service metrics



## GetApplicationMetrics

> ApplicationMetricResult GetApplicationMetrics(ctx, optional)

Get Application Metrics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetApplicationMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getApplications** | [**optional.Interface of GetApplications**](GetApplications.md)|  | 

### Return type

[**ApplicationMetricResult**](ApplicationMetricResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointsMetrics

> EndpointMetricResult GetEndpointsMetrics(ctx, optional)

Get Endpoint metrics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEndpointsMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEndpointsMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEndpoints** | [**optional.Interface of GetEndpoints**](GetEndpoints.md)|  | 

### Return type

[**EndpointMetricResult**](EndpointMetricResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicesMetrics

> ServiceMetricResult GetServicesMetrics(ctx, optional)

Get Service metrics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetServicesMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetServicesMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getServices** | [**optional.Interface of GetServices**](GetServices.md)|  | 

### Return type

[**ServiceMetricResult**](ServiceMetricResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

