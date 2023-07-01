# \UserAPI

All URIs are relative to *https://quoteapi.webull.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUser**](UserAPI.md#GetUser) | **Get** /user | getUser



## GetUser

> GetUserDetailsResponse GetUser(ctx).Did(did).AccessToken(accessToken).Execute()

getUser



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
    resp, r, err := apiClient.UserAPI.GetUser(context.Background()).Did(did).AccessToken(accessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: GetUserDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** | Device ID | [default to &quot;your_device&quot;]
 **accessToken** | **string** | Access token | 

### Return type

[**GetUserDetailsResponse**](GetUserDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

