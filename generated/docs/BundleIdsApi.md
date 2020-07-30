# \BundleIdsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BundleIdsAppGetToOneRelated**](BundleIdsApi.md#BundleIdsAppGetToOneRelated) | **Get** /v1/bundleIds/{id}/app | 
[**BundleIdsBundleIdCapabilitiesGetToManyRelated**](BundleIdsApi.md#BundleIdsBundleIdCapabilitiesGetToManyRelated) | **Get** /v1/bundleIds/{id}/bundleIdCapabilities | 
[**BundleIdsCreateInstance**](BundleIdsApi.md#BundleIdsCreateInstance) | **Post** /v1/bundleIds | 
[**BundleIdsDeleteInstance**](BundleIdsApi.md#BundleIdsDeleteInstance) | **Delete** /v1/bundleIds/{id} | 
[**BundleIdsGetCollection**](BundleIdsApi.md#BundleIdsGetCollection) | **Get** /v1/bundleIds | 
[**BundleIdsGetInstance**](BundleIdsApi.md#BundleIdsGetInstance) | **Get** /v1/bundleIds/{id} | 
[**BundleIdsProfilesGetToManyRelated**](BundleIdsApi.md#BundleIdsProfilesGetToManyRelated) | **Get** /v1/bundleIds/{id}/profiles | 
[**BundleIdsUpdateInstance**](BundleIdsApi.md#BundleIdsUpdateInstance) | **Patch** /v1/bundleIds/{id} | 



## BundleIdsAppGetToOneRelated

> AppResponse BundleIdsAppGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BundleIdsAppGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BundleIdsAppGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 

### Return type

[**AppResponse**](AppResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsBundleIdCapabilitiesGetToManyRelated

> BundleIdCapabilitiesResponse BundleIdsBundleIdCapabilitiesGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BundleIdsBundleIdCapabilitiesGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BundleIdsBundleIdCapabilitiesGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBundleIdCapabilities** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type bundleIdCapabilities | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**BundleIdCapabilitiesResponse**](BundleIdCapabilitiesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsCreateInstance

> BundleIdResponse BundleIdsCreateInstance(ctx, bundleIdCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleIdCreateRequest** | [**BundleIdCreateRequest**](BundleIdCreateRequest.md)| BundleId representation | 

### Return type

[**BundleIdResponse**](BundleIdResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsDeleteInstance

> BundleIdsDeleteInstance(ctx, id)



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


## BundleIdsGetCollection

> BundleIdsResponse BundleIdsGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BundleIdsGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BundleIdsGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterIdentifier** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;identifier&#39; | 
 **filterName** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;name&#39; | 
 **filterPlatform** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;platform&#39; | 
 **filterSeedId** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;seedId&#39; | 
 **filterId** | [**optional.Interface of []string**](string.md)| filter by id(s) | 
 **sort** | [**optional.Interface of []string**](string.md)| comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsBundleIds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type bundleIds | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBundleIdCapabilities** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type bundleIdCapabilities | 
 **fieldsProfiles** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type profiles | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limitBundleIdCapabilities** | **optional.Int32**| maximum number of related bundleIdCapabilities returned (when they are included) | 
 **limitProfiles** | **optional.Int32**| maximum number of related profiles returned (when they are included) | 

### Return type

[**BundleIdsResponse**](BundleIdsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsGetInstance

> BundleIdResponse BundleIdsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BundleIdsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BundleIdsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBundleIds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type bundleIds | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBundleIdCapabilities** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type bundleIdCapabilities | 
 **fieldsProfiles** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type profiles | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limitBundleIdCapabilities** | **optional.Int32**| maximum number of related bundleIdCapabilities returned (when they are included) | 
 **limitProfiles** | **optional.Int32**| maximum number of related profiles returned (when they are included) | 

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


## BundleIdsProfilesGetToManyRelated

> ProfilesResponse BundleIdsProfilesGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BundleIdsProfilesGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BundleIdsProfilesGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsProfiles** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type profiles | 
 **limit** | **optional.Int32**| maximum resources per page | 

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


## BundleIdsUpdateInstance

> BundleIdResponse BundleIdsUpdateInstance(ctx, id, bundleIdUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**bundleIdUpdateRequest** | [**BundleIdUpdateRequest**](BundleIdUpdateRequest.md)| BundleId representation | 

### Return type

[**BundleIdResponse**](BundleIdResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

