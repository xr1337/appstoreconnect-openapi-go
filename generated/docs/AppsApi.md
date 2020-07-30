# \AppsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsAppInfosGetToManyRelated**](AppsApi.md#AppsAppInfosGetToManyRelated) | **Get** /v1/apps/{id}/appInfos | 
[**AppsAppStoreVersionsGetToManyRelated**](AppsApi.md#AppsAppStoreVersionsGetToManyRelated) | **Get** /v1/apps/{id}/appStoreVersions | 
[**AppsAvailableTerritoriesGetToManyRelated**](AppsApi.md#AppsAvailableTerritoriesGetToManyRelated) | **Get** /v1/apps/{id}/availableTerritories | 
[**AppsBetaAppLocalizationsGetToManyRelated**](AppsApi.md#AppsBetaAppLocalizationsGetToManyRelated) | **Get** /v1/apps/{id}/betaAppLocalizations | 
[**AppsBetaAppReviewDetailGetToOneRelated**](AppsApi.md#AppsBetaAppReviewDetailGetToOneRelated) | **Get** /v1/apps/{id}/betaAppReviewDetail | 
[**AppsBetaGroupsGetToManyRelated**](AppsApi.md#AppsBetaGroupsGetToManyRelated) | **Get** /v1/apps/{id}/betaGroups | 
[**AppsBetaLicenseAgreementGetToOneRelated**](AppsApi.md#AppsBetaLicenseAgreementGetToOneRelated) | **Get** /v1/apps/{id}/betaLicenseAgreement | 
[**AppsBetaTestersDeleteToManyRelationship**](AppsApi.md#AppsBetaTestersDeleteToManyRelationship) | **Delete** /v1/apps/{id}/relationships/betaTesters | 
[**AppsBuildsGetToManyRelated**](AppsApi.md#AppsBuildsGetToManyRelated) | **Get** /v1/apps/{id}/builds | 
[**AppsEndUserLicenseAgreementGetToOneRelated**](AppsApi.md#AppsEndUserLicenseAgreementGetToOneRelated) | **Get** /v1/apps/{id}/endUserLicenseAgreement | 
[**AppsGameCenterEnabledVersionsGetToManyRelated**](AppsApi.md#AppsGameCenterEnabledVersionsGetToManyRelated) | **Get** /v1/apps/{id}/gameCenterEnabledVersions | 
[**AppsGetCollection**](AppsApi.md#AppsGetCollection) | **Get** /v1/apps | 
[**AppsGetInstance**](AppsApi.md#AppsGetInstance) | **Get** /v1/apps/{id} | 
[**AppsInAppPurchasesGetToManyRelated**](AppsApi.md#AppsInAppPurchasesGetToManyRelated) | **Get** /v1/apps/{id}/inAppPurchases | 
[**AppsPerfPowerMetricsGetToManyRelated**](AppsApi.md#AppsPerfPowerMetricsGetToManyRelated) | **Get** /v1/apps/{id}/perfPowerMetrics | 
[**AppsPreOrderGetToOneRelated**](AppsApi.md#AppsPreOrderGetToOneRelated) | **Get** /v1/apps/{id}/preOrder | 
[**AppsPreReleaseVersionsGetToManyRelated**](AppsApi.md#AppsPreReleaseVersionsGetToManyRelated) | **Get** /v1/apps/{id}/preReleaseVersions | 
[**AppsPricesGetToManyRelated**](AppsApi.md#AppsPricesGetToManyRelated) | **Get** /v1/apps/{id}/prices | 
[**AppsUpdateInstance**](AppsApi.md#AppsUpdateInstance) | **Patch** /v1/apps/{id} | 



## AppsAppInfosGetToManyRelated

> AppInfosResponse AppsAppInfosGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsAppInfosGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsAppInfosGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppInfos** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appInfos | 
 **fieldsAppCategories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appCategories | 
 **fieldsAppInfoLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appInfoLocalizations | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppInfosResponse**](AppInfosResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAppStoreVersionsGetToManyRelated

> AppStoreVersionsResponse AppsAppStoreVersionsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsAppStoreVersionsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsAppStoreVersionsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterAppStoreState** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;appStoreState&#39; | 
 **filterPlatform** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;platform&#39; | 
 **filterVersionString** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;versionString&#39; | 
 **filterId** | [**optional.Interface of []string**](string.md)| filter by id(s) | 
 **fieldsIdfaDeclarations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type idfaDeclarations | 
 **fieldsAppStoreVersionLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersionLocalizations | 
 **fieldsRoutingAppCoverages** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type routingAppCoverages | 
 **fieldsAppStoreVersionPhasedReleases** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersionPhasedReleases | 
 **fieldsAgeRatingDeclarations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type ageRatingDeclarations | 
 **fieldsAppStoreReviewDetails** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreReviewDetails | 
 **fieldsAppStoreVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersions | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 
 **fieldsAppStoreVersionSubmissions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersionSubmissions | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppStoreVersionsResponse**](AppStoreVersionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAvailableTerritoriesGetToManyRelated

> TerritoriesResponse AppsAvailableTerritoriesGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsAvailableTerritoriesGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsAvailableTerritoriesGetToManyRelatedOpts struct


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


## AppsBetaAppLocalizationsGetToManyRelated

> BetaAppLocalizationsResponse AppsBetaAppLocalizationsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsBetaAppLocalizationsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsBetaAppLocalizationsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaAppLocalizations | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**BetaAppLocalizationsResponse**](BetaAppLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsBetaAppReviewDetailGetToOneRelated

> BetaAppReviewDetailResponse AppsBetaAppReviewDetailGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsBetaAppReviewDetailGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsBetaAppReviewDetailGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppReviewDetails** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaAppReviewDetails | 

### Return type

[**BetaAppReviewDetailResponse**](BetaAppReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsBetaGroupsGetToManyRelated

> BetaGroupsResponse AppsBetaGroupsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsBetaGroupsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsBetaGroupsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaGroups** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaGroups | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**BetaGroupsResponse**](BetaGroupsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsBetaLicenseAgreementGetToOneRelated

> BetaLicenseAgreementResponse AppsBetaLicenseAgreementGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsBetaLicenseAgreementGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsBetaLicenseAgreementGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaLicenseAgreements** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaLicenseAgreements | 

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


## AppsBetaTestersDeleteToManyRelationship

> AppsBetaTestersDeleteToManyRelationship(ctx, id, appBetaTestersLinkagesRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appBetaTestersLinkagesRequest** | [**AppBetaTestersLinkagesRequest**](AppBetaTestersLinkagesRequest.md)| List of related linkages | 

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


## AppsBuildsGetToManyRelated

> BuildsResponse AppsBuildsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsBuildsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsBuildsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 
 **limit** | **optional.Int32**| maximum resources per page | 

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


## AppsEndUserLicenseAgreementGetToOneRelated

> EndUserLicenseAgreementResponse AppsEndUserLicenseAgreementGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsEndUserLicenseAgreementGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsEndUserLicenseAgreementGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsEndUserLicenseAgreements** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type endUserLicenseAgreements | 

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


## AppsGameCenterEnabledVersionsGetToManyRelated

> GameCenterEnabledVersionsResponse AppsGameCenterEnabledVersionsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsGameCenterEnabledVersionsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsGameCenterEnabledVersionsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterPlatform** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;platform&#39; | 
 **filterVersionString** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;versionString&#39; | 
 **filterId** | [**optional.Interface of []string**](string.md)| filter by id(s) | 
 **sort** | [**optional.Interface of []string**](string.md)| comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsGameCenterEnabledVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type gameCenterEnabledVersions | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**GameCenterEnabledVersionsResponse**](GameCenterEnabledVersionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGetCollection

> AppsResponse AppsGetCollection(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppsGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterAppStoreVersionsAppStoreState** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;appStoreVersions.appStoreState&#39; | 
 **filterAppStoreVersionsPlatform** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;appStoreVersions.platform&#39; | 
 **filterBundleId** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;bundleId&#39; | 
 **filterName** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;name&#39; | 
 **filterSku** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;sku&#39; | 
 **filterAppStoreVersions** | [**optional.Interface of []string**](string.md)| filter by id(s) of related &#39;appStoreVersions&#39; | 
 **filterId** | [**optional.Interface of []string**](string.md)| filter by id(s) | 
 **existsGameCenterEnabledVersions** | [**optional.Interface of []string**](string.md)| filter by existence or non-existence of related &#39;gameCenterEnabledVersions&#39; | 
 **sort** | [**optional.Interface of []string**](string.md)| comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBetaGroups** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaGroups | 
 **fieldsPerfPowerMetrics** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type perfPowerMetrics | 
 **fieldsAppInfos** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appInfos | 
 **fieldsAppPreOrders** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPreOrders | 
 **fieldsPreReleaseVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type preReleaseVersions | 
 **fieldsAppPrices** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPrices | 
 **fieldsInAppPurchases** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type inAppPurchases | 
 **fieldsBetaAppReviewDetails** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaAppReviewDetails | 
 **fieldsTerritories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type territories | 
 **fieldsGameCenterEnabledVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type gameCenterEnabledVersions | 
 **fieldsAppStoreVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersions | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 
 **fieldsBetaAppLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaAppLocalizations | 
 **fieldsBetaLicenseAgreements** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaLicenseAgreements | 
 **fieldsEndUserLicenseAgreements** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type endUserLicenseAgreements | 
 **limitAppInfos** | **optional.Int32**| maximum number of related appInfos returned (when they are included) | 
 **limitAppStoreVersions** | **optional.Int32**| maximum number of related appStoreVersions returned (when they are included) | 
 **limitAvailableTerritories** | **optional.Int32**| maximum number of related availableTerritories returned (when they are included) | 
 **limitBetaAppLocalizations** | **optional.Int32**| maximum number of related betaAppLocalizations returned (when they are included) | 
 **limitBetaGroups** | **optional.Int32**| maximum number of related betaGroups returned (when they are included) | 
 **limitBuilds** | **optional.Int32**| maximum number of related builds returned (when they are included) | 
 **limitGameCenterEnabledVersions** | **optional.Int32**| maximum number of related gameCenterEnabledVersions returned (when they are included) | 
 **limitInAppPurchases** | **optional.Int32**| maximum number of related inAppPurchases returned (when they are included) | 
 **limitPreReleaseVersions** | **optional.Int32**| maximum number of related preReleaseVersions returned (when they are included) | 
 **limitPrices** | **optional.Int32**| maximum number of related prices returned (when they are included) | 

### Return type

[**AppsResponse**](AppsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGetInstance

> AppResponse AppsGetInstance(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsGetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsGetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 
 **fieldsBetaGroups** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaGroups | 
 **fieldsPerfPowerMetrics** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type perfPowerMetrics | 
 **fieldsAppInfos** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appInfos | 
 **fieldsAppPreOrders** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPreOrders | 
 **fieldsPreReleaseVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type preReleaseVersions | 
 **fieldsAppPrices** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPrices | 
 **fieldsInAppPurchases** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type inAppPurchases | 
 **fieldsBetaAppReviewDetails** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaAppReviewDetails | 
 **fieldsTerritories** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type territories | 
 **fieldsGameCenterEnabledVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type gameCenterEnabledVersions | 
 **fieldsAppStoreVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appStoreVersions | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type builds | 
 **fieldsBetaAppLocalizations** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaAppLocalizations | 
 **fieldsBetaLicenseAgreements** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type betaLicenseAgreements | 
 **fieldsEndUserLicenseAgreements** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type endUserLicenseAgreements | 
 **limitAppInfos** | **optional.Int32**| maximum number of related appInfos returned (when they are included) | 
 **limitAppStoreVersions** | **optional.Int32**| maximum number of related appStoreVersions returned (when they are included) | 
 **limitAvailableTerritories** | **optional.Int32**| maximum number of related availableTerritories returned (when they are included) | 
 **limitBetaAppLocalizations** | **optional.Int32**| maximum number of related betaAppLocalizations returned (when they are included) | 
 **limitBetaGroups** | **optional.Int32**| maximum number of related betaGroups returned (when they are included) | 
 **limitBuilds** | **optional.Int32**| maximum number of related builds returned (when they are included) | 
 **limitGameCenterEnabledVersions** | **optional.Int32**| maximum number of related gameCenterEnabledVersions returned (when they are included) | 
 **limitInAppPurchases** | **optional.Int32**| maximum number of related inAppPurchases returned (when they are included) | 
 **limitPreReleaseVersions** | **optional.Int32**| maximum number of related preReleaseVersions returned (when they are included) | 
 **limitPrices** | **optional.Int32**| maximum number of related prices returned (when they are included) | 

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


## AppsInAppPurchasesGetToManyRelated

> InAppPurchasesResponse AppsInAppPurchasesGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsInAppPurchasesGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsInAppPurchasesGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterInAppPurchaseType** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;inAppPurchaseType&#39; | 
 **filterCanBeSubmitted** | [**optional.Interface of []string**](string.md)| filter by canBeSubmitted | 
 **sort** | [**optional.Interface of []string**](string.md)| comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsInAppPurchases** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type inAppPurchases | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**InAppPurchasesResponse**](InAppPurchasesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsPerfPowerMetricsGetToManyRelated

> PerfPowerMetricsResponse AppsPerfPowerMetricsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsPerfPowerMetricsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsPerfPowerMetricsGetToManyRelatedOpts struct


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


## AppsPreOrderGetToOneRelated

> AppPreOrderResponse AppsPreOrderGetToOneRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsPreOrderGetToOneRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsPreOrderGetToOneRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPreOrders** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPreOrders | 

### Return type

[**AppPreOrderResponse**](AppPreOrderResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsPreReleaseVersionsGetToManyRelated

> PreReleaseVersionsResponse AppsPreReleaseVersionsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsPreReleaseVersionsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsPreReleaseVersionsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsPreReleaseVersions** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type preReleaseVersions | 
 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**PreReleaseVersionsResponse**](PreReleaseVersionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsPricesGetToManyRelated

> AppPricesResponse AppsPricesGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***AppsPricesGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsPricesGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPrices** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPrices | 
 **fieldsAppPriceTiers** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type appPriceTiers | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| the fields to include for returned resources of type apps | 
 **limit** | **optional.Int32**| maximum resources per page | 
 **include** | [**optional.Interface of []string**](string.md)| comma-separated list of relationships to include | 

### Return type

[**AppPricesResponse**](AppPricesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsUpdateInstance

> AppResponse AppsUpdateInstance(ctx, id, appUpdateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
**appUpdateRequest** | [**AppUpdateRequest**](AppUpdateRequest.md)| App representation | 

### Return type

[**AppResponse**](AppResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

