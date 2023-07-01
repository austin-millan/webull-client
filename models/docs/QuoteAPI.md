# \QuoteAPI

All URIs are relative to *https://quoteapi.webull.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStockQuote**](QuoteAPI.md#GetStockQuote) | **Get** /quote/tickerRealTimes/v5/{stock} | getStockQuote



## GetStockQuote

> GetStockQuoteResponse GetStockQuote(ctx, stock).Did(did).AccessToken(accessToken).Execute()

getStockQuote



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    did := "did_example" // string | Device ID (default to "your_device")
    accessToken := "accessToken_example" // string | Access token
    stock := int32(56) // int32 | Internal stock ticker ID (default to 913243251)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuoteAPI.GetStockQuote(context.Background(), stock).Did(did).AccessToken(accessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuoteAPI.GetStockQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStockQuote`: GetStockQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `QuoteAPI.GetStockQuote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stock** | **int32** | Internal stock ticker ID | [default to 913243251]

### Other Parameters

Other parameters are passed through a pointer to a apiGetStockQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 


### Return type

[**GetStockQuoteResponse**](GetStockQuoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

