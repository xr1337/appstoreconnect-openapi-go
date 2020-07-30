# \PreReleaseVersionsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreReleaseVersionsAppGetToOneRelated**](PreReleaseVersionsApi.md#PreReleaseVersionsAppGetToOneRelated) | **Get** /v1/preReleaseVersions/{id}/app | 
[**PreReleaseVersionsBuildsGetToManyRelated**](PreReleaseVersionsApi.md#PreReleaseVersionsBuildsGetToManyRelated) | **Get** /v1/preReleaseVersions/{id}/builds | 
[**PreReleaseVersionsGetCollection**](PreReleaseVersionsApi.md#PreReleaseVersionsGetCollection) | **Get** /v1/preReleaseVersions | 
[**PreReleaseVersionsGetInstance**](PreReleaseVersionsApi.md#PreReleaseVersionsGetInstance) | **Get** /v1/preReleaseVersions/{id} | 



## PreReleaseVersionsAppGetToOneRelated

> AppResponse PreReleaseVersionsAppGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***PreReleaseVersionsAppGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PreReleaseVersionsAppGetToOneRelatedOpts struct


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


## PreReleaseVersionsBuildsGetToManyRelated

> BuildsResponse PreReleaseVersionsBuildsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***PreReleaseVersionsBuildsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PreReleaseVersionsBuildsGetToManyRelatedOpts struct


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


## PreReleaseVersionsGetCollection

> PreReleaseVersionsResponse PreReleaseVersionsGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PreReleaseVersionsGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PreReleaseVersionsGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterBuildsExpired** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;builds.expired&#39; | 
 **filterBuildsProcessingState** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;builds.processingState&#39; | 
 **filterPlatform** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;platform&#39; | 
 **filterVersion** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;version&#39; | 
 **filterApp** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;app&#39; | 
 **filterBuilds** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;builds&#39; | 
 **sort** | [**optional.Interface of []string**](string.md)| comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsPreReleaseVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type preReleaseVersions | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limitBuilds** | **optional.Int32**| maximum number of related builds returned (when they are included) | 

### Return type

[**PreReleaseVersionsResponse**](PreReleaseVersionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreReleaseVersionsGetInstance

> PrereleaseVersionResponse PreReleaseVersionsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***PreReleaseVersionsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PreReleaseVersionsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsPreReleaseVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type preReleaseVersions | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limitBuilds** | **optional.Int32**| maximum number of related builds returned (when they are included) | 

### Return type

[**PrereleaseVersionResponse**](PrereleaseVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

