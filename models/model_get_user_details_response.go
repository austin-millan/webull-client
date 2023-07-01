/*
Webull API

Webull API Documentation

API version: 3.0.1
Contact: austin.millan@protonmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetUserDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserDetailsResponse{}

// GetUserDetailsResponse struct for GetUserDetailsResponse
type GetUserDetailsResponse struct {
	AvatarUrl *string `json:"avatarUrl,omitempty"`
	Email *string `json:"email,omitempty"`
	NickName *string `json:"nickName,omitempty"`
	PwdFlag *int32 `json:"pwdFlag,omitempty"`
	RegionId *int32 `json:"regionId,omitempty"`
	RegionName *string `json:"regionName,omitempty"`
	Sex *int32 `json:"sex,omitempty"`
	ThirdAccounts []map[string]interface{} `json:"thirdAccounts,omitempty"`
	TradeLockemail *string `json:"tradeLockemail,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

// NewGetUserDetailsResponse instantiates a new GetUserDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserDetailsResponse() *GetUserDetailsResponse {
	this := GetUserDetailsResponse{}
	return &this
}

// NewGetUserDetailsResponseWithDefaults instantiates a new GetUserDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserDetailsResponseWithDefaults() *GetUserDetailsResponse {
	this := GetUserDetailsResponse{}
	return &this
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *GetUserDetailsResponse) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserDetailsResponse) GetAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *GetUserDetailsResponse) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *GetUserDetailsResponse) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetUserDetailsResponse) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserDetailsResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetUserDetailsResponse) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetUserDetailsResponse) SetEmail(v string) {
	o.Email = &v
}

// GetNickName returns the NickName field value if set, zero value otherwise.
func (o *GetUserDetailsResponse) GetNickName() string {
	if o == nil || IsNil(o.NickName) {
		var ret string
		return ret
	}
	return *o.NickName
}

// GetNickNameOk returns a tuple with the NickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserDetailsResponse) GetNickNameOk() (*string, bool) {
	if o == nil || IsNil(o.NickName) {
		return nil, false
	}
	return o.NickName, true
}

// HasNickName returns a boolean if a field has been set.
func (o *GetUserDetailsResponse) HasNickName() bool {
	if o != nil && !IsNil(o.NickName) {
		return true
	}

	return false
}

// SetNickName gets a reference to the given string and assigns it to the NickName field.
func (o *GetUserDetailsResponse) SetNickName(v string) {
	o.NickName = &v
}

// GetPwdFlag returns the PwdFlag field value if set, zero value otherwise.
func (o *GetUserDetailsResponse) GetPwdFlag() int32 {
	if o == nil || IsNil(o.PwdFlag) {
		var ret int32
		return ret
	}
	return *o.PwdFlag
}

// GetPwdFlagOk returns a tuple with the PwdFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserDetailsResponse) GetPwdFlagOk() (*int32, bool) {
	if o == nil || IsNil(o.PwdFlag) {
		return nil, false
	}
	return o.PwdFlag, true
}

// HasPwdFlag returns a boolean if a field has been set.
func (o *GetUserDetailsResponse) HasPwdFlag() bool {
	if o != nil && !IsNil(o.PwdFlag) {
		return true
	}

	return false
}

// SetPwdFlag gets a reference to the given int32 and assigns it to the PwdFlag field.
func (o *GetUserDetailsResponse) SetPwdFlag(v int32) {
	o.PwdFlag = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *GetUserDetailsResponse) GetRegionId() int32 {
	if o == nil || IsNil(o.RegionId) {
		var ret int32
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserDetailsResponse) GetRegionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RegionId) {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *GetUserDetailsResponse) HasRegionId() bool {
	if o != nil && !IsNil(o.RegionId) {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given int32 and assigns it to the RegionId field.
func (o *GetUserDetailsResponse) SetRegionId(v int32) {
	o.RegionId = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *GetUserDetailsResponse) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserDetailsResponse) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *GetUserDetailsResponse) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *GetUserDetailsResponse) SetRegionName(v string) {
	o.RegionName = &v
}

// GetSex returns the Sex field value if set, zero value otherwise.
func (o *GetUserDetailsResponse) GetSex() int32 {
	if o == nil || IsNil(o.Sex) {
		var ret int32
		return ret
	}
	return *o.Sex
}

// GetSexOk returns a tuple with the Sex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserDetailsResponse) GetSexOk() (*int32, bool) {
	if o == nil || IsNil(o.Sex) {
		return nil, false
	}
	return o.Sex, true
}

// HasSex returns a boolean if a field has been set.
func (o *GetUserDetailsResponse) HasSex() bool {
	if o != nil && !IsNil(o.Sex) {
		return true
	}

	return false
}

// SetSex gets a reference to the given int32 and assigns it to the Sex field.
func (o *GetUserDetailsResponse) SetSex(v int32) {
	o.Sex = &v
}

// GetThirdAccounts returns the ThirdAccounts field value if set, zero value otherwise.
func (o *GetUserDetailsResponse) GetThirdAccounts() []map[string]interface{} {
	if o == nil || IsNil(o.ThirdAccounts) {
		var ret []map[string]interface{}
		return ret
	}
	return o.ThirdAccounts
}

// GetThirdAccountsOk returns a tuple with the ThirdAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserDetailsResponse) GetThirdAccountsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ThirdAccounts) {
		return nil, false
	}
	return o.ThirdAccounts, true
}

// HasThirdAccounts returns a boolean if a field has been set.
func (o *GetUserDetailsResponse) HasThirdAccounts() bool {
	if o != nil && !IsNil(o.ThirdAccounts) {
		return true
	}

	return false
}

// SetThirdAccounts gets a reference to the given []map[string]interface{} and assigns it to the ThirdAccounts field.
func (o *GetUserDetailsResponse) SetThirdAccounts(v []map[string]interface{}) {
	o.ThirdAccounts = v
}

// GetTradeLockemail returns the TradeLockemail field value if set, zero value otherwise.
func (o *GetUserDetailsResponse) GetTradeLockemail() string {
	if o == nil || IsNil(o.TradeLockemail) {
		var ret string
		return ret
	}
	return *o.TradeLockemail
}

// GetTradeLockemailOk returns a tuple with the TradeLockemail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserDetailsResponse) GetTradeLockemailOk() (*string, bool) {
	if o == nil || IsNil(o.TradeLockemail) {
		return nil, false
	}
	return o.TradeLockemail, true
}

// HasTradeLockemail returns a boolean if a field has been set.
func (o *GetUserDetailsResponse) HasTradeLockemail() bool {
	if o != nil && !IsNil(o.TradeLockemail) {
		return true
	}

	return false
}

// SetTradeLockemail gets a reference to the given string and assigns it to the TradeLockemail field.
func (o *GetUserDetailsResponse) SetTradeLockemail(v string) {
	o.TradeLockemail = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *GetUserDetailsResponse) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserDetailsResponse) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *GetUserDetailsResponse) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *GetUserDetailsResponse) SetUuid(v string) {
	o.Uuid = &v
}

func (o GetUserDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvatarUrl) {
		toSerialize["avatarUrl"] = o.AvatarUrl
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.NickName) {
		toSerialize["nickName"] = o.NickName
	}
	if !IsNil(o.PwdFlag) {
		toSerialize["pwdFlag"] = o.PwdFlag
	}
	if !IsNil(o.RegionId) {
		toSerialize["regionId"] = o.RegionId
	}
	if !IsNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	if !IsNil(o.Sex) {
		toSerialize["sex"] = o.Sex
	}
	if !IsNil(o.ThirdAccounts) {
		toSerialize["thirdAccounts"] = o.ThirdAccounts
	}
	if !IsNil(o.TradeLockemail) {
		toSerialize["tradeLockemail"] = o.TradeLockemail
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableGetUserDetailsResponse struct {
	value *GetUserDetailsResponse
	isSet bool
}

func (v NullableGetUserDetailsResponse) Get() *GetUserDetailsResponse {
	return v.value
}

func (v *NullableGetUserDetailsResponse) Set(val *GetUserDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserDetailsResponse(val *GetUserDetailsResponse) *NullableGetUserDetailsResponse {
	return &NullableGetUserDetailsResponse{value: val, isSet: true}
}

func (v NullableGetUserDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


