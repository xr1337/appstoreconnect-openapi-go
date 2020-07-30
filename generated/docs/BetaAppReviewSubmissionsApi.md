# \BetaAppReviewSubmissionsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaAppReviewSubmissionsBuildGetToOneRelated**](BetaAppReviewSubmissionsApi.md#BetaAppReviewSubmissionsBuildGetToOneRelated) | **Get** /v1/betaAppReviewSubmissions/{id}/build | 
[**BetaAppReviewSubmissionsCreateInstance**](BetaAppReviewSubmissionsApi.md#BetaAppReviewSubmissionsCreateInstance) | **Post** /v1/betaAppReviewSubmissions | 
[**BetaAppReviewSubmissionsGetCollection**](BetaAppReviewSubmissionsApi.md#BetaAppReviewSubmissionsGetCollection) | **Get** /v1/betaAppReviewSubmissions | 
[**BetaAppReviewSubmissionsGetInstance**](BetaAppReviewSubmissionsApi.md#BetaAppReviewSubmissionsGetInstance) | **Get** /v1/betaAppReviewSubmissions/{id} | 



## BetaAppReviewSubmissionsBuildGetToOneRelated

> BuildResponse BetaAppReviewSubmissionsBuildGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaAppReviewSubmissionsBuildGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaAppReviewSubmissionsBuildGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 

### Return type

[**BuildResponse**](BuildResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppReviewSubmissionsCreateInstance

> BetaAppReviewSubmissionResponse BetaAppReviewSubmissionsCreateInstance(ctx, betaAppReviewSubmissionCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**betaAppReviewSubmissionCreateRequest** | [**BetaAppReviewSubmissionCreateRequest**](BetaAppReviewSubmissionCreateRequest.md)| BetaAppReviewSubmission representation | 

### Return type

[**BetaAppReviewSubmissionResponse**](BetaAppReviewSubmissionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppReviewSubmissionsGetCollection

> BetaAppReviewSubmissionsResponse BetaAppReviewSubmissionsGetCollection(ctx, filterBuild, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterBuild** | [**[]string**](string.md)| filter by id(s) of related &#39;build&#39; | 
 **optional** | ***BetaAppReviewSubmissionsGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaAppReviewSubmissionsGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterBetaReviewState** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;betaReviewState&#39; | 
 **fieldsBetaAppReviewSubmissions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaAppReviewSubmissions | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 

### Return type

[**BetaAppReviewSubmissionsResponse**](BetaAppReviewSubmissionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppReviewSubmissionsGetInstance

> BetaAppReviewSubmissionResponse BetaAppReviewSubmissionsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaAppReviewSubmissionsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaAppReviewSubmissionsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppReviewSubmissions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaAppReviewSubmissions | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 

### Return type

[**BetaAppReviewSubmissionResponse**](BetaAppReviewSubmissionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

