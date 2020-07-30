# \AppCategoriesApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCategoriesGetCollection**](AppCategoriesApi.md#AppCategoriesGetCollection) | **Get** /v1/appCategories | 
[**AppCategoriesGetInstance**](AppCategoriesApi.md#AppCategoriesGetInstance) | **Get** /v1/appCategories/{id} | 
[**AppCategoriesParentGetToOneRelated**](AppCategoriesApi.md#AppCategoriesParentGetToOneRelated) | **Get** /v1/appCategories/{id}/parent | 
[**AppCategoriesSubcategoriesGetToManyRelated**](AppCategoriesApi.md#AppCategoriesSubcategoriesGetToManyRelated) | **Get** /v1/appCategories/{id}/subcategories | 



## AppCategoriesGetCollection

> AppCategoriesResponse AppCategoriesGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppCategoriesGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppCategoriesGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterPlatforms** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;platforms&#39; | 
 **existsParent** | [**optional.Interface of []string**](string.md)| filter by existence or non-existence of related &#39;parent&#39; | 
 **fieldsAppCategories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appCategories | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **limitSubcategories** | **optional.Int32**| maximum number of related subcategories returned (when they are included) | 

### Return type

[**AppCategoriesResponse**](AppCategoriesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCategoriesGetInstance

> AppCategoryResponse AppCategoriesGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppCategoriesGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppCategoriesGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppCategories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appCategories | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **limitSubcategories** | **optional.Int32**| maximum number of related subcategories returned (when they are included) | 

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


## AppCategoriesParentGetToOneRelated

> AppCategoryResponse AppCategoriesParentGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppCategoriesParentGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppCategoriesParentGetToOneRelatedOpts struct


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


## AppCategoriesSubcategoriesGetToManyRelated

> AppCategoriesResponse AppCategoriesSubcategoriesGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppCategoriesSubcategoriesGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppCategoriesSubcategoriesGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppCategories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appCategories | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**AppCategoriesResponse**](AppCategoriesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

