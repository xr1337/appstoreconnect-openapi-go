# \BetaBuildLocalizationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaBuildLocalizationsBuildGetToOneRelated**](BetaBuildLocalizationsApi.md#BetaBuildLocalizationsBuildGetToOneRelated) | **Get** /v1/betaBuildLocalizations/{id}/build | 
[**BetaBuildLocalizationsCreateInstance**](BetaBuildLocalizationsApi.md#BetaBuildLocalizationsCreateInstance) | **Post** /v1/betaBuildLocalizations | 
[**BetaBuildLocalizationsDeleteInstance**](BetaBuildLocalizationsApi.md#BetaBuildLocalizationsDeleteInstance) | **Delete** /v1/betaBuildLocalizations/{id} | 
[**BetaBuildLocalizationsGetCollection**](BetaBuildLocalizationsApi.md#BetaBuildLocalizationsGetCollection) | **Get** /v1/betaBuildLocalizations | 
[**BetaBuildLocalizationsGetInstance**](BetaBuildLocalizationsApi.md#BetaBuildLocalizationsGetInstance) | **Get** /v1/betaBuildLocalizations/{id} | 
[**BetaBuildLocalizationsUpdateInstance**](BetaBuildLocalizationsApi.md#BetaBuildLocalizationsUpdateInstance) | **Patch** /v1/betaBuildLocalizations/{id} | 



## BetaBuildLocalizationsBuildGetToOneRelated

> BuildResponse BetaBuildLocalizationsBuildGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaBuildLocalizationsBuildGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaBuildLocalizationsBuildGetToOneRelatedOpts struct


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


## BetaBuildLocalizationsCreateInstance

> BetaBuildLocalizationResponse BetaBuildLocalizationsCreateInstance(ctx, betaBuildLocalizationCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**betaBuildLocalizationCreateRequest** | [**BetaBuildLocalizationCreateRequest**](BetaBuildLocalizationCreateRequest.md)| BetaBuildLocalization representation | 

### Return type

[**BetaBuildLocalizationResponse**](BetaBuildLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaBuildLocalizationsDeleteInstance

> BetaBuildLocalizationsDeleteInstance(ctx, id)



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


## BetaBuildLocalizationsGetCollection

> BetaBuildLocalizationsResponse BetaBuildLocalizationsGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BetaBuildLocalizationsGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaBuildLocalizationsGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterLocale** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;locale&#39; | 
 **filterBuild** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;build&#39; | 
 **fieldsBetaBuildLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaBuildLocalizations | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 

### Return type

[**BetaBuildLocalizationsResponse**](BetaBuildLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaBuildLocalizationsGetInstance

> BetaBuildLocalizationResponse BetaBuildLocalizationsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaBuildLocalizationsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaBuildLocalizationsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaBuildLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaBuildLocalizations | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 

### Return type

[**BetaBuildLocalizationResponse**](BetaBuildLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaBuildLocalizationsUpdateInstance

> BetaBuildLocalizationResponse BetaBuildLocalizationsUpdateInstance(ctx, id, betaBuildLocalizationUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**betaBuildLocalizationUpdateRequest** | [**BetaBuildLocalizationUpdateRequest**](BetaBuildLocalizationUpdateRequest.md)| BetaBuildLocalization representation | 

### Return type

[**BetaBuildLocalizationResponse**](BetaBuildLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

