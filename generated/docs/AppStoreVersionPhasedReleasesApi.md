# \AppStoreVersionPhasedReleasesApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreVersionPhasedReleasesCreateInstance**](AppStoreVersionPhasedReleasesApi.md#AppStoreVersionPhasedReleasesCreateInstance) | **Post** /v1/appStoreVersionPhasedReleases | 
[**AppStoreVersionPhasedReleasesDeleteInstance**](AppStoreVersionPhasedReleasesApi.md#AppStoreVersionPhasedReleasesDeleteInstance) | **Delete** /v1/appStoreVersionPhasedReleases/{id} | 
[**AppStoreVersionPhasedReleasesUpdateInstance**](AppStoreVersionPhasedReleasesApi.md#AppStoreVersionPhasedReleasesUpdateInstance) | **Patch** /v1/appStoreVersionPhasedReleases/{id} | 



## AppStoreVersionPhasedReleasesCreateInstance

> AppStoreVersionPhasedReleaseResponse AppStoreVersionPhasedReleasesCreateInstance(ctx, appStoreVersionPhasedReleaseCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appStoreVersionPhasedReleaseCreateRequest** | [**AppStoreVersionPhasedReleaseCreateRequest**](AppStoreVersionPhasedReleaseCreateRequest.md)| AppStoreVersionPhasedRelease representation | 

### Return type

[**AppStoreVersionPhasedReleaseResponse**](AppStoreVersionPhasedReleaseResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionPhasedReleasesDeleteInstance

> AppStoreVersionPhasedReleasesDeleteInstance(ctx, id)



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


## AppStoreVersionPhasedReleasesUpdateInstance

> AppStoreVersionPhasedReleaseResponse AppStoreVersionPhasedReleasesUpdateInstance(ctx, id, appStoreVersionPhasedReleaseUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appStoreVersionPhasedReleaseUpdateRequest** | [**AppStoreVersionPhasedReleaseUpdateRequest**](AppStoreVersionPhasedReleaseUpdateRequest.md)| AppStoreVersionPhasedRelease representation | 

### Return type

[**AppStoreVersionPhasedReleaseResponse**](AppStoreVersionPhasedReleaseResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

