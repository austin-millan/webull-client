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

// checks if the GetOrdersItemOrdersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrdersItemOrdersInner{}

// GetOrdersItemOrdersInner struct for GetOrdersItemOrdersInner
type GetOrdersItemOrdersInner struct {
	Action *OrderSide `json:"action,omitempty"`
	AssetType *string `json:"assetType,omitempty"`
	AvgFilledPrice *string `json:"avgFilledPrice,omitempty"`
	BrokerId *int32 `json:"brokerId,omitempty"`
	CanCancel *bool `json:"canCancel,omitempty"`
	CanModify *bool `json:"canModify,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	CreateTime0 *int64 `json:"createTime0,omitempty"`
	FilledQuantity *string `json:"filledQuantity,omitempty"`
	FilledTime *string `json:"filledTime,omitempty"`
	FilledTime0 *float32 `json:"filledTime0,omitempty"`
	FilledValue *string `json:"filledValue,omitempty"`
	OptionContractMultiplier *string `json:"optionContractMultiplier,omitempty"`
	OptionExercisePrice *string `json:"optionExercisePrice,omitempty"`
	OptionCycle *int64 `json:"optionCycle,omitempty"`
	OrderId *int32 `json:"orderId,omitempty"`
	OrderType *OrderType `json:"orderType,omitempty"`
	OptionType *OptionType `json:"optionType,omitempty"`
	OptionExpirationDate *string `json:"optionExpirationDate,omitempty"`
	Relation *string `json:"relation,omitempty"`
	OptionCategory *string `json:"optionCategory,omitempty"`
	RemainQuantity *string `json:"remainQuantity,omitempty"`
	StatusCode *string `json:"statusCode,omitempty"`
	StatusStr *string `json:"statusStr,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Ticker *GetOrdersItemOrdersInnerTicker `json:"ticker,omitempty"`
	TickerId *int32 `json:"tickerId,omitempty"`
	TickerPriceDefineList []GetOrdersItemOrdersInnerTickerPriceDefineListInner `json:"tickerPriceDefineList,omitempty"`
	TimeInForce *TimeInForce `json:"timeInForce,omitempty"`
	TotalQuantity *string `json:"totalQuantity,omitempty"`
	UpdateTime *string `json:"updateTime,omitempty"`
	UpdateTime0 *float32 `json:"updateTime0,omitempty"`
}

// NewGetOrdersItemOrdersInner instantiates a new GetOrdersItemOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrdersItemOrdersInner() *GetOrdersItemOrdersInner {
	this := GetOrdersItemOrdersInner{}
	return &this
}

// NewGetOrdersItemOrdersInnerWithDefaults instantiates a new GetOrdersItemOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrdersItemOrdersInnerWithDefaults() *GetOrdersItemOrdersInner {
	this := GetOrdersItemOrdersInner{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetAction() OrderSide {
	if o == nil || IsNil(o.Action) {
		var ret OrderSide
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetActionOk() (*OrderSide, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given OrderSide and assigns it to the Action field.
func (o *GetOrdersItemOrdersInner) SetAction(v OrderSide) {
	o.Action = &v
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetAssetType() string {
	if o == nil || IsNil(o.AssetType) {
		var ret string
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetAssetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssetType) {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasAssetType() bool {
	if o != nil && !IsNil(o.AssetType) {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given string and assigns it to the AssetType field.
func (o *GetOrdersItemOrdersInner) SetAssetType(v string) {
	o.AssetType = &v
}

// GetAvgFilledPrice returns the AvgFilledPrice field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetAvgFilledPrice() string {
	if o == nil || IsNil(o.AvgFilledPrice) {
		var ret string
		return ret
	}
	return *o.AvgFilledPrice
}

// GetAvgFilledPriceOk returns a tuple with the AvgFilledPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetAvgFilledPriceOk() (*string, bool) {
	if o == nil || IsNil(o.AvgFilledPrice) {
		return nil, false
	}
	return o.AvgFilledPrice, true
}

// HasAvgFilledPrice returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasAvgFilledPrice() bool {
	if o != nil && !IsNil(o.AvgFilledPrice) {
		return true
	}

	return false
}

// SetAvgFilledPrice gets a reference to the given string and assigns it to the AvgFilledPrice field.
func (o *GetOrdersItemOrdersInner) SetAvgFilledPrice(v string) {
	o.AvgFilledPrice = &v
}

// GetBrokerId returns the BrokerId field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetBrokerId() int32 {
	if o == nil || IsNil(o.BrokerId) {
		var ret int32
		return ret
	}
	return *o.BrokerId
}

// GetBrokerIdOk returns a tuple with the BrokerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetBrokerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BrokerId) {
		return nil, false
	}
	return o.BrokerId, true
}

// HasBrokerId returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasBrokerId() bool {
	if o != nil && !IsNil(o.BrokerId) {
		return true
	}

	return false
}

// SetBrokerId gets a reference to the given int32 and assigns it to the BrokerId field.
func (o *GetOrdersItemOrdersInner) SetBrokerId(v int32) {
	o.BrokerId = &v
}

// GetCanCancel returns the CanCancel field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetCanCancel() bool {
	if o == nil || IsNil(o.CanCancel) {
		var ret bool
		return ret
	}
	return *o.CanCancel
}

// GetCanCancelOk returns a tuple with the CanCancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetCanCancelOk() (*bool, bool) {
	if o == nil || IsNil(o.CanCancel) {
		return nil, false
	}
	return o.CanCancel, true
}

