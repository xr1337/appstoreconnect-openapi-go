# \AppStoreReviewAttachmentsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreReviewAttachmentsCreateInstance**](AppStoreReviewAttachmentsApi.md#AppStoreReviewAttachmentsCreateInstance) | **Post** /v1/appStoreReviewAttachments | 
[**AppStoreReviewAttachmentsDeleteInstance**](AppStoreReviewAttachmentsApi.md#AppStoreReviewAttachmentsDeleteInstance) | **Delete** /v1/appStoreReviewAttachments/{id} | 
[**AppStoreReviewAttachmentsGetInstance**](AppStoreReviewAttachmentsApi.md#AppStoreReviewAttachmentsGetInstance) | **Get** /v1/appStoreReviewAttachments/{id} | 
[**AppStoreReviewAttachmentsUpdateInstance**](AppStoreReviewAttachmentsApi.md#AppStoreReviewAttachmentsUpdateInstance) | **Patch** /v1/appStoreReviewAttachments/{id} | 



## AppStoreReviewAttachmentsCreateInstance

> AppStoreReviewAttachmentResponse AppStoreReviewAttachmentsCreateInstance(ctx, appStoreReviewAttachmentCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appStoreReviewAttachmentCreateRequest** | [**AppStoreReviewAttachmentCreateRequest**](AppStoreReviewAttachmentCreateRequest.md)| AppStoreReviewAttachment representation | 

### Return type

[**AppStoreReviewAttachmentResponse**](AppStoreReviewAttachmentResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreReviewAttachmentsDeleteInstance

> AppStoreReviewAttachmentsDeleteInstance(ctx, id)



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


## AppStoreReviewAttachmentsGetInstance

> AppStoreReviewAttachmentResponse AppStoreReviewAttachmentsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppStoreReviewAttachmentsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppStoreReviewAttachmentsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreReviewAttachments** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreReviewAttachments | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppStoreReviewAttachmentResponse**](AppStoreReviewAttachmentResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreReviewAttachmentsUpdateInstance

> AppStoreReviewAttachmentResponse AppStoreReviewAttachmentsUpdateInstance(ctx, id, appStoreReviewAttachmentUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appStoreReviewAttachmentUpdateRequest** | [**AppStoreReviewAttachmentUpdateRequest**](AppStoreReviewAttachmentUpdateRequest.md)| AppStoreReviewAttachment representation | 

### Return type

[**AppStoreReviewAttachmentResponse**](AppStoreReviewAttachmentResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

