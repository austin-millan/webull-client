# GetSecurityAccountsResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountTypes** | Pointer to **[]string** |  | [optional] 
**AllowDeposit** | Pointer to **bool** |  | [optional] 
**BrokerAccountId** | Pointer to **string** |  | [optional] 
**BrokerId** | Pointer to **int32** |  | [optional] 
**BrokerName** | Pointer to **string** |  | [optional] 
**ComboTypes** | Pointer to **[]string** |  | [optional] 
**CustomerType** | Pointer to **string** |  | [optional] 
**Deposit** | Pointer to **bool** |  | [optional] 
**DepositStatus** | Pointer to **string** |  | [optional] 
**GiftStockStatus** | Pointer to **int32** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**IsDefaultChecked** | Pointer to **bool** |  | [optional] 
**OpenAccountUrl** | Pointer to **string** |  | [optional] 
**OptionOpenStatus** | Pointer to **string** |  | [optional] 
**RegisterRegionId** | Pointer to **int32** |  | [optional] 
**RegisterTradeUrl** | Pointer to **string** |  | [optional] 
**SecAccountId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SupportClickIPO** | Pointer to **bool** |  | [optional] 
**SupportOpenOption** | Pointer to **bool** |  | [optional] 
**SupportOutsideRth** | Pointer to **bool** |  | [optional] 
**TickerTypes** | Pointer to [**[]GetSecurityAccountsResponseDataInnerTickerTypesInner**](GetSecurityAccountsResponseDataInnerTickerTypesInner.md) |  | [optional] 
**TimeInForces** | Pointer to [**[]GetSecurityAccountsResponseDataInnerTimeInForcesInner**](GetSecurityAccountsResponseDataInnerTimeInForcesInner.md) |  | [optional] 
**UserTradePermissionVOs** | Pointer to [**[]GetSecurityAccountsResponseDataInnerUserTradePermissionVOsInner**](GetSecurityAccountsResponseDataInnerUserTradePermissionVOsInner.md) |  | [optional] 

## Methods

### NewGetSecurityAccountsResponseDataInner

`func NewGetSecurityAccountsResponseDataInner() *GetSecurityAccountsResponseDataInner`

NewGetSecurityAccountsResponseDataInner instantiates a new GetSecurityAccountsResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecurityAccountsResponseDataInnerWithDefaults

`func NewGetSecurityAccountsResponseDataInnerWithDefaults() *GetSecurityAccountsResponseDataInner`

NewGetSecurityAccountsResponseDataInnerWithDefaults instantiates a new GetSecurityAccountsResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountTypes

`func (o *GetSecurityAccountsResponseDataInner) GetAccountTypes() []string`

GetAccountTypes returns the AccountTypes field if non-nil, zero value otherwise.

### GetAccountTypesOk

`func (o *GetSecurityAccountsResponseDataInner) GetAccountTypesOk() (*[]string, bool)`

GetAccountTypesOk returns a tuple with the AccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypes

`func (o *GetSecurityAccountsResponseDataInner) SetAccountTypes(v []string)`

SetAccountTypes sets AccountTypes field to given value.

### HasAccountTypes

`func (o *GetSecurityAccountsResponseDataInner) HasAccountTypes() bool`

HasAccountTypes returns a boolean if a field has been set.

### GetAllowDeposit

`func (o *GetSecurityAccountsResponseDataInner) GetAllowDeposit() bool`

GetAllowDeposit returns the AllowDeposit field if non-nil, zero value otherwise.

### GetAllowDepositOk

`func (o *GetSecurityAccountsResponseDataInner) GetAllowDepositOk() (*bool, bool)`

GetAllowDepositOk returns a tuple with the AllowDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDeposit

`func (o *GetSecurityAccountsResponseDataInner) SetAllowDeposit(v bool)`

SetAllowDeposit sets AllowDeposit field to given value.

### HasAllowDeposit

`func (o *GetSecurityAccountsResponseDataInner) HasAllowDeposit() bool`

