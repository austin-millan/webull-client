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

// checks if the ActiveGainersLosers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveGainersLosers{}

// ActiveGainersLosers struct for ActiveGainersLosers
type ActiveGainersLosers struct {
	Change *string `json:"change,omitempty"`
	ChangeRatio *string `json:"changeRatio,omitempty"`
	Close *string `json:"close,omitempty"`
	CurrencyId *int32 `json:"currencyId,omitempty"`
	DerivativeSupport *int32 `json:"derivativeSupport,omitempty"`
	DisExchangeCode *string `json:"disExchangeCode,omitempty"`
	DisSymbol *string `json:"disSymbol,omitempty"`
	ExchangeCode *string `json:"exchangeCode,omitempty"`
	ExchangeId *int32 `json:"exchangeId,omitempty"`
	ListStatus *int32 `json:"listStatus,omitempty"`
	MarketValue *string `json:"marketValue,omitempty"`
	Name *string `json:"name,omitempty"`
	PChRatio *string `json:"pChRatio,omitempty"`
	PChange *string `json:"pChange,omitempty"`
	PPrice *string `json:"pPrice,omitempty"`
	RegionCode *string `json:"regionCode,omitempty"`
	RegionId *int32 `json:"regionId,omitempty"`
	SecType []int32 `json:"secType,omitempty"`
	Status *string `json:"status,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Template *string `json:"template,omitempty"`
	TickerId *int32 `json:"tickerId,omitempty"`
	TurnoverRate *string `json:"turnoverRate,omitempty"`
	Type *int32 `json:"type,omitempty"`
	Volume *string `json:"volume,omitempty"`
}

// NewActiveGainersLosers instantiates a new ActiveGainersLosers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveGainersLosers() *ActiveGainersLosers {
	this := ActiveGainersLosers{}
	return &this
}

// NewActiveGainersLosersWithDefaults instantiates a new ActiveGainersLosers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveGainersLosersWithDefaults() *ActiveGainersLosers {
	this := ActiveGainersLosers{}
	return &this
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetChange() string {
	if o == nil || IsNil(o.Change) {
		var ret string
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetChangeOk() (*string, bool) {
	if o == nil || IsNil(o.Change) {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasChange() bool {
	if o != nil && !IsNil(o.Change) {
		return true
	}

	return false
}

// SetChange gets a reference to the given string and assigns it to the Change field.
func (o *ActiveGainersLosers) SetChange(v string) {
	o.Change = &v
}

// GetChangeRatio returns the ChangeRatio field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetChangeRatio() string {
	if o == nil || IsNil(o.ChangeRatio) {
		var ret string
		return ret
	}
	return *o.ChangeRatio
}

// GetChangeRatioOk returns a tuple with the ChangeRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetChangeRatioOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeRatio) {
		return nil, false
	}
	return o.ChangeRatio, true
}

// HasChangeRatio returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasChangeRatio() bool {
	if o != nil && !IsNil(o.ChangeRatio) {
		return true
	}

	return false
}

// SetChangeRatio gets a reference to the given string and assigns it to the ChangeRatio field.
func (o *ActiveGainersLosers) SetChangeRatio(v string) {
	o.ChangeRatio = &v
}

// GetClose returns the Close field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetClose() string {
	if o == nil || IsNil(o.Close) {
		var ret string
		return ret
	}
	return *o.Close
}

// GetCloseOk returns a tuple with the Close field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetCloseOk() (*string, bool) {
	if o == nil || IsNil(o.Close) {
		return nil, false
	}
	return o.Close, true
}

// HasClose returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasClose() bool {
	if o != nil && !IsNil(o.Close) {
		return true
	}

	return false
}

// SetClose gets a reference to the given string and assigns it to the Close field.
func (o *ActiveGainersLosers) SetClose(v string) {
	o.Close = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetCurrencyId() int32 {
	if o == nil || IsNil(o.CurrencyId) {
		var ret int32
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetCurrencyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given int32 and assigns it to the CurrencyId field.
func (o *ActiveGainersLosers) SetCurrencyId(v int32) {
	o.CurrencyId = &v
}

// GetDerivativeSupport returns the DerivativeSupport field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetDerivativeSupport() int32 {
	if o == nil || IsNil(o.DerivativeSupport) {
		var ret int32
		return ret
	}
	return *o.DerivativeSupport
}

// GetDerivativeSupportOk returns a tuple with the DerivativeSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetDerivativeSupportOk() (*int32, bool) {
	if o == nil || IsNil(o.DerivativeSupport) {
		return nil, false
	}
	return o.DerivativeSupport, true
}

// HasDerivativeSupport returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasDerivativeSupport() bool {
	if o != nil && !IsNil(o.DerivativeSupport) {
		return true
	}

	return false
}

// SetDerivativeSupport gets a reference to the given int32 and assigns it to the DerivativeSupport field.
func (o *ActiveGainersLosers) SetDerivativeSupport(v int32) {
	o.DerivativeSupport = &v
}

// GetDisExchangeCode returns the DisExchangeCode field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetDisExchangeCode() string {
	if o == nil || IsNil(o.DisExchangeCode) {
		var ret string
		return ret
	}
	return *o.DisExchangeCode
}

// GetDisExchangeCodeOk returns a tuple with the DisExchangeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetDisExchangeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DisExchangeCode) {
		return nil, false
	}
	return o.DisExchangeCode, true
}

// HasDisExchangeCode returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasDisExchangeCode() bool {
	if o != nil && !IsNil(o.DisExchangeCode) {
		return true
	}

	return false
}

// SetDisExchangeCode gets a reference to the given string and assigns it to the DisExchangeCode field.
func (o *ActiveGainersLosers) SetDisExchangeCode(v string) {
	o.DisExchangeCode = &v
}

// GetDisSymbol returns the DisSymbol field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetDisSymbol() string {
	if o == nil || IsNil(o.DisSymbol) {
		var ret string
		return ret
	}
	return *o.DisSymbol
}

// GetDisSymbolOk returns a tuple with the DisSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetDisSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.DisSymbol) {
		return nil, false
	}
	return o.DisSymbol, true
}

// HasDisSymbol returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasDisSymbol() bool {
	if o != nil && !IsNil(o.DisSymbol) {
		return true
	}

	return false
}

// SetDisSymbol gets a reference to the given string and assigns it to the DisSymbol field.
func (o *ActiveGainersLosers) SetDisSymbol(v string) {
	o.DisSymbol = &v
}

// GetExchangeCode returns the ExchangeCode field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetExchangeCode() string {
	if o == nil || IsNil(o.ExchangeCode) {
		var ret string
		return ret
	}
	return *o.ExchangeCode
}

// GetExchangeCodeOk returns a tuple with the ExchangeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetExchangeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ExchangeCode) {
		return nil, false
	}
	return o.ExchangeCode, true
}

// HasExchangeCode returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasExchangeCode() bool {
	if o != nil && !IsNil(o.ExchangeCode) {
		return true
	}

	return false
}

// SetExchangeCode gets a reference to the given string and assigns it to the ExchangeCode field.
func (o *ActiveGainersLosers) SetExchangeCode(v string) {
	o.ExchangeCode = &v
}

// GetExchangeId returns the ExchangeId field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetExchangeId() int32 {
	if o == nil || IsNil(o.ExchangeId) {
		var ret int32
		return ret
	}
	return *o.ExchangeId
}

// GetExchangeIdOk returns a tuple with the ExchangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetExchangeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ExchangeId) {
		return nil, false
	}
	return o.ExchangeId, true
}

// HasExchangeId returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasExchangeId() bool {
	if o != nil && !IsNil(o.ExchangeId) {
		return true
	}

	return false
}

// SetExchangeId gets a reference to the given int32 and assigns it to the ExchangeId field.
func (o *ActiveGainersLosers) SetExchangeId(v int32) {
	o.ExchangeId = &v
}

// GetListStatus returns the ListStatus field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetListStatus() int32 {
	if o == nil || IsNil(o.ListStatus) {
		var ret int32
		return ret
	}
	return *o.ListStatus
}

// GetListStatusOk returns a tuple with the ListStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetListStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.ListStatus) {
		return nil, false
	}
	return o.ListStatus, true
}

// HasListStatus returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasListStatus() bool {
	if o != nil && !IsNil(o.ListStatus) {
		return true
	}

	return false
}

// SetListStatus gets a reference to the given int32 and assigns it to the ListStatus field.
func (o *ActiveGainersLosers) SetListStatus(v int32) {
	o.ListStatus = &v
}

// GetMarketValue returns the MarketValue field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetMarketValue() string {
	if o == nil || IsNil(o.MarketValue) {
		var ret string
		return ret
	}
	return *o.MarketValue
}

// GetMarketValueOk returns a tuple with the MarketValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetMarketValueOk() (*string, bool) {
	if o == nil || IsNil(o.MarketValue) {
		return nil, false
	}
	return o.MarketValue, true
}

// HasMarketValue returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasMarketValue() bool {
	if o != nil && !IsNil(o.MarketValue) {
		return true
	}

	return false
}

// SetMarketValue gets a reference to the given string and assigns it to the MarketValue field.
func (o *ActiveGainersLosers) SetMarketValue(v string) {
	o.MarketValue = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ActiveGainersLosers) SetName(v string) {
	o.Name = &v
}

// GetPChRatio returns the PChRatio field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetPChRatio() string {
	if o == nil || IsNil(o.PChRatio) {
		var ret string
		return ret
	}
	return *o.PChRatio
}

// GetPChRatioOk returns a tuple with the PChRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetPChRatioOk() (*string, bool) {
	if o == nil || IsNil(o.PChRatio) {
		return nil, false
	}
	return o.PChRatio, true
}

// HasPChRatio returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasPChRatio() bool {
	if o != nil && !IsNil(o.PChRatio) {
		return true
	}

	return false
}

// SetPChRatio gets a reference to the given string and assigns it to the PChRatio field.
func (o *ActiveGainersLosers) SetPChRatio(v string) {
	o.PChRatio = &v
}

// GetPChange returns the PChange field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetPChange() string {
	if o == nil || IsNil(o.PChange) {
		var ret string
		return ret
	}
	return *o.PChange
}

// GetPChangeOk returns a tuple with the PChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetPChangeOk() (*string, bool) {
	if o == nil || IsNil(o.PChange) {
		return nil, false
	}
	return o.PChange, true
}

// HasPChange returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasPChange() bool {
	if o != nil && !IsNil(o.PChange) {
		return true
	}

	return false
}

// SetPChange gets a reference to the given string and assigns it to the PChange field.
func (o *ActiveGainersLosers) SetPChange(v string) {
	o.PChange = &v
}

// GetPPrice returns the PPrice field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetPPrice() string {
	if o == nil || IsNil(o.PPrice) {
		var ret string
		return ret
	}
	return *o.PPrice
}

// GetPPriceOk returns a tuple with the PPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetPPriceOk() (*string, bool) {
	if o == nil || IsNil(o.PPrice) {
		return nil, false
	}
	return o.PPrice, true
}

// HasPPrice returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasPPrice() bool {
	if o != nil && !IsNil(o.PPrice) {
		return true
	}

	return false
}

// SetPPrice gets a reference to the given string and assigns it to the PPrice field.
func (o *ActiveGainersLosers) SetPPrice(v string) {
	o.PPrice = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetRegionCode() string {
	if o == nil || IsNil(o.RegionCode) {
		var ret string
		return ret
	}
	return *o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetRegionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RegionCode) {
		return nil, false
	}
	return o.RegionCode, true
}

// HasRegionCode returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasRegionCode() bool {
	if o != nil && !IsNil(o.RegionCode) {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given string and assigns it to the RegionCode field.
func (o *ActiveGainersLosers) SetRegionCode(v string) {
	o.RegionCode = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetRegionId() int32 {
	if o == nil || IsNil(o.RegionId) {
		var ret int32
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetRegionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RegionId) {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasRegionId() bool {
	if o != nil && !IsNil(o.RegionId) {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given int32 and assigns it to the RegionId field.
func (o *ActiveGainersLosers) SetRegionId(v int32) {
	o.RegionId = &v
}

// GetSecType returns the SecType field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetSecType() []int32 {
	if o == nil || IsNil(o.SecType) {
		var ret []int32
		return ret
	}
	return o.SecType
}

// GetSecTypeOk returns a tuple with the SecType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetSecTypeOk() ([]int32, bool) {
	if o == nil || IsNil(o.SecType) {
		return nil, false
	}
	return o.SecType, true
}

// HasSecType returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasSecType() bool {
	if o != nil && !IsNil(o.SecType) {
		return true
	}

	return false
}

// SetSecType gets a reference to the given []int32 and assigns it to the SecType field.
func (o *ActiveGainersLosers) SetSecType(v []int32) {
	o.SecType = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ActiveGainersLosers) SetStatus(v string) {
	o.Status = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ActiveGainersLosers) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *ActiveGainersLosers) SetTemplate(v string) {
	o.Template = &v
}

// GetTickerId returns the TickerId field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetTickerId() int32 {
	if o == nil || IsNil(o.TickerId) {
		var ret int32
		return ret
	}
	return *o.TickerId
}

// GetTickerIdOk returns a tuple with the TickerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetTickerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TickerId) {
		return nil, false
	}
	return o.TickerId, true
}

// HasTickerId returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasTickerId() bool {
	if o != nil && !IsNil(o.TickerId) {
		return true
	}

	return false
}

// SetTickerId gets a reference to the given int32 and assigns it to the TickerId field.
func (o *ActiveGainersLosers) SetTickerId(v int32) {
	o.TickerId = &v
}

// GetTurnoverRate returns the TurnoverRate field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetTurnoverRate() string {
	if o == nil || IsNil(o.TurnoverRate) {
		var ret string
		return ret
	}
	return *o.TurnoverRate
}

// GetTurnoverRateOk returns a tuple with the TurnoverRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetTurnoverRateOk() (*string, bool) {
	if o == nil || IsNil(o.TurnoverRate) {
		return nil, false
	}
	return o.TurnoverRate, true
}

// HasTurnoverRate returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasTurnoverRate() bool {
	if o != nil && !IsNil(o.TurnoverRate) {
		return true
	}

	return false
}

// SetTurnoverRate gets a reference to the given string and assigns it to the TurnoverRate field.
func (o *ActiveGainersLosers) SetTurnoverRate(v string) {
	o.TurnoverRate = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *ActiveGainersLosers) SetType(v int32) {
	o.Type = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *ActiveGainersLosers) GetVolume() string {
	if o == nil || IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGainersLosers) GetVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *ActiveGainersLosers) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *ActiveGainersLosers) SetVolume(v string) {
	o.Volume = &v
}

func (o ActiveGainersLosers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveGainersLosers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Change) {
		toSerialize["change"] = o.Change
	}
	if !IsNil(o.ChangeRatio) {
		toSerialize["changeRatio"] = o.ChangeRatio
	}
	if !IsNil(o.Close) {
		toSerialize["close"] = o.Close
	}
	if !IsNil(o.CurrencyId) {
		toSerialize["currencyId"] = o.CurrencyId
	}
	if !IsNil(o.DerivativeSupport) {
		toSerialize["derivativeSupport"] = o.DerivativeSupport
	}
	if !IsNil(o.DisExchangeCode) {
		toSerialize["disExchangeCode"] = o.DisExchangeCode
	}
	if !IsNil(o.DisSymbol) {
		toSerialize["disSymbol"] = o.DisSymbol
	}
	if !IsNil(o.ExchangeCode) {
		toSerialize["exchangeCode"] = o.ExchangeCode
	}
	if !IsNil(o.ExchangeId) {
		toSerialize["exchangeId"] = o.ExchangeId
	}
	if !IsNil(o.ListStatus) {
		toSerialize["listStatus"] = o.ListStatus
	}
	if !IsNil(o.MarketValue) {
		toSerialize["marketValue"] = o.MarketValue
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PChRatio) {
		toSerialize["pChRatio"] = o.PChRatio
	}
	if !IsNil(o.PChange) {
		toSerialize["pChange"] = o.PChange
	}
	if !IsNil(o.PPrice) {
		toSerialize["pPrice"] = o.PPrice
	}
	if !IsNil(o.RegionCode) {
		toSerialize["regionCode"] = o.RegionCode
	}
	if !IsNil(o.RegionId) {
		toSerialize["regionId"] = o.RegionId
	}
	if !IsNil(o.SecType) {
		toSerialize["secType"] = o.SecType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.TickerId) {
		toSerialize["tickerId"] = o.TickerId
	}
	if !IsNil(o.TurnoverRate) {
		toSerialize["turnoverRate"] = o.TurnoverRate
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

type NullableActiveGainersLosers struct {
	value *ActiveGainersLosers
	isSet bool
}

func (v NullableActiveGainersLosers) Get() *ActiveGainersLosers {
	return v.value
}

func (v *NullableActiveGainersLosers) Set(val *ActiveGainersLosers) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveGainersLosers) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveGainersLosers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveGainersLosers(val *ActiveGainersLosers) *NullableActiveGainersLosers {
	return &NullableActiveGainersLosers{value: val, isSet: true}
}

func (v NullableActiveGainersLosers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveGainersLosers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


