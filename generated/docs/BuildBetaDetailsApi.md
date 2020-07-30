# \BuildBetaDetailsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildBetaDetailsBuildGetToOneRelated**](BuildBetaDetailsApi.md#BuildBetaDetailsBuildGetToOneRelated) | **Get** /v1/buildBetaDetails/{id}/build | 
[**BuildBetaDetailsGetCollection**](BuildBetaDetailsApi.md#BuildBetaDetailsGetCollection) | **Get** /v1/buildBetaDetails | 
[**BuildBetaDetailsGetInstance**](BuildBetaDetailsApi.md#BuildBetaDetailsGetInstance) | **Get** /v1/buildBetaDetails/{id} | 
[**BuildBetaDetailsUpdateInstance**](BuildBetaDetailsApi.md#BuildBetaDetailsUpdateInstance) | **Patch** /v1/buildBetaDetails/{id} | 



## BuildBetaDetailsBuildGetToOneRelated

> BuildResponse BuildBetaDetailsBuildGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BuildBetaDetailsBuildGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildBetaDetailsBuildGetToOneRelatedOpts struct


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


## BuildBetaDetailsGetCollection

> BuildBetaDetailsResponse BuildBetaDetailsGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BuildBetaDetailsGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildBetaDetailsGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterBuild** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;build&#39; | 
 **filterId** | [**optional.Interface of []string**](string.md)| filter by id(s) | 
 **fieldsBuildBetaDetails** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type buildBetaDetails | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 

### Return type

[**BuildBetaDetailsResponse**](BuildBetaDetailsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildBetaDetailsGetInstance

> BuildBetaDetailResponse BuildBetaDetailsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BuildBetaDetailsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildBetaDetailsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuildBetaDetails** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type buildBetaDetails | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 

### Return type

[**BuildBetaDetailResponse**](BuildBetaDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildBetaDetailsUpdateInstance

> BuildBetaDetailResponse BuildBetaDetailsUpdateInstance(ctx, id, buildBetaDetailUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**buildBetaDetailUpdateRequest** | [**BuildBetaDetailUpdateRequest**](BuildBetaDetailUpdateRequest.md)| BuildBetaDetail representation | 

### Return type

[**BuildBetaDetailResponse**](BuildBetaDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

