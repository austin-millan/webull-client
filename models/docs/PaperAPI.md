# \PaperAPI

All URIs are relative to *https://quoteapi.webull.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPaperTradeOrder**](PaperAPI.md#CancelPaperTradeOrder) | **Post** /paper/1/acc/{paper_account_id}/orderop/cancel/{order_id} | cancelPaperTradeOrder
[**GetPaperOrders**](PaperAPI.md#GetPaperOrders) | **Get** /paper/1/acc/{paper_account_id}/order | getPaperOrders
[**GetPaperTradingAccountID**](PaperAPI.md#GetPaperTradingAccountID) | **Get** /myaccounts/true | getPaperTradingAccountID
[**ModifyPaperTradeOrder**](PaperAPI.md#ModifyPaperTradeOrder) | **Post** /paper/1/acc/{paper_account_id}/orderop/modify/{order_id} | modifyPaperTradeOrder
[**PlacePaperTradeOrder**](PaperAPI.md#PlacePaperTradeOrder) | **Post** /paper/1/acc/{paper_account_id}/orderop/place/{stock} | placePaperTradeOrder



## CancelPaperTradeOrder

> map[string]interface{} CancelPaperTradeOrder(ctx, paperAccountId, orderId).Did(did).AccessToken(accessToken).Execute()

cancelPaperTradeOrder



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
    paperAccountId := "paperAccountId_example" // string | 
    orderId := "orderId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaperAPI.CancelPaperTradeOrder(context.Background(), paperAccountId, orderId).Did(did).AccessToken(accessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaperAPI.CancelPaperTradeOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelPaperTradeOrder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PaperAPI.CancelPaperTradeOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paperAccountId** | **string** |  | 
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPaperTradeOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 



### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaperOrders

> []PaperOrder GetPaperOrders(ctx, paperAccountId).Did(did).AccessToken(accessToken).DateType(dateType).StartTime(startTime).PageSize(pageSize).Status(status).Execute()

getPaperOrders



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
    paperAccountId := int32(56) // int32 | 
    dateType := "dateType_example" // string |  (default to "ORDER")
    startTime := "startTime_example" // string |  (optional) (default to "1970-0-1")
    pageSize := float32(8.14) // float32 |  (optional) (default to 256)
    status := openapiclient.OrderStatus("Queued") // OrderStatus |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaperAPI.GetPaperOrders(context.Background(), paperAccountId).Did(did).AccessToken(accessToken).DateType(dateType).StartTime(startTime).PageSize(pageSize).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaperAPI.GetPaperOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaperOrders`: []PaperOrder
    fmt.Fprintf(os.Stdout, "Response from `PaperAPI.GetPaperOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paperAccountId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaperOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 

 **dateType** | **string** |  | [default to &quot;ORDER&quot;]
 **startTime** | **string** |  | [default to &quot;1970-0-1&quot;]
 **pageSize** | **float32** |  | [default to 256]
 **status** | [**OrderStatus**](OrderStatus.md) |  | 

### Return type

[**[]PaperOrder**](PaperOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaperTradingAccountID

> []PaperAccount GetPaperTradingAccountID(ctx).Did(did).AccessToken(accessToken).Execute()

getPaperTradingAccountID



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaperAPI.GetPaperTradingAccountID(context.Background()).Did(did).AccessToken(accessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaperAPI.GetPaperTradingAccountID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaperTradingAccountID`: []PaperAccount
    fmt.Fprintf(os.Stdout, "Response from `PaperAPI.GetPaperTradingAccountID`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaperTradingAccountIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 

### Return type

[**[]PaperAccount**](PaperAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyPaperTradeOrder

> map[string]interface{} ModifyPaperTradeOrder(ctx, paperAccountId, orderId).Did(did).AccessToken(accessToken).PostStockOrderRequest(postStockOrderRequest).Execute()

modifyPaperTradeOrder



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
    paperAccountId := int32(56) // int32 | 
    orderId := "orderId_example" // string | 
    postStockOrderRequest := *openapiclient.NewPostStockOrderRequest() // PostStockOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaperAPI.ModifyPaperTradeOrder(context.Background(), paperAccountId, orderId).Did(did).AccessToken(accessToken).PostStockOrderRequest(postStockOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaperAPI.ModifyPaperTradeOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyPaperTradeOrder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PaperAPI.ModifyPaperTradeOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paperAccountId** | **int32** |  | 
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPaperTradeOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 


 **postStockOrderRequest** | [**PostStockOrderRequest**](PostStockOrderRequest.md) |  | 

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


## PlacePaperTradeOrder

> PostPaperOrderResponse PlacePaperTradeOrder(ctx, paperAccountId, stock).Did(did).AccessToken(accessToken).PostStockOrderRequest(postStockOrderRequest).Execute()

placePaperTradeOrder



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
    paperAccountId := int32(56) // int32 | 
    stock := int32(56) // int32 | Internal stock ticker ID (default to 913243251)
    postStockOrderRequest := *openapiclient.NewPostStockOrderRequest() // PostStockOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaperAPI.PlacePaperTradeOrder(context.Background(), paperAccountId, stock).Did(did).AccessToken(accessToken).PostStockOrderRequest(postStockOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaperAPI.PlacePaperTradeOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlacePaperTradeOrder`: PostPaperOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `PaperAPI.PlacePaperTradeOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paperAccountId** | **int32** |  | 
**stock** | **int32** | Internal stock ticker ID | [default to 913243251]

### Other Parameters

Other parameters are passed through a pointer to a apiPlacePaperTradeOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 


 **postStockOrderRequest** | [**PostStockOrderRequest**](PostStockOrderRequest.md) |  | 

### Return type

[**PostPaperOrderResponse**](PostPaperOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

