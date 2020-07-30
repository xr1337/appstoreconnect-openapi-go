# \BetaGroupsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaGroupsAppGetToOneRelated**](BetaGroupsApi.md#BetaGroupsAppGetToOneRelated) | **Get** /v1/betaGroups/{id}/app | 
[**BetaGroupsBetaTestersCreateToManyRelationship**](BetaGroupsApi.md#BetaGroupsBetaTestersCreateToManyRelationship) | **Post** /v1/betaGroups/{id}/relationships/betaTesters | 
[**BetaGroupsBetaTestersDeleteToManyRelationship**](BetaGroupsApi.md#BetaGroupsBetaTestersDeleteToManyRelationship) | **Delete** /v1/betaGroups/{id}/relationships/betaTesters | 
[**BetaGroupsBetaTestersGetToManyRelated**](BetaGroupsApi.md#BetaGroupsBetaTestersGetToManyRelated) | **Get** /v1/betaGroups/{id}/betaTesters | 
[**BetaGroupsBetaTestersGetToManyRelationship**](BetaGroupsApi.md#BetaGroupsBetaTestersGetToManyRelationship) | **Get** /v1/betaGroups/{id}/relationships/betaTesters | 
[**BetaGroupsBuildsCreateToManyRelationship**](BetaGroupsApi.md#BetaGroupsBuildsCreateToManyRelationship) | **Post** /v1/betaGroups/{id}/relationships/builds | 
[**BetaGroupsBuildsDeleteToManyRelationship**](BetaGroupsApi.md#BetaGroupsBuildsDeleteToManyRelationship) | **Delete** /v1/betaGroups/{id}/relationships/builds | 
[**BetaGroupsBuildsGetToManyRelated**](BetaGroupsApi.md#BetaGroupsBuildsGetToManyRelated) | **Get** /v1/betaGroups/{id}/builds | 
[**BetaGroupsBuildsGetToManyRelationship**](BetaGroupsApi.md#BetaGroupsBuildsGetToManyRelationship) | **Get** /v1/betaGroups/{id}/relationships/builds | 
[**BetaGroupsCreateInstance**](BetaGroupsApi.md#BetaGroupsCreateInstance) | **Post** /v1/betaGroups | 
[**BetaGroupsDeleteInstance**](BetaGroupsApi.md#BetaGroupsDeleteInstance) | **Delete** /v1/betaGroups/{id} | 
[**BetaGroupsGetCollection**](BetaGroupsApi.md#BetaGroupsGetCollection) | **Get** /v1/betaGroups | 
[**BetaGroupsGetInstance**](BetaGroupsApi.md#BetaGroupsGetInstance) | **Get** /v1/betaGroups/{id} | 
[**BetaGroupsUpdateInstance**](BetaGroupsApi.md#BetaGroupsUpdateInstance) | **Patch** /v1/betaGroups/{id} | 



## BetaGroupsAppGetToOneRelated

> AppResponse BetaGroupsAppGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaGroupsAppGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaGroupsAppGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 

### Return type

[**AppResponse**](AppResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsBetaTestersCreateToManyRelationship

> BetaGroupsBetaTestersCreateToManyRelationship(ctx, id, betaGroupBetaTestersLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**betaGroupBetaTestersLinkagesRequest** | [**BetaGroupBetaTestersLinkagesRequest**](BetaGroupBetaTestersLinkagesRequest.md)| List of related linkages | 

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


## BetaGroupsBetaTestersDeleteToManyRelationship

> BetaGroupsBetaTestersDeleteToManyRelationship(ctx, id, betaGroupBetaTestersLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**betaGroupBetaTestersLinkagesRequest** | [**BetaGroupBetaTestersLinkagesRequest**](BetaGroupBetaTestersLinkagesRequest.md)| List of related linkages | 

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


## BetaGroupsBetaTestersGetToManyRelated

> BetaTestersResponse BetaGroupsBetaTestersGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaGroupsBetaTestersGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaGroupsBetaTestersGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaTesters** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaTesters | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**BetaTestersResponse**](BetaTestersResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsBetaTestersGetToManyRelationship

> BetaGroupBetaTestersLinkagesResponse BetaGroupsBetaTestersGetToManyRelationship(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaGroupsBetaTestersGetToManyRelationshipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaGroupsBetaTestersGetToManyRelationshipOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**BetaGroupBetaTestersLinkagesResponse**](BetaGroupBetaTestersLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsBuildsCreateToManyRelationship

> BetaGroupsBuildsCreateToManyRelationship(ctx, id, betaGroupBuildsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**betaGroupBuildsLinkagesRequest** | [**BetaGroupBuildsLinkagesRequest**](BetaGroupBuildsLinkagesRequest.md)| List of related linkages | 

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


## BetaGroupsBuildsDeleteToManyRelationship

> BetaGroupsBuildsDeleteToManyRelationship(ctx, id, betaGroupBuildsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**betaGroupBuildsLinkagesRequest** | [**BetaGroupBuildsLinkagesRequest**](BetaGroupBuildsLinkagesRequest.md)| List of related linkages | 

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


## BetaGroupsBuildsGetToManyRelated

> BuildsResponse BetaGroupsBuildsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaGroupsBuildsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaGroupsBuildsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**BuildsResponse**](BuildsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsBuildsGetToManyRelationship

> BetaGroupBuildsLinkagesResponse BetaGroupsBuildsGetToManyRelationship(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaGroupsBuildsGetToManyRelationshipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaGroupsBuildsGetToManyRelationshipOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**BetaGroupBuildsLinkagesResponse**](BetaGroupBuildsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsCreateInstance

> BetaGroupResponse BetaGroupsCreateInstance(ctx, betaGroupCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**betaGroupCreateRequest** | [**BetaGroupCreateRequest**](BetaGroupCreateRequest.md)| BetaGroup representation | 

### Return type

[**BetaGroupResponse**](BetaGroupResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsDeleteInstance

> BetaGroupsDeleteInstance(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsGetCollection

> BetaGroupsResponse BetaGroupsGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BetaGroupsGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaGroupsGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterIsInternalGroup** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;isInternalGroup&#39; | 
 **filterName** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;name&#39; | 
 **filterPublicLink** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;publicLink&#39; | 
 **filterPublicLinkEnabled** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;publicLinkEnabled&#39; | 
 **filterPublicLinkLimitEnabled** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;publicLinkLimitEnabled&#39; | 
 **filterApp** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;app&#39; | 
 **filterBuilds** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;builds&#39; | 
 **filterId** | [**optional.Interface of []string**](string.md)| filter by id(s) | 
 **sort** | [**optional.Interface of []string**](string.md)| comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsBetaGroups** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaGroups | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 
 **fieldsBetaTesters** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaTesters | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limitBetaTesters** | **optional.Int32**| maximum number of related betaTesters returned (when they are included) | 
 **limitBuilds** | **optional.Int32**| maximum number of related builds returned (when they are included) | 

### Return type

[**BetaGroupsResponse**](BetaGroupsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsGetInstance

> BetaGroupResponse BetaGroupsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaGroupsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaGroupsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaGroups** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaGroups | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 
 **fieldsBetaTesters** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaTesters | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limitBetaTesters** | **optional.Int32**| maximum number of related betaTesters returned (when they are included) | 
 **limitBuilds** | **optional.Int32**| maximum number of related builds returned (when they are included) | 

### Return type

[**BetaGroupResponse**](BetaGroupResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsUpdateInstance

> BetaGroupResponse BetaGroupsUpdateInstance(ctx, id, betaGroupUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**betaGroupUpdateRequest** | [**BetaGroupUpdateRequest**](BetaGroupUpdateRequest.md)| BetaGroup representation | 

### Return type

[**BetaGroupResponse**](BetaGroupResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

