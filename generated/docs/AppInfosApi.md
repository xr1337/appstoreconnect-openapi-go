# \AppInfosApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppInfosAppInfoLocalizationsGetToManyRelated**](AppInfosApi.md#AppInfosAppInfoLocalizationsGetToManyRelated) | **Get** /v1/appInfos/{id}/appInfoLocalizations | 
[**AppInfosGetInstance**](AppInfosApi.md#AppInfosGetInstance) | **Get** /v1/appInfos/{id} | 
[**AppInfosPrimaryCategoryGetToOneRelated**](AppInfosApi.md#AppInfosPrimaryCategoryGetToOneRelated) | **Get** /v1/appInfos/{id}/primaryCategory | 
[**AppInfosPrimarySubcategoryOneGetToOneRelated**](AppInfosApi.md#AppInfosPrimarySubcategoryOneGetToOneRelated) | **Get** /v1/appInfos/{id}/primarySubcategoryOne | 
[**AppInfosPrimarySubcategoryTwoGetToOneRelated**](AppInfosApi.md#AppInfosPrimarySubcategoryTwoGetToOneRelated) | **Get** /v1/appInfos/{id}/primarySubcategoryTwo | 
[**AppInfosSecondaryCategoryGetToOneRelated**](AppInfosApi.md#AppInfosSecondaryCategoryGetToOneRelated) | **Get** /v1/appInfos/{id}/secondaryCategory | 
[**AppInfosSecondarySubcategoryOneGetToOneRelated**](AppInfosApi.md#AppInfosSecondarySubcategoryOneGetToOneRelated) | **Get** /v1/appInfos/{id}/secondarySubcategoryOne | 
[**AppInfosSecondarySubcategoryTwoGetToOneRelated**](AppInfosApi.md#AppInfosSecondarySubcategoryTwoGetToOneRelated) | **Get** /v1/appInfos/{id}/secondarySubcategoryTwo | 
[**AppInfosUpdateInstance**](AppInfosApi.md#AppInfosUpdateInstance) | **Patch** /v1/appInfos/{id} | 



## AppInfosAppInfoLocalizationsGetToManyRelated

> AppInfoLocalizationsResponse AppInfosAppInfoLocalizationsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppInfosAppInfoLocalizationsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppInfosAppInfoLocalizationsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterLocale** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;locale&#39; | 
 **fieldsAppInfos** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appInfos | 
 **fieldsAppInfoLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appInfoLocalizations | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppInfoLocalizationsResponse**](AppInfoLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppInfosGetInstance

> AppInfoResponse AppInfosGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppInfosGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppInfosGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppInfos** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appInfos | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsAppCategories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appCategories | 
 **fieldsAppInfoLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appInfoLocalizations | 
 **limitAppInfoLocalizations** | **optional.Int32**| maximum number of related appInfoLocalizations returned (when they are included) | 

### Return type

[**AppInfoResponse**](AppInfoResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppInfosPrimaryCategoryGetToOneRelated

> AppCategoryResponse AppInfosPrimaryCategoryGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppInfosPrimaryCategoryGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppInfosPrimaryCategoryGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppCategories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appCategories | 

### Return type

[**AppCategoryResponse**](AppCategoryResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppInfosPrimarySubcategoryOneGetToOneRelated

> AppCategoryResponse AppInfosPrimarySubcategoryOneGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppInfosPrimarySubcategoryOneGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppInfosPrimarySubcategoryOneGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppCategories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appCategories | 

### Return type

[**AppCategoryResponse**](AppCategoryResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppInfosPrimarySubcategoryTwoGetToOneRelated

> AppCategoryResponse AppInfosPrimarySubcategoryTwoGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppInfosPrimarySubcategoryTwoGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppInfosPrimarySubcategoryTwoGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppCategories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appCategories | 

### Return type

[**AppCategoryResponse**](AppCategoryResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppInfosSecondaryCategoryGetToOneRelated

> AppCategoryResponse AppInfosSecondaryCategoryGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppInfosSecondaryCategoryGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppInfosSecondaryCategoryGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppCategories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appCategories | 

### Return type

[**AppCategoryResponse**](AppCategoryResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppInfosSecondarySubcategoryOneGetToOneRelated

> AppCategoryResponse AppInfosSecondarySubcategoryOneGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppInfosSecondarySubcategoryOneGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppInfosSecondarySubcategoryOneGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppCategories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appCategories | 

### Return type

[**AppCategoryResponse**](AppCategoryResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppInfosSecondarySubcategoryTwoGetToOneRelated

> AppCategoryResponse AppInfosSecondarySubcategoryTwoGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppInfosSecondarySubcategoryTwoGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppInfosSecondarySubcategoryTwoGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppCategories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appCategories | 

### Return type

[**AppCategoryResponse**](AppCategoryResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppInfosUpdateInstance

> AppInfoResponse AppInfosUpdateInstance(ctx, id, appInfoUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appInfoUpdateRequest** | [**AppInfoUpdateRequest**](AppInfoUpdateRequest.md)| AppInfo representation | 

### Return type

[**AppInfoResponse**](AppInfoResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

