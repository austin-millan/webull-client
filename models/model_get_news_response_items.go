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

// checks if the GetNewsResponseItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNewsResponseItems{}

// GetNewsResponseItems struct for GetNewsResponseItems
type GetNewsResponseItems struct {
	CollectSource *string `json:"collectSource,omitempty"`
	Id *int32 `json:"id,omitempty"`
	NewsTime *string `json:"newsTime,omitempty"`
	NewsUrl *string `json:"newsUrl,omitempty"`
	SiteType *int32 `json:"siteType,omitempty"`
	SourceName *string `json:"sourceName,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewGetNewsResponseItems instantiates a new GetNewsResponseItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNewsResponseItems() *GetNewsResponseItems {
	this := GetNewsResponseItems{}
	return &this
}

// NewGetNewsResponseItemsWithDefaults instantiates a new GetNewsResponseItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNewsResponseItemsWithDefaults() *GetNewsResponseItems {
	this := GetNewsResponseItems{}
	return &this
}

// GetCollectSource returns the CollectSource field value if set, zero value otherwise.
func (o *GetNewsResponseItems) GetCollectSource() string {
	if o == nil || IsNil(o.CollectSource) {
		var ret string
		return ret
	}
	return *o.CollectSource
}

// GetCollectSourceOk returns a tuple with the CollectSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNewsResponseItems) GetCollectSourceOk() (*string, bool) {
	if o == nil || IsNil(o.CollectSource) {
		return nil, false
	}
	return o.CollectSource, true
}

// HasCollectSource returns a boolean if a field has been set.
func (o *GetNewsResponseItems) HasCollectSource() bool {
	if o != nil && !IsNil(o.CollectSource) {
		return true
	}

	return false
}

// SetCollectSource gets a reference to the given string and assigns it to the CollectSource field.
func (o *GetNewsResponseItems) SetCollectSource(v string) {
	o.CollectSource = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNewsResponseItems) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNewsResponseItems) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNewsResponseItems) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetNewsResponseItems) SetId(v int32) {
	o.Id = &v
}

// GetNewsTime returns the NewsTime field value if set, zero value otherwise.
func (o *GetNewsResponseItems) GetNewsTime() string {
	if o == nil || IsNil(o.NewsTime) {
		var ret string
		return ret
	}
	return *o.NewsTime
}

// GetNewsTimeOk returns a tuple with the NewsTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNewsResponseItems) GetNewsTimeOk() (*string, bool) {
	if o == nil || IsNil(o.NewsTime) {
		return nil, false
	}
	return o.NewsTime, true
}

// HasNewsTime returns a boolean if a field has been set.
func (o *GetNewsResponseItems) HasNewsTime() bool {
	if o != nil && !IsNil(o.NewsTime) {
		return true
	}

	return false
}

// SetNewsTime gets a reference to the given string and assigns it to the NewsTime field.
func (o *GetNewsResponseItems) SetNewsTime(v string) {
	o.NewsTime = &v
}

// GetNewsUrl returns the NewsUrl field value if set, zero value otherwise.
func (o *GetNewsResponseItems) GetNewsUrl() string {
	if o == nil || IsNil(o.NewsUrl) {
		var ret string
		return ret
	}
	return *o.NewsUrl
}

// GetNewsUrlOk returns a tuple with the NewsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNewsResponseItems) GetNewsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.NewsUrl) {
		return nil, false
	}
	return o.NewsUrl, true
}

// HasNewsUrl returns a boolean if a field has been set.
func (o *GetNewsResponseItems) HasNewsUrl() bool {
	if o != nil && !IsNil(o.NewsUrl) {
		return true
	}

	return false
}

// SetNewsUrl gets a reference to the given string and assigns it to the NewsUrl field.
func (o *GetNewsResponseItems) SetNewsUrl(v string) {
	o.NewsUrl = &v
}

// GetSiteType returns the SiteType field value if set, zero value otherwise.
func (o *GetNewsResponseItems) GetSiteType() int32 {
	if o == nil || IsNil(o.SiteType) {
		var ret int32
		return ret
	}
	return *o.SiteType
}

// GetSiteTypeOk returns a tuple with the SiteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNewsResponseItems) GetSiteTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.SiteType) {
		return nil, false
	}
	return o.SiteType, true
}

// HasSiteType returns a boolean if a field has been set.
func (o *GetNewsResponseItems) HasSiteType() bool {
	if o != nil && !IsNil(o.SiteType) {
		return true
	}

	return false
}

// SetSiteType gets a reference to the given int32 and assigns it to the SiteType field.
func (o *GetNewsResponseItems) SetSiteType(v int32) {
	o.SiteType = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *GetNewsResponseItems) GetSourceName() string {
	if o == nil || IsNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNewsResponseItems) GetSourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *GetNewsResponseItems) HasSourceName() bool {
	if o != nil && !IsNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *GetNewsResponseItems) SetSourceName(v string) {
	o.SourceName = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *GetNewsResponseItems) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNewsResponseItems) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *GetNewsResponseItems) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *GetNewsResponseItems) SetSummary(v string) {
	o.Summary = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetNewsResponseItems) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNewsResponseItems) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetNewsResponseItems) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetNewsResponseItems) SetTitle(v string) {
	o.Title = &v
}

func (o GetNewsResponseItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNewsResponseItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollectSource) {
		toSerialize["collectSource"] = o.CollectSource
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.NewsTime) {
		toSerialize["newsTime"] = o.NewsTime
	}
	if !IsNil(o.NewsUrl) {
		toSerialize["newsUrl"] = o.NewsUrl
	}
	if !IsNil(o.SiteType) {
		toSerialize["siteType"] = o.SiteType
	}
	if !IsNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableGetNewsResponseItems struct {
	value *GetNewsResponseItems
	isSet bool
}

func (v NullableGetNewsResponseItems) Get() *GetNewsResponseItems {
	return v.value
}

func (v *NullableGetNewsResponseItems) Set(val *GetNewsResponseItems) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNewsResponseItems) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNewsResponseItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNewsResponseItems(val *GetNewsResponseItems) *NullableGetNewsResponseItems {
	return &NullableGetNewsResponseItems{value: val, isSet: true}
}

func (v NullableGetNewsResponseItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNewsResponseItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


