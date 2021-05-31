# \QuoteApi

All URIs are relative to *https://query1.finance.yahoo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQuote**](QuoteApi.md#GetQuote) | **Get** /v7/finance/quote | Returns quotes for the specified symbols



## GetQuote

> QuoteResponse GetQuote(ctx).Symbols(symbols).Formatted(formatted).Region(region).Lang(lang).IncludePrePost(includePrePost).Fields(fields).CorsDomain(corsDomain).Execute()

Returns quotes for the specified symbols



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    symbols := "symbols_example" // string | 
    formatted := true // bool |  (optional) (default to true)
    region := "region_example" // string |  (optional)
    lang := "lang_example" // string |  (optional)
    includePrePost := true // bool |  (optional) (default to false)
    fields := "fields_example" // string |  (optional)
    corsDomain := "corsDomain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuoteApi.GetQuote(context.Background()).Symbols(symbols).Formatted(formatted).Region(region).Lang(lang).IncludePrePost(includePrePost).Fields(fields).CorsDomain(corsDomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuoteApi.GetQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQuote`: QuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `QuoteApi.GetQuote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbols** | **string** |  | 
 **formatted** | **bool** |  | [default to true]
 **region** | **string** |  | 
 **lang** | **string** |  | 
 **includePrePost** | **bool** |  | [default to false]
 **fields** | **string** |  | 
 **corsDomain** | **string** |  | 

### Return type

[**QuoteResponse**](QuoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

