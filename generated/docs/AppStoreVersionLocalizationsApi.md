# \AppStoreVersionLocalizationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelated**](AppStoreVersionLocalizationsApi.md#AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelated) | **Get** /v1/appStoreVersionLocalizations/{id}/appPreviewSets | 
[**AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelated**](AppStoreVersionLocalizationsApi.md#AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelated) | **Get** /v1/appStoreVersionLocalizations/{id}/appScreenshotSets | 
[**AppStoreVersionLocalizationsCreateInstance**](AppStoreVersionLocalizationsApi.md#AppStoreVersionLocalizationsCreateInstance) | **Post** /v1/appStoreVersionLocalizations | 
[**AppStoreVersionLocalizationsDeleteInstance**](AppStoreVersionLocalizationsApi.md#AppStoreVersionLocalizationsDeleteInstance) | **Delete** /v1/appStoreVersionLocalizations/{id} | 
[**AppStoreVersionLocalizationsGetInstance**](AppStoreVersionLocalizationsApi.md#AppStoreVersionLocalizationsGetInstance) | **Get** /v1/appStoreVersionLocalizations/{id} | 
[**AppStoreVersionLocalizationsUpdateInstance**](AppStoreVersionLocalizationsApi.md#AppStoreVersionLocalizationsUpdateInstance) | **Patch** /v1/appStoreVersionLocalizations/{id} | 



## AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelated

> AppPreviewSetsResponse AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterPreviewType** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;previewType&#39; | 
 **fieldsAppStoreVersionLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersionLocalizations | 
 **fieldsAppPreviews** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPreviews | 
 **fieldsAppPreviewSets** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPreviewSets | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppPreviewSetsResponse**](AppPreviewSetsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelated

> AppScreenshotSetsResponse AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterScreenshotDisplayType** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;screenshotDisplayType&#39; | 
 **fieldsAppStoreVersionLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersionLocalizations | 
 **fieldsAppScreenshotSets** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appScreenshotSets | 
 **fieldsAppScreenshots** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appScreenshots | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppScreenshotSetsResponse**](AppScreenshotSetsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionLocalizationsCreateInstance

> AppStoreVersionLocalizationResponse AppStoreVersionLocalizationsCreateInstance(ctx, appStoreVersionLocalizationCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appStoreVersionLocalizationCreateRequest** | [**AppStoreVersionLocalizationCreateRequest**](AppStoreVersionLocalizationCreateRequest.md)| AppStoreVersionLocalization representation | 

### Return type

[**AppStoreVersionLocalizationResponse**](AppStoreVersionLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionLocalizationsDeleteInstance

> AppStoreVersionLocalizationsDeleteInstance(ctx, id)



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


## AppStoreVersionLocalizationsGetInstance

> AppStoreVersionLocalizationResponse AppStoreVersionLocalizationsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppStoreVersionLocalizationsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppStoreVersionLocalizationsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersionLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersionLocalizations | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsAppScreenshotSets** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appScreenshotSets | 
 **fieldsAppPreviewSets** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPreviewSets | 
 **limitAppPreviewSets** | **optional.Int32**| maximum number of related appPreviewSets returned (when they are included) | 
 **limitAppScreenshotSets** | **optional.Int32**| maximum number of related appScreenshotSets returned (when they are included) | 

### Return type

[**AppStoreVersionLocalizationResponse**](AppStoreVersionLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionLocalizationsUpdateInstance

> AppStoreVersionLocalizationResponse AppStoreVersionLocalizationsUpdateInstance(ctx, id, appStoreVersionLocalizationUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appStoreVersionLocalizationUpdateRequest** | [**AppStoreVersionLocalizationUpdateRequest**](AppStoreVersionLocalizationUpdateRequest.md)| AppStoreVersionLocalization representation | 

### Return type

[**AppStoreVersionLocalizationResponse**](AppStoreVersionLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

