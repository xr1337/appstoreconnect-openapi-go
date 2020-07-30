# \FinanceReportsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FinanceReportsGetCollection**](FinanceReportsApi.md#FinanceReportsGetCollection) | **Get** /v1/financeReports | 



## FinanceReportsGetCollection

> *os.File FinanceReportsGetCollection(ctx, filterRegionCode, filterReportDate, filterReportType, filterVendorNumber)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterRegionCode** | [**[]string**](string.md)| filter by attribute &#39;regionCode&#39; | 
**filterReportDate** | [**[]string**](string.md)| filter by attribute &#39;reportDate&#39; | 
**filterReportType** | [**[]string**](string.md)| filter by attribute &#39;reportType&#39; | 
**filterVendorNumber** | [**[]string**](string.md)| filter by attribute &#39;vendorNumber&#39; | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

