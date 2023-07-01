# \OrderAPI

All URIs are relative to *https://quoteapi.webull.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](OrderAPI.md#CancelOrder) | **Post** /order/{account_id}/cancelStockOrder/ | cancelOrder
[**CancelOtocoOrder**](OrderAPI.md#CancelOtocoOrder) | **Post** /v2/corder/stock/modify/{account_id} | cancelOtocoOrder
[**CheckOtocoOrder**](OrderAPI.md#CheckOtocoOrder) | **Post** /v2/corder/stock/check/{account_id} | checkOtocoOrder
[**GetOrders**](OrderAPI.md#GetOrders) | **Get** /v2/option/list | getOrders
[**IsTradeable**](OrderAPI.md#IsTradeable) | **Get** /ticker/broker/permissionV2 | isTradeable
[**ModifyOrder**](OrderAPI.md#ModifyOrder) | **Post** /order/{account_id}/modifyStockOrder/{order_id} | modifyOrder
[**PlaceOrder**](OrderAPI.md#PlaceOrder) | **Post** /order/{account_id}/placeStockOrder | placeOrder
[**PlaceOtocoOrder**](OrderAPI.md#PlaceOtocoOrder) | **Post** /v2/corder/stock/place/{account_id} | placeOtocoOrder



## CancelOrder

> []map[string]interface{} CancelOrder(ctx, accountId).Did(did).AccessToken(accessToken).TToken(tToken).TTime(tTime).Execute()

cancelOrder



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
    accountId := "accountId_example" // string | Account ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.CancelOrder(context.Background(), accountId).Did(did).AccessToken(accessToken).TToken(tToken).TTime(tTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CancelOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelOrder`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CancelOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 
 **tToken** | **string** | Trade token | 
 **tTime** | **string** | Time | 


### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelOtocoOrder

> []map[string]interface{} CancelOtocoOrder(ctx, accountId).Did(did).AccessToken(accessToken).TToken(tToken).TTime(tTime).CancelOtocoOrderRequest(cancelOtocoOrderRequest).Execute()

cancelOtocoOrder



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
    accountId := "accountId_example" // string | Account ID
    cancelOtocoOrderRequest := *openapiclient.NewCancelOtocoOrderRequest() // CancelOtocoOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.CancelOtocoOrder(context.Background(), accountId).Did(did).AccessToken(accessToken).TToken(tToken).TTime(tTime).CancelOtocoOrderRequest(cancelOtocoOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CancelOtocoOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelOtocoOrder`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CancelOtocoOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOtocoOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 
 **tToken** | **string** | Trade token | 
 **tTime** | **string** | Time | 

 **cancelOtocoOrderRequest** | [**CancelOtocoOrderRequest**](CancelOtocoOrderRequest.md) |  | 

### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckOtocoOrder

> []map[string]interface{} CheckOtocoOrder(ctx, accountId).Did(did).AccessToken(accessToken).PostOtocoOrderRequest(postOtocoOrderRequest).Execute()

checkOtocoOrder



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
    postOtocoOrderRequest := *openapiclient.NewPostOtocoOrderRequest() // PostOtocoOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.CheckOtocoOrder(context.Background(), accountId).Did(did).AccessToken(accessToken).PostOtocoOrderRequest(postOtocoOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CheckOtocoOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckOtocoOrder`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CheckOtocoOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckOtocoOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 

 **postOtocoOrderRequest** | [**PostOtocoOrderRequest**](PostOtocoOrderRequest.md) |  | 

### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrders

> []GetOrdersResponse GetOrders(ctx).Did(did).AccessToken(accessToken).SecAccountId(secAccountId).TTime(tTime).TToken(tToken).StartTime(startTime).DateType(dateType).PageSize(pageSize).LastCreateTime(lastCreateTime).Status(status).Execute()

getOrders



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
    secAccountId := "secAccountId_example" // string | Numeric ID of the user to get
    tTime := "tTime_example" // string | Time (optional)
    tToken := "tToken_example" // string | Trade token (optional)
    startTime := "startTime_example" // string | Start date for orders (optional)
    dateType := "dateType_example" // string | Order type (optional) (default to "ORDER")
    pageSize := int32(56) // int32 | Page size (optional) (default to 256)
    lastCreateTime := "lastCreateTime_example" // string | Last create time (optional)
    status := openapiclient.OrderStatus("Queued") // OrderStatus | Status of order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.GetOrders(context.Background()).Did(did).AccessToken(accessToken).SecAccountId(secAccountId).TTime(tTime).TToken(tToken).StartTime(startTime).DateType(dateType).PageSize(pageSize).LastCreateTime(lastCreateTime).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrders`: []GetOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 
 **secAccountId** | **string** | Numeric ID of the user to get | 
 **tTime** | **string** | Time | 
 **tToken** | **string** | Trade token | 
 **startTime** | **string** | Start date for orders | 
 **dateType** | **string** | Order type | [default to &quot;ORDER&quot;]
 **pageSize** | **int32** | Page size | [default to 256]
 **lastCreateTime** | **string** | Last create time | 
 **status** | [**OrderStatus**](OrderStatus.md) | Status of order | 

### Return type

[**[]GetOrdersResponse**](GetOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsTradeable

> []GetIsTradeableResponse IsTradeable(ctx).Did(did).AccessToken(accessToken).TickerId(tickerId).Execute()

isTradeable



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
    tickerId := int32(56) // int32 | Ticker ID is a Stock to query for (default to 913243251)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.IsTradeable(context.Background()).Did(did).AccessToken(accessToken).TickerId(tickerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.IsTradeable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IsTradeable`: []GetIsTradeableResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.IsTradeable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIsTradeableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 
 **tickerId** | **int32** | Ticker ID is a Stock to query for | [default to 913243251]

### Return type

[**[]GetIsTradeableResponse**](GetIsTradeableResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyOrder

> map[string]interface{} ModifyOrder(ctx, accountId, orderId).Did(did).AccessToken(accessToken).TToken(tToken).TTime(tTime).PostStockOrderRequest(postStockOrderRequest).Execute()

modifyOrder

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
    orderId := "orderId_example" // string | order_id
    postStockOrderRequest := *openapiclient.NewPostStockOrderRequest() // PostStockOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.ModifyOrder(context.Background(), accountId, orderId).Did(did).AccessToken(accessToken).TToken(tToken).TTime(tTime).PostStockOrderRequest(postStockOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ModifyOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyOrder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ModifyOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account_id | 
**orderId** | **string** | order_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 
 **tToken** | **string** | Trade token | 
 **tTime** | **string** | Time | 


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


## PlaceOrder

> PostOrderResponse PlaceOrder(ctx, accountId).Did(did).AccessToken(accessToken).TToken(tToken).TTime(tTime).PostStockOrderRequest(postStockOrderRequest).Execute()

placeOrder

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
    postStockOrderRequest := *openapiclient.NewPostStockOrderRequest() // PostStockOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.PlaceOrder(context.Background(), accountId).Did(did).AccessToken(accessToken).TToken(tToken).TTime(tTime).PostStockOrderRequest(postStockOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.PlaceOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlaceOrder`: PostOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.PlaceOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlaceOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 
 **tToken** | **string** | Trade token | 
 **tTime** | **string** | Time | 

 **postStockOrderRequest** | [**PostStockOrderRequest**](PostStockOrderRequest.md) |  | 

### Return type

[**PostOrderResponse**](PostOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaceOtocoOrder

> map[string]interface{} PlaceOtocoOrder(ctx, accountId).Did(did).AccessToken(accessToken).TTime(tTime).PostOtocoOrderRequest(postOtocoOrderRequest).Execute()

placeOtocoOrder

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
    tTime := "tTime_example" // string | Time
    accountId := "accountId_example" // string | account_id
    postOtocoOrderRequest := *openapiclient.NewPostOtocoOrderRequest() // PostOtocoOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.PlaceOtocoOrder(context.Background(), accountId).Did(did).AccessToken(accessToken).TTime(tTime).PostOtocoOrderRequest(postOtocoOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.PlaceOtocoOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlaceOtocoOrder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.PlaceOtocoOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlaceOtocoOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 
 **tTime** | **string** | Time | 

 **postOtocoOrderRequest** | [**PostOtocoOrderRequest**](PostOtocoOrderRequest.md) |  | 

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

