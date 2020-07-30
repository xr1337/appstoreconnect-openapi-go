# \AppStoreReviewDetailsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated**](AppStoreReviewDetailsApi.md#AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated) | **Get** /v1/appStoreReviewDetails/{id}/appStoreReviewAttachments | 
[**AppStoreReviewDetailsCreateInstance**](AppStoreReviewDetailsApi.md#AppStoreReviewDetailsCreateInstance) | **Post** /v1/appStoreReviewDetails | 
[**AppStoreReviewDetailsGetInstance**](AppStoreReviewDetailsApi.md#AppStoreReviewDetailsGetInstance) | **Get** /v1/appStoreReviewDetails/{id} | 
[**AppStoreReviewDetailsUpdateInstance**](AppStoreReviewDetailsApi.md#AppStoreReviewDetailsUpdateInstance) | **Patch** /v1/appStoreReviewDetails/{id} | 



## AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated

> AppStoreReviewAttachmentsResponse AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreReviewDetails** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreReviewDetails | 
 **fieldsAppStoreReviewAttachments** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreReviewAttachments | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppStoreReviewAttachmentsResponse**](AppStoreReviewAttachmentsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreReviewDetailsCreateInstance

> AppStoreReviewDetailResponse AppStoreReviewDetailsCreateInstance(ctx, appStoreReviewDetailCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appStoreReviewDetailCreateRequest** | [**AppStoreReviewDetailCreateRequest**](AppStoreReviewDetailCreateRequest.md)| AppStoreReviewDetail representation | 

### Return type

[**AppStoreReviewDetailResponse**](AppStoreReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreReviewDetailsGetInstance

> AppStoreReviewDetailResponse AppStoreReviewDetailsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppStoreReviewDetailsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppStoreReviewDetailsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreReviewDetails** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreReviewDetails | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsAppStoreReviewAttachments** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreReviewAttachments | 
 **limitAppStoreReviewAttachments** | **optional.Int32**| maximum number of related appStoreReviewAttachments returned (when they are included) | 

### Return type

[**AppStoreReviewDetailResponse**](AppStoreReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreReviewDetailsUpdateInstance

> AppStoreReviewDetailResponse AppStoreReviewDetailsUpdateInstance(ctx, id, appStoreReviewDetailUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appStoreReviewDetailUpdateRequest** | [**AppStoreReviewDetailUpdateRequest**](AppStoreReviewDetailUpdateRequest.md)| AppStoreReviewDetail representation | 

### Return type

[**AppStoreReviewDetailResponse**](AppStoreReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

