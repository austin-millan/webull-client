# \OptionsAPI

All URIs are relative to *https://quoteapi.webull.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOptionQuotes**](OptionsAPI.md#GetOptionQuotes) | **Get** /quote/option/query/list | getOptionQuotes
[**GetStockOptions**](OptionsAPI.md#GetStockOptions) | **Get** /quote/option/{stock}/list | getStockOptions
[**PlaceOptionOrder**](OptionsAPI.md#PlaceOptionOrder) | **Post** /v2/option/placeOrder/{account_id} | placeOptionOrder
[**ReplaceOptionOrder**](OptionsAPI.md#ReplaceOptionOrder) | **Post** /v2/option/replaceOrder/{account_id} | replaceOptionOrder



## GetOptionQuotes

> GetOptionQuotesRequest GetOptionQuotes(ctx).Did(did).AccessToken(accessToken).TickerId(tickerId).DerivativeIds(derivativeIds).Execute()

getOptionQuotes

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
    tickerId := int32(56) // int32 | tickerId (default to 913243251)
    derivativeIds := float32(8.14) // float32 | derivativeIds (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptionsAPI.GetOptionQuotes(context.Background()).Did(did).AccessToken(accessToken).TickerId(tickerId).DerivativeIds(derivativeIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetOptionQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOptionQuotes`: GetOptionQuotesRequest
    fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetOptionQuotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOptionQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 
 **tickerId** | **int32** | tickerId | [default to 913243251]
 **derivativeIds** | **float32** | derivativeIds | 

### Return type

[**GetOptionQuotesRequest**](GetOptionQuotesRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStockOptions

> GetStockOptionsResponse GetStockOptions(ctx, stock).Did(did).AccessToken(accessToken).Count(count).IncludeWeekly(includeWeekly).Direction(direction).ExpireDate(expireDate).QueryAll(queryAll).Execute()

getStockOptions

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
    count := int32(56) // int32 |  (default to -1)
    includeWeekly := int32(56) // int32 |  (optional) (default to 1)
    direction := openapiclient.OptionDirection("all") // OptionDirection |  (optional)
    expireDate := "expireDate_example" // string |  (optional) (default to "12/16/2022")
    queryAll := int32(56) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptionsAPI.GetStockOptions(context.Background(), stock).Did(did).AccessToken(accessToken).Count(count).IncludeWeekly(includeWeekly).Direction(direction).ExpireDate(expireDate).QueryAll(queryAll).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetStockOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStockOptions`: GetStockOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetStockOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stock** | **int32** | Internal stock ticker ID | [default to 913243251]

### Other Parameters

Other parameters are passed through a pointer to a apiGetStockOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 

 **count** | **int32** |  | [default to -1]
 **includeWeekly** | **int32** |  | [default to 1]
 **direction** | [**OptionDirection**](OptionDirection.md) |  | 
 **expireDate** | **string** |  | [default to &quot;12/16/2022&quot;]
 **queryAll** | **int32** |  | [default to 0]

### Return type

[**GetStockOptionsResponse**](GetStockOptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaceOptionOrder

> map[string]interface{} PlaceOptionOrder(ctx, accountId).Did(did).AccessToken(accessToken).TToken(tToken).TTime(tTime).PostOptionOrderRequest(postOptionOrderRequest).Execute()

placeOptionOrder



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
    tToken := "tToken_example" // string | Trade token
    tTime := "tTime_example" // string | Time
    accountId := "accountId_example" // string | account_id
    postOptionOrderRequest := *openapiclient.NewPostOptionOrderRequest() // PostOptionOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptionsAPI.PlaceOptionOrder(context.Background(), accountId).Did(did).AccessToken(accessToken).TToken(tToken).TTime(tTime).PostOptionOrderRequest(postOptionOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.PlaceOptionOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlaceOptionOrder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.PlaceOptionOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlaceOptionOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 
 **tToken** | **string** | Trade token | 
 **tTime** | **string** | Time | 

 **postOptionOrderRequest** | [**PostOptionOrderRequest**](PostOptionOrderRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOptionOrder

> map[string]interface{} ReplaceOptionOrder(ctx, accountId).Did(did).AccessToken(accessToken).TToken(tToken).TTime(tTime).ReplaceOptionOrderRequest(replaceOptionOrderRequest).Execute()

replaceOptionOrder

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
    tToken := "tToken_example" // string | Trade token
    tTime := "tTime_example" // string | Time
    accountId := "accountId_example" // string | account_id
    replaceOptionOrderRequest := *openapiclient.NewReplaceOptionOrderRequest() // ReplaceOptionOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptionsAPI.ReplaceOptionOrder(context.Background(), accountId).Did(did).AccessToken(accessToken).TToken(tToken).TTime(tTime).ReplaceOptionOrderRequest(replaceOptionOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.ReplaceOptionOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceOptionOrder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.ReplaceOptionOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOptionOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 
 **tToken** | **string** | Trade token | 
 **tTime** | **string** | Time | 

 **replaceOptionOrderRequest** | [**ReplaceOptionOrderRequest**](ReplaceOptionOrderRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

