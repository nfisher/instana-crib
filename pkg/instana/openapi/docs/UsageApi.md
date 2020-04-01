# \UsageApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllUsage**](UsageApi.md#GetAllUsage) | **Get** /api/instana/usage/api | API usage by customer
[**GetHostsPerDay**](UsageApi.md#GetHostsPerDay) | **Get** /api/instana/usage/hosts/{day}/{month}/{year} | Host count day / month / year
[**GetHostsPerMonth**](UsageApi.md#GetHostsPerMonth) | **Get** /api/instana/usage/hosts/{month}/{year} | Host count month / year
[**GetUsagePerDay**](UsageApi.md#GetUsagePerDay) | **Get** /api/instana/usage/api/{day}/{month}/{year} | API usage day / month / year
[**GetUsagePerMonth**](UsageApi.md#GetUsagePerMonth) | **Get** /api/instana/usage/api/{month}/{year} | API usage month / year



## GetAllUsage

> []UsageResult GetAllUsage(ctx, )

API usage by customer

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]UsageResult**](UsageResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostsPerDay

> []UsageResult GetHostsPerDay(ctx, day, month, year)

Host count day / month / year

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**day** | **int32**|  | 
**month** | **int32**|  | 
**year** | **int32**|  | 

### Return type

[**[]UsageResult**](UsageResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostsPerMonth

> []UsageResult GetHostsPerMonth(ctx, month, year)

Host count month / year

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**month** | **int32**|  | 
**year** | **int32**|  | 

### Return type

[**[]UsageResult**](UsageResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsagePerDay

> []UsageResult GetUsagePerDay(ctx, day, month, year)

API usage day / month / year

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**day** | **int32**|  | 
**month** | **int32**|  | 
**year** | **int32**|  | 

### Return type

[**[]UsageResult**](UsageResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsagePerMonth

> []UsageResult GetUsagePerMonth(ctx, month, year)

API usage month / year

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**month** | **int32**|  | 
**year** | **int32**|  | 

### Return type

[**[]UsageResult**](UsageResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

