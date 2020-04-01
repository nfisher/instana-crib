# \APITokenApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiToken**](APITokenApi.md#DeleteApiToken) | **Delete** /api/settings/api-tokens/{apiTokenId} | Delete API token
[**GetApiToken**](APITokenApi.md#GetApiToken) | **Get** /api/settings/api-tokens/{apiTokenId} | API token
[**GetApiTokens**](APITokenApi.md#GetApiTokens) | **Get** /api/settings/api-tokens | All API tokes
[**PutApiToken**](APITokenApi.md#PutApiToken) | **Put** /api/settings/api-tokens/{apiTokenId} | Create or update an API token



## DeleteApiToken

> DeleteApiToken(ctx, apiTokenId)

Delete API token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiTokenId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiToken

> ApiToken GetApiToken(ctx, apiTokenId)

API token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiTokenId** | **string**|  | 

### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiTokens

> []ApiToken GetApiTokens(ctx, )

All API tokes

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ApiToken**](ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApiToken

> ApiToken PutApiToken(ctx, apiTokenId, apiToken)

Create or update an API token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiTokenId** | **string**|  | 
**apiToken** | [**ApiToken**](ApiToken.md)|  | 

### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

