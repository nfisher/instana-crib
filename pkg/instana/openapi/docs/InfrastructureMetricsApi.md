# \InfrastructureMetricsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInfrastructureMetrics**](InfrastructureMetricsApi.md#GetInfrastructureMetrics) | **Post** /api/infrastructure-monitoring/metrics | Get infrastructure metrics
[**GetSnapshot**](InfrastructureMetricsApi.md#GetSnapshot) | **Get** /api/infrastructure-monitoring/snapshots/{id} | Get snapshot details
[**GetSnapshots**](InfrastructureMetricsApi.md#GetSnapshots) | **Get** /api/infrastructure-monitoring/snapshots | Search snapshots



## GetInfrastructureMetrics

> InfrastructureMetricResult GetInfrastructureMetrics(ctx, optional)

Get infrastructure metrics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetInfrastructureMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInfrastructureMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **optional.Bool**|  | 
 **offline** | **optional.Bool**|  | 
 **getCombinedMetrics** | [**optional.Interface of GetCombinedMetrics**](GetCombinedMetrics.md)|  | 

### Return type

[**InfrastructureMetricResult**](InfrastructureMetricResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshot

> SnapshotItem GetSnapshot(ctx, id, optional)

Get snapshot details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***GetSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSnapshotOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **optional.Int64**|  | 
 **windowSize** | **optional.Int64**|  | 

### Return type

[**SnapshotItem**](SnapshotItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshots

> SnapshotResult GetSnapshots(ctx, optional)

Search snapshots

These APIs can be used to retrieve information about hosts, processes, JVMs and other entities that we are calling snapshots. A snapshot represents static information about an entity as it was at a specific point in time. To clarify:  **Static information** is any information which is seldom changing, e.g. process IDs, host FQDNs or a list of host hard disks. The counterpart to static information are metrics which have a much higher change rate, e.g. host CPU usage or JVM garbage collection activity. Snapshots only contain static information.  - Snapshots are **versioned** and represent an entity's state for a specific point in time. While snapshots only contain static information, even that information may change. For example you may add another hard disk to a server. For such a change, a new snapshot would be created.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSnapshotsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSnapshotsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**|  | 
 **to** | **optional.Int64**|  | 
 **windowSize** | **optional.Int64**|  | 
 **size** | **optional.Int32**|  | 
 **plugin** | **optional.String**|  | 
 **offline** | **optional.Bool**|  | 

### Return type

[**SnapshotResult**](SnapshotResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

