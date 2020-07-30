# \BuildBetaNotificationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildBetaNotificationsCreateInstance**](BuildBetaNotificationsApi.md#BuildBetaNotificationsCreateInstance) | **Post** /v1/buildBetaNotifications | 



## BuildBetaNotificationsCreateInstance

> BuildBetaNotificationResponse BuildBetaNotificationsCreateInstance(ctx, buildBetaNotificationCreateRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildBetaNotificationCreateRequest** | [**BuildBetaNotificationCreateRequest**](BuildBetaNotificationCreateRequest.md)| BuildBetaNotification representation | 

### Return type

[**BuildBetaNotificationResponse**](BuildBetaNotificationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

