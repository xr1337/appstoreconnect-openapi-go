# \AppPriceTiersApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppPriceTiersGetCollection**](AppPriceTiersApi.md#AppPriceTiersGetCollection) | **Get** /v1/appPriceTiers | 
[**AppPriceTiersGetInstance**](AppPriceTiersApi.md#AppPriceTiersGetInstance) | **Get** /v1/appPriceTiers/{id} | 
[**AppPriceTiersPricePointsGetToManyRelated**](AppPriceTiersApi.md#AppPriceTiersPricePointsGetToManyRelated) | **Get** /v1/appPriceTiers/{id}/pricePoints | 



## AppPriceTiersGetCollection

> AppPriceTiersResponse AppPriceTiersGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppPriceTiersGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppPriceTiersGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterId** | [**optional.Interface of []string**](string.md)| filter by id(s) | 
 **fieldsAppPriceTiers** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPriceTiers | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsAppPricePoints** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPricePoints | 
 **limitPricePoints** | **optional.Int32**| maximum number of related pricePoints returned (when they are included) | 

### Return type

[**AppPriceTiersResponse**](AppPriceTiersResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPriceTiersGetInstance

> AppPriceTierResponse AppPriceTiersGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppPriceTiersGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppPriceTiersGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPriceTiers** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPriceTiers | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsAppPricePoints** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPricePoints | 
 **limitPricePoints** | **optional.Int32**| maximum number of related pricePoints returned (when they are included) | 

### Return type

[**AppPriceTierResponse**](AppPriceTierResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPriceTiersPricePointsGetToManyRelated

> AppPricePointsResponse AppPriceTiersPricePointsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppPriceTiersPricePointsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppPriceTiersPricePointsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPricePoints** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPricePoints | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**AppPricePointsResponse**](AppPricePointsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

