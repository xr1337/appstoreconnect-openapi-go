# \AppStoreVersionSubmissionsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreVersionSubmissionsCreateInstance**](AppStoreVersionSubmissionsApi.md#AppStoreVersionSubmissionsCreateInstance) | **Post** /v1/appStoreVersionSubmissions | 
[**AppStoreVersionSubmissionsDeleteInstance**](AppStoreVersionSubmissionsApi.md#AppStoreVersionSubmissionsDeleteInstance) | **Delete** /v1/appStoreVersionSubmissions/{id} | 



## AppStoreVersionSubmissionsCreateInstance

> AppStoreVersionSubmissionResponse AppStoreVersionSubmissionsCreateInstance(ctx, appStoreVersionSubmissionCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appStoreVersionSubmissionCreateRequest** | [**AppStoreVersionSubmissionCreateRequest**](AppStoreVersionSubmissionCreateRequest.md)| AppStoreVersionSubmission representation | 

### Return type

[**AppStoreVersionSubmissionResponse**](AppStoreVersionSubmissionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionSubmissionsDeleteInstance

> AppStoreVersionSubmissionsDeleteInstance(ctx, id)



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

