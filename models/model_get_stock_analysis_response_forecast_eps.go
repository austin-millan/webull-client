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

// checks if the GetStockAnalysisResponseForecastEps type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStockAnalysisResponseForecastEps{}

// GetStockAnalysisResponseForecastEps struct for GetStockAnalysisResponseForecastEps
type GetStockAnalysisResponseForecastEps struct {
	CurrencyName *string `json:"currencyName,omitempty"`
	Id *string `json:"id,omitempty"`
	Points []GetStockAnalysisResponseForecastEpsPointsInner `json:"points,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewGetStockAnalysisResponseForecastEps instantiates a new GetStockAnalysisResponseForecastEps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStockAnalysisResponseForecastEps() *GetStockAnalysisResponseForecastEps {
	this := GetStockAnalysisResponseForecastEps{}
	return &this
}

// NewGetStockAnalysisResponseForecastEpsWithDefaults instantiates a new GetStockAnalysisResponseForecastEps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStockAnalysisResponseForecastEpsWithDefaults() *GetStockAnalysisResponseForecastEps {
	this := GetStockAnalysisResponseForecastEps{}
	return &this
}

// GetCurrencyName returns the CurrencyName field value if set, zero value otherwise.
func (o *GetStockAnalysisResponseForecastEps) GetCurrencyName() string {
	if o == nil || IsNil(o.CurrencyName) {
		var ret string
		return ret
	}
	return *o.CurrencyName
}

// GetCurrencyNameOk returns a tuple with the CurrencyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStockAnalysisResponseForecastEps) GetCurrencyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyName) {
		return nil, false
	}
	return o.CurrencyName, true
}

// HasCurrencyName returns a boolean if a field has been set.
func (o *GetStockAnalysisResponseForecastEps) HasCurrencyName() bool {
	if o != nil && !IsNil(o.CurrencyName) {
		return true
	}

	return false
}

// SetCurrencyName gets a reference to the given string and assigns it to the CurrencyName field.
func (o *GetStockAnalysisResponseForecastEps) SetCurrencyName(v string) {
	o.CurrencyName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetStockAnalysisResponseForecastEps) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStockAnalysisResponseForecastEps) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetStockAnalysisResponseForecastEps) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetStockAnalysisResponseForecastEps) SetId(v string) {
	o.Id = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *GetStockAnalysisResponseForecastEps) GetPoints() []GetStockAnalysisResponseForecastEpsPointsInner {
	if o == nil || IsNil(o.Points) {
		var ret []GetStockAnalysisResponseForecastEpsPointsInner
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStockAnalysisResponseForecastEps) GetPointsOk() ([]GetStockAnalysisResponseForecastEpsPointsInner, bool) {
	if o == nil || IsNil(o.Points) {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *GetStockAnalysisResponseForecastEps) HasPoints() bool {
	if o != nil && !IsNil(o.Points) {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []GetStockAnalysisResponseForecastEpsPointsInner and assigns it to the Points field.
func (o *GetStockAnalysisResponseForecastEps) SetPoints(v []GetStockAnalysisResponseForecastEpsPointsInner) {
	o.Points = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetStockAnalysisResponseForecastEps) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStockAnalysisResponseForecastEps) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetStockAnalysisResponseForecastEps) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetStockAnalysisResponseForecastEps) SetTitle(v string) {
	o.Title = &v
}

func (o GetStockAnalysisResponseForecastEps) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStockAnalysisResponseForecastEps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrencyName) {
		toSerialize["currencyName"] = o.CurrencyName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Points) {
		toSerialize["points"] = o.Points
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableGetStockAnalysisResponseForecastEps struct {
	value *GetStockAnalysisResponseForecastEps
	isSet bool
}

func (v NullableGetStockAnalysisResponseForecastEps) Get() *GetStockAnalysisResponseForecastEps {
	return v.value
}

func (v *NullableGetStockAnalysisResponseForecastEps) Set(val *GetStockAnalysisResponseForecastEps) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStockAnalysisResponseForecastEps) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStockAnalysisResponseForecastEps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStockAnalysisResponseForecastEps(val *GetStockAnalysisResponseForecastEps) *NullableGetStockAnalysisResponseForecastEps {
	return &NullableGetStockAnalysisResponseForecastEps{value: val, isSet: true}
}

func (v NullableGetStockAnalysisResponseForecastEps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStockAnalysisResponseForecastEps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


