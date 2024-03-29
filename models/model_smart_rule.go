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

// checks if the SmartRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartRule{}

// SmartRule struct for SmartRule
type SmartRule struct {
	Active *string `json:"active,omitempty"`
	AlertRuleKey *string `json:"alertRuleKey,omitempty"`
	Field *string `json:"field,omitempty"`
	Remark *string `json:"remark,omitempty"`
	// Example: 'earnPre', 'fastUp', 'fastDown', 'week52Up', 'week52Down', 'day5Down'
	Type *string `json:"type,omitempty"`
}

// NewSmartRule instantiates a new SmartRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartRule() *SmartRule {
	this := SmartRule{}
	return &this
}

// NewSmartRuleWithDefaults instantiates a new SmartRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartRuleWithDefaults() *SmartRule {
	this := SmartRule{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SmartRule) GetActive() string {
	if o == nil || IsNil(o.Active) {
		var ret string
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartRule) GetActiveOk() (*string, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SmartRule) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given string and assigns it to the Active field.
func (o *SmartRule) SetActive(v string) {
	o.Active = &v
}

// GetAlertRuleKey returns the AlertRuleKey field value if set, zero value otherwise.
func (o *SmartRule) GetAlertRuleKey() string {
	if o == nil || IsNil(o.AlertRuleKey) {
		var ret string
		return ret
	}
	return *o.AlertRuleKey
}

// GetAlertRuleKeyOk returns a tuple with the AlertRuleKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartRule) GetAlertRuleKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AlertRuleKey) {
		return nil, false
	}
	return o.AlertRuleKey, true
}

// HasAlertRuleKey returns a boolean if a field has been set.
func (o *SmartRule) HasAlertRuleKey() bool {
	if o != nil && !IsNil(o.AlertRuleKey) {
		return true
	}

	return false
}

// SetAlertRuleKey gets a reference to the given string and assigns it to the AlertRuleKey field.
func (o *SmartRule) SetAlertRuleKey(v string) {
	o.AlertRuleKey = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *SmartRule) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartRule) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *SmartRule) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *SmartRule) SetField(v string) {
	o.Field = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *SmartRule) GetRemark() string {
	if o == nil || IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartRule) GetRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *SmartRule) HasRemark() bool {
	if o != nil && !IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *SmartRule) SetRemark(v string) {
	o.Remark = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SmartRule) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartRule) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SmartRule) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SmartRule) SetType(v string) {
	o.Type = &v
}

func (o SmartRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.AlertRuleKey) {
		toSerialize["alertRuleKey"] = o.AlertRuleKey
	}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSmartRule struct {
	value *SmartRule
	isSet bool
}

func (v NullableSmartRule) Get() *SmartRule {
	return v.value
}

func (v *NullableSmartRule) Set(val *SmartRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartRule(val *SmartRule) *NullableSmartRule {
	return &NullableSmartRule{value: val, isSet: true}
}

func (v NullableSmartRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