// HasCanCancel returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasCanCancel() bool {
	if o != nil && !IsNil(o.CanCancel) {
		return true
	}

	return false
}

// SetCanCancel gets a reference to the given bool and assigns it to the CanCancel field.
func (o *GetOrdersItemOrdersInner) SetCanCancel(v bool) {
	o.CanCancel = &v
}

// GetCanModify returns the CanModify field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetCanModify() bool {
	if o == nil || IsNil(o.CanModify) {
		var ret bool
		return ret
	}
	return *o.CanModify
}

// GetCanModifyOk returns a tuple with the CanModify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetCanModifyOk() (*bool, bool) {
	if o == nil || IsNil(o.CanModify) {
		return nil, false
	}
	return o.CanModify, true
}

// HasCanModify returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasCanModify() bool {
	if o != nil && !IsNil(o.CanModify) {
		return true
	}

	return false
}

// SetCanModify gets a reference to the given bool and assigns it to the CanModify field.
func (o *GetOrdersItemOrdersInner) SetCanModify(v bool) {
	o.CanModify = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *GetOrdersItemOrdersInner) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetCreateTime0 returns the CreateTime0 field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetCreateTime0() int64 {
	if o == nil || IsNil(o.CreateTime0) {
		var ret int64
		return ret
	}
	return *o.CreateTime0
}

// GetCreateTime0Ok returns a tuple with the CreateTime0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetCreateTime0Ok() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime0) {
		return nil, false
	}
	return o.CreateTime0, true
}

// HasCreateTime0 returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasCreateTime0() bool {
	if o != nil && !IsNil(o.CreateTime0) {
		return true
	}

	return false
}

// SetCreateTime0 gets a reference to the given int64 and assigns it to the CreateTime0 field.
func (o *GetOrdersItemOrdersInner) SetCreateTime0(v int64) {
	o.CreateTime0 = &v
}

// GetFilledQuantity returns the FilledQuantity field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetFilledQuantity() string {
	if o == nil || IsNil(o.FilledQuantity) {
		var ret string
		return ret
	}
	return *o.FilledQuantity
}

// GetFilledQuantityOk returns a tuple with the FilledQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetFilledQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.FilledQuantity) {
		return nil, false
	}
	return o.FilledQuantity, true
}

// HasFilledQuantity returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasFilledQuantity() bool {
	if o != nil && !IsNil(o.FilledQuantity) {
		return true
	}

	return false
}

// SetFilledQuantity gets a reference to the given string and assigns it to the FilledQuantity field.
func (o *GetOrdersItemOrdersInner) SetFilledQuantity(v string) {
	o.FilledQuantity = &v
}

// GetFilledTime returns the FilledTime field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetFilledTime() string {
	if o == nil || IsNil(o.FilledTime) {
		var ret string
		return ret
	}
	return *o.FilledTime
}

// GetFilledTimeOk returns a tuple with the FilledTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetFilledTimeOk() (*string, bool) {
	if o == nil || IsNil(o.FilledTime) {
		return nil, false
	}
	return o.FilledTime, true
}

// HasFilledTime returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasFilledTime() bool {
	if o != nil && !IsNil(o.FilledTime) {
		return true
	}

	return false
}

// SetFilledTime gets a reference to the given string and assigns it to the FilledTime field.
func (o *GetOrdersItemOrdersInner) SetFilledTime(v string) {
	o.FilledTime = &v
}

