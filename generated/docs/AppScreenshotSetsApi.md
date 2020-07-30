# \AppScreenshotSetsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppScreenshotSetsAppScreenshotsGetToManyRelated**](AppScreenshotSetsApi.md#AppScreenshotSetsAppScreenshotsGetToManyRelated) | **Get** /v1/appScreenshotSets/{id}/appScreenshots | 
[**AppScreenshotSetsAppScreenshotsGetToManyRelationship**](AppScreenshotSetsApi.md#AppScreenshotSetsAppScreenshotsGetToManyRelationship) | **Get** /v1/appScreenshotSets/{id}/relationships/appScreenshots | 
[**AppScreenshotSetsAppScreenshotsReplaceToManyRelationship**](AppScreenshotSetsApi.md#AppScreenshotSetsAppScreenshotsReplaceToManyRelationship) | **Patch** /v1/appScreenshotSets/{id}/relationships/appScreenshots | 
[**AppScreenshotSetsCreateInstance**](AppScreenshotSetsApi.md#AppScreenshotSetsCreateInstance) | **Post** /v1/appScreenshotSets | 
[**AppScreenshotSetsDeleteInstance**](AppScreenshotSetsApi.md#AppScreenshotSetsDeleteInstance) | **Delete** /v1/appScreenshotSets/{id} | 
[**AppScreenshotSetsGetInstance**](AppScreenshotSetsApi.md#AppScreenshotSetsGetInstance) | **Get** /v1/appScreenshotSets/{id} | 



## AppScreenshotSetsAppScreenshotsGetToManyRelated

> AppScreenshotsResponse AppScreenshotSetsAppScreenshotsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppScreenshotSetsAppScreenshotsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppScreenshotSetsAppScreenshotsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppScreenshotSets** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appScreenshotSets | 
 **fieldsAppScreenshots** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appScreenshots | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppScreenshotsResponse**](AppScreenshotsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotSetsAppScreenshotsGetToManyRelationship

> AppScreenshotSetAppScreenshotsLinkagesResponse AppScreenshotSetsAppScreenshotsGetToManyRelationship(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppScreenshotSetsAppScreenshotsGetToManyRelationshipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppScreenshotSetsAppScreenshotsGetToManyRelationshipOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**AppScreenshotSetAppScreenshotsLinkagesResponse**](AppScreenshotSetAppScreenshotsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotSetsAppScreenshotsReplaceToManyRelationship

> AppScreenshotSetsAppScreenshotsReplaceToManyRelationship(ctx, id, appScreenshotSetAppScreenshotsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appScreenshotSetAppScreenshotsLinkagesRequest** | [**AppScreenshotSetAppScreenshotsLinkagesRequest**](AppScreenshotSetAppScreenshotsLinkagesRequest.md)| List of related linkages | 

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


## AppScreenshotSetsCreateInstance

> AppScreenshotSetResponse AppScreenshotSetsCreateInstance(ctx, appScreenshotSetCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appScreenshotSetCreateRequest** | [**AppScreenshotSetCreateRequest**](AppScreenshotSetCreateRequest.md)| AppScreenshotSet representation | 

### Return type

[**AppScreenshotSetResponse**](AppScreenshotSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotSetsDeleteInstance

> AppScreenshotSetsDeleteInstance(ctx, id)



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


## AppScreenshotSetsGetInstance

> AppScreenshotSetResponse AppScreenshotSetsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppScreenshotSetsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppScreenshotSetsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppScreenshotSets** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appScreenshotSets | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsAppScreenshots** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appScreenshots | 
 **limitAppScreenshots** | **optional.Int32**| maximum number of related appScreenshots returned (when they are included) | 

### Return type

[**AppScreenshotSetResponse**](AppScreenshotSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

