# Span

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ParentId** | **string** |  | [optional] 
**CallId** | **string** |  | 
**Name** | **string** |  | 
**Label** | **string** |  | 
**Start** | **int64** |  | [optional] 
**Duration** | **int64** |  | [optional] 
**CalculatedSelfTime** | **int64** |  | [optional] 
**ErrorCount** | **int32** |  | [optional] 
**BatchSize** | **int32** |  | [optional] 
**BatchSelfTime** | **int64** |  | [optional] 
**Kind** | **string** |  | 
**IsSynthetic** | **bool** |  | [optional] 
**Data** | **map[string]map[string]interface{}** |  | 
**Source** | [**SpanRelation**](SpanRelation.md) |  | [optional] 
**Destination** | [**SpanRelation**](SpanRelation.md) |  | [optional] 
**StackTrace** | [**[]StackTraceItem**](StackTraceItem.md) |  | 
**ChildSpans** | [**[]Span**](Span.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


