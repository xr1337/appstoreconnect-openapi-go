# \IdfaDeclarationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdfaDeclarationsCreateInstance**](IdfaDeclarationsApi.md#IdfaDeclarationsCreateInstance) | **Post** /v1/idfaDeclarations | 
[**IdfaDeclarationsDeleteInstance**](IdfaDeclarationsApi.md#IdfaDeclarationsDeleteInstance) | **Delete** /v1/idfaDeclarations/{id} | 
[**IdfaDeclarationsUpdateInstance**](IdfaDeclarationsApi.md#IdfaDeclarationsUpdateInstance) | **Patch** /v1/idfaDeclarations/{id} | 



## IdfaDeclarationsCreateInstance

> IdfaDeclarationResponse IdfaDeclarationsCreateInstance(ctx, idfaDeclarationCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idfaDeclarationCreateRequest** | [**IdfaDeclarationCreateRequest**](IdfaDeclarationCreateRequest.md)| IdfaDeclaration representation | 

### Return type

[**IdfaDeclarationResponse**](IdfaDeclarationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdfaDeclarationsDeleteInstance

> IdfaDeclarationsDeleteInstance(ctx, id)



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


## IdfaDeclarationsUpdateInstance

> IdfaDeclarationResponse IdfaDeclarationsUpdateInstance(ctx, id, idfaDeclarationUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**idfaDeclarationUpdateRequest** | [**IdfaDeclarationUpdateRequest**](IdfaDeclarationUpdateRequest.md)| IdfaDeclaration representation | 

### Return type

[**IdfaDeclarationResponse**](IdfaDeclarationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

