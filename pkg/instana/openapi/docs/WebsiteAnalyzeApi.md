# \WebsiteAnalyzeApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBeaconGroups**](WebsiteAnalyzeApi.md#GetBeaconGroups) | **Post** /api/website-monitoring/analyze/beacon-groups | Get grouped beacon metrics
[**GetBeacons**](WebsiteAnalyzeApi.md#GetBeacons) | **Post** /api/website-monitoring/analyze/beacons | Get all beacons



## GetBeaconGroups

> BeaconGroupsResult GetBeaconGroups(ctx, optional)

Get grouped beacon metrics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBeaconGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBeaconGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fillTimeSeries** | **optional.Bool**|  | 
 **getWebsiteBeaconGroups** | [**optional.Interface of GetWebsiteBeaconGroups**](GetWebsiteBeaconGroups.md)|  | 

### Return type

[**BeaconGroupsResult**](BeaconGroupsResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBeacons

> BeaconResult GetBeacons(ctx, optional)

Get all beacons

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBeaconsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBeaconsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getWebsiteBeacons** | [**optional.Interface of GetWebsiteBeacons**](GetWebsiteBeacons.md)|  | 

### Return type

[**BeaconResult**](BeaconResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