HasAllowDeposit returns a boolean if a field has been set.

### GetBrokerAccountId

`func (o *GetSecurityAccountsResponseDataInner) GetBrokerAccountId() string`

GetBrokerAccountId returns the BrokerAccountId field if non-nil, zero value otherwise.

### GetBrokerAccountIdOk

`func (o *GetSecurityAccountsResponseDataInner) GetBrokerAccountIdOk() (*string, bool)`

GetBrokerAccountIdOk returns a tuple with the BrokerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerAccountId

`func (o *GetSecurityAccountsResponseDataInner) SetBrokerAccountId(v string)`

SetBrokerAccountId sets BrokerAccountId field to given value.

### HasBrokerAccountId

`func (o *GetSecurityAccountsResponseDataInner) HasBrokerAccountId() bool`

HasBrokerAccountId returns a boolean if a field has been set.

### GetBrokerId

`func (o *GetSecurityAccountsResponseDataInner) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *GetSecurityAccountsResponseDataInner) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *GetSecurityAccountsResponseDataInner) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *GetSecurityAccountsResponseDataInner) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetBrokerName

`func (o *GetSecurityAccountsResponseDataInner) GetBrokerName() string`

GetBrokerName returns the BrokerName field if non-nil, zero value otherwise.

### GetBrokerNameOk

`func (o *GetSecurityAccountsResponseDataInner) GetBrokerNameOk() (*string, bool)`

GetBrokerNameOk returns a tuple with the BrokerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerName

`func (o *GetSecurityAccountsResponseDataInner) SetBrokerName(v string)`

SetBrokerName sets BrokerName field to given value.

### HasBrokerName

`func (o *GetSecurityAccountsResponseDataInner) HasBrokerName() bool`

HasBrokerName returns a boolean if a field has been set.

### GetComboTypes

`func (o *GetSecurityAccountsResponseDataInner) GetComboTypes() []string`

GetComboTypes returns the ComboTypes field if non-nil, zero value otherwise.

### GetComboTypesOk

`func (o *GetSecurityAccountsResponseDataInner) GetComboTypesOk() (*[]string, bool)`

GetComboTypesOk returns a tuple with the ComboTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComboTypes

`func (o *GetSecurityAccountsResponseDataInner) SetComboTypes(v []string)`

SetComboTypes sets ComboTypes field to given value.

### HasComboTypes

`func (o *GetSecurityAccountsResponseDataInner) HasComboTypes() bool`

HasComboTypes returns a boolean if a field has been set.

### GetCustomerType

`func (o *GetSecurityAccountsResponseDataInner) GetCustomerType() string`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *GetSecurityAccountsResponseDataInner) GetCustomerTypeOk() (*string, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *GetSecurityAccountsResponseDataInner) SetCustomerType(v string)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *GetSecurityAccountsResponseDataInner) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### GetDeposit

`func (o *GetSecurityAccountsResponseDataInner) GetDeposit() bool`

GetDeposit returns the Deposit field if non-nil, zero value otherwise.

### GetDepositOk

`func (o *GetSecurityAccountsResponseDataInner) GetDepositOk() (*bool, bool)`

GetDepositOk returns a tuple with the Deposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposit

`func (o *GetSecurityAccountsResponseDataInner) SetDeposit(v bool)`

SetDeposit sets Deposit field to given value.

### HasDeposit

`func (o *GetSecurityAccountsResponseDataInner) HasDeposit() bool`

HasDeposit returns a boolean if a field has been set.

### GetDepositStatus

`func (o *GetSecurityAccountsResponseDataInner) GetDepositStatus() string`

GetDepositStatus returns the DepositStatus field if non-nil, zero value otherwise.

### GetDepositStatusOk

`func (o *GetSecurityAccountsResponseDataInner) GetDepositStatusOk() (*string, bool)`

GetDepositStatusOk returns a tuple with the DepositStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositStatus

`func (o *GetSecurityAccountsResponseDataInner) SetDepositStatus(v string)`

SetDepositStatus sets DepositStatus field to given value.

### HasDepositStatus

`func (o *GetSecurityAccountsResponseDataInner) HasDepositStatus() bool`

HasDepositStatus returns a boolean if a field has been set.

### GetGiftStockStatus

`func (o *GetSecurityAccountsResponseDataInner) GetGiftStockStatus() int32`

GetGiftStockStatus returns the GiftStockStatus field if non-nil, zero value otherwise.

### GetGiftStockStatusOk

`func (o *GetSecurityAccountsResponseDataInner) GetGiftStockStatusOk() (*int32, bool)`

GetGiftStockStatusOk returns a tuple with the GiftStockStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftStockStatus

`func (o *GetSecurityAccountsResponseDataInner) SetGiftStockStatus(v int32)`

SetGiftStockStatus sets GiftStockStatus field to given value.

### HasGiftStockStatus

`func (o *GetSecurityAccountsResponseDataInner) HasGiftStockStatus() bool`

HasGiftStockStatus returns a boolean if a field has been set.

### GetIsDefault

`func (o *GetSecurityAccountsResponseDataInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GetSecurityAccountsResponseDataInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GetSecurityAccountsResponseDataInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GetSecurityAccountsResponseDataInner) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsDefaultChecked

