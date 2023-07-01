# PostRefreshTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**TokenExpireTime** | Pointer to **string** |  | [optional] 

## Methods

### NewPostRefreshTokenResponse

`func NewPostRefreshTokenResponse() *PostRefreshTokenResponse`

NewPostRefreshTokenResponse instantiates a new PostRefreshTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRefreshTokenResponseWithDefaults

`func NewPostRefreshTokenResponseWithDefaults() *PostRefreshTokenResponse`

NewPostRefreshTokenResponseWithDefaults instantiates a new PostRefreshTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *PostRefreshTokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *PostRefreshTokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *PostRefreshTokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *PostRefreshTokenResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *PostRefreshTokenResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *PostRefreshTokenResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *PostRefreshTokenResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *PostRefreshTokenResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetTokenExpireTime

`func (o *PostRefreshTokenResponse) GetTokenExpireTime() string`

GetTokenExpireTime returns the TokenExpireTime field if non-nil, zero value otherwise.

### GetTokenExpireTimeOk

`func (o *PostRefreshTokenResponse) GetTokenExpireTimeOk() (*string, bool)`

GetTokenExpireTimeOk returns a tuple with the TokenExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpireTime

`func (o *PostRefreshTokenResponse) SetTokenExpireTime(v string)`

SetTokenExpireTime sets TokenExpireTime field to given value.

### HasTokenExpireTime

`func (o *PostRefreshTokenResponse) HasTokenExpireTime() bool`

HasTokenExpireTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


