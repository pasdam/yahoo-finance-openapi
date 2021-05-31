# \ChartApi

All URIs are relative to *https://query1.finance.yahoo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChart**](ChartApi.md#GetChart) | **Get** /v8/finance/chart/{symbol} | 



## GetChart

> ChartResponse GetChart(ctx, symbol).Interval(interval).Period1(period1).Period2(period2).Region(region).IncludePrePost(includePrePost).Lang(lang).UseYfid(useYfid).CorsDomain(corsDomain).Tsrc(tsrc).Execute()



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
    symbol := "symbol_example" // string | 
    interval := openapiclient.Interval("1m") // Interval | 
    period1 := int32(56) // int32 | 
    period2 := int32(56) // int32 | 
    region := "region_example" // string |  (optional)
    includePrePost := true // bool |  (optional) (default to false)
    lang := "lang_example" // string |  (optional)
    useYfid := true // bool |  (optional) (default to true)
    corsDomain := "corsDomain_example" // string |  (optional)
    tsrc := "tsrc_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChartApi.GetChart(context.Background(), symbol).Interval(interval).Period1(period1).Period2(period2).Region(region).IncludePrePost(includePrePost).Lang(lang).UseYfid(useYfid).CorsDomain(corsDomain).Tsrc(tsrc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartApi.GetChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChart`: ChartResponse
    fmt.Fprintf(os.Stdout, "Response from `ChartApi.GetChart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **interval** | [**Interval**](Interval.md) |  | 
 **period1** | **int32** |  | 
 **period2** | **int32** |  | 
 **region** | **string** |  | 
 **includePrePost** | **bool** |  | [default to false]
 **lang** | **string** |  | 
 **useYfid** | **bool** |  | [default to true]
 **corsDomain** | **string** |  | 
 **tsrc** | **string** |  | 

### Return type

[**ChartResponse**](ChartResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

