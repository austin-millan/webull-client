# Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**SecAccountId** | Pointer to **float32** |  | [optional] 
**BrokerAccountId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 
**AmountStr** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**TransferId** | Pointer to **string** |  | [optional] 
**ExternalTransferId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubStatus** | Pointer to **string** |  | [optional] 
**AchId** | Pointer to **string** |  | [optional] 
**BankId** | Pointer to **string** |  | [optional] 
**BankType** | Pointer to **string** |  | [optional] 
**BankAccount** | Pointer to **string** |  | [optional] 
**BankAccountName** | Pointer to **string** |  | [optional] 
**CustomerType** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **string** |  | [optional] 
**CreateTimeStr** | Pointer to **string** |  | [optional] 
**FinishTime** | Pointer to **string** |  | [optional] 
**AvailableTime** | Pointer to **string** |  | [optional] 
**AvailableLampSendStatus** | Pointer to **string** |  | [optional] 
**AvailableTimeStr** | Pointer to **string** |  | [optional] 
**ReturnSendStatus** | Pointer to **string** |  | [optional] 
**RecordSendStatus** | Pointer to **string** |  | [optional] 
**SerialId** | Pointer to **string** |  | [optional] 
**ArriveTime** | Pointer to **string** |  | [optional] 
**FirstGift** | Pointer to **bool** |  | [optional] 
**WaitCardBinding** | Pointer to **bool** |  | [optional] 
**TipsInfo** | Pointer to **string** |  | [optional] 

## Methods

### NewTransfer

`func NewTransfer() *Transfer`

NewTransfer instantiates a new Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWithDefaults

`func NewTransferWithDefaults() *Transfer`

NewTransferWithDefaults instantiates a new Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transfer) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transfer) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transfer) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *Transfer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSecAccountId

`func (o *Transfer) GetSecAccountId() float32`

GetSecAccountId returns the SecAccountId field if non-nil, zero value otherwise.

### GetSecAccountIdOk

`func (o *Transfer) GetSecAccountIdOk() (*float32, bool)`

GetSecAccountIdOk returns a tuple with the SecAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecAccountId

`func (o *Transfer) SetSecAccountId(v float32)`

SetSecAccountId sets SecAccountId field to given value.

### HasSecAccountId

`func (o *Transfer) HasSecAccountId() bool`

HasSecAccountId returns a boolean if a field has been set.

### GetBrokerAccountId

`func (o *Transfer) GetBrokerAccountId() string`

GetBrokerAccountId returns the BrokerAccountId field if non-nil, zero value otherwise.

### GetBrokerAccountIdOk

`func (o *Transfer) GetBrokerAccountIdOk() (*string, bool)`

GetBrokerAccountIdOk returns a tuple with the BrokerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerAccountId

`func (o *Transfer) SetBrokerAccountId(v string)`

SetBrokerAccountId sets BrokerAccountId field to given value.

### HasBrokerAccountId

`func (o *Transfer) HasBrokerAccountId() bool`

HasBrokerAccountId returns a boolean if a field has been set.

### GetType

`func (o *Transfer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transfer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transfer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Transfer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDirection

`func (o *Transfer) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Transfer) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Transfer) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *Transfer) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetAmount

`func (o *Transfer) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transfer) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transfer) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Transfer) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountStr

`func (o *Transfer) GetAmountStr() string`

GetAmountStr returns the AmountStr field if non-nil, zero value otherwise.

### GetAmountStrOk

`func (o *Transfer) GetAmountStrOk() (*string, bool)`

GetAmountStrOk returns a tuple with the AmountStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountStr

`func (o *Transfer) SetAmountStr(v string)`

SetAmountStr sets AmountStr field to given value.

### HasAmountStr

`func (o *Transfer) HasAmountStr() bool`

HasAmountStr returns a boolean if a field has been set.

### GetCurrency

`func (o *Transfer) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Transfer) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Transfer) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Transfer) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTransferId

`func (o *Transfer) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *Transfer) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *Transfer) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.

### HasTransferId

`func (o *Transfer) HasTransferId() bool`

HasTransferId returns a boolean if a field has been set.

### GetExternalTransferId

`func (o *Transfer) GetExternalTransferId() string`

GetExternalTransferId returns the ExternalTransferId field if non-nil, zero value otherwise.

### GetExternalTransferIdOk

`func (o *Transfer) GetExternalTransferIdOk() (*string, bool)`

GetExternalTransferIdOk returns a tuple with the ExternalTransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTransferId

`func (o *Transfer) SetExternalTransferId(v string)`

SetExternalTransferId sets ExternalTransferId field to given value.

### HasExternalTransferId

`func (o *Transfer) HasExternalTransferId() bool`

HasExternalTransferId returns a boolean if a field has been set.

### GetStatus

`func (o *Transfer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transfer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transfer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Transfer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubStatus

`func (o *Transfer) GetSubStatus() string`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *Transfer) GetSubStatusOk() (*string, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *Transfer) SetSubStatus(v string)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *Transfer) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### GetAchId

