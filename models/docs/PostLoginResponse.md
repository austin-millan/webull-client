# PostLoginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**ExtInfo** | Pointer to [**PostLoginResponseExtInfo**](PostLoginResponseExtInfo.md) |  | [optional] 
**FirstTimeOfThird** | Pointer to **bool** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**RegisterAddress** | Pointer to **int32** |  | [optional] 
**Settings** | Pointer to [**PostLoginResponseSettings**](PostLoginResponseSettings.md) |  | [optional] 
**TokenExpireTime** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewPostLoginResponse

`func NewPostLoginResponse() *PostLoginResponse`

NewPostLoginResponse instantiates a new PostLoginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLoginResponseWithDefaults

`func NewPostLoginResponseWithDefaults() *PostLoginResponse`

NewPostLoginResponseWithDefaults instantiates a new PostLoginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *PostLoginResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *PostLoginResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *PostLoginResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *PostLoginResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetExtInfo

`func (o *PostLoginResponse) GetExtInfo() PostLoginResponseExtInfo`

GetExtInfo returns the ExtInfo field if non-nil, zero value otherwise.

### GetExtInfoOk

`func (o *PostLoginResponse) GetExtInfoOk() (*PostLoginResponseExtInfo, bool)`

GetExtInfoOk returns a tuple with the ExtInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtInfo

`func (o *PostLoginResponse) SetExtInfo(v PostLoginResponseExtInfo)`

SetExtInfo sets ExtInfo field to given value.

### HasExtInfo

`func (o *PostLoginResponse) HasExtInfo() bool`

HasExtInfo returns a boolean if a field has been set.

### GetFirstTimeOfThird

`func (o *PostLoginResponse) GetFirstTimeOfThird() bool`

GetFirstTimeOfThird returns the FirstTimeOfThird field if non-nil, zero value otherwise.

### GetFirstTimeOfThirdOk

`func (o *PostLoginResponse) GetFirstTimeOfThirdOk() (*bool, bool)`

GetFirstTimeOfThirdOk returns a tuple with the FirstTimeOfThird field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTimeOfThird

`func (o *PostLoginResponse) SetFirstTimeOfThird(v bool)`

SetFirstTimeOfThird sets FirstTimeOfThird field to given value.

### HasFirstTimeOfThird

`func (o *PostLoginResponse) HasFirstTimeOfThird() bool`

HasFirstTimeOfThird returns a boolean if a field has been set.

### GetRefreshToken

`func (o *PostLoginResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *PostLoginResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *PostLoginResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *PostLoginResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetRegisterAddress

`func (o *PostLoginResponse) GetRegisterAddress() int32`

GetRegisterAddress returns the RegisterAddress field if non-nil, zero value otherwise.

### GetRegisterAddressOk

`func (o *PostLoginResponse) GetRegisterAddressOk() (*int32, bool)`

GetRegisterAddressOk returns a tuple with the RegisterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterAddress

`func (o *PostLoginResponse) SetRegisterAddress(v int32)`

SetRegisterAddress sets RegisterAddress field to given value.

### HasRegisterAddress

`func (o *PostLoginResponse) HasRegisterAddress() bool`

HasRegisterAddress returns a boolean if a field has been set.

### GetSettings

`func (o *PostLoginResponse) GetSettings() PostLoginResponseSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PostLoginResponse) GetSettingsOk() (*PostLoginResponseSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PostLoginResponse) SetSettings(v PostLoginResponseSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PostLoginResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTokenExpireTime

`func (o *PostLoginResponse) GetTokenExpireTime() string`

GetTokenExpireTime returns the TokenExpireTime field if non-nil, zero value otherwise.

### GetTokenExpireTimeOk

`func (o *PostLoginResponse) GetTokenExpireTimeOk() (*string, bool)`

GetTokenExpireTimeOk returns a tuple with the TokenExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpireTime

`func (o *PostLoginResponse) SetTokenExpireTime(v string)`

SetTokenExpireTime sets TokenExpireTime field to given value.

### HasTokenExpireTime

`func (o *PostLoginResponse) HasTokenExpireTime() bool`

HasTokenExpireTime returns a boolean if a field has been set.

### GetUuid

`func (o *PostLoginResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PostLoginResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PostLoginResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PostLoginResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


