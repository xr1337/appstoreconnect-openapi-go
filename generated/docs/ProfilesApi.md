# \ProfilesApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProfilesBundleIdGetToOneRelated**](ProfilesApi.md#ProfilesBundleIdGetToOneRelated) | **Get** /v1/profiles/{id}/bundleId | 
[**ProfilesCertificatesGetToManyRelated**](ProfilesApi.md#ProfilesCertificatesGetToManyRelated) | **Get** /v1/profiles/{id}/certificates | 
[**ProfilesCreateInstance**](ProfilesApi.md#ProfilesCreateInstance) | **Post** /v1/profiles | 
[**ProfilesDeleteInstance**](ProfilesApi.md#ProfilesDeleteInstance) | **Delete** /v1/profiles/{id} | 
[**ProfilesDevicesGetToManyRelated**](ProfilesApi.md#ProfilesDevicesGetToManyRelated) | **Get** /v1/profiles/{id}/devices | 
[**ProfilesGetCollection**](ProfilesApi.md#ProfilesGetCollection) | **Get** /v1/profiles | 
[**ProfilesGetInstance**](ProfilesApi.md#ProfilesGetInstance) | **Get** /v1/profiles/{id} | 



## ProfilesBundleIdGetToOneRelated

> BundleIdResponse ProfilesBundleIdGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***ProfilesBundleIdGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProfilesBundleIdGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBundleIds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type bundleIds | 

### Return type

[**BundleIdResponse**](BundleIdResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfilesCertificatesGetToManyRelated

> CertificatesResponse ProfilesCertificatesGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***ProfilesCertificatesGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProfilesCertificatesGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## ProfilesCreateInstance

> ProfileResponse ProfilesCreateInstance(ctx, profileCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileCreateRequest** | [**ProfileCreateRequest**](ProfileCreateRequest.md)| Profile representation | 

### Return type

[**ProfileResponse**](ProfileResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfilesDeleteInstance

> ProfilesDeleteInstance(ctx, id)



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


## ProfilesDevicesGetToManyRelated

> DevicesResponse ProfilesDevicesGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***ProfilesDevicesGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProfilesDevicesGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## ProfilesGetCollection

> ProfilesResponse ProfilesGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProfilesGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProfilesGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterName** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;name&#39; | 
 **filterProfileState** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;profileState&#39; | 
 **filterProfileType** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;profileType&#39; | 
 **filterId** | [**optional.Interface of []string**](string.md)| filter by id(s) | 
 **sort** | [**optional.Interface of []string**](string.md)| comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsProfiles** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type profiles | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsCertificates** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type certificates | 
 **fieldsDevices** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type devices | 
 **fieldsBundleIds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type bundleIds | 
 **limitCertificates** | **optional.Int32**| maximum number of related certificates returned (when they are included) | 
 **limitDevices** | **optional.Int32**| maximum number of related devices returned (when they are included) | 

### Return type

[**ProfilesResponse**](ProfilesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfilesGetInstance

> ProfileResponse ProfilesGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***ProfilesGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProfilesGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsProfiles** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type profiles | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsCertificates** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type certificates | 
 **fieldsDevices** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type devices | 
 **fieldsBundleIds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type bundleIds | 
 **limitCertificates** | **optional.Int32**| maximum number of related certificates returned (when they are included) | 
 **limitDevices** | **optional.Int32**| maximum number of related devices returned (when they are included) | 

### Return type

[**ProfileResponse**](ProfileResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