// GetFilledTime0 returns the FilledTime0 field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetFilledTime0() float32 {
	if o == nil || IsNil(o.FilledTime0) {
		var ret float32
		return ret
	}
	return *o.FilledTime0
}

// GetFilledTime0Ok returns a tuple with the FilledTime0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetFilledTime0Ok() (*float32, bool) {
	if o == nil || IsNil(o.FilledTime0) {
		return nil, false
	}
	return o.FilledTime0, true
}

// HasFilledTime0 returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasFilledTime0() bool {
	if o != nil && !IsNil(o.FilledTime0) {
		return true
	}

	return false
}

// SetFilledTime0 gets a reference to the given float32 and assigns it to the FilledTime0 field.
func (o *GetOrdersItemOrdersInner) SetFilledTime0(v float32) {
	o.FilledTime0 = &v
}

// GetFilledValue returns the FilledValue field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetFilledValue() string {
	if o == nil || IsNil(o.FilledValue) {
		var ret string
		return ret
	}
	return *o.FilledValue
}

// GetFilledValueOk returns a tuple with the FilledValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetFilledValueOk() (*string, bool) {
	if o == nil || IsNil(o.FilledValue) {
		return nil, false
	}
	return o.FilledValue, true
}

// HasFilledValue returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasFilledValue() bool {
	if o != nil && !IsNil(o.FilledValue) {
		return true
	}

	return false
}

// SetFilledValue gets a reference to the given string and assigns it to the FilledValue field.
func (o *GetOrdersItemOrdersInner) SetFilledValue(v string) {
	o.FilledValue = &v
}

// GetOptionContractMultiplier returns the OptionContractMultiplier field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetOptionContractMultiplier() string {
	if o == nil || IsNil(o.OptionContractMultiplier) {
		var ret string
		return ret
	}
	return *o.OptionContractMultiplier
}

// GetOptionContractMultiplierOk returns a tuple with the OptionContractMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetOptionContractMultiplierOk() (*string, bool) {
	if o == nil || IsNil(o.OptionContractMultiplier) {
		return nil, false
	}
	return o.OptionContractMultiplier, true
}

// HasOptionContractMultiplier returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasOptionContractMultiplier() bool {
	if o != nil && !IsNil(o.OptionContractMultiplier) {
		return true
	}

	return false
}

// SetOptionContractMultiplier gets a reference to the given string and assigns it to the OptionContractMultiplier field.
func (o *GetOrdersItemOrdersInner) SetOptionContractMultiplier(v string) {
	o.OptionContractMultiplier = &v
}

// GetOptionExercisePrice returns the OptionExercisePrice field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetOptionExercisePrice() string {
	if o == nil || IsNil(o.OptionExercisePrice) {
		var ret string
		return ret
	}
	return *o.OptionExercisePrice
}

// GetOptionExercisePriceOk returns a tuple with the OptionExercisePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetOptionExercisePriceOk() (*string, bool) {
	if o == nil || IsNil(o.OptionExercisePrice) {
		return nil, false
	}
	return o.OptionExercisePrice, true
}

// HasOptionExercisePrice returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasOptionExercisePrice() bool {
	if o != nil && !IsNil(o.OptionExercisePrice) {
		return true
	}

	return false
}

// SetOptionExercisePrice gets a reference to the given string and assigns it to the OptionExercisePrice field.
func (o *GetOrdersItemOrdersInner) SetOptionExercisePrice(v string) {
	o.OptionExercisePrice = &v
}

// GetOptionCycle returns the OptionCycle field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetOptionCycle() int64 {
	if o == nil || IsNil(o.OptionCycle) {
		var ret int64
		return ret
	}
	return *o.OptionCycle
}

// GetOptionCycleOk returns a tuple with the OptionCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetOptionCycleOk() (*int64, bool) {
	if o == nil || IsNil(o.OptionCycle) {
		return nil, false
	}
	return o.OptionCycle, true
}

// HasOptionCycle returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasOptionCycle() bool {
	if o != nil && !IsNil(o.OptionCycle) {
		return true
	}

	return false
}

