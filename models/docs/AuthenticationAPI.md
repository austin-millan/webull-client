# \AuthenticationAPI

All URIs are relative to *https://quoteapi.webull.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMultiFactorAuth**](AuthenticationAPI.md#GetMultiFactorAuth) | **Post** /passport/verificationCode/sendCode | getMultiFactorAuth
[**GetTradeToken**](AuthenticationAPI.md#GetTradeToken) | **Post** /login | getTradeToken
[**Login**](AuthenticationAPI.md#Login) | **Post** /passport/login/v3/account | login
[**Logout**](AuthenticationAPI.md#Logout) | **Get** /passport/login/logout | logout
[**RefreshToken**](AuthenticationAPI.md#RefreshToken) | **Post** /passport/refreshToken | refreshToken



## GetMultiFactorAuth

> map[string]interface{} GetMultiFactorAuth(ctx).Did(did).Account(account).AccountType(accountType).DeviceId(deviceId).CodeType(codeType).RegionCode(regionCode).Execute()

getMultiFactorAuth

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
    account := "account_example" // string | account
    accountType := openapiclient.AccountType(1) // AccountType | phone = 1, email = 2 (default to 2)
    deviceId := "deviceId_example" // string | deviceId
    codeType := "codeType_example" // string | codeType (default to "5")
    regionCode := "regionCode_example" // string | regionCode (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationAPI.GetMultiFactorAuth(context.Background()).Did(did).Account(account).AccountType(accountType).DeviceId(deviceId).CodeType(codeType).RegionCode(regionCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetMultiFactorAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMultiFactorAuth`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetMultiFactorAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMultiFactorAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **account** | **string** | account | 
 **accountType** | [**AccountType**](AccountType.md) | phone &#x3D; 1, email &#x3D; 2 | [default to 2]
 **deviceId** | **string** | deviceId | 
 **codeType** | **string** | codeType | [default to &quot;5&quot;]
 **regionCode** | **string** | regionCode | [default to &quot;1&quot;]

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


## GetTradeToken

> []PostTradeTokenResponse GetTradeToken(ctx).Did(did).AccessToken(accessToken).PostLoginParametersRequest(postLoginParametersRequest).Execute()

getTradeToken



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
    postLoginParametersRequest := *openapiclient.NewPostLoginParametersRequest() // PostLoginParametersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationAPI.GetTradeToken(context.Background()).Did(did).AccessToken(accessToken).PostLoginParametersRequest(postLoginParametersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetTradeToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTradeToken`: []PostTradeTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetTradeToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 
 **postLoginParametersRequest** | [**PostLoginParametersRequest**](PostLoginParametersRequest.md) |  | 

### Return type

[**[]PostTradeTokenResponse**](PostTradeTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> PostLoginResponse Login(ctx).Did(did).PostLoginParametersRequest(postLoginParametersRequest).Execute()

login

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
    postLoginParametersRequest := *openapiclient.NewPostLoginParametersRequest() // PostLoginParametersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationAPI.Login(context.Background()).Did(did).PostLoginParametersRequest(postLoginParametersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Login``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Login`: PostLoginResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **postLoginParametersRequest** | [**PostLoginParametersRequest**](PostLoginParametersRequest.md) |  | 

### Return type

[**PostLoginResponse**](PostLoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> map[string]interface{} Logout(ctx).Did(did).AccessToken(accessToken).Execute()

logout



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
    resp, r, err := apiClient.AuthenticationAPI.Logout(context.Background()).Did(did).AccessToken(accessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Logout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Logout`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Logout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


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


## RefreshToken

> []PostRefreshTokenResponse RefreshToken(ctx).Did(did).AccessToken(accessToken).RefreshToken(refreshToken).Execute()

refreshToken



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
    refreshToken := "refreshToken_example" // string | Refresh token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationAPI.RefreshToken(context.Background()).Did(did).AccessToken(accessToken).RefreshToken(refreshToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.RefreshToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshToken`: []PostRefreshTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.RefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 
 **refreshToken** | **string** | Refresh token | 

### Return type

[**[]PostRefreshTokenResponse**](PostRefreshTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

