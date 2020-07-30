# \RoutingAppCoveragesApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RoutingAppCoveragesCreateInstance**](RoutingAppCoveragesApi.md#RoutingAppCoveragesCreateInstance) | **Post** /v1/routingAppCoverages | 
[**RoutingAppCoveragesDeleteInstance**](RoutingAppCoveragesApi.md#RoutingAppCoveragesDeleteInstance) | **Delete** /v1/routingAppCoverages/{id} | 
[**RoutingAppCoveragesGetInstance**](RoutingAppCoveragesApi.md#RoutingAppCoveragesGetInstance) | **Get** /v1/routingAppCoverages/{id} | 
[**RoutingAppCoveragesUpdateInstance**](RoutingAppCoveragesApi.md#RoutingAppCoveragesUpdateInstance) | **Patch** /v1/routingAppCoverages/{id} | 



## RoutingAppCoveragesCreateInstance

> RoutingAppCoverageResponse RoutingAppCoveragesCreateInstance(ctx, routingAppCoverageCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routingAppCoverageCreateRequest** | [**RoutingAppCoverageCreateRequest**](RoutingAppCoverageCreateRequest.md)| RoutingAppCoverage representation | 

### Return type

[**RoutingAppCoverageResponse**](RoutingAppCoverageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoutingAppCoveragesDeleteInstance

> RoutingAppCoveragesDeleteInstance(ctx, id)



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


## RoutingAppCoveragesGetInstance

> RoutingAppCoverageResponse RoutingAppCoveragesGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***RoutingAppCoveragesGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RoutingAppCoveragesGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsRoutingAppCoverages** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type routingAppCoverages | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**RoutingAppCoverageResponse**](RoutingAppCoverageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoutingAppCoveragesUpdateInstance

> RoutingAppCoverageResponse RoutingAppCoveragesUpdateInstance(ctx, id, routingAppCoverageUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**routingAppCoverageUpdateRequest** | [**RoutingAppCoverageUpdateRequest**](RoutingAppCoverageUpdateRequest.md)| RoutingAppCoverage representation | 

### Return type

[**RoutingAppCoverageResponse**](RoutingAppCoverageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

