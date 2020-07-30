# \BetaTestersApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaTestersAppsDeleteToManyRelationship**](BetaTestersApi.md#BetaTestersAppsDeleteToManyRelationship) | **Delete** /v1/betaTesters/{id}/relationships/apps | 
[**BetaTestersAppsGetToManyRelated**](BetaTestersApi.md#BetaTestersAppsGetToManyRelated) | **Get** /v1/betaTesters/{id}/apps | 
[**BetaTestersAppsGetToManyRelationship**](BetaTestersApi.md#BetaTestersAppsGetToManyRelationship) | **Get** /v1/betaTesters/{id}/relationships/apps | 
[**BetaTestersBetaGroupsCreateToManyRelationship**](BetaTestersApi.md#BetaTestersBetaGroupsCreateToManyRelationship) | **Post** /v1/betaTesters/{id}/relationships/betaGroups | 
[**BetaTestersBetaGroupsDeleteToManyRelationship**](BetaTestersApi.md#BetaTestersBetaGroupsDeleteToManyRelationship) | **Delete** /v1/betaTesters/{id}/relationships/betaGroups | 
[**BetaTestersBetaGroupsGetToManyRelated**](BetaTestersApi.md#BetaTestersBetaGroupsGetToManyRelated) | **Get** /v1/betaTesters/{id}/betaGroups | 
[**BetaTestersBetaGroupsGetToManyRelationship**](BetaTestersApi.md#BetaTestersBetaGroupsGetToManyRelationship) | **Get** /v1/betaTesters/{id}/relationships/betaGroups | 
[**BetaTestersBuildsCreateToManyRelationship**](BetaTestersApi.md#BetaTestersBuildsCreateToManyRelationship) | **Post** /v1/betaTesters/{id}/relationships/builds | 
[**BetaTestersBuildsDeleteToManyRelationship**](BetaTestersApi.md#BetaTestersBuildsDeleteToManyRelationship) | **Delete** /v1/betaTesters/{id}/relationships/builds | 
[**BetaTestersBuildsGetToManyRelated**](BetaTestersApi.md#BetaTestersBuildsGetToManyRelated) | **Get** /v1/betaTesters/{id}/builds | 
[**BetaTestersBuildsGetToManyRelationship**](BetaTestersApi.md#BetaTestersBuildsGetToManyRelationship) | **Get** /v1/betaTesters/{id}/relationships/builds | 
[**BetaTestersCreateInstance**](BetaTestersApi.md#BetaTestersCreateInstance) | **Post** /v1/betaTesters | 
[**BetaTestersDeleteInstance**](BetaTestersApi.md#BetaTestersDeleteInstance) | **Delete** /v1/betaTesters/{id} | 
[**BetaTestersGetCollection**](BetaTestersApi.md#BetaTestersGetCollection) | **Get** /v1/betaTesters | 
[**BetaTestersGetInstance**](BetaTestersApi.md#BetaTestersGetInstance) | **Get** /v1/betaTesters/{id} | 



## BetaTestersAppsDeleteToManyRelationship

> BetaTestersAppsDeleteToManyRelationship(ctx, id, betaTesterAppsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**betaTesterAppsLinkagesRequest** | [**BetaTesterAppsLinkagesRequest**](BetaTesterAppsLinkagesRequest.md)| List of related linkages | 

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


## BetaTestersAppsGetToManyRelated

> AppsResponse BetaTestersAppsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaTestersAppsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaTestersAppsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**AppsResponse**](AppsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersAppsGetToManyRelationship

> BetaTesterAppsLinkagesResponse BetaTestersAppsGetToManyRelationship(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaTestersAppsGetToManyRelationshipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaTestersAppsGetToManyRelationshipOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**BetaTesterAppsLinkagesResponse**](BetaTesterAppsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersBetaGroupsCreateToManyRelationship

> BetaTestersBetaGroupsCreateToManyRelationship(ctx, id, betaTesterBetaGroupsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**betaTesterBetaGroupsLinkagesRequest** | [**BetaTesterBetaGroupsLinkagesRequest**](BetaTesterBetaGroupsLinkagesRequest.md)| List of related linkages | 

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


## BetaTestersBetaGroupsDeleteToManyRelationship

> BetaTestersBetaGroupsDeleteToManyRelationship(ctx, id, betaTesterBetaGroupsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**betaTesterBetaGroupsLinkagesRequest** | [**BetaTesterBetaGroupsLinkagesRequest**](BetaTesterBetaGroupsLinkagesRequest.md)| List of related linkages | 

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


## BetaTestersBetaGroupsGetToManyRelated

> BetaGroupsResponse BetaTestersBetaGroupsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaTestersBetaGroupsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaTestersBetaGroupsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaGroups** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaGroups | 
 **limit** | **optional.Int32**| maximum resources per page | 

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


## BetaTestersBetaGroupsGetToManyRelationship

> BetaTesterBetaGroupsLinkagesResponse BetaTestersBetaGroupsGetToManyRelationship(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaTestersBetaGroupsGetToManyRelationshipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaTestersBetaGroupsGetToManyRelationshipOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**BetaTesterBetaGroupsLinkagesResponse**](BetaTesterBetaGroupsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersBuildsCreateToManyRelationship

> BetaTestersBuildsCreateToManyRelationship(ctx, id, betaTesterBuildsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**betaTesterBuildsLinkagesRequest** | [**BetaTesterBuildsLinkagesRequest**](BetaTesterBuildsLinkagesRequest.md)| List of related linkages | 

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


## BetaTestersBuildsDeleteToManyRelationship

> BetaTestersBuildsDeleteToManyRelationship(ctx, id, betaTesterBuildsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**betaTesterBuildsLinkagesRequest** | [**BetaTesterBuildsLinkagesRequest**](BetaTesterBuildsLinkagesRequest.md)| List of related linkages | 

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


## BetaTestersBuildsGetToManyRelated

> BuildsResponse BetaTestersBuildsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaTestersBuildsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaTestersBuildsGetToManyRelatedOpts struct


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


## BetaTestersBuildsGetToManyRelationship

> BetaTesterBuildsLinkagesResponse BetaTestersBuildsGetToManyRelationship(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaTestersBuildsGetToManyRelationshipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaTestersBuildsGetToManyRelationshipOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**BetaTesterBuildsLinkagesResponse**](BetaTesterBuildsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersCreateInstance

> BetaTesterResponse BetaTestersCreateInstance(ctx, betaTesterCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**betaTesterCreateRequest** | [**BetaTesterCreateRequest**](BetaTesterCreateRequest.md)| BetaTester representation | 

### Return type

[**BetaTesterResponse**](BetaTesterResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersDeleteInstance

> BetaTestersDeleteInstance(ctx, id)



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


## BetaTestersGetCollection

> BetaTestersResponse BetaTestersGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BetaTestersGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaTestersGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterEmail** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;email&#39; | 
 **filterFirstName** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;firstName&#39; | 
 **filterInviteType** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;inviteType&#39; | 
 **filterLastName** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;lastName&#39; | 
 **filterApps** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;apps&#39; | 
 **filterBetaGroups** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;betaGroups&#39; | 
 **filterBuilds** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;builds&#39; | 
 **sort** | [**optional.Interface of []string**](string.md)| comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsBetaTesters** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaTesters | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBetaGroups** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaGroups | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limitApps** | **optional.Int32**| maximum number of related apps returned (when they are included) | 
 **limitBetaGroups** | **optional.Int32**| maximum number of related betaGroups returned (when they are included) | 
 **limitBuilds** | **optional.Int32**| maximum number of related builds returned (when they are included) | 

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


## BetaTestersGetInstance

> BetaTesterResponse BetaTestersGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaTestersGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaTestersGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaTesters** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaTesters | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBetaGroups** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaGroups | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limitApps** | **optional.Int32**| maximum number of related apps returned (when they are included) | 
 **limitBetaGroups** | **optional.Int32**| maximum number of related betaGroups returned (when they are included) | 
 **limitBuilds** | **optional.Int32**| maximum number of related builds returned (when they are included) | 

### Return type

[**BetaTesterResponse**](BetaTesterResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