`func (o *GetSecurityAccountsResponseDataInner) GetIsDefaultChecked() bool`

GetIsDefaultChecked returns the IsDefaultChecked field if non-nil, zero value otherwise.

### GetIsDefaultCheckedOk

`func (o *GetSecurityAccountsResponseDataInner) GetIsDefaultCheckedOk() (*bool, bool)`

GetIsDefaultCheckedOk returns a tuple with the IsDefaultChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultChecked

`func (o *GetSecurityAccountsResponseDataInner) SetIsDefaultChecked(v bool)`

SetIsDefaultChecked sets IsDefaultChecked field to given value.

### HasIsDefaultChecked

`func (o *GetSecurityAccountsResponseDataInner) HasIsDefaultChecked() bool`

HasIsDefaultChecked returns a boolean if a field has been set.

### GetOpenAccountUrl

`func (o *GetSecurityAccountsResponseDataInner) GetOpenAccountUrl() string`

GetOpenAccountUrl returns the OpenAccountUrl field if non-nil, zero value otherwise.

### GetOpenAccountUrlOk

`func (o *GetSecurityAccountsResponseDataInner) GetOpenAccountUrlOk() (*string, bool)`

GetOpenAccountUrlOk returns a tuple with the OpenAccountUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAccountUrl

`func (o *GetSecurityAccountsResponseDataInner) SetOpenAccountUrl(v string)`

SetOpenAccountUrl sets OpenAccountUrl field to given value.

### HasOpenAccountUrl

`func (o *GetSecurityAccountsResponseDataInner) HasOpenAccountUrl() bool`

HasOpenAccountUrl returns a boolean if a field has been set.

### GetOptionOpenStatus

`func (o *GetSecurityAccountsResponseDataInner) GetOptionOpenStatus() string`

GetOptionOpenStatus returns the OptionOpenStatus field if non-nil, zero value otherwise.

### GetOptionOpenStatusOk

`func (o *GetSecurityAccountsResponseDataInner) GetOptionOpenStatusOk() (*string, bool)`

GetOptionOpenStatusOk returns a tuple with the OptionOpenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionOpenStatus

`func (o *GetSecurityAccountsResponseDataInner) SetOptionOpenStatus(v string)`

SetOptionOpenStatus sets OptionOpenStatus field to given value.

### HasOptionOpenStatus

`func (o *GetSecurityAccountsResponseDataInner) HasOptionOpenStatus() bool`

HasOptionOpenStatus returns a boolean if a field has been set.

### GetRegisterRegionId

`func (o *GetSecurityAccountsResponseDataInner) GetRegisterRegionId() int32`

