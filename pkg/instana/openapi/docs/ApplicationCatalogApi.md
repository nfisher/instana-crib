# \ApplicationCatalogApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationCatalogMetrics**](ApplicationCatalogApi.md#GetApplicationCatalogMetrics) | **Get** /api/application-monitoring/catalog/metrics | Get Metric catalog
[**GetApplicationCatalogTags**](ApplicationCatalogApi.md#GetApplicationCatalogTags) | **Get** /api/application-monitoring/catalog/tags | Get filter tag catalog



## GetApplicationCatalogMetrics

> []MetricDescription GetApplicationCatalogMetrics(ctx, )

Get Metric catalog

This endpoint retrieves all available metric definitions for application monitoring. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]MetricDescription**](MetricDescription.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCatalogTags

> []Tag GetApplicationCatalogTags(ctx, )

Get filter tag catalog

This endpoint retrieves all available tags for your monitored system.  These tags can be used to group metric results. ``` \"group\": {   \"groupbyTag\": \"service.name\" } ```  These tags can be used to filter metric results. ``` \"tagFilters\": [{  \"name\": \"application.name\",  \"operator\": \"EQUALS\",  \"value\": \"example\" }] ``` 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Tag**](Tag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

