# \EndUserLicenseAgreementsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndUserLicenseAgreementsCreateInstance**](EndUserLicenseAgreementsApi.md#EndUserLicenseAgreementsCreateInstance) | **Post** /v1/endUserLicenseAgreements | 
[**EndUserLicenseAgreementsDeleteInstance**](EndUserLicenseAgreementsApi.md#EndUserLicenseAgreementsDeleteInstance) | **Delete** /v1/endUserLicenseAgreements/{id} | 
[**EndUserLicenseAgreementsGetInstance**](EndUserLicenseAgreementsApi.md#EndUserLicenseAgreementsGetInstance) | **Get** /v1/endUserLicenseAgreements/{id} | 
[**EndUserLicenseAgreementsTerritoriesGetToManyRelated**](EndUserLicenseAgreementsApi.md#EndUserLicenseAgreementsTerritoriesGetToManyRelated) | **Get** /v1/endUserLicenseAgreements/{id}/territories | 
[**EndUserLicenseAgreementsUpdateInstance**](EndUserLicenseAgreementsApi.md#EndUserLicenseAgreementsUpdateInstance) | **Patch** /v1/endUserLicenseAgreements/{id} | 



## EndUserLicenseAgreementsCreateInstance

> EndUserLicenseAgreementResponse EndUserLicenseAgreementsCreateInstance(ctx, endUserLicenseAgreementCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endUserLicenseAgreementCreateRequest** | [**EndUserLicenseAgreementCreateRequest**](EndUserLicenseAgreementCreateRequest.md)| EndUserLicenseAgreement representation | 

### Return type

[**EndUserLicenseAgreementResponse**](EndUserLicenseAgreementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUserLicenseAgreementsDeleteInstance

> EndUserLicenseAgreementsDeleteInstance(ctx, id)



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


## EndUserLicenseAgreementsGetInstance

> EndUserLicenseAgreementResponse EndUserLicenseAgreementsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***EndUserLicenseAgreementsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EndUserLicenseAgreementsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsEndUserLicenseAgreements** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type endUserLicenseAgreements | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsTerritories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type territories | 
 **limitTerritories** | **optional.Int32**| maximum number of related territories returned (when they are included) | 

### Return type

[**EndUserLicenseAgreementResponse**](EndUserLicenseAgreementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUserLicenseAgreementsTerritoriesGetToManyRelated

> TerritoriesResponse EndUserLicenseAgreementsTerritoriesGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***EndUserLicenseAgreementsTerritoriesGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EndUserLicenseAgreementsTerritoriesGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsTerritories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type territories | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**TerritoriesResponse**](TerritoriesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUserLicenseAgreementsUpdateInstance

> EndUserLicenseAgreementResponse EndUserLicenseAgreementsUpdateInstance(ctx, id, endUserLicenseAgreementUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**endUserLicenseAgreementUpdateRequest** | [**EndUserLicenseAgreementUpdateRequest**](EndUserLicenseAgreementUpdateRequest.md)| EndUserLicenseAgreement representation | 

### Return type

[**EndUserLicenseAgreementResponse**](EndUserLicenseAgreementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

