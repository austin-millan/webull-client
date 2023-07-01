# AccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | Pointer to [**AccountType**](AccountType.md) |  | [optional] [default to _2]
**Currency** | Pointer to **string** |  | [optional] 
**NetLiquidation** | Pointer to **float32** |  | [optional] 
**Pdt** | Pointer to **bool** |  | [optional] 
**UnrealizedProfitLoss** | Pointer to **float32** |  | [optional] 
**UnrealizedProfitLossBase** | Pointer to **float32** |  | [optional] 
**Warning** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccountInfo

`func NewAccountInfo() *AccountInfo`

NewAccountInfo instantiates a new AccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInfoWithDefaults

`func NewAccountInfoWithDefaults() *AccountInfo`

NewAccountInfoWithDefaults instantiates a new AccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *AccountInfo) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountInfo) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountInfo) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountInfo) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetCurrency

`func (o *AccountInfo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountInfo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountInfo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AccountInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetNetLiquidation

`func (o *AccountInfo) GetNetLiquidation() float32`

GetNetLiquidation returns the NetLiquidation field if non-nil, zero value otherwise.

### GetNetLiquidationOk

`func (o *AccountInfo) GetNetLiquidationOk() (*float32, bool)`

GetNetLiquidationOk returns a tuple with the NetLiquidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetLiquidation

`func (o *AccountInfo) SetNetLiquidation(v float32)`

SetNetLiquidation sets NetLiquidation field to given value.

### HasNetLiquidation

`func (o *AccountInfo) HasNetLiquidation() bool`

HasNetLiquidation returns a boolean if a field has been set.

### GetPdt

`func (o *AccountInfo) GetPdt() bool`

GetPdt returns the Pdt field if non-nil, zero value otherwise.

### GetPdtOk

`func (o *AccountInfo) GetPdtOk() (*bool, bool)`

GetPdtOk returns a tuple with the Pdt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdt

`func (o *AccountInfo) SetPdt(v bool)`

SetPdt sets Pdt field to given value.

### HasPdt

`func (o *AccountInfo) HasPdt() bool`

HasPdt returns a boolean if a field has been set.

### GetUnrealizedProfitLoss

`func (o *AccountInfo) GetUnrealizedProfitLoss() float32`

GetUnrealizedProfitLoss returns the UnrealizedProfitLoss field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossOk

`func (o *AccountInfo) GetUnrealizedProfitLossOk() (*float32, bool)`

GetUnrealizedProfitLossOk returns a tuple with the UnrealizedProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLoss

`func (o *AccountInfo) SetUnrealizedProfitLoss(v float32)`

SetUnrealizedProfitLoss sets UnrealizedProfitLoss field to given value.

### HasUnrealizedProfitLoss

`func (o *AccountInfo) HasUnrealizedProfitLoss() bool`

HasUnrealizedProfitLoss returns a boolean if a field has been set.

### GetUnrealizedProfitLossBase

`func (o *AccountInfo) GetUnrealizedProfitLossBase() float32`

GetUnrealizedProfitLossBase returns the UnrealizedProfitLossBase field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossBaseOk

`func (o *AccountInfo) GetUnrealizedProfitLossBaseOk() (*float32, bool)`

GetUnrealizedProfitLossBaseOk returns a tuple with the UnrealizedProfitLossBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLossBase

`func (o *AccountInfo) SetUnrealizedProfitLossBase(v float32)`

SetUnrealizedProfitLossBase sets UnrealizedProfitLossBase field to given value.

### HasUnrealizedProfitLossBase

`func (o *AccountInfo) HasUnrealizedProfitLossBase() bool`

HasUnrealizedProfitLossBase returns a boolean if a field has been set.

### GetWarning

`func (o *AccountInfo) GetWarning() bool`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *AccountInfo) GetWarningOk() (*bool, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *AccountInfo) SetWarning(v bool)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *AccountInfo) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


