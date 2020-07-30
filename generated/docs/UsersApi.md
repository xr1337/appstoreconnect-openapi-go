# \UsersApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersDeleteInstance**](UsersApi.md#UsersDeleteInstance) | **Delete** /v1/users/{id} | 
[**UsersGetCollection**](UsersApi.md#UsersGetCollection) | **Get** /v1/users | 
[**UsersGetInstance**](UsersApi.md#UsersGetInstance) | **Get** /v1/users/{id} | 
[**UsersUpdateInstance**](UsersApi.md#UsersUpdateInstance) | **Patch** /v1/users/{id} | 
[**UsersVisibleAppsCreateToManyRelationship**](UsersApi.md#UsersVisibleAppsCreateToManyRelationship) | **Post** /v1/users/{id}/relationships/visibleApps | 
[**UsersVisibleAppsDeleteToManyRelationship**](UsersApi.md#UsersVisibleAppsDeleteToManyRelationship) | **Delete** /v1/users/{id}/relationships/visibleApps | 
[**UsersVisibleAppsGetToManyRelated**](UsersApi.md#UsersVisibleAppsGetToManyRelated) | **Get** /v1/users/{id}/visibleApps | 
[**UsersVisibleAppsGetToManyRelationship**](UsersApi.md#UsersVisibleAppsGetToManyRelationship) | **Get** /v1/users/{id}/relationships/visibleApps | 
[**UsersVisibleAppsReplaceToManyRelationship**](UsersApi.md#UsersVisibleAppsReplaceToManyRelationship) | **Patch** /v1/users/{id}/relationships/visibleApps | 



## UsersDeleteInstance

> UsersDeleteInstance(ctx, id)



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


## UsersGetCollection

> UsersResponse UsersGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterRoles** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;roles&#39; | 
 **filterUsername** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;username&#39; | 
 **filterVisibleApps** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;visibleApps&#39; | 
 **sort** | [**optional.Interface of []string**](string.md)| comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsUsers** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type users | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limitVisibleApps** | **optional.Int32**| maximum number of related visibleApps returned (when they are included) | 

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetInstance

> UserResponse UsersGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***UsersGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsUsers** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type users | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limitVisibleApps** | **optional.Int32**| maximum number of related visibleApps returned (when they are included) | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateInstance

> UserResponse UsersUpdateInstance(ctx, id, userUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**userUpdateRequest** | [**UserUpdateRequest**](UserUpdateRequest.md)| User representation | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersVisibleAppsCreateToManyRelationship

> UsersVisibleAppsCreateToManyRelationship(ctx, id, userVisibleAppsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**userVisibleAppsLinkagesRequest** | [**UserVisibleAppsLinkagesRequest**](UserVisibleAppsLinkagesRequest.md)| List of related linkages | 

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


## UsersVisibleAppsDeleteToManyRelationship

> UsersVisibleAppsDeleteToManyRelationship(ctx, id, userVisibleAppsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**userVisibleAppsLinkagesRequest** | [**UserVisibleAppsLinkagesRequest**](UserVisibleAppsLinkagesRequest.md)| List of related linkages | 

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


## UsersVisibleAppsGetToManyRelated

> AppsResponse UsersVisibleAppsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***UsersVisibleAppsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersVisibleAppsGetToManyRelatedOpts struct


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


## UsersVisibleAppsGetToManyRelationship

> UserVisibleAppsLinkagesResponse UsersVisibleAppsGetToManyRelationship(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***UsersVisibleAppsGetToManyRelationshipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersVisibleAppsGetToManyRelationshipOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**UserVisibleAppsLinkagesResponse**](UserVisibleAppsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersVisibleAppsReplaceToManyRelationship

> UsersVisibleAppsReplaceToManyRelationship(ctx, id, userVisibleAppsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**userVisibleAppsLinkagesRequest** | [**UserVisibleAppsLinkagesRequest**](UserVisibleAppsLinkagesRequest.md)| List of related linkages | 

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

