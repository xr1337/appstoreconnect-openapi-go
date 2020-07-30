# \InAppPurchasesApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InAppPurchasesGetInstance**](InAppPurchasesApi.md#InAppPurchasesGetInstance) | **Get** /v1/inAppPurchases/{id} | 



## InAppPurchasesGetInstance

> InAppPurchaseResponse InAppPurchasesGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***InAppPurchasesGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InAppPurchasesGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsInAppPurchases** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type inAppPurchases | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **limitApps** | **optional.Int32**| maximum number of related apps returned (when they are included) | 

### Return type

[**InAppPurchaseResponse**](InAppPurchaseResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

