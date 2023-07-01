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

// checks if the Transfer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transfer{}

// Transfer struct for Transfer
type Transfer struct {
	Id *float32 `json:"id,omitempty"`
	SecAccountId *float32 `json:"secAccountId,omitempty"`
	BrokerAccountId *string `json:"brokerAccountId,omitempty"`
	Type *string `json:"type,omitempty"`
	Direction *string `json:"direction,omitempty"`
	Amount *float32 `json:"amount,omitempty"`
	AmountStr *string `json:"amountStr,omitempty"`
	Currency *string `json:"currency,omitempty"`
	TransferId *string `json:"transferId,omitempty"`
	ExternalTransferId *string `json:"externalTransferId,omitempty"`
	Status *string `json:"status,omitempty"`
	SubStatus *string `json:"subStatus,omitempty"`
	AchId *string `json:"achId,omitempty"`
	BankId *string `json:"bankId,omitempty"`
	BankType *string `json:"bankType,omitempty"`
	BankAccount *string `json:"bankAccount,omitempty"`
	BankAccountName *string `json:"bankAccountName,omitempty"`
	CustomerType *string `json:"customerType,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	UpdateTime *string `json:"updateTime,omitempty"`
	CreateTimeStr *string `json:"createTimeStr,omitempty"`
	FinishTime *string `json:"finishTime,omitempty"`
	AvailableTime *string `json:"availableTime,omitempty"`
	AvailableLampSendStatus *string `json:"availableLampSendStatus,omitempty"`
	AvailableTimeStr *string `json:"availableTimeStr,omitempty"`
	ReturnSendStatus *string `json:"returnSendStatus,omitempty"`
	RecordSendStatus *string `json:"recordSendStatus,omitempty"`
	SerialId *string `json:"serialId,omitempty"`
	ArriveTime *string `json:"arriveTime,omitempty"`
	FirstGift *bool `json:"firstGift,omitempty"`
	WaitCardBinding *bool `json:"waitCardBinding,omitempty"`
	TipsInfo *string `json:"tipsInfo,omitempty"`
}

// NewTransfer instantiates a new Transfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransfer() *Transfer {
	this := Transfer{}
	return &this
}

// NewTransferWithDefaults instantiates a new Transfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferWithDefaults() *Transfer {
	this := Transfer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Transfer) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Transfer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *Transfer) SetId(v float32) {
	o.Id = &v
}

// GetSecAccountId returns the SecAccountId field value if set, zero value otherwise.
func (o *Transfer) GetSecAccountId() float32 {
	if o == nil || IsNil(o.SecAccountId) {
		var ret float32
		return ret
	}
	return *o.SecAccountId
}

// GetSecAccountIdOk returns a tuple with the SecAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetSecAccountIdOk() (*float32, bool) {
	if o == nil || IsNil(o.SecAccountId) {
		return nil, false
	}
	return o.SecAccountId, true
}

// HasSecAccountId returns a boolean if a field has been set.
func (o *Transfer) HasSecAccountId() bool {
	if o != nil && !IsNil(o.SecAccountId) {
		return true
	}

	return false
}

// SetSecAccountId gets a reference to the given float32 and assigns it to the SecAccountId field.
func (o *Transfer) SetSecAccountId(v float32) {
	o.SecAccountId = &v
}

// GetBrokerAccountId returns the BrokerAccountId field value if set, zero value otherwise.
func (o *Transfer) GetBrokerAccountId() string {
	if o == nil || IsNil(o.BrokerAccountId) {
		var ret string
		return ret
	}
	return *o.BrokerAccountId
}

// GetBrokerAccountIdOk returns a tuple with the BrokerAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetBrokerAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrokerAccountId) {
		return nil, false
	}
	return o.BrokerAccountId, true
}

// HasBrokerAccountId returns a boolean if a field has been set.
func (o *Transfer) HasBrokerAccountId() bool {
	if o != nil && !IsNil(o.BrokerAccountId) {
		return true
	}

	return false
}

// SetBrokerAccountId gets a reference to the given string and assigns it to the BrokerAccountId field.
func (o *Transfer) SetBrokerAccountId(v string) {
	o.BrokerAccountId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Transfer) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Transfer) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Transfer) SetType(v string) {
	o.Type = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *Transfer) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *Transfer) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *Transfer) SetDirection(v string) {
	o.Direction = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Transfer) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Transfer) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *Transfer) SetAmount(v float32) {
	o.Amount = &v
}

// GetAmountStr returns the AmountStr field value if set, zero value otherwise.
func (o *Transfer) GetAmountStr() string {
	if o == nil || IsNil(o.AmountStr) {
		var ret string
		return ret
	}
	return *o.AmountStr
}

// GetAmountStrOk returns a tuple with the AmountStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetAmountStrOk() (*string, bool) {
	if o == nil || IsNil(o.AmountStr) {
		return nil, false
	}
	return o.AmountStr, true
}

// HasAmountStr returns a boolean if a field has been set.
func (o *Transfer) HasAmountStr() bool {
	if o != nil && !IsNil(o.AmountStr) {
		return true
	}

	return false
}

// SetAmountStr gets a reference to the given string and assigns it to the AmountStr field.
func (o *Transfer) SetAmountStr(v string) {
	o.AmountStr = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Transfer) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Transfer) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Transfer) SetCurrency(v string) {
	o.Currency = &v
}

// GetTransferId returns the TransferId field value if set, zero value otherwise.
func (o *Transfer) GetTransferId() string {
	if o == nil || IsNil(o.TransferId) {
		var ret string
		return ret
	}
	return *o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetTransferIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransferId) {
		return nil, false
	}
	return o.TransferId, true
}

// HasTransferId returns a boolean if a field has been set.
func (o *Transfer) HasTransferId() bool {
	if o != nil && !IsNil(o.TransferId) {
		return true
	}

	return false
}

// SetTransferId gets a reference to the given string and assigns it to the TransferId field.
func (o *Transfer) SetTransferId(v string) {
	o.TransferId = &v
}

// GetExternalTransferId returns the ExternalTransferId field value if set, zero value otherwise.
func (o *Transfer) GetExternalTransferId() string {
	if o == nil || IsNil(o.ExternalTransferId) {
		var ret string
		return ret
	}
	return *o.ExternalTransferId
}

// GetExternalTransferIdOk returns a tuple with the ExternalTransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetExternalTransferIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalTransferId) {
		return nil, false
	}
	return o.ExternalTransferId, true
}

// HasExternalTransferId returns a boolean if a field has been set.
func (o *Transfer) HasExternalTransferId() bool {
	if o != nil && !IsNil(o.ExternalTransferId) {
		return true
	}

	return false
}

// SetExternalTransferId gets a reference to the given string and assigns it to the ExternalTransferId field.
func (o *Transfer) SetExternalTransferId(v string) {
	o.ExternalTransferId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Transfer) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Transfer) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Transfer) SetStatus(v string) {
	o.Status = &v
}

// GetSubStatus returns the SubStatus field value if set, zero value otherwise.
func (o *Transfer) GetSubStatus() string {
	if o == nil || IsNil(o.SubStatus) {
		var ret string
		return ret
	}
	return *o.SubStatus
}

// GetSubStatusOk returns a tuple with the SubStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetSubStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SubStatus) {
		return nil, false
	}
	return o.SubStatus, true
}

// HasSubStatus returns a boolean if a field has been set.
func (o *Transfer) HasSubStatus() bool {
	if o != nil && !IsNil(o.SubStatus) {
		return true
	}

	return false
}

// SetSubStatus gets a reference to the given string and assigns it to the SubStatus field.
func (o *Transfer) SetSubStatus(v string) {
	o.SubStatus = &v
}

// GetAchId returns the AchId field value if set, zero value otherwise.
func (o *Transfer) GetAchId() string {
	if o == nil || IsNil(o.AchId) {
		var ret string
		return ret
	}
	return *o.AchId
}

// GetAchIdOk returns a tuple with the AchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetAchIdOk() (*string, bool) {
	if o == nil || IsNil(o.AchId) {
		return nil, false
	}
	return o.AchId, true
}

// HasAchId returns a boolean if a field has been set.
func (o *Transfer) HasAchId() bool {
	if o != nil && !IsNil(o.AchId) {
		return true
	}

	return false
}

// SetAchId gets a reference to the given string and assigns it to the AchId field.
func (o *Transfer) SetAchId(v string) {
	o.AchId = &v
}

// GetBankId returns the BankId field value if set, zero value otherwise.
func (o *Transfer) GetBankId() string {
	if o == nil || IsNil(o.BankId) {
		var ret string
		return ret
	}
	return *o.BankId
}

// GetBankIdOk returns a tuple with the BankId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetBankIdOk() (*string, bool) {
	if o == nil || IsNil(o.BankId) {
		return nil, false
	}
	return o.BankId, true
}

// HasBankId returns a boolean if a field has been set.
func (o *Transfer) HasBankId() bool {
	if o != nil && !IsNil(o.BankId) {
		return true
	}

	return false
}

// SetBankId gets a reference to the given string and assigns it to the BankId field.
func (o *Transfer) SetBankId(v string) {
	o.BankId = &v
}

// GetBankType returns the BankType field value if set, zero value otherwise.
func (o *Transfer) GetBankType() string {
	if o == nil || IsNil(o.BankType) {
		var ret string
		return ret
	}
	return *o.BankType
}

// GetBankTypeOk returns a tuple with the BankType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetBankTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BankType) {
		return nil, false
	}
	return o.BankType, true
}

// HasBankType returns a boolean if a field has been set.
func (o *Transfer) HasBankType() bool {
	if o != nil && !IsNil(o.BankType) {
		return true
	}

	return false
}

// SetBankType gets a reference to the given string and assigns it to the BankType field.
func (o *Transfer) SetBankType(v string) {
	o.BankType = &v
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise.
func (o *Transfer) GetBankAccount() string {
	if o == nil || IsNil(o.BankAccount) {
		var ret string
		return ret
	}
	return *o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetBankAccountOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccount) {
		return nil, false
	}
	return o.BankAccount, true
}

// HasBankAccount returns a boolean if a field has been set.
func (o *Transfer) HasBankAccount() bool {
	if o != nil && !IsNil(o.BankAccount) {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given string and assigns it to the BankAccount field.
func (o *Transfer) SetBankAccount(v string) {
	o.BankAccount = &v
}

// GetBankAccountName returns the BankAccountName field value if set, zero value otherwise.
func (o *Transfer) GetBankAccountName() string {
	if o == nil || IsNil(o.BankAccountName) {
		var ret string
		return ret
	}
	return *o.BankAccountName
}

// GetBankAccountNameOk returns a tuple with the BankAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetBankAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountName) {
		return nil, false
	}
	return o.BankAccountName, true
}

// HasBankAccountName returns a boolean if a field has been set.
func (o *Transfer) HasBankAccountName() bool {
	if o != nil && !IsNil(o.BankAccountName) {
		return true
	}

	return false
}

// SetBankAccountName gets a reference to the given string and assigns it to the BankAccountName field.
func (o *Transfer) SetBankAccountName(v string) {
	o.BankAccountName = &v
}

// GetCustomerType returns the CustomerType field value if set, zero value otherwise.
func (o *Transfer) GetCustomerType() string {
	if o == nil || IsNil(o.CustomerType) {
		var ret string
		return ret
	}
	return *o.CustomerType
}

// GetCustomerTypeOk returns a tuple with the CustomerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetCustomerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerType) {
		return nil, false
	}
	return o.CustomerType, true
}

// HasCustomerType returns a boolean if a field has been set.
func (o *Transfer) HasCustomerType() bool {
	if o != nil && !IsNil(o.CustomerType) {
		return true
	}

	return false
}

// SetCustomerType gets a reference to the given string and assigns it to the CustomerType field.
func (o *Transfer) SetCustomerType(v string) {
	o.CustomerType = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *Transfer) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *Transfer) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *Transfer) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *Transfer) GetUpdateTime() string {
	if o == nil || IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *Transfer) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *Transfer) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// GetCreateTimeStr returns the CreateTimeStr field value if set, zero value otherwise.
func (o *Transfer) GetCreateTimeStr() string {
	if o == nil || IsNil(o.CreateTimeStr) {
		var ret string
		return ret
	}
	return *o.CreateTimeStr
}

// GetCreateTimeStrOk returns a tuple with the CreateTimeStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetCreateTimeStrOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTimeStr) {
		return nil, false
	}
	return o.CreateTimeStr, true
}

// HasCreateTimeStr returns a boolean if a field has been set.
func (o *Transfer) HasCreateTimeStr() bool {
	if o != nil && !IsNil(o.CreateTimeStr) {
		return true
	}

	return false
}

// SetCreateTimeStr gets a reference to the given string and assigns it to the CreateTimeStr field.
func (o *Transfer) SetCreateTimeStr(v string) {
	o.CreateTimeStr = &v
}

// GetFinishTime returns the FinishTime field value if set, zero value otherwise.
func (o *Transfer) GetFinishTime() string {
	if o == nil || IsNil(o.FinishTime) {
		var ret string
		return ret
	}
	return *o.FinishTime
}

// GetFinishTimeOk returns a tuple with the FinishTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetFinishTimeOk() (*string, bool) {
	if o == nil || IsNil(o.FinishTime) {
		return nil, false
	}
	return o.FinishTime, true
}

// HasFinishTime returns a boolean if a field has been set.
func (o *Transfer) HasFinishTime() bool {
	if o != nil && !IsNil(o.FinishTime) {
		return true
	}

	return false
}

// SetFinishTime gets a reference to the given string and assigns it to the FinishTime field.
func (o *Transfer) SetFinishTime(v string) {
	o.FinishTime = &v
}

// GetAvailableTime returns the AvailableTime field value if set, zero value otherwise.
func (o *Transfer) GetAvailableTime() string {
	if o == nil || IsNil(o.AvailableTime) {
		var ret string
		return ret
	}
	return *o.AvailableTime
}

// GetAvailableTimeOk returns a tuple with the AvailableTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetAvailableTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AvailableTime) {
		return nil, false
	}
	return o.AvailableTime, true
}

// HasAvailableTime returns a boolean if a field has been set.
func (o *Transfer) HasAvailableTime() bool {
	if o != nil && !IsNil(o.AvailableTime) {
		return true
	}

	return false
}

// SetAvailableTime gets a reference to the given string and assigns it to the AvailableTime field.
func (o *Transfer) SetAvailableTime(v string) {
	o.AvailableTime = &v
}

// GetAvailableLampSendStatus returns the AvailableLampSendStatus field value if set, zero value otherwise.
func (o *Transfer) GetAvailableLampSendStatus() string {
	if o == nil || IsNil(o.AvailableLampSendStatus) {
		var ret string
		return ret
	}
	return *o.AvailableLampSendStatus
}

// GetAvailableLampSendStatusOk returns a tuple with the AvailableLampSendStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetAvailableLampSendStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AvailableLampSendStatus) {
		return nil, false
	}
	return o.AvailableLampSendStatus, true
}

// HasAvailableLampSendStatus returns a boolean if a field has been set.
func (o *Transfer) HasAvailableLampSendStatus() bool {
	if o != nil && !IsNil(o.AvailableLampSendStatus) {
		return true
	}

	return false
}

// SetAvailableLampSendStatus gets a reference to the given string and assigns it to the AvailableLampSendStatus field.
func (o *Transfer) SetAvailableLampSendStatus(v string) {
	o.AvailableLampSendStatus = &v
}

// GetAvailableTimeStr returns the AvailableTimeStr field value if set, zero value otherwise.
func (o *Transfer) GetAvailableTimeStr() string {
	if o == nil || IsNil(o.AvailableTimeStr) {
		var ret string
		return ret
	}
	return *o.AvailableTimeStr
}

// GetAvailableTimeStrOk returns a tuple with the AvailableTimeStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetAvailableTimeStrOk() (*string, bool) {
	if o == nil || IsNil(o.AvailableTimeStr) {
		return nil, false
	}
	return o.AvailableTimeStr, true
}

// HasAvailableTimeStr returns a boolean if a field has been set.
func (o *Transfer) HasAvailableTimeStr() bool {
	if o != nil && !IsNil(o.AvailableTimeStr) {
		return true
	}

	return false
}

// SetAvailableTimeStr gets a reference to the given string and assigns it to the AvailableTimeStr field.
func (o *Transfer) SetAvailableTimeStr(v string) {
	o.AvailableTimeStr = &v
}

// GetReturnSendStatus returns the ReturnSendStatus field value if set, zero value otherwise.
func (o *Transfer) GetReturnSendStatus() string {
	if o == nil || IsNil(o.ReturnSendStatus) {
		var ret string
		return ret
	}
	return *o.ReturnSendStatus
}

// GetReturnSendStatusOk returns a tuple with the ReturnSendStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetReturnSendStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnSendStatus) {
		return nil, false
	}
	return o.ReturnSendStatus, true
}

// HasReturnSendStatus returns a boolean if a field has been set.
func (o *Transfer) HasReturnSendStatus() bool {
	if o != nil && !IsNil(o.ReturnSendStatus) {
		return true
	}

	return false
}

// SetReturnSendStatus gets a reference to the given string and assigns it to the ReturnSendStatus field.
func (o *Transfer) SetReturnSendStatus(v string) {
	o.ReturnSendStatus = &v
}

// GetRecordSendStatus returns the RecordSendStatus field value if set, zero value otherwise.
func (o *Transfer) GetRecordSendStatus() string {
	if o == nil || IsNil(o.RecordSendStatus) {
		var ret string
		return ret
	}
	return *o.RecordSendStatus
}

// GetRecordSendStatusOk returns a tuple with the RecordSendStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetRecordSendStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RecordSendStatus) {
		return nil, false
	}
	return o.RecordSendStatus, true
}

// HasRecordSendStatus returns a boolean if a field has been set.
func (o *Transfer) HasRecordSendStatus() bool {
	if o != nil && !IsNil(o.RecordSendStatus) {
		return true
	}

	return false
}

// SetRecordSendStatus gets a reference to the given string and assigns it to the RecordSendStatus field.
func (o *Transfer) SetRecordSendStatus(v string) {
	o.RecordSendStatus = &v
}

// GetSerialId returns the SerialId field value if set, zero value otherwise.
func (o *Transfer) GetSerialId() string {
	if o == nil || IsNil(o.SerialId) {
		var ret string
		return ret
	}
	return *o.SerialId
}

// GetSerialIdOk returns a tuple with the SerialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetSerialIdOk() (*string, bool) {
	if o == nil || IsNil(o.SerialId) {
		return nil, false
	}
	return o.SerialId, true
}

// HasSerialId returns a boolean if a field has been set.
func (o *Transfer) HasSerialId() bool {
	if o != nil && !IsNil(o.SerialId) {
		return true
	}

	return false
}

// SetSerialId gets a reference to the given string and assigns it to the SerialId field.
func (o *Transfer) SetSerialId(v string) {
	o.SerialId = &v
}

// GetArriveTime returns the ArriveTime field value if set, zero value otherwise.
func (o *Transfer) GetArriveTime() string {
	if o == nil || IsNil(o.ArriveTime) {
		var ret string
		return ret
	}
	return *o.ArriveTime
}

// GetArriveTimeOk returns a tuple with the ArriveTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetArriveTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ArriveTime) {
		return nil, false
	}
	return o.ArriveTime, true
}

// HasArriveTime returns a boolean if a field has been set.
func (o *Transfer) HasArriveTime() bool {
	if o != nil && !IsNil(o.ArriveTime) {
		return true
	}

	return false
}

// SetArriveTime gets a reference to the given string and assigns it to the ArriveTime field.
func (o *Transfer) SetArriveTime(v string) {
	o.ArriveTime = &v
}

// GetFirstGift returns the FirstGift field value if set, zero value otherwise.
func (o *Transfer) GetFirstGift() bool {
	if o == nil || IsNil(o.FirstGift) {
		var ret bool
		return ret
	}
	return *o.FirstGift
}

// GetFirstGiftOk returns a tuple with the FirstGift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetFirstGiftOk() (*bool, bool) {
	if o == nil || IsNil(o.FirstGift) {
		return nil, false
	}
	return o.FirstGift, true
}

// HasFirstGift returns a boolean if a field has been set.
func (o *Transfer) HasFirstGift() bool {
	if o != nil && !IsNil(o.FirstGift) {
		return true
	}

	return false
}

// SetFirstGift gets a reference to the given bool and assigns it to the FirstGift field.
func (o *Transfer) SetFirstGift(v bool) {
	o.FirstGift = &v
}

// GetWaitCardBinding returns the WaitCardBinding field value if set, zero value otherwise.
func (o *Transfer) GetWaitCardBinding() bool {
	if o == nil || IsNil(o.WaitCardBinding) {
		var ret bool
		return ret
	}
	return *o.WaitCardBinding
}

// GetWaitCardBindingOk returns a tuple with the WaitCardBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetWaitCardBindingOk() (*bool, bool) {
	if o == nil || IsNil(o.WaitCardBinding) {
		return nil, false
	}
	return o.WaitCardBinding, true
}

// HasWaitCardBinding returns a boolean if a field has been set.
func (o *Transfer) HasWaitCardBinding() bool {
	if o != nil && !IsNil(o.WaitCardBinding) {
		return true
	}

	return false
}

// SetWaitCardBinding gets a reference to the given bool and assigns it to the WaitCardBinding field.
func (o *Transfer) SetWaitCardBinding(v bool) {
	o.WaitCardBinding = &v
}

// GetTipsInfo returns the TipsInfo field value if set, zero value otherwise.
func (o *Transfer) GetTipsInfo() string {
	if o == nil || IsNil(o.TipsInfo) {
		var ret string
		return ret
	}
	return *o.TipsInfo
}

// GetTipsInfoOk returns a tuple with the TipsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetTipsInfoOk() (*string, bool) {
	if o == nil || IsNil(o.TipsInfo) {
		return nil, false
	}
	return o.TipsInfo, true
}

// HasTipsInfo returns a boolean if a field has been set.
func (o *Transfer) HasTipsInfo() bool {
	if o != nil && !IsNil(o.TipsInfo) {
		return true
	}

	return false
}

// SetTipsInfo gets a reference to the given string and assigns it to the TipsInfo field.
func (o *Transfer) SetTipsInfo(v string) {
	o.TipsInfo = &v
}

func (o Transfer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transfer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SecAccountId) {
		toSerialize["secAccountId"] = o.SecAccountId
	}
	if !IsNil(o.BrokerAccountId) {
		toSerialize["brokerAccountId"] = o.BrokerAccountId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.AmountStr) {
		toSerialize["amountStr"] = o.AmountStr
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.TransferId) {
		toSerialize["transferId"] = o.TransferId
	}
	if !IsNil(o.ExternalTransferId) {
		toSerialize["externalTransferId"] = o.ExternalTransferId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubStatus) {
		toSerialize["subStatus"] = o.SubStatus
	}
	if !IsNil(o.AchId) {
		toSerialize["achId"] = o.AchId
	}
	if !IsNil(o.BankId) {
		toSerialize["bankId"] = o.BankId
	}
	if !IsNil(o.BankType) {
		toSerialize["bankType"] = o.BankType
	}
	if !IsNil(o.BankAccount) {
		toSerialize["bankAccount"] = o.BankAccount
	}
	if !IsNil(o.BankAccountName) {
		toSerialize["bankAccountName"] = o.BankAccountName
	}
	if !IsNil(o.CustomerType) {
		toSerialize["customerType"] = o.CustomerType
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !IsNil(o.CreateTimeStr) {
		toSerialize["createTimeStr"] = o.CreateTimeStr
	}
	if !IsNil(o.FinishTime) {
		toSerialize["finishTime"] = o.FinishTime
	}
	if !IsNil(o.AvailableTime) {
		toSerialize["availableTime"] = o.AvailableTime
	}
	if !IsNil(o.AvailableLampSendStatus) {
		toSerialize["availableLampSendStatus"] = o.AvailableLampSendStatus
	}
	if !IsNil(o.AvailableTimeStr) {
		toSerialize["availableTimeStr"] = o.AvailableTimeStr
	}
	if !IsNil(o.ReturnSendStatus) {
		toSerialize["returnSendStatus"] = o.ReturnSendStatus
	}
	if !IsNil(o.RecordSendStatus) {
		toSerialize["recordSendStatus"] = o.RecordSendStatus
	}
	if !IsNil(o.SerialId) {
		toSerialize["serialId"] = o.SerialId
	}
	if !IsNil(o.ArriveTime) {
		toSerialize["arriveTime"] = o.ArriveTime
	}
	if !IsNil(o.FirstGift) {
		toSerialize["firstGift"] = o.FirstGift
	}
	if !IsNil(o.WaitCardBinding) {
		toSerialize["waitCardBinding"] = o.WaitCardBinding
	}
	if !IsNil(o.TipsInfo) {
		toSerialize["tipsInfo"] = o.TipsInfo
	}
	return toSerialize, nil
}

type NullableTransfer struct {
	value *Transfer
	isSet bool
}

func (v NullableTransfer) Get() *Transfer {
	return v.value
}

func (v *NullableTransfer) Set(val *Transfer) {
	v.value = val
	v.isSet = true
}

func (v NullableTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransfer(val *Transfer) *NullableTransfer {
	return &NullableTransfer{value: val, isSet: true}
}

func (v NullableTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


