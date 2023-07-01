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

// checks if the GetStockAnalysisResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStockAnalysisResponse{}

// GetStockAnalysisResponse struct for GetStockAnalysisResponse
type GetStockAnalysisResponse struct {
	ForecastEps *GetStockAnalysisResponseForecastEps `json:"forecastEps,omitempty"`
	Rating *GetStockAnalysisResponseRating `json:"rating,omitempty"`
	TargetPrice *GetStockAnalysisResponseTargetPrice `json:"targetPrice,omitempty"`
}

// NewGetStockAnalysisResponse instantiates a new GetStockAnalysisResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStockAnalysisResponse() *GetStockAnalysisResponse {
	this := GetStockAnalysisResponse{}
	return &this
}

// NewGetStockAnalysisResponseWithDefaults instantiates a new GetStockAnalysisResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStockAnalysisResponseWithDefaults() *GetStockAnalysisResponse {
	this := GetStockAnalysisResponse{}
	return &this
}

// GetForecastEps returns the ForecastEps field value if set, zero value otherwise.
func (o *GetStockAnalysisResponse) GetForecastEps() GetStockAnalysisResponseForecastEps {
	if o == nil || IsNil(o.ForecastEps) {
		var ret GetStockAnalysisResponseForecastEps
		return ret
	}
	return *o.ForecastEps
}

// GetForecastEpsOk returns a tuple with the ForecastEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStockAnalysisResponse) GetForecastEpsOk() (*GetStockAnalysisResponseForecastEps, bool) {
	if o == nil || IsNil(o.ForecastEps) {
		return nil, false
	}
	return o.ForecastEps, true
}

// HasForecastEps returns a boolean if a field has been set.
func (o *GetStockAnalysisResponse) HasForecastEps() bool {
	if o != nil && !IsNil(o.ForecastEps) {
		return true
	}

	return false
}

// SetForecastEps gets a reference to the given GetStockAnalysisResponseForecastEps and assigns it to the ForecastEps field.
func (o *GetStockAnalysisResponse) SetForecastEps(v GetStockAnalysisResponseForecastEps) {
	o.ForecastEps = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *GetStockAnalysisResponse) GetRating() GetStockAnalysisResponseRating {
	if o == nil || IsNil(o.Rating) {
		var ret GetStockAnalysisResponseRating
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStockAnalysisResponse) GetRatingOk() (*GetStockAnalysisResponseRating, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *GetStockAnalysisResponse) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given GetStockAnalysisResponseRating and assigns it to the Rating field.
func (o *GetStockAnalysisResponse) SetRating(v GetStockAnalysisResponseRating) {
	o.Rating = &v
}

// GetTargetPrice returns the TargetPrice field value if set, zero value otherwise.
func (o *GetStockAnalysisResponse) GetTargetPrice() GetStockAnalysisResponseTargetPrice {
	if o == nil || IsNil(o.TargetPrice) {
		var ret GetStockAnalysisResponseTargetPrice
		return ret
	}
	return *o.TargetPrice
}

// GetTargetPriceOk returns a tuple with the TargetPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStockAnalysisResponse) GetTargetPriceOk() (*GetStockAnalysisResponseTargetPrice, bool) {
	if o == nil || IsNil(o.TargetPrice) {
		return nil, false
	}
	return o.TargetPrice, true
}

// HasTargetPrice returns a boolean if a field has been set.
func (o *GetStockAnalysisResponse) HasTargetPrice() bool {
	if o != nil && !IsNil(o.TargetPrice) {
		return true
	}

	return false
}

// SetTargetPrice gets a reference to the given GetStockAnalysisResponseTargetPrice and assigns it to the TargetPrice field.
func (o *GetStockAnalysisResponse) SetTargetPrice(v GetStockAnalysisResponseTargetPrice) {
	o.TargetPrice = &v
}

func (o GetStockAnalysisResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStockAnalysisResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForecastEps) {
		toSerialize["forecastEps"] = o.ForecastEps
	}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if !IsNil(o.TargetPrice) {
		toSerialize["targetPrice"] = o.TargetPrice
	}
	return toSerialize, nil
}

type NullableGetStockAnalysisResponse struct {
	value *GetStockAnalysisResponse
	isSet bool
}

func (v NullableGetStockAnalysisResponse) Get() *GetStockAnalysisResponse {
	return v.value
}

func (v *NullableGetStockAnalysisResponse) Set(val *GetStockAnalysisResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStockAnalysisResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStockAnalysisResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStockAnalysisResponse(val *GetStockAnalysisResponse) *NullableGetStockAnalysisResponse {
	return &NullableGetStockAnalysisResponse{value: val, isSet: true}
}

func (v NullableGetStockAnalysisResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStockAnalysisResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


