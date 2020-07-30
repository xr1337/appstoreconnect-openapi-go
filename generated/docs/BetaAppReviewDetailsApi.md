# \BetaAppReviewDetailsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaAppReviewDetailsAppGetToOneRelated**](BetaAppReviewDetailsApi.md#BetaAppReviewDetailsAppGetToOneRelated) | **Get** /v1/betaAppReviewDetails/{id}/app | 
[**BetaAppReviewDetailsGetCollection**](BetaAppReviewDetailsApi.md#BetaAppReviewDetailsGetCollection) | **Get** /v1/betaAppReviewDetails | 
[**BetaAppReviewDetailsGetInstance**](BetaAppReviewDetailsApi.md#BetaAppReviewDetailsGetInstance) | **Get** /v1/betaAppReviewDetails/{id} | 
[**BetaAppReviewDetailsUpdateInstance**](BetaAppReviewDetailsApi.md#BetaAppReviewDetailsUpdateInstance) | **Patch** /v1/betaAppReviewDetails/{id} | 



## BetaAppReviewDetailsAppGetToOneRelated

> AppResponse BetaAppReviewDetailsAppGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaAppReviewDetailsAppGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaAppReviewDetailsAppGetToOneRelatedOpts struct


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


## BetaAppReviewDetailsGetCollection

> BetaAppReviewDetailsResponse BetaAppReviewDetailsGetCollection(ctx, filterApp, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterApp** | [**[]string**](string.md)| filter by id(s) of related &#39;app&#39; | 
 **optional** | ***BetaAppReviewDetailsGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaAppReviewDetailsGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppReviewDetails** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaAppReviewDetails | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 

### Return type

[**BetaAppReviewDetailsResponse**](BetaAppReviewDetailsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppReviewDetailsGetInstance

> BetaAppReviewDetailResponse BetaAppReviewDetailsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaAppReviewDetailsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaAppReviewDetailsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppReviewDetails** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaAppReviewDetails | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 

### Return type

[**BetaAppReviewDetailResponse**](BetaAppReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppReviewDetailsUpdateInstance

> BetaAppReviewDetailResponse BetaAppReviewDetailsUpdateInstance(ctx, id, betaAppReviewDetailUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**betaAppReviewDetailUpdateRequest** | [**BetaAppReviewDetailUpdateRequest**](BetaAppReviewDetailUpdateRequest.md)| BetaAppReviewDetail representation | 

### Return type

[**BetaAppReviewDetailResponse**](BetaAppReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

