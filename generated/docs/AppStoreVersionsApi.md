# \AppStoreVersionsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreVersionsAgeRatingDeclarationGetToOneRelated**](AppStoreVersionsApi.md#AppStoreVersionsAgeRatingDeclarationGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/ageRatingDeclaration | 
[**AppStoreVersionsAppStoreReviewDetailGetToOneRelated**](AppStoreVersionsApi.md#AppStoreVersionsAppStoreReviewDetailGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/appStoreReviewDetail | 
[**AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated**](AppStoreVersionsApi.md#AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated) | **Get** /v1/appStoreVersions/{id}/appStoreVersionLocalizations | 
[**AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated**](AppStoreVersionsApi.md#AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/appStoreVersionPhasedRelease | 
[**AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated**](AppStoreVersionsApi.md#AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/appStoreVersionSubmission | 
[**AppStoreVersionsBuildGetToOneRelated**](AppStoreVersionsApi.md#AppStoreVersionsBuildGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/build | 
[**AppStoreVersionsBuildGetToOneRelationship**](AppStoreVersionsApi.md#AppStoreVersionsBuildGetToOneRelationship) | **Get** /v1/appStoreVersions/{id}/relationships/build | 
[**AppStoreVersionsBuildUpdateToOneRelationship**](AppStoreVersionsApi.md#AppStoreVersionsBuildUpdateToOneRelationship) | **Patch** /v1/appStoreVersions/{id}/relationships/build | 
[**AppStoreVersionsCreateInstance**](AppStoreVersionsApi.md#AppStoreVersionsCreateInstance) | **Post** /v1/appStoreVersions | 
[**AppStoreVersionsDeleteInstance**](AppStoreVersionsApi.md#AppStoreVersionsDeleteInstance) | **Delete** /v1/appStoreVersions/{id} | 
[**AppStoreVersionsGetInstance**](AppStoreVersionsApi.md#AppStoreVersionsGetInstance) | **Get** /v1/appStoreVersions/{id} | 
[**AppStoreVersionsIdfaDeclarationGetToOneRelated**](AppStoreVersionsApi.md#AppStoreVersionsIdfaDeclarationGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/idfaDeclaration | 
[**AppStoreVersionsRoutingAppCoverageGetToOneRelated**](AppStoreVersionsApi.md#AppStoreVersionsRoutingAppCoverageGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/routingAppCoverage | 
[**AppStoreVersionsUpdateInstance**](AppStoreVersionsApi.md#AppStoreVersionsUpdateInstance) | **Patch** /v1/appStoreVersions/{id} | 



## AppStoreVersionsAgeRatingDeclarationGetToOneRelated

> AgeRatingDeclarationResponse AppStoreVersionsAgeRatingDeclarationGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppStoreVersionsAgeRatingDeclarationGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppStoreVersionsAgeRatingDeclarationGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAgeRatingDeclarations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type ageRatingDeclarations | 

### Return type

[**AgeRatingDeclarationResponse**](AgeRatingDeclarationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppStoreReviewDetailGetToOneRelated

> AppStoreReviewDetailResponse AppStoreVersionsAppStoreReviewDetailGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppStoreVersionsAppStoreReviewDetailGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppStoreVersionsAppStoreReviewDetailGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreReviewDetails** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreReviewDetails | 
 **fieldsAppStoreVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersions | 
 **fieldsAppStoreReviewAttachments** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreReviewAttachments | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppStoreReviewDetailResponse**](AppStoreReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated

> AppStoreVersionLocalizationsResponse AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersionLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersionLocalizations | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**AppStoreVersionLocalizationsResponse**](AppStoreVersionLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated

> AppStoreVersionPhasedReleaseResponse AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersionPhasedReleases** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersionPhasedReleases | 

### Return type

[**AppStoreVersionPhasedReleaseResponse**](AppStoreVersionPhasedReleaseResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated

> AppStoreVersionSubmissionResponse AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppStoreVersionsAppStoreVersionSubmissionGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppStoreVersionsAppStoreVersionSubmissionGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersions | 
 **fieldsAppStoreVersionSubmissions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersionSubmissions | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppStoreVersionSubmissionResponse**](AppStoreVersionSubmissionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsBuildGetToOneRelated

> BuildResponse AppStoreVersionsBuildGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppStoreVersionsBuildGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppStoreVersionsBuildGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 

### Return type

[**BuildResponse**](BuildResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsBuildGetToOneRelationship

> AppStoreVersionBuildLinkageResponse AppStoreVersionsBuildGetToOneRelationship(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 

### Return type

[**AppStoreVersionBuildLinkageResponse**](AppStoreVersionBuildLinkageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsBuildUpdateToOneRelationship

> AppStoreVersionsBuildUpdateToOneRelationship(ctx, id, appStoreVersionBuildLinkageRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appStoreVersionBuildLinkageRequest** | [**AppStoreVersionBuildLinkageRequest**](AppStoreVersionBuildLinkageRequest.md)| Related linkage | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsCreateInstance

> AppStoreVersionResponse AppStoreVersionsCreateInstance(ctx, appStoreVersionCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appStoreVersionCreateRequest** | [**AppStoreVersionCreateRequest**](AppStoreVersionCreateRequest.md)| AppStoreVersion representation | 

### Return type

[**AppStoreVersionResponse**](AppStoreVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsDeleteInstance

> AppStoreVersionsDeleteInstance(ctx, id)



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


## AppStoreVersionsGetInstance

> AppStoreVersionResponse AppStoreVersionsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppStoreVersionsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppStoreVersionsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersions | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsAppStoreVersionLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersionLocalizations | 
 **fieldsIdfaDeclarations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type idfaDeclarations | 
 **fieldsRoutingAppCoverages** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type routingAppCoverages | 
 **fieldsAppStoreVersionPhasedReleases** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersionPhasedReleases | 
 **fieldsAgeRatingDeclarations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type ageRatingDeclarations | 
 **fieldsAppStoreReviewDetails** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreReviewDetails | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 
 **fieldsAppStoreVersionSubmissions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersionSubmissions | 
 **limitAppStoreVersionLocalizations** | **optional.Int32**| maximum number of related appStoreVersionLocalizations returned (when they are included) | 

### Return type

[**AppStoreVersionResponse**](AppStoreVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsIdfaDeclarationGetToOneRelated

> IdfaDeclarationResponse AppStoreVersionsIdfaDeclarationGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppStoreVersionsIdfaDeclarationGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppStoreVersionsIdfaDeclarationGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsIdfaDeclarations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type idfaDeclarations | 

### Return type

[**IdfaDeclarationResponse**](IdfaDeclarationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsRoutingAppCoverageGetToOneRelated

> RoutingAppCoverageResponse AppStoreVersionsRoutingAppCoverageGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppStoreVersionsRoutingAppCoverageGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppStoreVersionsRoutingAppCoverageGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsRoutingAppCoverages** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type routingAppCoverages | 

### Return type

[**RoutingAppCoverageResponse**](RoutingAppCoverageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsUpdateInstance

> AppStoreVersionResponse AppStoreVersionsUpdateInstance(ctx, id, appStoreVersionUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appStoreVersionUpdateRequest** | [**AppStoreVersionUpdateRequest**](AppStoreVersionUpdateRequest.md)| AppStoreVersion representation | 

### Return type

[**AppStoreVersionResponse**](AppStoreVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

