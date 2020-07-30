# \BetaAppLocalizationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaAppLocalizationsAppGetToOneRelated**](BetaAppLocalizationsApi.md#BetaAppLocalizationsAppGetToOneRelated) | **Get** /v1/betaAppLocalizations/{id}/app | 
[**BetaAppLocalizationsCreateInstance**](BetaAppLocalizationsApi.md#BetaAppLocalizationsCreateInstance) | **Post** /v1/betaAppLocalizations | 
[**BetaAppLocalizationsDeleteInstance**](BetaAppLocalizationsApi.md#BetaAppLocalizationsDeleteInstance) | **Delete** /v1/betaAppLocalizations/{id} | 
[**BetaAppLocalizationsGetCollection**](BetaAppLocalizationsApi.md#BetaAppLocalizationsGetCollection) | **Get** /v1/betaAppLocalizations | 
[**BetaAppLocalizationsGetInstance**](BetaAppLocalizationsApi.md#BetaAppLocalizationsGetInstance) | **Get** /v1/betaAppLocalizations/{id} | 
[**BetaAppLocalizationsUpdateInstance**](BetaAppLocalizationsApi.md#BetaAppLocalizationsUpdateInstance) | **Patch** /v1/betaAppLocalizations/{id} | 



## BetaAppLocalizationsAppGetToOneRelated

> AppResponse BetaAppLocalizationsAppGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaAppLocalizationsAppGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaAppLocalizationsAppGetToOneRelatedOpts struct


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


## BetaAppLocalizationsCreateInstance

> BetaAppLocalizationResponse BetaAppLocalizationsCreateInstance(ctx, betaAppLocalizationCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**betaAppLocalizationCreateRequest** | [**BetaAppLocalizationCreateRequest**](BetaAppLocalizationCreateRequest.md)| BetaAppLocalization representation | 

### Return type

[**BetaAppLocalizationResponse**](BetaAppLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppLocalizationsDeleteInstance

> BetaAppLocalizationsDeleteInstance(ctx, id)



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


## BetaAppLocalizationsGetCollection

> BetaAppLocalizationsResponse BetaAppLocalizationsGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BetaAppLocalizationsGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaAppLocalizationsGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterLocale** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;locale&#39; | 
 **filterApp** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;app&#39; | 
 **fieldsBetaAppLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaAppLocalizations | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 

### Return type

[**BetaAppLocalizationsResponse**](BetaAppLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppLocalizationsGetInstance

> BetaAppLocalizationResponse BetaAppLocalizationsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaAppLocalizationsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaAppLocalizationsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaAppLocalizations | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 

### Return type

[**BetaAppLocalizationResponse**](BetaAppLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppLocalizationsUpdateInstance

> BetaAppLocalizationResponse BetaAppLocalizationsUpdateInstance(ctx, id, betaAppLocalizationUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**betaAppLocalizationUpdateRequest** | [**BetaAppLocalizationUpdateRequest**](BetaAppLocalizationUpdateRequest.md)| BetaAppLocalization representation | 

### Return type

[**BetaAppLocalizationResponse**](BetaAppLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