GetRegisterRegionId returns the RegisterRegionId field if non-nil, zero value otherwise.

### GetRegisterRegionIdOk

`func (o *GetSecurityAccountsResponseDataInner) GetRegisterRegionIdOk() (*int32, bool)`

GetRegisterRegionIdOk returns a tuple with the RegisterRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterRegionId

`func (o *GetSecurityAccountsResponseDataInner) SetRegisterRegionId(v int32)`

SetRegisterRegionId sets RegisterRegionId field to given value.

### HasRegisterRegionId

`func (o *GetSecurityAccountsResponseDataInner) HasRegisterRegionId() bool`

HasRegisterRegionId returns a boolean if a field has been set.

### GetRegisterTradeUrl

`func (o *GetSecurityAccountsResponseDataInner) GetRegisterTradeUrl() string`

GetRegisterTradeUrl returns the RegisterTradeUrl field if non-nil, zero value otherwise.

### GetRegisterTradeUrlOk

`func (o *GetSecurityAccountsResponseDataInner) GetRegisterTradeUrlOk() (*string, bool)`

GetRegisterTradeUrlOk returns a tuple with the RegisterTradeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterTradeUrl

`func (o *GetSecurityAccountsResponseDataInner) SetRegisterTradeUrl(v string)`

SetRegisterTradeUrl sets RegisterTradeUrl field to given value.

### HasRegisterTradeUrl

`func (o *GetSecurityAccountsResponseDataInner) HasRegisterTradeUrl() bool`

HasRegisterTradeUrl returns a boolean if a field has been set.

### GetSecAccountId

`func (o *GetSecurityAccountsResponseDataInner) GetSecAccountId() int32`

GetSecAccountId returns the SecAccountId field if non-nil, zero value otherwise.

### GetSecAccountIdOk

`func (o *GetSecurityAccountsResponseDataInner) GetSecAccountIdOk() (*int32, bool)`

GetSecAccountIdOk returns a tuple with the SecAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecAccountId

`func (o *GetSecurityAccountsResponseDataInner) SetSecAccountId(v int32)`

SetSecAccountId sets SecAccountId field to given value.

### HasSecAccountId

`func (o *GetSecurityAccountsResponseDataInner) HasSecAccountId() bool`

HasSecAccountId returns a boolean if a field has been set.

### GetStatus