// SetOptionCycle gets a reference to the given int64 and assigns it to the OptionCycle field.
func (o *GetOrdersItemOrdersInner) SetOptionCycle(v int64) {
	o.OptionCycle = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetOrderId() int32 {
	if o == nil || IsNil(o.OrderId) {
		var ret int32
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetOrderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int32 and assigns it to the OrderId field.
func (o *GetOrdersItemOrdersInner) SetOrderId(v int32) {
	o.OrderId = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetOrderType() OrderType {
	if o == nil || IsNil(o.OrderType) {
		var ret OrderType
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetOrderTypeOk() (*OrderType, bool) {
	if o == nil || IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasOrderType() bool {
	if o != nil && !IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given OrderType and assigns it to the OrderType field.
func (o *GetOrdersItemOrdersInner) SetOrderType(v OrderType) {
	o.OrderType = &v
}

// GetOptionType returns the OptionType field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetOptionType() OptionType {
	if o == nil || IsNil(o.OptionType) {
		var ret OptionType
		return ret
	}
	return *o.OptionType
}

// GetOptionTypeOk returns a tuple with the OptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetOptionTypeOk() (*OptionType, bool) {
	if o == nil || IsNil(o.OptionType) {
		return nil, false
	}
	return o.OptionType, true
}

// HasOptionType returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasOptionType() bool {
	if o != nil && !IsNil(o.OptionType) {
		return true
	}

	return false
}

// SetOptionType gets a reference to the given OptionType and assigns it to the OptionType field.
func (o *GetOrdersItemOrdersInner) SetOptionType(v OptionType) {
	o.OptionType = &v
}

// GetOptionExpirationDate returns the OptionExpirationDate field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetOptionExpirationDate() string {
	if o == nil || IsNil(o.OptionExpirationDate) {
		var ret string
		return ret
	}
	return *o.OptionExpirationDate
}

// GetOptionExpirationDateOk returns a tuple with the OptionExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetOptionExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.OptionExpirationDate) {
		return nil, false
	}
	return o.OptionExpirationDate, true
}

// HasOptionExpirationDate returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasOptionExpirationDate() bool {
	if o != nil && !IsNil(o.OptionExpirationDate) {
		return true
	}

	return false
}

// SetOptionExpirationDate gets a reference to the given string and assigns it to the OptionExpirationDate field.
func (o *GetOrdersItemOrdersInner) SetOptionExpirationDate(v string) {
	o.OptionExpirationDate = &v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetRelation() string {
	if o == nil || IsNil(o.Relation) {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetRelationOk() (*string, bool) {
	if o == nil || IsNil(o.Relation) {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasRelation() bool {
	if o != nil && !IsNil(o.Relation) {
		return true
	}

	return false
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *GetOrdersItemOrdersInner) SetRelation(v string) {
	o.Relation = &v
}

// GetOptionCategory returns the OptionCategory field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetOptionCategory() string {
	if o == nil || IsNil(o.OptionCategory) {
		var ret string
		return ret
	}
	return *o.OptionCategory
}

// GetOptionCategoryOk returns a tuple with the OptionCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetOptionCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.OptionCategory) {
		return nil, false
	}
	return o.OptionCategory, true
}

// HasOptionCategory returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasOptionCategory() bool {
	if o != nil && !IsNil(o.OptionCategory) {
		return true
	}

	return false
}

// SetOptionCategory gets a reference to the given string and assigns it to the OptionCategory field.
func (o *GetOrdersItemOrdersInner) SetOptionCategory(v string) {
	o.OptionCategory = &v
}

// GetRemainQuantity returns the RemainQuantity field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetRemainQuantity() string {
	if o == nil || IsNil(o.RemainQuantity) {
		var ret string
		return ret
	}
	return *o.RemainQuantity
}

// GetRemainQuantityOk returns a tuple with the RemainQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetRemainQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.RemainQuantity) {
		return nil, false
	}
	return o.RemainQuantity, true
}

// HasRemainQuantity returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasRemainQuantity() bool {
	if o != nil && !IsNil(o.RemainQuantity) {
		return true
	}

	return false
}

