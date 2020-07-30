# \AppPreviewsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppPreviewsCreateInstance**](AppPreviewsApi.md#AppPreviewsCreateInstance) | **Post** /v1/appPreviews | 
[**AppPreviewsDeleteInstance**](AppPreviewsApi.md#AppPreviewsDeleteInstance) | **Delete** /v1/appPreviews/{id} | 
[**AppPreviewsGetInstance**](AppPreviewsApi.md#AppPreviewsGetInstance) | **Get** /v1/appPreviews/{id} | 
[**AppPreviewsUpdateInstance**](AppPreviewsApi.md#AppPreviewsUpdateInstance) | **Patch** /v1/appPreviews/{id} | 



## AppPreviewsCreateInstance

> AppPreviewResponse AppPreviewsCreateInstance(ctx, appPreviewCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appPreviewCreateRequest** | [**AppPreviewCreateRequest**](AppPreviewCreateRequest.md)| AppPreview representation | 

### Return type

[**AppPreviewResponse**](AppPreviewResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPreviewsDeleteInstance

> AppPreviewsDeleteInstance(ctx, id)



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


## AppPreviewsGetInstance

> AppPreviewResponse AppPreviewsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppPreviewsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppPreviewsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPreviews** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPreviews | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppPreviewResponse**](AppPreviewResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPreviewsUpdateInstance

> AppPreviewResponse AppPreviewsUpdateInstance(ctx, id, appPreviewUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appPreviewUpdateRequest** | [**AppPreviewUpdateRequest**](AppPreviewUpdateRequest.md)| AppPreview representation | 

### Return type

[**AppPreviewResponse**](AppPreviewResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

