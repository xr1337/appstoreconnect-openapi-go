# \DiagnosticSignaturesApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiagnosticSignaturesLogsGetToManyRelated**](DiagnosticSignaturesApi.md#DiagnosticSignaturesLogsGetToManyRelated) | **Get** /v1/diagnosticSignatures/{id}/logs | 



## DiagnosticSignaturesLogsGetToManyRelated

> DiagnosticLogsResponse DiagnosticSignaturesLogsGetToManyRelated(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the id of the requested resource | 
 **optional** | ***DiagnosticSignaturesLogsGetToManyRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DiagnosticSignaturesLogsGetToManyRelatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| maximum resources per page | 

### Return type

[**DiagnosticLogsResponse**](DiagnosticLogsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

