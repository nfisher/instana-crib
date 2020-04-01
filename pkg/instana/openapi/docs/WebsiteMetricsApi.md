# \WebsiteMetricsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBeaconMetrics**](WebsiteMetricsApi.md#GetBeaconMetrics) | **Post** /api/website-monitoring/metrics | Get beacon metrics
[**GetPageLoad**](WebsiteMetricsApi.md#GetPageLoad) | **Get** /api/website-monitoring/page-load | Get page load



## GetBeaconMetrics

> WebsiteMetricResult GetBeaconMetrics(ctx, optional)

Get beacon metrics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBeaconMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBeaconMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getWebsiteMetrics** | [**optional.Interface of GetWebsiteMetrics**](GetWebsiteMetrics.md)|  | 

### Return type

[**WebsiteMetricResult**](WebsiteMetricResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageLoad

> []WebsiteMonitoringBeacon GetPageLoad(ctx, id, timestamp)

Get page load

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**timestamp** | **int64**|  | 

### Return type

[**[]WebsiteMonitoringBeacon**](WebsiteMonitoringBeacon.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

