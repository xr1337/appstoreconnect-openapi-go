# \AppInfoLocalizationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppInfoLocalizationsCreateInstance**](AppInfoLocalizationsApi.md#AppInfoLocalizationsCreateInstance) | **Post** /v1/appInfoLocalizations | 
[**AppInfoLocalizationsDeleteInstance**](AppInfoLocalizationsApi.md#AppInfoLocalizationsDeleteInstance) | **Delete** /v1/appInfoLocalizations/{id} | 
[**AppInfoLocalizationsGetInstance**](AppInfoLocalizationsApi.md#AppInfoLocalizationsGetInstance) | **Get** /v1/appInfoLocalizations/{id} | 
[**AppInfoLocalizationsUpdateInstance**](AppInfoLocalizationsApi.md#AppInfoLocalizationsUpdateInstance) | **Patch** /v1/appInfoLocalizations/{id} | 



## AppInfoLocalizationsCreateInstance

> AppInfoLocalizationResponse AppInfoLocalizationsCreateInstance(ctx, appInfoLocalizationCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appInfoLocalizationCreateRequest** | [**AppInfoLocalizationCreateRequest**](AppInfoLocalizationCreateRequest.md)| AppInfoLocalization representation | 

### Return type

[**AppInfoLocalizationResponse**](AppInfoLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppInfoLocalizationsDeleteInstance

> AppInfoLocalizationsDeleteInstance(ctx, id)



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


## AppInfoLocalizationsGetInstance

> AppInfoLocalizationResponse AppInfoLocalizationsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppInfoLocalizationsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppInfoLocalizationsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppInfoLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appInfoLocalizations | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppInfoLocalizationResponse**](AppInfoLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppInfoLocalizationsUpdateInstance

> AppInfoLocalizationResponse AppInfoLocalizationsUpdateInstance(ctx, id, appInfoLocalizationUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appInfoLocalizationUpdateRequest** | [**AppInfoLocalizationUpdateRequest**](AppInfoLocalizationUpdateRequest.md)| AppInfoLocalization representation | 

### Return type

[**AppInfoLocalizationResponse**](AppInfoLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

