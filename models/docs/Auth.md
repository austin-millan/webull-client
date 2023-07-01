# Auth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MFA** | Pointer to **string** |  | [optional] 
**AccessToken** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**TokenExpireTime** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewAuth

`func NewAuth() *Auth`

NewAuth instantiates a new Auth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthWithDefaults

`func NewAuthWithDefaults() *Auth`

NewAuthWithDefaults instantiates a new Auth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMFA

`func (o *Auth) GetMFA() string`

GetMFA returns the MFA field if non-nil, zero value otherwise.

### GetMFAOk

`func (o *Auth) GetMFAOk() (*string, bool)`

GetMFAOk returns a tuple with the MFA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMFA

`func (o *Auth) SetMFA(v string)`

SetMFA sets MFA field to given value.

### HasMFA

`func (o *Auth) HasMFA() bool`

HasMFA returns a boolean if a field has been set.

### GetAccessToken

`func (o *Auth) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Auth) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Auth) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *Auth) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetDeviceId

`func (o *Auth) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *Auth) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *Auth) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *Auth) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetRefreshToken

`func (o *Auth) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *Auth) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *Auth) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *Auth) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetTokenExpireTime

`func (o *Auth) GetTokenExpireTime() string`

GetTokenExpireTime returns the TokenExpireTime field if non-nil, zero value otherwise.

### GetTokenExpireTimeOk

`func (o *Auth) GetTokenExpireTimeOk() (*string, bool)`

GetTokenExpireTimeOk returns a tuple with the TokenExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpireTime

`func (o *Auth) SetTokenExpireTime(v string)`

SetTokenExpireTime sets TokenExpireTime field to given value.

### HasTokenExpireTime

`func (o *Auth) HasTokenExpireTime() bool`

HasTokenExpireTime returns a boolean if a field has been set.

### GetUsername

`func (o *Auth) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Auth) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Auth) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Auth) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUuid

`func (o *Auth) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Auth) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Auth) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Auth) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


