# \ApplicationAnalyzeApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCallGroup**](ApplicationAnalyzeApi.md#GetCallGroup) | **Post** /api/application-monitoring/analyze/call-groups | Get grouped call metrics
[**GetTrace**](ApplicationAnalyzeApi.md#GetTrace) | **Get** /api/application-monitoring/analyze/traces/{id} | Get trace detail
[**GetTraceGroups**](ApplicationAnalyzeApi.md#GetTraceGroups) | **Post** /api/application-monitoring/analyze/trace-groups | Get grouped trace metrics
[**GetTraces**](ApplicationAnalyzeApi.md#GetTraces) | **Post** /api/application-monitoring/analyze/traces | Get all traces



## GetCallGroup

> CallGroupsResult GetCallGroup(ctx, optional)

Get grouped call metrics

This endpoint retrieves the metrics for calls.    **Manditory Paramters:**    **Optional Paramters:**    **Defaults:**    **Limits:**    **Tips:**  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCallGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCallGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fillTimeSeries** | **optional.Bool**|  | 
 **getCallGroups** | [**optional.Interface of GetCallGroups**](GetCallGroups.md)|  | 

### Return type

[**CallGroupsResult**](CallGroupsResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrace

> FullTrace GetTrace(ctx, id)

Get trace detail

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**FullTrace**](FullTrace.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceGroups

> TraceGroupsResult GetTraceGroups(ctx, optional)

Get grouped trace metrics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTraceGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTraceGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fillTimeSeries** | **optional.Bool**|  | 
 **getTraceGroups** | [**optional.Interface of GetTraceGroups**](GetTraceGroups.md)|  | 

### Return type

[**TraceGroupsResult**](TraceGroupsResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraces

> TraceResult GetTraces(ctx, optional)

Get all traces

This endpoint retrieves the metrics for traces.    **Manditory Paramters:**    **Optional Paramters:**    **Defaults:**    **Limits:**    **Tips:**  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTracesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTracesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getTraces** | [**optional.Interface of GetTraces**](GetTraces.md)|  | 

### Return type

[**TraceResult**](TraceResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

