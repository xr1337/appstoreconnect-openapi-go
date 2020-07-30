# \BundleIdCapabilitiesApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BundleIdCapabilitiesCreateInstance**](BundleIdCapabilitiesApi.md#BundleIdCapabilitiesCreateInstance) | **Post** /v1/bundleIdCapabilities | 
[**BundleIdCapabilitiesDeleteInstance**](BundleIdCapabilitiesApi.md#BundleIdCapabilitiesDeleteInstance) | **Delete** /v1/bundleIdCapabilities/{id} | 
[**BundleIdCapabilitiesUpdateInstance**](BundleIdCapabilitiesApi.md#BundleIdCapabilitiesUpdateInstance) | **Patch** /v1/bundleIdCapabilities/{id} | 



## BundleIdCapabilitiesCreateInstance

> BundleIdCapabilityResponse BundleIdCapabilitiesCreateInstance(ctx, bundleIdCapabilityCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleIdCapabilityCreateRequest** | [**BundleIdCapabilityCreateRequest**](BundleIdCapabilityCreateRequest.md)| BundleIdCapability representation | 

### Return type

[**BundleIdCapabilityResponse**](BundleIdCapabilityResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdCapabilitiesDeleteInstance

> BundleIdCapabilitiesDeleteInstance(ctx, id)



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


## BundleIdCapabilitiesUpdateInstance

> BundleIdCapabilityResponse BundleIdCapabilitiesUpdateInstance(ctx, id, bundleIdCapabilityUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**bundleIdCapabilityUpdateRequest** | [**BundleIdCapabilityUpdateRequest**](BundleIdCapabilityUpdateRequest.md)| BundleIdCapability representation | 

### Return type

[**BundleIdCapabilityResponse**](BundleIdCapabilityResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