`func (o *Transfer) GetAchId() string`

GetAchId returns the AchId field if non-nil, zero value otherwise.

### GetAchIdOk

`func (o *Transfer) GetAchIdOk() (*string, bool)`

GetAchIdOk returns a tuple with the AchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchId

`func (o *Transfer) SetAchId(v string)`

SetAchId sets AchId field to given value.

### HasAchId

`func (o *Transfer) HasAchId() bool`

HasAchId returns a boolean if a field has been set.

### GetBankId

`func (o *Transfer) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *Transfer) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *Transfer) SetBankId(v string)`

SetBankId sets BankId field to given value.

### HasBankId

`func (o *Transfer) HasBankId() bool`

HasBankId returns a boolean if a field has been set.

### GetBankType

`func (o *Transfer) GetBankType() string`

GetBankType returns the BankType field if non-nil, zero value otherwise.

### GetBankTypeOk

`func (o *Transfer) GetBankTypeOk() (*string, bool)`

GetBankTypeOk returns a tuple with the BankType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankType

`func (o *Transfer) SetBankType(v string)`

SetBankType sets BankType field to given value.

### HasBankType

`func (o *Transfer) HasBankType() bool`

HasBankType returns a boolean if a field has been set.

### GetBankAccount

`func (o *Transfer) GetBankAccount() string`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *Transfer) GetBankAccountOk() (*string, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *Transfer) SetBankAccount(v string)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *Transfer) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### GetBankAccountName

`func (o *Transfer) GetBankAccountName() string`

GetBankAccountName returns the BankAccountName field if non-nil, zero value otherwise.

### GetBankAccountNameOk

`func (o *Transfer) GetBankAccountNameOk() (*string, bool)`

GetBankAccountNameOk returns a tuple with the BankAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountName

`func (o *Transfer) SetBankAccountName(v string)`

SetBankAccountName sets BankAccountName field to given value.

### HasBankAccountName

`func (o *Transfer) HasBankAccountName() bool`

HasBankAccountName returns a boolean if a field has been set.

### GetCustomerType

`func (o *Transfer) GetCustomerType() string`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *Transfer) GetCustomerTypeOk() (*string, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *Transfer) SetCustomerType(v string)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *Transfer) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### GetCreateTime

