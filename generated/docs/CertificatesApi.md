# \CertificatesApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificatesCreateInstance**](CertificatesApi.md#CertificatesCreateInstance) | **Post** /v1/certificates | 
[**CertificatesDeleteInstance**](CertificatesApi.md#CertificatesDeleteInstance) | **Delete** /v1/certificates/{id} | 
[**CertificatesGetCollection**](CertificatesApi.md#CertificatesGetCollection) | **Get** /v1/certificates | 
[**CertificatesGetInstance**](CertificatesApi.md#CertificatesGetInstance) | **Get** /v1/certificates/{id} | 



## CertificatesCreateInstance

> CertificateResponse CertificatesCreateInstance(ctx, certificateCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateCreateRequest** | [**CertificateCreateRequest**](CertificateCreateRequest.md)| Certificate representation | 

### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificatesDeleteInstance

> CertificatesDeleteInstance(ctx, id)



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


## CertificatesGetCollection

> CertificatesResponse CertificatesGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CertificatesGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CertificatesGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCertificateType** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;certificateType&#39; | 
 **filterDisplayName** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;displayName&#39; | 
 **filterSerialNumber** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;serialNumber&#39; | 
 **filterId** | [**optional.Interface of []string**](string.md)| filter by id(s) | 
 **sort** | [**optional.Interface of []string**](string.md)| comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsCertificates** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type certificates | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**CertificatesResponse**](CertificatesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificatesGetInstance

> CertificateResponse CertificatesGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***CertificatesGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CertificatesGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCertificates** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type certificates | 

### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

