# \GameCenterEnabledVersionsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterEnabledVersionsCompatibleVersionsCreateToManyRelationship**](GameCenterEnabledVersionsApi.md#GameCenterEnabledVersionsCompatibleVersionsCreateToManyRelationship) | **Post** /v1/gameCenterEnabledVersions/{id}/relationships/compatibleVersions | 
[**GameCenterEnabledVersionsCompatibleVersionsDeleteToManyRelationship**](GameCenterEnabledVersionsApi.md#GameCenterEnabledVersionsCompatibleVersionsDeleteToManyRelationship) | **Delete** /v1/gameCenterEnabledVersions/{id}/relationships/compatibleVersions | 
[**GameCenterEnabledVersionsCompatibleVersionsGetToManyRelated**](GameCenterEnabledVersionsApi.md#GameCenterEnabledVersionsCompatibleVersionsGetToManyRelated) | **Get** /v1/gameCenterEnabledVersions/{id}/compatibleVersions | 
[**GameCenterEnabledVersionsCompatibleVersionsGetToManyRelationship**](GameCenterEnabledVersionsApi.md#GameCenterEnabledVersionsCompatibleVersionsGetToManyRelationship) | **Get** /v1/gameCenterEnabledVersions/{id}/relationships/compatibleVersions | 
[**GameCenterEnabledVersionsCompatibleVersionsReplaceToManyRelationship**](GameCenterEnabledVersionsApi.md#GameCenterEnabledVersionsCompatibleVersionsReplaceToManyRelationship) | **Patch** /v1/gameCenterEnabledVersions/{id}/relationships/compatibleVersions | 



## GameCenterEnabledVersionsCompatibleVersionsCreateToManyRelationship

> GameCenterEnabledVersionsCompatibleVersionsCreateToManyRelationship(ctx, id, gameCenterEnabledVersionCompatibleVersionsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**gameCenterEnabledVersionCompatibleVersionsLinkagesRequest** | [**GameCenterEnabledVersionCompatibleVersionsLinkagesRequest**](GameCenterEnabledVersionCompatibleVersionsLinkagesRequest.md)| List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterEnabledVersionsCompatibleVersionsDeleteToManyRelationship

> GameCenterEnabledVersionsCompatibleVersionsDeleteToManyRelationship(ctx, id, gameCenterEnabledVersionCompatibleVersionsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**gameCenterEnabledVersionCompatibleVersionsLinkagesRequest** | [**GameCenterEnabledVersionCompatibleVersionsLinkagesRequest**](GameCenterEnabledVersionCompatibleVersionsLinkagesRequest.md)| List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterEnabledVersionsCompatibleVersionsGetToManyRelated

> GameCenterEnabledVersionsResponse GameCenterEnabledVersionsCompatibleVersionsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***GameCenterEnabledVersionsCompatibleVersionsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GameCenterEnabledVersionsCompatibleVersionsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterPlatform** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;platform&#39; | 
 **filterVersionString** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;versionString&#39; | 
 **filterApp** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;app&#39; | 
 **filterId** | [**optional.Interface of []string**](string.md)| filter by id(s) | 
 **sort** | [**optional.Interface of []string**](string.md)| comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsGameCenterEnabledVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type gameCenterEnabledVersions | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**GameCenterEnabledVersionsResponse**](GameCenterEnabledVersionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterEnabledVersionsCompatibleVersionsGetToManyRelationship

> GameCenterEnabledVersionCompatibleVersionsLinkagesResponse GameCenterEnabledVersionsCompatibleVersionsGetToManyRelationship(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***GameCenterEnabledVersionsCompatibleVersionsGetToManyRelationshipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GameCenterEnabledVersionsCompatibleVersionsGetToManyRelationshipOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**GameCenterEnabledVersionCompatibleVersionsLinkagesResponse**](GameCenterEnabledVersionCompatibleVersionsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterEnabledVersionsCompatibleVersionsReplaceToManyRelationship

> GameCenterEnabledVersionsCompatibleVersionsReplaceToManyRelationship(ctx, id, gameCenterEnabledVersionCompatibleVersionsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**gameCenterEnabledVersionCompatibleVersionsLinkagesRequest** | [**GameCenterEnabledVersionCompatibleVersionsLinkagesRequest**](GameCenterEnabledVersionCompatibleVersionsLinkagesRequest.md)| List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

