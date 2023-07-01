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

// checks if the GetTickerChartResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTickerChartResponse{}

// GetTickerChartResponse struct for GetTickerChartResponse
type GetTickerChartResponse struct {
	Data []string `json:"data,omitempty"`
	Dates []GetTickerChartResponseDatesInner `json:"dates,omitempty"`
	ExchangeStatus *bool `json:"exchangeStatus,omitempty"`
	HasMore *int32 `json:"hasMore,omitempty"`
	PreClose *string `json:"preClose,omitempty"`
	TickerId *int32 `json:"tickerId,omitempty"`
	TimeZone *string `json:"timeZone,omitempty"`
}

// NewGetTickerChartResponse instantiates a new GetTickerChartResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTickerChartResponse() *GetTickerChartResponse {
	this := GetTickerChartResponse{}
	return &this
}

// NewGetTickerChartResponseWithDefaults instantiates a new GetTickerChartResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTickerChartResponseWithDefaults() *GetTickerChartResponse {
	this := GetTickerChartResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetTickerChartResponse) GetData() []string {
	if o == nil || IsNil(o.Data) {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTickerChartResponse) GetDataOk() ([]string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetTickerChartResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *GetTickerChartResponse) SetData(v []string) {
	o.Data = v
}

// GetDates returns the Dates field value if set, zero value otherwise.
func (o *GetTickerChartResponse) GetDates() []GetTickerChartResponseDatesInner {
	if o == nil || IsNil(o.Dates) {
		var ret []GetTickerChartResponseDatesInner
		return ret
	}
	return o.Dates
}

// GetDatesOk returns a tuple with the Dates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTickerChartResponse) GetDatesOk() ([]GetTickerChartResponseDatesInner, bool) {
	if o == nil || IsNil(o.Dates) {
		return nil, false
	}
	return o.Dates, true
}

// HasDates returns a boolean if a field has been set.
func (o *GetTickerChartResponse) HasDates() bool {
	if o != nil && !IsNil(o.Dates) {
		return true
	}

	return false
}

// SetDates gets a reference to the given []GetTickerChartResponseDatesInner and assigns it to the Dates field.
func (o *GetTickerChartResponse) SetDates(v []GetTickerChartResponseDatesInner) {
	o.Dates = v
}

// GetExchangeStatus returns the ExchangeStatus field value if set, zero value otherwise.
func (o *GetTickerChartResponse) GetExchangeStatus() bool {
	if o == nil || IsNil(o.ExchangeStatus) {
		var ret bool
		return ret
	}
	return *o.ExchangeStatus
}

// GetExchangeStatusOk returns a tuple with the ExchangeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTickerChartResponse) GetExchangeStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.ExchangeStatus) {
		return nil, false
	}
	return o.ExchangeStatus, true
}

// HasExchangeStatus returns a boolean if a field has been set.
func (o *GetTickerChartResponse) HasExchangeStatus() bool {
	if o != nil && !IsNil(o.ExchangeStatus) {
		return true
	}

	return false
}

// SetExchangeStatus gets a reference to the given bool and assigns it to the ExchangeStatus field.
func (o *GetTickerChartResponse) SetExchangeStatus(v bool) {
	o.ExchangeStatus = &v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *GetTickerChartResponse) GetHasMore() int32 {
	if o == nil || IsNil(o.HasMore) {
		var ret int32
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTickerChartResponse) GetHasMoreOk() (*int32, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *GetTickerChartResponse) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given int32 and assigns it to the HasMore field.
func (o *GetTickerChartResponse) SetHasMore(v int32) {
	o.HasMore = &v
}

// GetPreClose returns the PreClose field value if set, zero value otherwise.
func (o *GetTickerChartResponse) GetPreClose() string {
	if o == nil || IsNil(o.PreClose) {
		var ret string
		return ret
	}
	return *o.PreClose
}

// GetPreCloseOk returns a tuple with the PreClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTickerChartResponse) GetPreCloseOk() (*string, bool) {
	if o == nil || IsNil(o.PreClose) {
		return nil, false
	}
	return o.PreClose, true
}

// HasPreClose returns a boolean if a field has been set.
func (o *GetTickerChartResponse) HasPreClose() bool {
	if o != nil && !IsNil(o.PreClose) {
		return true
	}

	return false
}

// SetPreClose gets a reference to the given string and assigns it to the PreClose field.
func (o *GetTickerChartResponse) SetPreClose(v string) {
	o.PreClose = &v
}

// GetTickerId returns the TickerId field value if set, zero value otherwise.
func (o *GetTickerChartResponse) GetTickerId() int32 {
	if o == nil || IsNil(o.TickerId) {
		var ret int32
		return ret
	}
	return *o.TickerId
}

// GetTickerIdOk returns a tuple with the TickerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTickerChartResponse) GetTickerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TickerId) {
		return nil, false
	}
	return o.TickerId, true
}

// HasTickerId returns a boolean if a field has been set.
func (o *GetTickerChartResponse) HasTickerId() bool {
	if o != nil && !IsNil(o.TickerId) {
		return true
	}

	return false
}

// SetTickerId gets a reference to the given int32 and assigns it to the TickerId field.
func (o *GetTickerChartResponse) SetTickerId(v int32) {
	o.TickerId = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *GetTickerChartResponse) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTickerChartResponse) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *GetTickerChartResponse) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *GetTickerChartResponse) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o GetTickerChartResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTickerChartResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Dates) {
		toSerialize["dates"] = o.Dates
	}
	if !IsNil(o.ExchangeStatus) {
		toSerialize["exchangeStatus"] = o.ExchangeStatus
	}
	if !IsNil(o.HasMore) {
		toSerialize["hasMore"] = o.HasMore
	}
	if !IsNil(o.PreClose) {
		toSerialize["preClose"] = o.PreClose
	}
	if !IsNil(o.TickerId) {
		toSerialize["tickerId"] = o.TickerId
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	return toSerialize, nil
}

type NullableGetTickerChartResponse struct {
	value *GetTickerChartResponse
	isSet bool
}

func (v NullableGetTickerChartResponse) Get() *GetTickerChartResponse {
	return v.value
}

func (v *NullableGetTickerChartResponse) Set(val *GetTickerChartResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTickerChartResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTickerChartResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTickerChartResponse(val *GetTickerChartResponse) *NullableGetTickerChartResponse {
	return &NullableGetTickerChartResponse{value: val, isSet: true}
}

func (v NullableGetTickerChartResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTickerChartResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


