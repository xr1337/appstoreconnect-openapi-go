# \AppPreviewSetsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppPreviewSetsAppPreviewsGetToManyRelated**](AppPreviewSetsApi.md#AppPreviewSetsAppPreviewsGetToManyRelated) | **Get** /v1/appPreviewSets/{id}/appPreviews | 
[**AppPreviewSetsAppPreviewsGetToManyRelationship**](AppPreviewSetsApi.md#AppPreviewSetsAppPreviewsGetToManyRelationship) | **Get** /v1/appPreviewSets/{id}/relationships/appPreviews | 
[**AppPreviewSetsAppPreviewsReplaceToManyRelationship**](AppPreviewSetsApi.md#AppPreviewSetsAppPreviewsReplaceToManyRelationship) | **Patch** /v1/appPreviewSets/{id}/relationships/appPreviews | 
[**AppPreviewSetsCreateInstance**](AppPreviewSetsApi.md#AppPreviewSetsCreateInstance) | **Post** /v1/appPreviewSets | 
[**AppPreviewSetsDeleteInstance**](AppPreviewSetsApi.md#AppPreviewSetsDeleteInstance) | **Delete** /v1/appPreviewSets/{id} | 
[**AppPreviewSetsGetInstance**](AppPreviewSetsApi.md#AppPreviewSetsGetInstance) | **Get** /v1/appPreviewSets/{id} | 



## AppPreviewSetsAppPreviewsGetToManyRelated

> AppPreviewsResponse AppPreviewSetsAppPreviewsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppPreviewSetsAppPreviewsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppPreviewSetsAppPreviewsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPreviews** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPreviews | 
 **fieldsAppPreviewSets** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPreviewSets | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppPreviewsResponse**](AppPreviewsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPreviewSetsAppPreviewsGetToManyRelationship

> AppPreviewSetAppPreviewsLinkagesResponse AppPreviewSetsAppPreviewsGetToManyRelationship(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppPreviewSetsAppPreviewsGetToManyRelationshipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppPreviewSetsAppPreviewsGetToManyRelationshipOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**AppPreviewSetAppPreviewsLinkagesResponse**](AppPreviewSetAppPreviewsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPreviewSetsAppPreviewsReplaceToManyRelationship

> AppPreviewSetsAppPreviewsReplaceToManyRelationship(ctx, id, appPreviewSetAppPreviewsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appPreviewSetAppPreviewsLinkagesRequest** | [**AppPreviewSetAppPreviewsLinkagesRequest**](AppPreviewSetAppPreviewsLinkagesRequest.md)| List of related linkages | 

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


## AppPreviewSetsCreateInstance

> AppPreviewSetResponse AppPreviewSetsCreateInstance(ctx, appPreviewSetCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appPreviewSetCreateRequest** | [**AppPreviewSetCreateRequest**](AppPreviewSetCreateRequest.md)| AppPreviewSet representation | 

### Return type

[**AppPreviewSetResponse**](AppPreviewSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPreviewSetsDeleteInstance

> AppPreviewSetsDeleteInstance(ctx, id)



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


## AppPreviewSetsGetInstance

> AppPreviewSetResponse AppPreviewSetsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppPreviewSetsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppPreviewSetsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPreviewSets** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPreviewSets | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsAppPreviews** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPreviews | 
 **limitAppPreviews** | **optional.Int32**| maximum number of related appPreviews returned (when they are included) | 

### Return type

[**AppPreviewSetResponse**](AppPreviewSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

