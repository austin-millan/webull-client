# \StocksAPI

All URIs are relative to *https://quoteapi.webull.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActiveGainersLosers**](StocksAPI.md#GetActiveGainersLosers) | **Get** /securities/market/v5/card/stockActivityPc.{direction}/list | getActiveGainersLosers
[**GetFundamentals**](StocksAPI.md#GetFundamentals) | **Get** /securities/financial/index/{stock} | getFundamentals
[**GetStockAnalysis**](StocksAPI.md#GetStockAnalysis) | **Get** /securities/ticker/v5/analysis/{stock} | getStockAnalysis
[**GetStockId**](StocksAPI.md#GetStockId) | **Get** /search/tickers5 | getStockID
[**GetStockNews**](StocksAPI.md#GetStockNews) | **Get** /information/news/v5/tickerNews/{stock} | getStockNews
[**GetTickerChart**](StocksAPI.md#GetTickerChart) | **Get** /quote/tickerChartDatas/v5/{stock} | getTickerChart
[**Screener**](StocksAPI.md#Screener) | **Get** /wlas/screener/ng/query | screener



## GetActiveGainersLosers

> []ActiveGainersLosers GetActiveGainersLosers(ctx, direction).Did(did).AccessToken(accessToken).RegionId(regionId).UserRegionId(userRegionId).Execute()

getActiveGainersLosers



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
    direction := openapiclient.AdvancedDeclinedDirection("advanced") // AdvancedDeclinedDirection | Direction
    regionId := "regionId_example" // string | regionId (default to "6")
    userRegionId := "userRegionId_example" // string | userRegionId (default to "6")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StocksAPI.GetActiveGainersLosers(context.Background(), direction).Did(did).AccessToken(accessToken).RegionId(regionId).UserRegionId(userRegionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StocksAPI.GetActiveGainersLosers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveGainersLosers`: []ActiveGainersLosers
    fmt.Fprintf(os.Stdout, "Response from `StocksAPI.GetActiveGainersLosers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**direction** | [**AdvancedDeclinedDirection**](.md) | Direction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveGainersLosersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 

 **regionId** | **string** | regionId | [default to &quot;6&quot;]
 **userRegionId** | **string** | userRegionId | [default to &quot;6&quot;]

### Return type

[**[]ActiveGainersLosers**](ActiveGainersLosers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundamentals

> []GetFundamentalsResponse GetFundamentals(ctx, stock).Did(did).AccessToken(accessToken).Execute()

getFundamentals



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
    resp, r, err := apiClient.StocksAPI.GetFundamentals(context.Background(), stock).Did(did).AccessToken(accessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StocksAPI.GetFundamentals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundamentals`: []GetFundamentalsResponse
    fmt.Fprintf(os.Stdout, "Response from `StocksAPI.GetFundamentals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stock** | **int32** | Internal stock ticker ID | [default to 913243251]

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundamentalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 


### Return type

[**[]GetFundamentalsResponse**](GetFundamentalsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStockAnalysis

> GetStockAnalysisResponse GetStockAnalysis(ctx, stock).Did(did).AccessToken(accessToken).Execute()

getStockAnalysis



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
    resp, r, err := apiClient.StocksAPI.GetStockAnalysis(context.Background(), stock).Did(did).AccessToken(accessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StocksAPI.GetStockAnalysis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStockAnalysis`: GetStockAnalysisResponse
    fmt.Fprintf(os.Stdout, "Response from `StocksAPI.GetStockAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stock** | **int32** | Internal stock ticker ID | [default to 913243251]

### Other Parameters

Other parameters are passed through a pointer to a apiGetStockAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 


### Return type

[**GetStockAnalysisResponse**](GetStockAnalysisResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStockId

> LookupTickerResponse GetStockId(ctx).Did(did).AccessToken(accessToken).Keys(keys).QueryNumber(queryNumber).Execute()

getStockID



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
    keys := "keys_example" // string | Stock ticker (default to "913243251")
    queryNumber := "queryNumber_example" // string | queryNumber (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StocksAPI.GetStockId(context.Background()).Did(did).AccessToken(accessToken).Keys(keys).QueryNumber(queryNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StocksAPI.GetStockId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStockId`: LookupTickerResponse
    fmt.Fprintf(os.Stdout, "Response from `StocksAPI.GetStockId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStockIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 
 **keys** | **string** | Stock ticker | [default to &quot;913243251&quot;]
 **queryNumber** | **string** | queryNumber | [default to &quot;1&quot;]

### Return type

[**LookupTickerResponse**](LookupTickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStockNews

> GetNewsResponse GetStockNews(ctx, stock).Did(did).AccessToken(accessToken).CurrentNewsId(currentNewsId).PageSize(pageSize).Execute()

getStockNews

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
    stock := int32(56) // int32 | Stocks internal ticker ID (default to 913243251)
    currentNewsId := float32(8.14) // float32 | 0 is the latest article (optional) (default to 0)
    pageSize := int32(56) // int32 | Number of articles (optional) (default to 256)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StocksAPI.GetStockNews(context.Background(), stock).Did(did).AccessToken(accessToken).CurrentNewsId(currentNewsId).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StocksAPI.GetStockNews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStockNews`: GetNewsResponse
    fmt.Fprintf(os.Stdout, "Response from `StocksAPI.GetStockNews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stock** | **int32** | Stocks internal ticker ID | [default to 913243251]

### Other Parameters

Other parameters are passed through a pointer to a apiGetStockNewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 

 **currentNewsId** | **float32** | 0 is the latest article | [default to 0]
 **pageSize** | **int32** | Number of articles | [default to 256]

### Return type

[**GetNewsResponse**](GetNewsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTickerChart

> GetTickerChartResponse GetTickerChart(ctx, stock).Did(did).AccessToken(accessToken).Count(count).ExtendTrading(extendTrading).Type_(type_).Execute()

getTickerChart



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
    stock := int32(56) // int32 | Internal stock symbol ID (default to 913243251)
    count := int32(56) // int32 | Number of bars to return (default to -1)
    extendTrading := int32(56) // int32 | Whether to include pre-market and afterhours bars. '1' is used for pre-market and after-hours bars. (default to 1)
    type_ := "type__example" // string | X (default to "m1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StocksAPI.GetTickerChart(context.Background(), stock).Did(did).AccessToken(accessToken).Count(count).ExtendTrading(extendTrading).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StocksAPI.GetTickerChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTickerChart`: GetTickerChartResponse
    fmt.Fprintf(os.Stdout, "Response from `StocksAPI.GetTickerChart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stock** | **int32** | Internal stock symbol ID | [default to 913243251]

### Other Parameters

Other parameters are passed through a pointer to a apiGetTickerChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 

 **count** | **int32** | Number of bars to return | [default to -1]
 **extendTrading** | **int32** | Whether to include pre-market and afterhours bars. &#39;1&#39; is used for pre-market and after-hours bars. | [default to 1]
 **type_** | **string** | X | [default to &quot;m1&quot;]

### Return type

[**GetTickerChartResponse**](GetTickerChartResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Screener

> map[string]interface{} Screener(ctx).Did(did).AccessToken(accessToken).Fetch(fetch).Rules(rules).Sort(sort).Attach(attach).Execute()

screener



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
    fetch := int32(56) // int32 | fetch (default to 200)
    rules := []openapiclient.SmartRule{*openapiclient.NewSmartRule()} // []SmartRule | rules
    sort := map[string]interface{}{ ... } // map[string]interface{} | sort
    attach := map[string][]openapiclient.ScreenerAttachParameter{ ... } // ScreenerAttachParameter | attach

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StocksAPI.Screener(context.Background()).Did(did).AccessToken(accessToken).Fetch(fetch).Rules(rules).Sort(sort).Attach(attach).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StocksAPI.Screener``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Screener`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StocksAPI.Screener`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScreenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 
 **fetch** | **int32** | fetch | [default to 200]
 **rules** | [**[]SmartRule**](SmartRule.md) | rules | 
 **sort** | [**map[string]interface{}**](map[string]interface{}.md) | sort | 
 **attach** | [**ScreenerAttachParameter**](ScreenerAttachParameter.md) | attach | 

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

