# \DevicesApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesCreateInstance**](DevicesApi.md#DevicesCreateInstance) | **Post** /v1/devices | 
[**DevicesGetCollection**](DevicesApi.md#DevicesGetCollection) | **Get** /v1/devices | 
[**DevicesGetInstance**](DevicesApi.md#DevicesGetInstance) | **Get** /v1/devices/{id} | 
[**DevicesUpdateInstance**](DevicesApi.md#DevicesUpdateInstance) | **Patch** /v1/devices/{id} | 



## DevicesCreateInstance

> DeviceResponse DevicesCreateInstance(ctx, deviceCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCreateRequest** | [**DeviceCreateRequest**](DeviceCreateRequest.md)| Device representation | 

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesGetCollection

> DevicesResponse DevicesGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DevicesGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterName** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;name&#39; | 
 **filterPlatform** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;platform&#39; | 
 **filterStatus** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;status&#39; | 
 **filterUdid** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;udid&#39; | 
 **filterId** | [**optional.Interface of []string**](string.md)| filter by id(s) | 
 **sort** | [**optional.Interface of []string**](string.md)| comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsDevices** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type devices | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**DevicesResponse**](DevicesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesGetInstance

> DeviceResponse DevicesGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***DevicesGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DevicesGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsDevices** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type devices | 

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesUpdateInstance

> DeviceResponse DevicesUpdateInstance(ctx, id, deviceUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**deviceUpdateRequest** | [**DeviceUpdateRequest**](DeviceUpdateRequest.md)| Device representation | 

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

