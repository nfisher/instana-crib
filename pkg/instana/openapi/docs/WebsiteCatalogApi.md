# \WebsiteCatalogApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebsiteCatalogMetrics**](WebsiteCatalogApi.md#GetWebsiteCatalogMetrics) | **Get** /api/website-monitoring/catalog/metrics | Metric catalog
[**GetWebsiteCatalogTags**](WebsiteCatalogApi.md#GetWebsiteCatalogTags) | **Get** /api/website-monitoring/catalog/tags | Filter tag catalog



## GetWebsiteCatalogMetrics

> []WebsiteMonitoringMetricDescription GetWebsiteCatalogMetrics(ctx, )

Metric catalog

This endpoint retrieves all available metric definitions for website monitoring. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]WebsiteMonitoringMetricDescription**](WebsiteMonitoringMetricDescription.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebsiteCatalogTags

> []Tag GetWebsiteCatalogTags(ctx, )

Filter tag catalog

This endpoint retrieves all available tags for your monitored system.  These tags can be used to group metric results. ``` \"group\": {   \"groupbyTag\": \"beacon.page.name\" } ```  These tags can be used to filter metric results. ``` \"tagFilters\": [{  \"name\": \"beacon.website.name\",  \"operator\": \"EQUALS\",  \"value\": \"example\" }] ``` 

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

