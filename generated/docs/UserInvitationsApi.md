# \UserInvitationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserInvitationsCreateInstance**](UserInvitationsApi.md#UserInvitationsCreateInstance) | **Post** /v1/userInvitations | 
[**UserInvitationsDeleteInstance**](UserInvitationsApi.md#UserInvitationsDeleteInstance) | **Delete** /v1/userInvitations/{id} | 
[**UserInvitationsGetCollection**](UserInvitationsApi.md#UserInvitationsGetCollection) | **Get** /v1/userInvitations | 
[**UserInvitationsGetInstance**](UserInvitationsApi.md#UserInvitationsGetInstance) | **Get** /v1/userInvitations/{id} | 
[**UserInvitationsVisibleAppsGetToManyRelated**](UserInvitationsApi.md#UserInvitationsVisibleAppsGetToManyRelated) | **Get** /v1/userInvitations/{id}/visibleApps | 



## UserInvitationsCreateInstance

> UserInvitationResponse UserInvitationsCreateInstance(ctx, userInvitationCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userInvitationCreateRequest** | [**UserInvitationCreateRequest**](UserInvitationCreateRequest.md)| UserInvitation representation | 

### Return type

[**UserInvitationResponse**](UserInvitationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInvitationsDeleteInstance

> UserInvitationsDeleteInstance(ctx, id)



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


## UserInvitationsGetCollection

> UserInvitationsResponse UserInvitationsGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserInvitationsGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserInvitationsGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterEmail** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;email&#39; | 
 **filterRoles** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;roles&#39; | 
 **filterVisibleApps** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;visibleApps&#39; | 
 **sort** | [**optional.Interface of []string**](string.md)| comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsUserInvitations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type userInvitations | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limitVisibleApps** | **optional.Int32**| maximum number of related visibleApps returned (when they are included) | 

### Return type

[**UserInvitationsResponse**](UserInvitationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInvitationsGetInstance

> UserInvitationResponse UserInvitationsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***UserInvitationsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserInvitationsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsUserInvitations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type userInvitations | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limitVisibleApps** | **optional.Int32**| maximum number of related visibleApps returned (when they are included) | 

### Return type

[**UserInvitationResponse**](UserInvitationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInvitationsVisibleAppsGetToManyRelated

> AppsResponse UserInvitationsVisibleAppsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***UserInvitationsVisibleAppsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserInvitationsVisibleAppsGetToManyRelatedOpts struct


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

