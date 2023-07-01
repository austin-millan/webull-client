# \DividendsAPI

All URIs are relative to *https://quoteapi.webull.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDividends**](DividendsAPI.md#GetDividends) | **Get** /v2/account/{account_id}/dividends | getDividends



## GetDividends

> []GetDividendsResponse GetDividends(ctx, accountId).Did(did).AccessToken(accessToken).Direct(direct).Execute()

getDividends



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
    accountId := "accountId_example" // string | Account ID
    direct := "direct_example" // string | Dividends (optional) (default to "in")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DividendsAPI.GetDividends(context.Background(), accountId).Did(did).AccessToken(accessToken).Direct(direct).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DividendsAPI.GetDividends``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDividends`: []GetDividendsResponse
    fmt.Fprintf(os.Stdout, "Response from `DividendsAPI.GetDividends`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDividendsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 

 **direct** | **string** | Dividends | [default to &quot;in&quot;]

### Return type

[**[]GetDividendsResponse**](GetDividendsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

