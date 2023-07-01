# GetUserDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarUrl** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**NickName** | Pointer to **string** |  | [optional] 
**PwdFlag** | Pointer to **int32** |  | [optional] 
**RegionId** | Pointer to **int32** |  | [optional] 
**RegionName** | Pointer to **string** |  | [optional] 
**Sex** | Pointer to **int32** |  | [optional] 
**ThirdAccounts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TradeLockemail** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewGetUserDetailsResponse

`func NewGetUserDetailsResponse() *GetUserDetailsResponse`

NewGetUserDetailsResponse instantiates a new GetUserDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserDetailsResponseWithDefaults

`func NewGetUserDetailsResponseWithDefaults() *GetUserDetailsResponse`

NewGetUserDetailsResponseWithDefaults instantiates a new GetUserDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrl

`func (o *GetUserDetailsResponse) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *GetUserDetailsResponse) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *GetUserDetailsResponse) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *GetUserDetailsResponse) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetEmail

`func (o *GetUserDetailsResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetUserDetailsResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetUserDetailsResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetUserDetailsResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetNickName

`func (o *GetUserDetailsResponse) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *GetUserDetailsResponse) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *GetUserDetailsResponse) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *GetUserDetailsResponse) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### GetPwdFlag

`func (o *GetUserDetailsResponse) GetPwdFlag() int32`

GetPwdFlag returns the PwdFlag field if non-nil, zero value otherwise.

### GetPwdFlagOk

`func (o *GetUserDetailsResponse) GetPwdFlagOk() (*int32, bool)`

GetPwdFlagOk returns a tuple with the PwdFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdFlag

`func (o *GetUserDetailsResponse) SetPwdFlag(v int32)`

SetPwdFlag sets PwdFlag field to given value.

### HasPwdFlag

`func (o *GetUserDetailsResponse) HasPwdFlag() bool`

HasPwdFlag returns a boolean if a field has been set.

### GetRegionId

`func (o *GetUserDetailsResponse) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *GetUserDetailsResponse) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *GetUserDetailsResponse) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *GetUserDetailsResponse) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetRegionName

`func (o *GetUserDetailsResponse) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *GetUserDetailsResponse) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *GetUserDetailsResponse) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *GetUserDetailsResponse) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetSex

`func (o *GetUserDetailsResponse) GetSex() int32`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *GetUserDetailsResponse) GetSexOk() (*int32, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *GetUserDetailsResponse) SetSex(v int32)`

SetSex sets Sex field to given value.

### HasSex

`func (o *GetUserDetailsResponse) HasSex() bool`

HasSex returns a boolean if a field has been set.

### GetThirdAccounts

`func (o *GetUserDetailsResponse) GetThirdAccounts() []map[string]interface{}`

GetThirdAccounts returns the ThirdAccounts field if non-nil, zero value otherwise.

### GetThirdAccountsOk

`func (o *GetUserDetailsResponse) GetThirdAccountsOk() (*[]map[string]interface{}, bool)`

GetThirdAccountsOk returns a tuple with the ThirdAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdAccounts

`func (o *GetUserDetailsResponse) SetThirdAccounts(v []map[string]interface{})`

SetThirdAccounts sets ThirdAccounts field to given value.

### HasThirdAccounts

`func (o *GetUserDetailsResponse) HasThirdAccounts() bool`

HasThirdAccounts returns a boolean if a field has been set.

### GetTradeLockemail

`func (o *GetUserDetailsResponse) GetTradeLockemail() string`

GetTradeLockemail returns the TradeLockemail field if non-nil, zero value otherwise.

### GetTradeLockemailOk

`func (o *GetUserDetailsResponse) GetTradeLockemailOk() (*string, bool)`

GetTradeLockemailOk returns a tuple with the TradeLockemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeLockemail

`func (o *GetUserDetailsResponse) SetTradeLockemail(v string)`

SetTradeLockemail sets TradeLockemail field to given value.

### HasTradeLockemail

`func (o *GetUserDetailsResponse) HasTradeLockemail() bool`

HasTradeLockemail returns a boolean if a field has been set.

### GetUuid

`func (o *GetUserDetailsResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetUserDetailsResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetUserDetailsResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetUserDetailsResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