`func (o *GetSecurityAccountsResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSecurityAccountsResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSecurityAccountsResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSecurityAccountsResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportClickIPO

`func (o *GetSecurityAccountsResponseDataInner) GetSupportClickIPO() bool`

GetSupportClickIPO returns the SupportClickIPO field if non-nil, zero value otherwise.

### GetSupportClickIPOOk

`func (o *GetSecurityAccountsResponseDataInner) GetSupportClickIPOOk() (*bool, bool)`

GetSupportClickIPOOk returns a tuple with the SupportClickIPO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportClickIPO

`func (o *GetSecurityAccountsResponseDataInner) SetSupportClickIPO(v bool)`

SetSupportClickIPO sets SupportClickIPO field to given value.

### HasSupportClickIPO

`func (o *GetSecurityAccountsResponseDataInner) HasSupportClickIPO() bool`

HasSupportClickIPO returns a boolean if a field has been set.

### GetSupportOpenOption

`func (o *GetSecurityAccountsResponseDataInner) GetSupportOpenOption() bool`

GetSupportOpenOption returns the SupportOpenOption field if non-nil, zero value otherwise.

### GetSupportOpenOptionOk

`func (o *GetSecurityAccountsResponseDataInner) GetSupportOpenOptionOk() (*bool, bool)`

GetSupportOpenOptionOk returns a tuple with the SupportOpenOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportOpenOption

`func (o *GetSecurityAccountsResponseDataInner) SetSupportOpenOption(v bool)`

SetSupportOpenOption sets SupportOpenOption field to given value.

### HasSupportOpenOption

`func (o *GetSecurityAccountsResponseDataInner) HasSupportOpenOption() bool`

HasSupportOpenOption returns a boolean if a field has been set.

### GetSupportOutsideRth

`func (o *GetSecurityAccountsResponseDataInner) GetSupportOutsideRth() bool`

GetSupportOutsideRth returns the SupportOutsideRth field if non-nil, zero value otherwise.

### GetSupportOutsideRthOk

`func (o *GetSecurityAccountsResponseDataInner) GetSupportOutsideRthOk() (*bool, bool)`

GetSupportOutsideRthOk returns a tuple with the SupportOutsideRth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportOutsideRth

`func (o *GetSecurityAccountsResponseDataInner) SetSupportOutsideRth(v bool)`

SetSupportOutsideRth sets SupportOutsideRth field to given value.

### HasSupportOutsideRth

`func (o *GetSecurityAccountsResponseDataInner) HasSupportOutsideRth() bool`

HasSupportOutsideRth returns a boolean if a field has been set.

### GetTickerTypes

`func (o *GetSecurityAccountsResponseDataInner) GetTickerTypes() []GetSecurityAccountsResponseDataInnerTickerTypesInner`

GetTickerTypes returns the TickerTypes field if non-nil, zero value otherwise.

### GetTickerTypesOk

`func (o *GetSecurityAccountsResponseDataInner) GetTickerTypesOk() (*[]GetSecurityAccountsResponseDataInnerTickerTypesInner, bool)`

GetTickerTypesOk returns a tuple with the TickerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerTypes

`func (o *GetSecurityAccountsResponseDataInner) SetTickerTypes(v []GetSecurityAccountsResponseDataInnerTickerTypesInner)`

SetTickerTypes sets TickerTypes field to given value.

### HasTickerTypes

`func (o *GetSecurityAccountsResponseDataInner) HasTickerTypes() bool`

HasTickerTypes returns a boolean if a field has been set.

### GetTimeInForces

`func (o *GetSecurityAccountsResponseDataInner) GetTimeInForces() []GetSecurityAccountsResponseDataInnerTimeInForcesInner`

GetTimeInForces returns the TimeInForces field if non-nil, zero value otherwise.

### GetTimeInForcesOk

`func (o *GetSecurityAccountsResponseDataInner) GetTimeInForcesOk() (*[]GetSecurityAccountsResponseDataInnerTimeInForcesInner, bool)`

GetTimeInForcesOk returns a tuple with the TimeInForces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForces

`func (o *GetSecurityAccountsResponseDataInner) SetTimeInForces(v []GetSecurityAccountsResponseDataInnerTimeInForcesInner)`

SetTimeInForces sets TimeInForces field to given value.

### HasTimeInForces

`func (o *GetSecurityAccountsResponseDataInner) HasTimeInForces() bool`

HasTimeInForces returns a boolean if a field has been set.

### GetUserTradePermissionVOs

`func (o *GetSecurityAccountsResponseDataInner) GetUserTradePermissionVOs() []GetSecurityAccountsResponseDataInnerUserTradePermissionVOsInner`

GetUserTradePermissionVOs returns the UserTradePermissionVOs field if non-nil, zero value otherwise.

### GetUserTradePermissionVOsOk

`func (o *GetSecurityAccountsResponseDataInner) GetUserTradePermissionVOsOk() (*[]GetSecurityAccountsResponseDataInnerUserTradePermissionVOsInner, bool)`

GetUserTradePermissionVOsOk returns a tuple with the UserTradePermissionVOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTradePermissionVOs

`func (o *GetSecurityAccountsResponseDataInner) SetUserTradePermissionVOs(v []GetSecurityAccountsResponseDataInnerUserTradePermissionVOsInner)`

SetUserTradePermissionVOs sets UserTradePermissionVOs field to given value.

### HasUserTradePermissionVOs

`func (o *GetSecurityAccountsResponseDataInner) HasUserTradePermissionVOs() bool`

HasUserTradePermissionVOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


