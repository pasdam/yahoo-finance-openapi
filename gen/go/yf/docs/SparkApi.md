# \SparkApi

All URIs are relative to *https://query1.finance.yahoo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Spark**](SparkApi.md#Spark) | **Get** /v7/finance/spark | 



## Spark

> QuoteResponse Spark(ctx).Interval(interval).Range_(range_).Symbols(symbols).Lang(lang).IncludePrePost(includePrePost).IncludeTimestamps(includeTimestamps).Indicators(indicators).CorsDomain(corsDomain).Execute()



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
    interval := openapiclient.Interval("1m") // Interval | 
    range_ := openapiclient.Range("1d") // Range | 
    symbols := "symbols_example" // string | 
    lang := "lang_example" // string |  (optional)
    includePrePost := true // bool |  (optional) (default to false)
    includeTimestamps := true // bool |  (optional) (default to false)
    indicators := "indicators_example" // string |  (optional)
    corsDomain := "corsDomain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SparkApi.Spark(context.Background()).Interval(interval).Range_(range_).Symbols(symbols).Lang(lang).IncludePrePost(includePrePost).IncludeTimestamps(includeTimestamps).Indicators(indicators).CorsDomain(corsDomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SparkApi.Spark``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Spark`: QuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `SparkApi.Spark`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSparkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | [**Interval**](Interval.md) |  | 
 **range_** | [**Range**](Range.md) |  | 
 **symbols** | **string** |  | 
 **lang** | **string** |  | 
 **includePrePost** | **bool** |  | [default to false]
 **includeTimestamps** | **bool** |  | [default to false]
 **indicators** | **string** |  | 
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