// SetRemainQuantity gets a reference to the given string and assigns it to the RemainQuantity field.
func (o *GetOrdersItemOrdersInner) SetRemainQuantity(v string) {
	o.RemainQuantity = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetStatusCode() string {
	if o == nil || IsNil(o.StatusCode) {
		var ret string
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetStatusCodeOk() (*string, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given string and assigns it to the StatusCode field.
func (o *GetOrdersItemOrdersInner) SetStatusCode(v string) {
	o.StatusCode = &v
}

// GetStatusStr returns the StatusStr field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetStatusStr() string {
	if o == nil || IsNil(o.StatusStr) {
		var ret string
		return ret
	}
	return *o.StatusStr
}

// GetStatusStrOk returns a tuple with the StatusStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetStatusStrOk() (*string, bool) {
	if o == nil || IsNil(o.StatusStr) {
		return nil, false
	}
	return o.StatusStr, true
}

// HasStatusStr returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasStatusStr() bool {
	if o != nil && !IsNil(o.StatusStr) {
		return true
	}

	return false
}

// SetStatusStr gets a reference to the given string and assigns it to the StatusStr field.
func (o *GetOrdersItemOrdersInner) SetStatusStr(v string) {
	o.StatusStr = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetOrdersItemOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTicker returns the Ticker field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetTicker() GetOrdersItemOrdersInnerTicker {
	if o == nil || IsNil(o.Ticker) {
		var ret GetOrdersItemOrdersInnerTicker
		return ret
	}
	return *o.Ticker
}

// GetTickerOk returns a tuple with the Ticker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetTickerOk() (*GetOrdersItemOrdersInnerTicker, bool) {
	if o == nil || IsNil(o.Ticker) {
		return nil, false
	}
	return o.Ticker, true
}

// HasTicker returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasTicker() bool {
	if o != nil && !IsNil(o.Ticker) {
		return true
	}

	return false
}

// SetTicker gets a reference to the given GetOrdersItemOrdersInnerTicker and assigns it to the Ticker field.
func (o *GetOrdersItemOrdersInner) SetTicker(v GetOrdersItemOrdersInnerTicker) {
	o.Ticker = &v
}

// GetTickerId returns the TickerId field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetTickerId() int32 {
	if o == nil || IsNil(o.TickerId) {
		var ret int32
		return ret
	}
	return *o.TickerId
}

// GetTickerIdOk returns a tuple with the TickerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetTickerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TickerId) {
		return nil, false
	}
	return o.TickerId, true
}

// HasTickerId returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasTickerId() bool {
	if o != nil && !IsNil(o.TickerId) {
		return true
	}

	return false
}

// SetTickerId gets a reference to the given int32 and assigns it to the TickerId field.
func (o *GetOrdersItemOrdersInner) SetTickerId(v int32) {
	o.TickerId = &v
}

// GetTickerPriceDefineList returns the TickerPriceDefineList field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetTickerPriceDefineList() []GetOrdersItemOrdersInnerTickerPriceDefineListInner {
	if o == nil || IsNil(o.TickerPriceDefineList) {
		var ret []GetOrdersItemOrdersInnerTickerPriceDefineListInner
		return ret
	}
	return o.TickerPriceDefineList
}

// GetTickerPriceDefineListOk returns a tuple with the TickerPriceDefineList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetTickerPriceDefineListOk() ([]GetOrdersItemOrdersInnerTickerPriceDefineListInner, bool) {
	if o == nil || IsNil(o.TickerPriceDefineList) {
		return nil, false
	}
	return o.TickerPriceDefineList, true
}

// HasTickerPriceDefineList returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasTickerPriceDefineList() bool {
	if o != nil && !IsNil(o.TickerPriceDefineList) {
		return true
	}

	return false
}

// SetTickerPriceDefineList gets a reference to the given []GetOrdersItemOrdersInnerTickerPriceDefineListInner and assigns it to the TickerPriceDefineList field.
func (o *GetOrdersItemOrdersInner) SetTickerPriceDefineList(v []GetOrdersItemOrdersInnerTickerPriceDefineListInner) {
	o.TickerPriceDefineList = v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetTimeInForce() TimeInForce {
	if o == nil || IsNil(o.TimeInForce) {
		var ret TimeInForce
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetTimeInForceOk() (*TimeInForce, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given TimeInForce and assigns it to the TimeInForce field.
func (o *GetOrdersItemOrdersInner) SetTimeInForce(v TimeInForce) {
	o.TimeInForce = &v
}

// GetTotalQuantity returns the TotalQuantity field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetTotalQuantity() string {
	if o == nil || IsNil(o.TotalQuantity) {
		var ret string
		return ret
	}
	return *o.TotalQuantity
}

// GetTotalQuantityOk returns a tuple with the TotalQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetTotalQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.TotalQuantity) {
		return nil, false
	}
	return o.TotalQuantity, true
}

// HasTotalQuantity returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasTotalQuantity() bool {
	if o != nil && !IsNil(o.TotalQuantity) {
		return true
	}

	return false
}

