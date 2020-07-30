# \SalesReportsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SalesReportsGetCollection**](SalesReportsApi.md#SalesReportsGetCollection) | **Get** /v1/salesReports | 



## SalesReportsGetCollection

> *os.File SalesReportsGetCollection(ctx, filterFrequency, filterReportSubType, filterReportType, filterVendorNumber, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterFrequency** | [**[]string**](string.md)| filter by attribute &#39;frequency&#39; | 
**filterReportSubType** | [**[]string**](string.md)| filter by attribute &#39;reportSubType&#39; | 
**filterReportType** | [**[]string**](string.md)| filter by attribute &#39;reportType&#39; | 
**filterVendorNumber** | [**[]string**](string.md)| filter by attribute &#39;vendorNumber&#39; | 
 **optional** | ***SalesReportsGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SalesReportsGetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **filterReportDate** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;reportDate&#39; | 
 **filterVersion** | [**optional.Interface of []string**](string.md)| filter by attribute &#39;version&#39; | 

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

