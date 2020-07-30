# \AppScreenshotsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppScreenshotsCreateInstance**](AppScreenshotsApi.md#AppScreenshotsCreateInstance) | **Post** /v1/appScreenshots | 
[**AppScreenshotsDeleteInstance**](AppScreenshotsApi.md#AppScreenshotsDeleteInstance) | **Delete** /v1/appScreenshots/{id} | 
[**AppScreenshotsGetInstance**](AppScreenshotsApi.md#AppScreenshotsGetInstance) | **Get** /v1/appScreenshots/{id} | 
[**AppScreenshotsUpdateInstance**](AppScreenshotsApi.md#AppScreenshotsUpdateInstance) | **Patch** /v1/appScreenshots/{id} | 



## AppScreenshotsCreateInstance

> AppScreenshotResponse AppScreenshotsCreateInstance(ctx, appScreenshotCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appScreenshotCreateRequest** | [**AppScreenshotCreateRequest**](AppScreenshotCreateRequest.md)| AppScreenshot representation | 

### Return type

[**AppScreenshotResponse**](AppScreenshotResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotsDeleteInstance

> AppScreenshotsDeleteInstance(ctx, id)



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


## AppScreenshotsGetInstance

> AppScreenshotResponse AppScreenshotsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppScreenshotsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppScreenshotsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppScreenshots** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appScreenshots | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppScreenshotResponse**](AppScreenshotResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotsUpdateInstance

> AppScreenshotResponse AppScreenshotsUpdateInstance(ctx, id, appScreenshotUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appScreenshotUpdateRequest** | [**AppScreenshotUpdateRequest**](AppScreenshotUpdateRequest.md)| AppScreenshot representation | 

### Return type

[**AppScreenshotResponse**](AppScreenshotResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

