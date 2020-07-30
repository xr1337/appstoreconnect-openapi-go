# \BetaLicenseAgreementsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaLicenseAgreementsAppGetToOneRelated**](BetaLicenseAgreementsApi.md#BetaLicenseAgreementsAppGetToOneRelated) | **Get** /v1/betaLicenseAgreements/{id}/app | 
[**BetaLicenseAgreementsGetCollection**](BetaLicenseAgreementsApi.md#BetaLicenseAgreementsGetCollection) | **Get** /v1/betaLicenseAgreements | 
[**BetaLicenseAgreementsGetInstance**](BetaLicenseAgreementsApi.md#BetaLicenseAgreementsGetInstance) | **Get** /v1/betaLicenseAgreements/{id} | 
[**BetaLicenseAgreementsUpdateInstance**](BetaLicenseAgreementsApi.md#BetaLicenseAgreementsUpdateInstance) | **Patch** /v1/betaLicenseAgreements/{id} | 



## BetaLicenseAgreementsAppGetToOneRelated

> AppResponse BetaLicenseAgreementsAppGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaLicenseAgreementsAppGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaLicenseAgreementsAppGetToOneRelatedOpts struct


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


## BetaLicenseAgreementsGetCollection

> BetaLicenseAgreementsResponse BetaLicenseAgreementsGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BetaLicenseAgreementsGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaLicenseAgreementsGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterApp** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;app&#39; | 
 **fieldsBetaLicenseAgreements** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaLicenseAgreements | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 

### Return type

[**BetaLicenseAgreementsResponse**](BetaLicenseAgreementsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaLicenseAgreementsGetInstance

> BetaLicenseAgreementResponse BetaLicenseAgreementsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BetaLicenseAgreementsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BetaLicenseAgreementsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaLicenseAgreements** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaLicenseAgreements | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 

### Return type

[**BetaLicenseAgreementResponse**](BetaLicenseAgreementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaLicenseAgreementsUpdateInstance

> BetaLicenseAgreementResponse BetaLicenseAgreementsUpdateInstance(ctx, id, betaLicenseAgreementUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**betaLicenseAgreementUpdateRequest** | [**BetaLicenseAgreementUpdateRequest**](BetaLicenseAgreementUpdateRequest.md)| BetaLicenseAgreement representation | 

### Return type

[**BetaLicenseAgreementResponse**](BetaLicenseAgreementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

