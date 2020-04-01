# \InfrastructureCatalogApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInfrastructureCatalogMetrics**](InfrastructureCatalogApi.md#GetInfrastructureCatalogMetrics) | **Get** /api/infrastructure-monitoring/catalog/metrics/{plugin} | Get metric catalog
[**GetInfrastructureCatalogPlugins**](InfrastructureCatalogApi.md#GetInfrastructureCatalogPlugins) | **Get** /api/infrastructure-monitoring/catalog/plugins | Get plugin catalog
[**GetInfrastructureCatalogSearchFields**](InfrastructureCatalogApi.md#GetInfrastructureCatalogSearchFields) | **Get** /api/infrastructure-monitoring/catalog/search | get search field catalog



## GetInfrastructureCatalogMetrics

> []MetricInstance GetInfrastructureCatalogMetrics(ctx, plugin, optional)

Get metric catalog

This endpoint retrieves all available metric definitions of the requested plugin.  ### Path Parameters:  **plugin** The plugin id from [available plugins](#operation/getInfrastructureCatalogPlugins)  ### Optional Parameters:  **filter** You can restrict the returned metric definitions by passing a filter.  * `custom` to retrieve custom metric definitions only. * `builtin` to retrieve built-in metric definitions only. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plugin** | **string**|  | 
 **optional** | ***GetInfrastructureCatalogMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInfrastructureCatalogMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**|  | 

### Return type

[**[]MetricInstance**](MetricInstance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureCatalogPlugins

> []PluginResult GetInfrastructureCatalogPlugins(ctx, )

Get plugin catalog

This endpoint retrieves all available plugin ids for your monitored system. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]PluginResult**](PluginResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureCatalogSearchFields

> []SearchFieldResult GetInfrastructureCatalogSearchFields(ctx, )

get search field catalog

This endpoint retrieves all available search keywords for dynamic focus queries.  These search fields can be accessed via lucene queries. Each field belongs to a context, e.g. to entity, trace or event data. Some fields contain a set of possible fixed values, in this case a deviant value is invalid.  ``` ?query={keyword}:{value} ``` 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SearchFieldResult**](SearchFieldResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

