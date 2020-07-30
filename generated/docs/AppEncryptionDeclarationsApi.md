# \AppEncryptionDeclarationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppEncryptionDeclarationsAppGetToOneRelated**](AppEncryptionDeclarationsApi.md#AppEncryptionDeclarationsAppGetToOneRelated) | **Get** /v1/appEncryptionDeclarations/{id}/app | 
[**AppEncryptionDeclarationsBuildsCreateToManyRelationship**](AppEncryptionDeclarationsApi.md#AppEncryptionDeclarationsBuildsCreateToManyRelationship) | **Post** /v1/appEncryptionDeclarations/{id}/relationships/builds | 
[**AppEncryptionDeclarationsGetCollection**](AppEncryptionDeclarationsApi.md#AppEncryptionDeclarationsGetCollection) | **Get** /v1/appEncryptionDeclarations | 
[**AppEncryptionDeclarationsGetInstance**](AppEncryptionDeclarationsApi.md#AppEncryptionDeclarationsGetInstance) | **Get** /v1/appEncryptionDeclarations/{id} | 



## AppEncryptionDeclarationsAppGetToOneRelated

> AppResponse AppEncryptionDeclarationsAppGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppEncryptionDeclarationsAppGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppEncryptionDeclarationsAppGetToOneRelatedOpts struct


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


## AppEncryptionDeclarationsBuildsCreateToManyRelationship

> AppEncryptionDeclarationsBuildsCreateToManyRelationship(ctx, id, appEncryptionDeclarationBuildsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appEncryptionDeclarationBuildsLinkagesRequest** | [**AppEncryptionDeclarationBuildsLinkagesRequest**](AppEncryptionDeclarationBuildsLinkagesRequest.md)| List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEncryptionDeclarationsGetCollection

> AppEncryptionDeclarationsResponse AppEncryptionDeclarationsGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppEncryptionDeclarationsGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppEncryptionDeclarationsGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterPlatform** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;platform&#39; | 
 **filterApp** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;app&#39; | 
 **filterBuilds** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;builds&#39; | 
 **fieldsAppEncryptionDeclarations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appEncryptionDeclarations | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 

### Return type

[**AppEncryptionDeclarationsResponse**](AppEncryptionDeclarationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEncryptionDeclarationsGetInstance

> AppEncryptionDeclarationResponse AppEncryptionDeclarationsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppEncryptionDeclarationsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppEncryptionDeclarationsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppEncryptionDeclarations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appEncryptionDeclarations | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 

### Return type

[**AppEncryptionDeclarationResponse**](AppEncryptionDeclarationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

