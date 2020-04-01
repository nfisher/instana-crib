# \EventSettingsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](EventSettingsApi.md#Create) | **Post** /api/events/settings/website-alert-configs | Create Website Alert Config
[**DeleteAlert**](EventSettingsApi.md#DeleteAlert) | **Delete** /api/events/settings/alerts/{id} | Delete alerting
[**DeleteAlertingChannel**](EventSettingsApi.md#DeleteAlertingChannel) | **Delete** /api/events/settings/alertingChannels/{id} | Delete alerting channel
[**DeleteBuiltInEventSpecification**](EventSettingsApi.md#DeleteBuiltInEventSpecification) | **Delete** /api/events/settings/event-specifications/built-in/{eventSpecificationId} | Delete built-in event specification
[**DeleteCustomEventSpecification**](EventSettingsApi.md#DeleteCustomEventSpecification) | **Delete** /api/events/settings/event-specifications/custom/{eventSpecificationId} | Delete custom event specification
[**DisableBuiltInEventSpecification**](EventSettingsApi.md#DisableBuiltInEventSpecification) | **Post** /api/events/settings/event-specifications/built-in/{eventSpecificationId}/disable | Disable built-in event specification
[**DisableCustomEventSpecification**](EventSettingsApi.md#DisableCustomEventSpecification) | **Post** /api/events/settings/event-specifications/custom/{eventSpecificationId}/disable | Disable custom event specification
[**EnableBuiltInEventSpecification**](EventSettingsApi.md#EnableBuiltInEventSpecification) | **Post** /api/events/settings/event-specifications/built-in/{eventSpecificationId}/enable | Enable built-in event specification
[**EnableCustomEventSpecification**](EventSettingsApi.md#EnableCustomEventSpecification) | **Post** /api/events/settings/event-specifications/custom/{eventSpecificationId}/enable | Enable custom event specification
[**Find**](EventSettingsApi.md#Find) | **Get** /api/events/settings/website-alert-configs/{id} | Get Website Alert Config
[**FindAllActive**](EventSettingsApi.md#FindAllActive) | **Get** /api/events/settings/website-alert-configs | All Website Alert Configs
[**FindVersions**](EventSettingsApi.md#FindVersions) | **Get** /api/events/settings/website-alert-configs/{id}/versions | Get versions of Website Alert Config
[**GetAlert**](EventSettingsApi.md#GetAlert) | **Get** /api/events/settings/alerts/{id} | Alerting
[**GetAlertingChannel**](EventSettingsApi.md#GetAlertingChannel) | **Get** /api/events/settings/alertingChannels/{id} | Alerting channel
[**GetAlertingChannels**](EventSettingsApi.md#GetAlertingChannels) | **Get** /api/events/settings/alertingChannels | All alerting channels
[**GetAlertingConfigurationInfos**](EventSettingsApi.md#GetAlertingConfigurationInfos) | **Get** /api/events/settings/alerts/infos | All alerting configuration info
[**GetAlerts**](EventSettingsApi.md#GetAlerts) | **Get** /api/events/settings/alerts | All Alerting
[**GetBuiltInEventSpecification**](EventSettingsApi.md#GetBuiltInEventSpecification) | **Get** /api/events/settings/event-specifications/built-in/{eventSpecificationId} | Built-in event specifications
[**GetBuiltInEventSpecifications**](EventSettingsApi.md#GetBuiltInEventSpecifications) | **Get** /api/events/settings/event-specifications/built-in | All built-in event specification
[**GetCustomEventSpecification**](EventSettingsApi.md#GetCustomEventSpecification) | **Get** /api/events/settings/event-specifications/custom/{eventSpecificationId} | Custom event specification
[**GetCustomEventSpecifications**](EventSettingsApi.md#GetCustomEventSpecifications) | **Get** /api/events/settings/event-specifications/custom | All custom event specifications
[**GetEventSpecificationInfos**](EventSettingsApi.md#GetEventSpecificationInfos) | **Get** /api/events/settings/event-specifications/infos | Summary of all built-in and custom event specifications
[**GetEventSpecificationInfosByIds**](EventSettingsApi.md#GetEventSpecificationInfosByIds) | **Post** /api/events/settings/event-specifications/infos | All built-in and custom event specifications
[**GetSystemRules**](EventSettingsApi.md#GetSystemRules) | **Get** /api/events/settings/event-specifications/custom/systemRules | All system rules for custom event specifications
[**PutAlert**](EventSettingsApi.md#PutAlert) | **Put** /api/events/settings/alerts/{id} | Update alerting
[**PutAlertingChannel**](EventSettingsApi.md#PutAlertingChannel) | **Put** /api/events/settings/alertingChannels/{id} | Update alerting channel
[**PutCustomEventSpecification**](EventSettingsApi.md#PutCustomEventSpecification) | **Put** /api/events/settings/event-specifications/custom/{eventSpecificationId} | Update custom event specification
[**SendTestAlerting**](EventSettingsApi.md#SendTestAlerting) | **Put** /api/events/settings/alertingChannels/test | Test alerting channel
[**Update**](EventSettingsApi.md#Update) | **Post** /api/events/settings/website-alert-configs/{id} | Update Website Alert Config



## Create

> []WebsiteAlertConfigWithMetadata Create(ctx, websiteAlertConfig)

Create Website Alert Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteAlertConfig** | [**WebsiteAlertConfig**](WebsiteAlertConfig.md)|  | 

### Return type

[**[]WebsiteAlertConfigWithMetadata**](WebsiteAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlert

> DeleteAlert(ctx, id)

Delete alerting

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


## DeleteAlertingChannel

> DeleteAlertingChannel(ctx, id)

Delete alerting channel

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


## DeleteBuiltInEventSpecification

> DeleteBuiltInEventSpecification(ctx, eventSpecificationId)

Delete built-in event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 

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


## DeleteCustomEventSpecification

> DeleteCustomEventSpecification(ctx, eventSpecificationId)

Delete custom event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 

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


## DisableBuiltInEventSpecification

> DisableBuiltInEventSpecification(ctx, eventSpecificationId)

Disable built-in event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 

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


## DisableCustomEventSpecification

> CustomEventSpecificationWithLastUpdated DisableCustomEventSpecification(ctx, eventSpecificationId)

Disable custom event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 

### Return type

[**CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableBuiltInEventSpecification

> EnableBuiltInEventSpecification(ctx, eventSpecificationId)

Enable built-in event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 

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


## EnableCustomEventSpecification

> CustomEventSpecificationWithLastUpdated EnableCustomEventSpecification(ctx, eventSpecificationId)

Enable custom event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 

### Return type

[**CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Find

> []WebsiteAlertConfigWithMetadata Find(ctx, id, optional)

Get Website Alert Config

Find a Website Alert Config by ID. This will deliver deleted configs too.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***FindOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validOn** | **optional.Int64**|  | 

### Return type

[**[]WebsiteAlertConfigWithMetadata**](WebsiteAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllActive

> []WebsiteAlertConfigWithMetadata FindAllActive(ctx, optional)

All Website Alert Configs

Configs are sorted descending by their created date.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindAllActiveOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindAllActiveOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websiteId** | **optional.String**|  | 

### Return type

[**[]WebsiteAlertConfigWithMetadata**](WebsiteAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindVersions

> []ConfigVersion FindVersions(ctx, id)

Get versions of Website Alert Config

Find all versions of a Website Alert Config by ID. This will deliver deleted configs too. Configs are sorted descending by their created date.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**[]ConfigVersion**](ConfigVersion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlert

> AlertingConfigurationWithLastUpdated GetAlert(ctx, id)

Alerting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**AlertingConfigurationWithLastUpdated**](AlertingConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertingChannel

> AbstractIntegration GetAlertingChannel(ctx, id)

Alerting channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**AbstractIntegration**](AbstractIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertingChannels

> []AbstractIntegration GetAlertingChannels(ctx, optional)

All alerting channels

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAlertingChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAlertingChannelsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**[]AbstractIntegration**](AbstractIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertingConfigurationInfos

> []ValidatedAlertingChannelInputInfo GetAlertingConfigurationInfos(ctx, optional)

All alerting configuration info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAlertingConfigurationInfosOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAlertingConfigurationInfosOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationId** | **optional.String**|  | 

### Return type

[**[]ValidatedAlertingChannelInputInfo**](ValidatedAlertingChannelInputInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlerts

> []ValidatedAlertingConfiguration GetAlerts(ctx, )

All Alerting

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ValidatedAlertingConfiguration**](ValidatedAlertingConfiguration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuiltInEventSpecification

> BuiltInEventSpecification GetBuiltInEventSpecification(ctx, eventSpecificationId)

Built-in event specifications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 

### Return type

[**BuiltInEventSpecification**](BuiltInEventSpecification.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuiltInEventSpecifications

> []BuiltInEventSpecificationWithLastUpdated GetBuiltInEventSpecifications(ctx, optional)

All built-in event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBuiltInEventSpecificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBuiltInEventSpecificationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**[]BuiltInEventSpecificationWithLastUpdated**](BuiltInEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomEventSpecification

> CustomEventSpecificationWithLastUpdated GetCustomEventSpecification(ctx, eventSpecificationId)

Custom event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 

### Return type

[**CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomEventSpecifications

> []CustomEventSpecificationWithLastUpdated GetCustomEventSpecifications(ctx, )

All custom event specifications

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventSpecificationInfos

> []EventSpecificationInfo GetEventSpecificationInfos(ctx, )

Summary of all built-in and custom event specifications

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]EventSpecificationInfo**](EventSpecificationInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventSpecificationInfosByIds

> []EventSpecificationInfo GetEventSpecificationInfosByIds(ctx, requestBody)

All built-in and custom event specifications

Summary of all built-in and custom event specifications by IDs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestBody** | [**[]string**](string.md)|  | 

### Return type

[**[]EventSpecificationInfo**](EventSpecificationInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemRules

> []SystemRuleLabel GetSystemRules(ctx, )

All system rules for custom event specifications

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SystemRuleLabel**](SystemRuleLabel.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAlert

> AlertingConfigurationWithLastUpdated PutAlert(ctx, id, alertingConfiguration)

Update alerting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**alertingConfiguration** | [**AlertingConfiguration**](AlertingConfiguration.md)|  | 

### Return type

[**AlertingConfigurationWithLastUpdated**](AlertingConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAlertingChannel

> PutAlertingChannel(ctx, id, abstractIntegration)

Update alerting channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**abstractIntegration** | [**AbstractIntegration**](AbstractIntegration.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCustomEventSpecification

> CustomEventSpecificationWithLastUpdated PutCustomEventSpecification(ctx, eventSpecificationId, customEventSpecification)

Update custom event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 
**customEventSpecification** | [**CustomEventSpecification**](CustomEventSpecification.md)|  | 

### Return type

[**CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTestAlerting

> SendTestAlerting(ctx, abstractIntegration)

Test alerting channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**abstractIntegration** | [**AbstractIntegration**](AbstractIntegration.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> []WebsiteAlertConfigWithMetadata Update(ctx, id, websiteAlertConfig)

Update Website Alert Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**websiteAlertConfig** | [**WebsiteAlertConfig**](WebsiteAlertConfig.md)|  | 

### Return type

[**[]WebsiteAlertConfigWithMetadata**](WebsiteAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

