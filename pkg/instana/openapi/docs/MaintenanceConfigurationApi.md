# \MaintenanceConfigurationApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMaintenanceConfig**](MaintenanceConfigurationApi.md#DeleteMaintenanceConfig) | **Delete** /api/settings/maintenance/{id} | Delete maintenance configuration
[**GetMaintenanceConfig**](MaintenanceConfigurationApi.md#GetMaintenanceConfig) | **Get** /api/settings/maintenance/{id} | Maintenance configuration
[**GetMaintenanceConfigs**](MaintenanceConfigurationApi.md#GetMaintenanceConfigs) | **Get** /api/settings/maintenance | All maintenance configurations
[**PutMaintenanceConfig**](MaintenanceConfigurationApi.md#PutMaintenanceConfig) | **Put** /api/settings/maintenance/{id} | Create or update maintenance configuration



## DeleteMaintenanceConfig

> DeleteMaintenanceConfig(ctx, id)

Delete maintenance configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## GetMaintenanceConfig

> MaintenanceConfigWithLastUpdated GetMaintenanceConfig(ctx, id)

Maintenance configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**MaintenanceConfigWithLastUpdated**](MaintenanceConfigWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaintenanceConfigs

> []ValidatedMaintenanceConfigWithStatus GetMaintenanceConfigs(ctx, )

All maintenance configurations

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ValidatedMaintenanceConfigWithStatus**](ValidatedMaintenanceConfigWithStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMaintenanceConfig

> PutMaintenanceConfig(ctx, id, maintenanceConfig)

Create or update maintenance configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**maintenanceConfig** | [**MaintenanceConfig**](MaintenanceConfig.md)|  | 

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

