# \AgeRatingDeclarationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgeRatingDeclarationsUpdateInstance**](AgeRatingDeclarationsApi.md#AgeRatingDeclarationsUpdateInstance) | **Patch** /v1/ageRatingDeclarations/{id} | 



## AgeRatingDeclarationsUpdateInstance

> AgeRatingDeclarationResponse AgeRatingDeclarationsUpdateInstance(ctx, id, ageRatingDeclarationUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**ageRatingDeclarationUpdateRequest** | [**AgeRatingDeclarationUpdateRequest**](AgeRatingDeclarationUpdateRequest.md)| AgeRatingDeclaration representation | 

### Return type

[**AgeRatingDeclarationResponse**](AgeRatingDeclarationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

