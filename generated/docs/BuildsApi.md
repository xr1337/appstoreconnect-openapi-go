# \BuildsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildsAppEncryptionDeclarationGetToOneRelated**](BuildsApi.md#BuildsAppEncryptionDeclarationGetToOneRelated) | **Get** /v1/builds/{id}/appEncryptionDeclaration | 
[**BuildsAppEncryptionDeclarationGetToOneRelationship**](BuildsApi.md#BuildsAppEncryptionDeclarationGetToOneRelationship) | **Get** /v1/builds/{id}/relationships/appEncryptionDeclaration | 
[**BuildsAppEncryptionDeclarationUpdateToOneRelationship**](BuildsApi.md#BuildsAppEncryptionDeclarationUpdateToOneRelationship) | **Patch** /v1/builds/{id}/relationships/appEncryptionDeclaration | 
[**BuildsAppGetToOneRelated**](BuildsApi.md#BuildsAppGetToOneRelated) | **Get** /v1/builds/{id}/app | 
[**BuildsAppStoreVersionGetToOneRelated**](BuildsApi.md#BuildsAppStoreVersionGetToOneRelated) | **Get** /v1/builds/{id}/appStoreVersion | 
[**BuildsBetaAppReviewSubmissionGetToOneRelated**](BuildsApi.md#BuildsBetaAppReviewSubmissionGetToOneRelated) | **Get** /v1/builds/{id}/betaAppReviewSubmission | 
[**BuildsBetaBuildLocalizationsGetToManyRelated**](BuildsApi.md#BuildsBetaBuildLocalizationsGetToManyRelated) | **Get** /v1/builds/{id}/betaBuildLocalizations | 
[**BuildsBetaGroupsCreateToManyRelationship**](BuildsApi.md#BuildsBetaGroupsCreateToManyRelationship) | **Post** /v1/builds/{id}/relationships/betaGroups | 
[**BuildsBetaGroupsDeleteToManyRelationship**](BuildsApi.md#BuildsBetaGroupsDeleteToManyRelationship) | **Delete** /v1/builds/{id}/relationships/betaGroups | 
[**BuildsBuildBetaDetailGetToOneRelated**](BuildsApi.md#BuildsBuildBetaDetailGetToOneRelated) | **Get** /v1/builds/{id}/buildBetaDetail | 
[**BuildsDiagnosticSignaturesGetToManyRelated**](BuildsApi.md#BuildsDiagnosticSignaturesGetToManyRelated) | **Get** /v1/builds/{id}/diagnosticSignatures | 
[**BuildsGetCollection**](BuildsApi.md#BuildsGetCollection) | **Get** /v1/builds | 
[**BuildsGetInstance**](BuildsApi.md#BuildsGetInstance) | **Get** /v1/builds/{id} | 
[**BuildsIconsGetToManyRelated**](BuildsApi.md#BuildsIconsGetToManyRelated) | **Get** /v1/builds/{id}/icons | 
[**BuildsIndividualTestersCreateToManyRelationship**](BuildsApi.md#BuildsIndividualTestersCreateToManyRelationship) | **Post** /v1/builds/{id}/relationships/individualTesters | 
[**BuildsIndividualTestersDeleteToManyRelationship**](BuildsApi.md#BuildsIndividualTestersDeleteToManyRelationship) | **Delete** /v1/builds/{id}/relationships/individualTesters | 
[**BuildsIndividualTestersGetToManyRelated**](BuildsApi.md#BuildsIndividualTestersGetToManyRelated) | **Get** /v1/builds/{id}/individualTesters | 
[**BuildsIndividualTestersGetToManyRelationship**](BuildsApi.md#BuildsIndividualTestersGetToManyRelationship) | **Get** /v1/builds/{id}/relationships/individualTesters | 
[**BuildsPerfPowerMetricsGetToManyRelated**](BuildsApi.md#BuildsPerfPowerMetricsGetToManyRelated) | **Get** /v1/builds/{id}/perfPowerMetrics | 
[**BuildsPreReleaseVersionGetToOneRelated**](BuildsApi.md#BuildsPreReleaseVersionGetToOneRelated) | **Get** /v1/builds/{id}/preReleaseVersion | 
[**BuildsUpdateInstance**](BuildsApi.md#BuildsUpdateInstance) | **Patch** /v1/builds/{id} | 



## BuildsAppEncryptionDeclarationGetToOneRelated

> AppEncryptionDeclarationResponse BuildsAppEncryptionDeclarationGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BuildsAppEncryptionDeclarationGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildsAppEncryptionDeclarationGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppEncryptionDeclarations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appEncryptionDeclarations | 

### Return type

[**AppEncryptionDeclarationResponse**](AppEncryptionDeclarationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsAppEncryptionDeclarationGetToOneRelationship

> BuildAppEncryptionDeclarationLinkageResponse BuildsAppEncryptionDeclarationGetToOneRelationship(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 

### Return type

[**BuildAppEncryptionDeclarationLinkageResponse**](BuildAppEncryptionDeclarationLinkageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsAppEncryptionDeclarationUpdateToOneRelationship

> BuildsAppEncryptionDeclarationUpdateToOneRelationship(ctx, id, buildAppEncryptionDeclarationLinkageRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**buildAppEncryptionDeclarationLinkageRequest** | [**BuildAppEncryptionDeclarationLinkageRequest**](BuildAppEncryptionDeclarationLinkageRequest.md)| Related linkage | 

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


## BuildsAppGetToOneRelated

> AppResponse BuildsAppGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BuildsAppGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildsAppGetToOneRelatedOpts struct


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


## BuildsAppStoreVersionGetToOneRelated

> AppStoreVersionResponse BuildsAppStoreVersionGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BuildsAppStoreVersionGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildsAppStoreVersionGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersions | 

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


## BuildsBetaAppReviewSubmissionGetToOneRelated

> BetaAppReviewSubmissionResponse BuildsBetaAppReviewSubmissionGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BuildsBetaAppReviewSubmissionGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildsBetaAppReviewSubmissionGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppReviewSubmissions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaAppReviewSubmissions | 

### Return type

[**BetaAppReviewSubmissionResponse**](BetaAppReviewSubmissionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsBetaBuildLocalizationsGetToManyRelated

> BetaBuildLocalizationsResponse BuildsBetaBuildLocalizationsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BuildsBetaBuildLocalizationsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildsBetaBuildLocalizationsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaBuildLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaBuildLocalizations | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**BetaBuildLocalizationsResponse**](BetaBuildLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsBetaGroupsCreateToManyRelationship

> BuildsBetaGroupsCreateToManyRelationship(ctx, id, buildBetaGroupsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**buildBetaGroupsLinkagesRequest** | [**BuildBetaGroupsLinkagesRequest**](BuildBetaGroupsLinkagesRequest.md)| List of related linkages | 

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


## BuildsBetaGroupsDeleteToManyRelationship

> BuildsBetaGroupsDeleteToManyRelationship(ctx, id, buildBetaGroupsLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**buildBetaGroupsLinkagesRequest** | [**BuildBetaGroupsLinkagesRequest**](BuildBetaGroupsLinkagesRequest.md)| List of related linkages | 

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


## BuildsBuildBetaDetailGetToOneRelated

> BuildBetaDetailResponse BuildsBuildBetaDetailGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BuildsBuildBetaDetailGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildsBuildBetaDetailGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuildBetaDetails** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type buildBetaDetails | 

### Return type

[**BuildBetaDetailResponse**](BuildBetaDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsDiagnosticSignaturesGetToManyRelated

> DiagnosticSignaturesResponse BuildsDiagnosticSignaturesGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BuildsDiagnosticSignaturesGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildsDiagnosticSignaturesGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterDiagnosticType** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;diagnosticType&#39; | 
 **fieldsDiagnosticSignatures** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type diagnosticSignatures | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**DiagnosticSignaturesResponse**](DiagnosticSignaturesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsGetCollection

> BuildsResponse BuildsGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BuildsGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildsGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterBetaAppReviewSubmissionBetaReviewState** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;betaAppReviewSubmission.betaReviewState&#39; | 
 **filterExpired** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;expired&#39; | 
 **filterPreReleaseVersionPlatform** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;preReleaseVersion.platform&#39; | 
 **filterPreReleaseVersionVersion** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;preReleaseVersion.version&#39; | 
 **filterProcessingState** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;processingState&#39; | 
 **filterUsesNonExemptEncryption** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;usesNonExemptEncryption&#39; | 
 **filterVersion** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;version&#39; | 
 **filterApp** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;app&#39; | 
 **filterAppStoreVersion** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;appStoreVersion&#39; | 
 **filterBetaGroups** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;betaGroups&#39; | 
 **filterPreReleaseVersion** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;preReleaseVersion&#39; | 
 **filterId** | [**optional.Interface of []string**](string.md)| filter by id(s) | 
 **sort** | [**optional.Interface of []string**](string.md)| comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsAppEncryptionDeclarations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appEncryptionDeclarations | 
 **fieldsBetaAppReviewSubmissions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaAppReviewSubmissions | 
 **fieldsBuildBetaDetails** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type buildBetaDetails | 
 **fieldsBuildIcons** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type buildIcons | 
 **fieldsPerfPowerMetrics** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type perfPowerMetrics | 
 **fieldsPreReleaseVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type preReleaseVersions | 
 **fieldsAppStoreVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersions | 
 **fieldsDiagnosticSignatures** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type diagnosticSignatures | 
 **fieldsBetaTesters** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaTesters | 
 **fieldsBetaBuildLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaBuildLocalizations | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limitBetaBuildLocalizations** | **optional.Int32**| maximum number of related betaBuildLocalizations returned (when they are included) | 
 **limitIcons** | **optional.Int32**| maximum number of related icons returned (when they are included) | 
 **limitIndividualTesters** | **optional.Int32**| maximum number of related individualTesters returned (when they are included) | 

### Return type

[**BuildsResponse**](BuildsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsGetInstance

> BuildResponse BuildsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BuildsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsAppEncryptionDeclarations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appEncryptionDeclarations | 
 **fieldsBetaAppReviewSubmissions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaAppReviewSubmissions | 
 **fieldsBuildBetaDetails** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type buildBetaDetails | 
 **fieldsBuildIcons** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type buildIcons | 
 **fieldsPerfPowerMetrics** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type perfPowerMetrics | 
 **fieldsPreReleaseVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type preReleaseVersions | 
 **fieldsAppStoreVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersions | 
 **fieldsDiagnosticSignatures** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type diagnosticSignatures | 
 **fieldsBetaTesters** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaTesters | 
 **fieldsBetaBuildLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaBuildLocalizations | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limitBetaBuildLocalizations** | **optional.Int32**| maximum number of related betaBuildLocalizations returned (when they are included) | 
 **limitIcons** | **optional.Int32**| maximum number of related icons returned (when they are included) | 
 **limitIndividualTesters** | **optional.Int32**| maximum number of related individualTesters returned (when they are included) | 

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


## BuildsIconsGetToManyRelated

> BuildIconsResponse BuildsIconsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BuildsIconsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildsIconsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuildIcons** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type buildIcons | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**BuildIconsResponse**](BuildIconsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsIndividualTestersCreateToManyRelationship

> BuildsIndividualTestersCreateToManyRelationship(ctx, id, buildIndividualTestersLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**buildIndividualTestersLinkagesRequest** | [**BuildIndividualTestersLinkagesRequest**](BuildIndividualTestersLinkagesRequest.md)| List of related linkages | 

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


## BuildsIndividualTestersDeleteToManyRelationship

> BuildsIndividualTestersDeleteToManyRelationship(ctx, id, buildIndividualTestersLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**buildIndividualTestersLinkagesRequest** | [**BuildIndividualTestersLinkagesRequest**](BuildIndividualTestersLinkagesRequest.md)| List of related linkages | 

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


## BuildsIndividualTestersGetToManyRelated

> BetaTestersResponse BuildsIndividualTestersGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BuildsIndividualTestersGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildsIndividualTestersGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaTesters** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaTesters | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**BetaTestersResponse**](BetaTestersResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsIndividualTestersGetToManyRelationship

> BuildIndividualTestersLinkagesResponse BuildsIndividualTestersGetToManyRelationship(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BuildsIndividualTestersGetToManyRelationshipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildsIndividualTestersGetToManyRelationshipOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**BuildIndividualTestersLinkagesResponse**](BuildIndividualTestersLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsPerfPowerMetricsGetToManyRelated

> PerfPowerMetricsResponse BuildsPerfPowerMetricsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BuildsPerfPowerMetricsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildsPerfPowerMetricsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterDeviceType** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;deviceType&#39; | 
 **filterMetricType** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;metricType&#39; | 
 **filterPlatform** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;platform&#39; | 

### Return type

[**PerfPowerMetricsResponse**](PerfPowerMetricsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsPreReleaseVersionGetToOneRelated

> PrereleaseVersionResponse BuildsPreReleaseVersionGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***BuildsPreReleaseVersionGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BuildsPreReleaseVersionGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsPreReleaseVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type preReleaseVersions | 

### Return type

[**PrereleaseVersionResponse**](PrereleaseVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsUpdateInstance

> BuildResponse BuildsUpdateInstance(ctx, id, buildUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**buildUpdateRequest** | [**BuildUpdateRequest**](BuildUpdateRequest.md)| Build representation | 

### Return type

[**BuildResponse**](BuildResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

