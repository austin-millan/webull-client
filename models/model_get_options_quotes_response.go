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

// checks if the GetOptionsQuotesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOptionsQuotesResponse{}

// GetOptionsQuotesResponse struct for GetOptionsQuotesResponse
type GetOptionsQuotesResponse struct {
	Data []map[string]interface{} `json:"data,omitempty"`
	DisExchangeCode *string `json:"disExchangeCode,omitempty"`
	DisSymbol *string `json:"disSymbol,omitempty"`
	Name *string `json:"name,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TickerId *int32 `json:"tickerId,omitempty"`
}

// NewGetOptionsQuotesResponse instantiates a new GetOptionsQuotesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOptionsQuotesResponse() *GetOptionsQuotesResponse {
	this := GetOptionsQuotesResponse{}
	return &this
}

// NewGetOptionsQuotesResponseWithDefaults instantiates a new GetOptionsQuotesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOptionsQuotesResponseWithDefaults() *GetOptionsQuotesResponse {
	this := GetOptionsQuotesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetOptionsQuotesResponse) GetData() []map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionsQuotesResponse) GetDataOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetOptionsQuotesResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []map[string]interface{} and assigns it to the Data field.
func (o *GetOptionsQuotesResponse) SetData(v []map[string]interface{}) {
	o.Data = v
}

// GetDisExchangeCode returns the DisExchangeCode field value if set, zero value otherwise.
func (o *GetOptionsQuotesResponse) GetDisExchangeCode() string {
	if o == nil || IsNil(o.DisExchangeCode) {
		var ret string
		return ret
	}
	return *o.DisExchangeCode
}

// GetDisExchangeCodeOk returns a tuple with the DisExchangeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionsQuotesResponse) GetDisExchangeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DisExchangeCode) {
		return nil, false
	}
	return o.DisExchangeCode, true
}

// HasDisExchangeCode returns a boolean if a field has been set.
func (o *GetOptionsQuotesResponse) HasDisExchangeCode() bool {
	if o != nil && !IsNil(o.DisExchangeCode) {
		return true
	}

	return false
}

// SetDisExchangeCode gets a reference to the given string and assigns it to the DisExchangeCode field.
func (o *GetOptionsQuotesResponse) SetDisExchangeCode(v string) {
	o.DisExchangeCode = &v
}

// GetDisSymbol returns the DisSymbol field value if set, zero value otherwise.
func (o *GetOptionsQuotesResponse) GetDisSymbol() string {
	if o == nil || IsNil(o.DisSymbol) {
		var ret string
		return ret
	}
	return *o.DisSymbol
}

// GetDisSymbolOk returns a tuple with the DisSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionsQuotesResponse) GetDisSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.DisSymbol) {
		return nil, false
	}
	return o.DisSymbol, true
}

// HasDisSymbol returns a boolean if a field has been set.
func (o *GetOptionsQuotesResponse) HasDisSymbol() bool {
	if o != nil && !IsNil(o.DisSymbol) {
		return true
	}

	return false
}

// SetDisSymbol gets a reference to the given string and assigns it to the DisSymbol field.
func (o *GetOptionsQuotesResponse) SetDisSymbol(v string) {
	o.DisSymbol = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOptionsQuotesResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionsQuotesResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOptionsQuotesResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOptionsQuotesResponse) SetName(v string) {
	o.Name = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetOptionsQuotesResponse) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionsQuotesResponse) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetOptionsQuotesResponse) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetOptionsQuotesResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTickerId returns the TickerId field value if set, zero value otherwise.
func (o *GetOptionsQuotesResponse) GetTickerId() int32 {
	if o == nil || IsNil(o.TickerId) {
		var ret int32
		return ret
	}
	return *o.TickerId
}

// GetTickerIdOk returns a tuple with the TickerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionsQuotesResponse) GetTickerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TickerId) {
		return nil, false
	}
	return o.TickerId, true
}

// HasTickerId returns a boolean if a field has been set.
func (o *GetOptionsQuotesResponse) HasTickerId() bool {
	if o != nil && !IsNil(o.TickerId) {
		return true
	}

	return false
}

// SetTickerId gets a reference to the given int32 and assigns it to the TickerId field.
func (o *GetOptionsQuotesResponse) SetTickerId(v int32) {
	o.TickerId = &v
}

func (o GetOptionsQuotesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOptionsQuotesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.DisExchangeCode) {
		toSerialize["disExchangeCode"] = o.DisExchangeCode
	}
	if !IsNil(o.DisSymbol) {
		toSerialize["disSymbol"] = o.DisSymbol
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.TickerId) {
		toSerialize["tickerId"] = o.TickerId
	}
	return toSerialize, nil
}

type NullableGetOptionsQuotesResponse struct {
	value *GetOptionsQuotesResponse
	isSet bool
}

func (v NullableGetOptionsQuotesResponse) Get() *GetOptionsQuotesResponse {
	return v.value
}

func (v *NullableGetOptionsQuotesResponse) Set(val *GetOptionsQuotesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOptionsQuotesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOptionsQuotesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOptionsQuotesResponse(val *GetOptionsQuotesResponse) *NullableGetOptionsQuotesResponse {
	return &NullableGetOptionsQuotesResponse{value: val, isSet: true}
}

func (v NullableGetOptionsQuotesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOptionsQuotesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


