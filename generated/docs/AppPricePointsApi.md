# \AppPricePointsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppPricePointsGetCollection**](AppPricePointsApi.md#AppPricePointsGetCollection) | **Get** /v1/appPricePoints | 
[**AppPricePointsGetInstance**](AppPricePointsApi.md#AppPricePointsGetInstance) | **Get** /v1/appPricePoints/{id} | 
[**AppPricePointsTerritoryGetToOneRelated**](AppPricePointsApi.md#AppPricePointsTerritoryGetToOneRelated) | **Get** /v1/appPricePoints/{id}/territory | 



## AppPricePointsGetCollection

> AppPricePointsResponse AppPricePointsGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppPricePointsGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppPricePointsGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterPriceTier** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;priceTier&#39; | 
 **filterTerritory** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;territory&#39; | 
 **fieldsAppPricePoints** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPricePoints | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsTerritories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type territories | 

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


## AppPricePointsGetInstance

> AppPricePointResponse AppPricePointsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppPricePointsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppPricePointsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPricePoints** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPricePoints | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsTerritories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type territories | 

### Return type

[**AppPricePointResponse**](AppPricePointResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPricePointsTerritoryGetToOneRelated

> TerritoryResponse AppPricePointsTerritoryGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppPricePointsTerritoryGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppPricePointsTerritoryGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsTerritories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type territories | 

### Return type

[**TerritoryResponse**](TerritoryResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

