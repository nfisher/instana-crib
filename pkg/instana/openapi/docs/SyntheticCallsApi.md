# \SyntheticCallsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSyntheticCall**](SyntheticCallsApi.md#DeleteSyntheticCall) | **Delete** /api/settings/synthetic-calls | Delete synthetic call configurations
[**GetSyntheticCalls**](SyntheticCallsApi.md#GetSyntheticCalls) | **Get** /api/settings/synthetic-calls | Synthetic call configurations
[**UpdateSyntheticCall**](SyntheticCallsApi.md#UpdateSyntheticCall) | **Put** /api/settings/synthetic-calls | Update synthetic call configurations



## DeleteSyntheticCall

> DeleteSyntheticCall(ctx, )

Delete synthetic call configurations

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyntheticCalls

> SyntheticCallWithDefaultsConfig GetSyntheticCalls(ctx, )

Synthetic call configurations

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SyntheticCallWithDefaultsConfig**](SyntheticCallWithDefaultsConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyntheticCall

> UpdateSyntheticCall(ctx, syntheticCallConfig)

Update synthetic call configurations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**syntheticCallConfig** | [**SyntheticCallConfig**](SyntheticCallConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

