# \AppPreOrdersApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppPreOrdersCreateInstance**](AppPreOrdersApi.md#AppPreOrdersCreateInstance) | **Post** /v1/appPreOrders | 
[**AppPreOrdersDeleteInstance**](AppPreOrdersApi.md#AppPreOrdersDeleteInstance) | **Delete** /v1/appPreOrders/{id} | 
[**AppPreOrdersGetInstance**](AppPreOrdersApi.md#AppPreOrdersGetInstance) | **Get** /v1/appPreOrders/{id} | 
[**AppPreOrdersUpdateInstance**](AppPreOrdersApi.md#AppPreOrdersUpdateInstance) | **Patch** /v1/appPreOrders/{id} | 



## AppPreOrdersCreateInstance

> AppPreOrderResponse AppPreOrdersCreateInstance(ctx, appPreOrderCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appPreOrderCreateRequest** | [**AppPreOrderCreateRequest**](AppPreOrderCreateRequest.md)| AppPreOrder representation | 

### Return type

[**AppPreOrderResponse**](AppPreOrderResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPreOrdersDeleteInstance

> AppPreOrdersDeleteInstance(ctx, id)



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


## AppPreOrdersGetInstance

> AppPreOrderResponse AppPreOrdersGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppPreOrdersGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppPreOrdersGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPreOrders** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPreOrders | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppPreOrderResponse**](AppPreOrderResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPreOrdersUpdateInstance

> AppPreOrderResponse AppPreOrdersUpdateInstance(ctx, id, appPreOrderUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appPreOrderUpdateRequest** | [**AppPreOrderUpdateRequest**](AppPreOrderUpdateRequest.md)| AppPreOrder representation | 

### Return type

[**AppPreOrderResponse**](AppPreOrderResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