`func (o *Transfer) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Transfer) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Transfer) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Transfer) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *Transfer) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Transfer) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Transfer) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *Transfer) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetCreateTimeStr

`func (o *Transfer) GetCreateTimeStr() string`

GetCreateTimeStr returns the CreateTimeStr field if non-nil, zero value otherwise.

### GetCreateTimeStrOk

`func (o *Transfer) GetCreateTimeStrOk() (*string, bool)`

GetCreateTimeStrOk returns a tuple with the CreateTimeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStr

`func (o *Transfer) SetCreateTimeStr(v string)`

SetCreateTimeStr sets CreateTimeStr field to given value.

### HasCreateTimeStr

`func (o *Transfer) HasCreateTimeStr() bool`

HasCreateTimeStr returns a boolean if a field has been set.

### GetFinishTime

`func (o *Transfer) GetFinishTime() string`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *Transfer) GetFinishTimeOk() (*string, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *Transfer) SetFinishTime(v string)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *Transfer) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetAvailableTime

`func (o *Transfer) GetAvailableTime() string`

GetAvailableTime returns the AvailableTime field if non-nil, zero value otherwise.

### GetAvailableTimeOk

`func (o *Transfer) GetAvailableTimeOk() (*string, bool)`

GetAvailableTimeOk returns a tuple with the AvailableTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTime

`func (o *Transfer) SetAvailableTime(v string)`

SetAvailableTime sets AvailableTime field to given value.

### HasAvailableTime

`func (o *Transfer) HasAvailableTime() bool`

HasAvailableTime returns a boolean if a field has been set.

### GetAvailableLampSendStatus

`func (o *Transfer) GetAvailableLampSendStatus() string`

GetAvailableLampSendStatus returns the AvailableLampSendStatus field if non-nil, zero value otherwise.

### GetAvailableLampSendStatusOk

`func (o *Transfer) GetAvailableLampSendStatusOk() (*string, bool)`

GetAvailableLampSendStatusOk returns a tuple with the AvailableLampSendStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableLampSendStatus

`func (o *Transfer) SetAvailableLampSendStatus(v string)`

SetAvailableLampSendStatus sets AvailableLampSendStatus field to given value.

### HasAvailableLampSendStatus

`func (o *Transfer) HasAvailableLampSendStatus() bool`

HasAvailableLampSendStatus returns a boolean if a field has been set.

### GetAvailableTimeStr

`func (o *Transfer) GetAvailableTimeStr() string`

GetAvailableTimeStr returns the AvailableTimeStr field if non-nil, zero value otherwise.

### GetAvailableTimeStrOk

`func (o *Transfer) GetAvailableTimeStrOk() (*string, bool)`

GetAvailableTimeStrOk returns a tuple with the AvailableTimeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTimeStr

`func (o *Transfer) SetAvailableTimeStr(v string)`

SetAvailableTimeStr sets AvailableTimeStr field to given value.

### HasAvailableTimeStr

`func (o *Transfer) HasAvailableTimeStr() bool`

HasAvailableTimeStr returns a boolean if a field has been set.

### GetReturnSendStatus

`func (o *Transfer) GetReturnSendStatus() string`

GetReturnSendStatus returns the ReturnSendStatus field if non-nil, zero value otherwise.

### GetReturnSendStatusOk

`func (o *Transfer) GetReturnSendStatusOk() (*string, bool)`

GetReturnSendStatusOk returns a tuple with the ReturnSendStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnSendStatus

`func (o *Transfer) SetReturnSendStatus(v string)`

SetReturnSendStatus sets ReturnSendStatus field to given value.

### HasReturnSendStatus

`func (o *Transfer) HasReturnSendStatus() bool`

HasReturnSendStatus returns a boolean if a field has been set.

### GetRecordSendStatus

`func (o *Transfer) GetRecordSendStatus() string`

GetRecordSendStatus returns the RecordSendStatus field if non-nil, zero value otherwise.

### GetRecordSendStatusOk

`func (o *Transfer) GetRecordSendStatusOk() (*string, bool)`

GetRecordSendStatusOk returns a tuple with the RecordSendStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordSendStatus

`func (o *Transfer) SetRecordSendStatus(v string)`

SetRecordSendStatus sets RecordSendStatus field to given value.

### HasRecordSendStatus

`func (o *Transfer) HasRecordSendStatus() bool`

HasRecordSendStatus returns a boolean if a field has been set.

### GetSerialId

`func (o *Transfer) GetSerialId() string`

GetSerialId returns the SerialId field if non-nil, zero value otherwise.

### GetSerialIdOk

`func (o *Transfer) GetSerialIdOk() (*string, bool)`

GetSerialIdOk returns a tuple with the SerialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialId

`func (o *Transfer) SetSerialId(v string)`

SetSerialId sets SerialId field to given value.

### HasSerialId

`func (o *Transfer) HasSerialId() bool`

HasSerialId returns a boolean if a field has been set.

### GetArriveTime

`func (o *Transfer) GetArriveTime() string`

GetArriveTime returns the ArriveTime field if non-nil, zero value otherwise.

### GetArriveTimeOk

`func (o *Transfer) GetArriveTimeOk() (*string, bool)`

GetArriveTimeOk returns a tuple with the ArriveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArriveTime

`func (o *Transfer) SetArriveTime(v string)`

SetArriveTime sets ArriveTime field to given value.

### HasArriveTime

`func (o *Transfer) HasArriveTime() bool`

HasArriveTime returns a boolean if a field has been set.

### GetFirstGift

`func (o *Transfer) GetFirstGift() bool`

GetFirstGift returns the FirstGift field if non-nil, zero value otherwise.

### GetFirstGiftOk

`func (o *Transfer) GetFirstGiftOk() (*bool, bool)`

GetFirstGiftOk returns a tuple with the FirstGift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstGift

`func (o *Transfer) SetFirstGift(v bool)`

SetFirstGift sets FirstGift field to given value.

### HasFirstGift

`func (o *Transfer) HasFirstGift() bool`

HasFirstGift returns a boolean if a field has been set.

### GetWaitCardBinding

`func (o *Transfer) GetWaitCardBinding() bool`

GetWaitCardBinding returns the WaitCardBinding field if non-nil, zero value otherwise.

### GetWaitCardBindingOk

`func (o *Transfer) GetWaitCardBindingOk() (*bool, bool)`

GetWaitCardBindingOk returns a tuple with the WaitCardBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitCardBinding

`func (o *Transfer) SetWaitCardBinding(v bool)`

SetWaitCardBinding sets WaitCardBinding field to given value.

### HasWaitCardBinding

`func (o *Transfer) HasWaitCardBinding() bool`

HasWaitCardBinding returns a boolean if a field has been set.

### GetTipsInfo

`func (o *Transfer) GetTipsInfo() string`

GetTipsInfo returns the TipsInfo field if non-nil, zero value otherwise.

### GetTipsInfoOk

`func (o *Transfer) GetTipsInfoOk() (*string, bool)`

GetTipsInfoOk returns a tuple with the TipsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipsInfo

`func (o *Transfer) SetTipsInfo(v string)`

SetTipsInfo sets TipsInfo field to given value.

### HasTipsInfo

`func (o *Transfer) HasTipsInfo() bool`

HasTipsInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


