# \HealthApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHealthState**](HealthApi.md#GetHealthState) | **Get** /api/instana/health | Basic health traffic light
[**GetVersion**](HealthApi.md#GetVersion) | **Get** /api/instana/version | API version information



## GetHealthState

> HealthState GetHealthState(ctx, )

Basic health traffic light

The returned JSON object will provide a health property which contains a simple traffic light (GREEN/YELLO/RED). For any non-Green-state a list  of reasons will be provided in the messages array.  Possible messages: * No data being processed * No data arriving from agents 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**HealthState**](HealthState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersion

> InstanaVersionInfo GetVersion(ctx, )

API version information

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InstanaVersionInfo**](InstanaVersionInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