// SetTotalQuantity gets a reference to the given string and assigns it to the TotalQuantity field.
func (o *GetOrdersItemOrdersInner) SetTotalQuantity(v string) {
	o.TotalQuantity = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetUpdateTime() string {
	if o == nil || IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *GetOrdersItemOrdersInner) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// GetUpdateTime0 returns the UpdateTime0 field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInner) GetUpdateTime0() float32 {
	if o == nil || IsNil(o.UpdateTime0) {
		var ret float32
		return ret
	}
	return *o.UpdateTime0
}

// GetUpdateTime0Ok returns a tuple with the UpdateTime0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInner) GetUpdateTime0Ok() (*float32, bool) {
	if o == nil || IsNil(o.UpdateTime0) {
		return nil, false
	}
	return o.UpdateTime0, true
}

// HasUpdateTime0 returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInner) HasUpdateTime0() bool {
	if o != nil && !IsNil(o.UpdateTime0) {
		return true
	}

	return false
}

// SetUpdateTime0 gets a reference to the given float32 and assigns it to the UpdateTime0 field.
func (o *GetOrdersItemOrdersInner) SetUpdateTime0(v float32) {
	o.UpdateTime0 = &v
}

func (o GetOrdersItemOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrdersItemOrdersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.AssetType) {
		toSerialize["assetType"] = o.AssetType
	}
	if !IsNil(o.AvgFilledPrice) {
		toSerialize["avgFilledPrice"] = o.AvgFilledPrice
	}
	if !IsNil(o.BrokerId) {
		toSerialize["brokerId"] = o.BrokerId
	}
	if !IsNil(o.CanCancel) {
		toSerialize["canCancel"] = o.CanCancel
	}
	if !IsNil(o.CanModify) {
		toSerialize["canModify"] = o.CanModify
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.CreateTime0) {
		toSerialize["createTime0"] = o.CreateTime0
	}
	if !IsNil(o.FilledQuantity) {
		toSerialize["filledQuantity"] = o.FilledQuantity
	}
	if !IsNil(o.FilledTime) {
		toSerialize["filledTime"] = o.FilledTime
	}
	if !IsNil(o.FilledTime0) {
		toSerialize["filledTime0"] = o.FilledTime0
	}
	if !IsNil(o.FilledValue) {
		toSerialize["filledValue"] = o.FilledValue
	}
	if !IsNil(o.OptionContractMultiplier) {
		toSerialize["optionContractMultiplier"] = o.OptionContractMultiplier
	}
	if !IsNil(o.OptionExercisePrice) {
		toSerialize["optionExercisePrice"] = o.OptionExercisePrice
	}
	if !IsNil(o.OptionCycle) {
		toSerialize["optionCycle"] = o.OptionCycle
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if !IsNil(o.OptionType) {
		toSerialize["optionType"] = o.OptionType
	}
	if !IsNil(o.OptionExpirationDate) {
		toSerialize["optionExpirationDate"] = o.OptionExpirationDate
	}
	if !IsNil(o.Relation) {
		toSerialize["relation"] = o.Relation
	}
	if !IsNil(o.OptionCategory) {
		toSerialize["optionCategory"] = o.OptionCategory
	}
	if !IsNil(o.RemainQuantity) {
		toSerialize["remainQuantity"] = o.RemainQuantity
	}
	if !IsNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	if !IsNil(o.StatusStr) {
		toSerialize["statusStr"] = o.StatusStr
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Ticker) {
		toSerialize["ticker"] = o.Ticker
	}
	if !IsNil(o.TickerId) {
		toSerialize["tickerId"] = o.TickerId
	}
	if !IsNil(o.TickerPriceDefineList) {
		toSerialize["tickerPriceDefineList"] = o.TickerPriceDefineList
	}
	if !IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !IsNil(o.TotalQuantity) {
		toSerialize["totalQuantity"] = o.TotalQuantity
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !IsNil(o.UpdateTime0) {
		toSerialize["updateTime0"] = o.UpdateTime0
	}
	return toSerialize, nil
}

type NullableGetOrdersItemOrdersInner struct {
	value *GetOrdersItemOrdersInner
	isSet bool
}

func (v NullableGetOrdersItemOrdersInner) Get() *GetOrdersItemOrdersInner {
	return v.value
}

func (v *NullableGetOrdersItemOrdersInner) Set(val *GetOrdersItemOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrdersItemOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrdersItemOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrdersItemOrdersInner(val *GetOrdersItemOrdersInner) *NullableGetOrdersItemOrdersInner {
	return &NullableGetOrdersItemOrdersInner{value: val, isSet: true}
}

func (v NullableGetOrdersItemOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrdersItemOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


