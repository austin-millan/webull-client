# GetAlertsResponseDataInnerTickerWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BizTimestamp** | Pointer to **float32** |  | [optional] 
**Del** | Pointer to **bool** |  | [optional] 
**DisExchangeCode** | Pointer to **string** |  | [optional] 
**DisSymbol** | Pointer to **string** |  | [optional] 
**ExchangeCode** | Pointer to **string** |  | [optional] 
**ExchangeTrade** | Pointer to **bool** |  | [optional] 
**RegionId** | Pointer to **int32** |  | [optional] 
**Rules** | Pointer to [**[]GetAlertsResponseDataInnerTickerWarningRulesInner**](GetAlertsResponseDataInnerTickerWarningRulesInner.md) |  | [optional] 
**ShowCode** | Pointer to **string** |  | [optional] 
**TickerId** | Pointer to **int32** |  | [optional] 
**TickerName** | Pointer to **string** |  | [optional] 
**TickerSymbol** | Pointer to **string** |  | [optional] 
**TickerType** | Pointer to **int32** |  | [optional] 
**TinyName** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **string** |  | [optional] 
**WarningFrequency** | Pointer to **int32** |  | [optional] 
**WarningInterval** | Pointer to **int32** |  | [optional] 
**WarningMode** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetAlertsResponseDataInnerTickerWarning

`func NewGetAlertsResponseDataInnerTickerWarning() *GetAlertsResponseDataInnerTickerWarning`

NewGetAlertsResponseDataInnerTickerWarning instantiates a new GetAlertsResponseDataInnerTickerWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlertsResponseDataInnerTickerWarningWithDefaults

`func NewGetAlertsResponseDataInnerTickerWarningWithDefaults() *GetAlertsResponseDataInnerTickerWarning`

NewGetAlertsResponseDataInnerTickerWarningWithDefaults instantiates a new GetAlertsResponseDataInnerTickerWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBizTimestamp

`func (o *GetAlertsResponseDataInnerTickerWarning) GetBizTimestamp() float32`

GetBizTimestamp returns the BizTimestamp field if non-nil, zero value otherwise.

### GetBizTimestampOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetBizTimestampOk() (*float32, bool)`

GetBizTimestampOk returns a tuple with the BizTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizTimestamp

`func (o *GetAlertsResponseDataInnerTickerWarning) SetBizTimestamp(v float32)`

SetBizTimestamp sets BizTimestamp field to given value.

### HasBizTimestamp

`func (o *GetAlertsResponseDataInnerTickerWarning) HasBizTimestamp() bool`

HasBizTimestamp returns a boolean if a field has been set.

### GetDel

`func (o *GetAlertsResponseDataInnerTickerWarning) GetDel() bool`

GetDel returns the Del field if non-nil, zero value otherwise.

### GetDelOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetDelOk() (*bool, bool)`

GetDelOk returns a tuple with the Del field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDel

`func (o *GetAlertsResponseDataInnerTickerWarning) SetDel(v bool)`

SetDel sets Del field to given value.

### HasDel

`func (o *GetAlertsResponseDataInnerTickerWarning) HasDel() bool`

HasDel returns a boolean if a field has been set.

### GetDisExchangeCode

`func (o *GetAlertsResponseDataInnerTickerWarning) GetDisExchangeCode() string`

GetDisExchangeCode returns the DisExchangeCode field if non-nil, zero value otherwise.

### GetDisExchangeCodeOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetDisExchangeCodeOk() (*string, bool)`

GetDisExchangeCodeOk returns a tuple with the DisExchangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisExchangeCode

`func (o *GetAlertsResponseDataInnerTickerWarning) SetDisExchangeCode(v string)`

SetDisExchangeCode sets DisExchangeCode field to given value.

### HasDisExchangeCode

`func (o *GetAlertsResponseDataInnerTickerWarning) HasDisExchangeCode() bool`

HasDisExchangeCode returns a boolean if a field has been set.

### GetDisSymbol

`func (o *GetAlertsResponseDataInnerTickerWarning) GetDisSymbol() string`

GetDisSymbol returns the DisSymbol field if non-nil, zero value otherwise.

### GetDisSymbolOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetDisSymbolOk() (*string, bool)`

GetDisSymbolOk returns a tuple with the DisSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisSymbol

`func (o *GetAlertsResponseDataInnerTickerWarning) SetDisSymbol(v string)`

SetDisSymbol sets DisSymbol field to given value.

### HasDisSymbol

`func (o *GetAlertsResponseDataInnerTickerWarning) HasDisSymbol() bool`

HasDisSymbol returns a boolean if a field has been set.

### GetExchangeCode

`func (o *GetAlertsResponseDataInnerTickerWarning) GetExchangeCode() string`

GetExchangeCode returns the ExchangeCode field if non-nil, zero value otherwise.

### GetExchangeCodeOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetExchangeCodeOk() (*string, bool)`

GetExchangeCodeOk returns a tuple with the ExchangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeCode

`func (o *GetAlertsResponseDataInnerTickerWarning) SetExchangeCode(v string)`

SetExchangeCode sets ExchangeCode field to given value.

### HasExchangeCode

`func (o *GetAlertsResponseDataInnerTickerWarning) HasExchangeCode() bool`

HasExchangeCode returns a boolean if a field has been set.

### GetExchangeTrade

`func (o *GetAlertsResponseDataInnerTickerWarning) GetExchangeTrade() bool`

GetExchangeTrade returns the ExchangeTrade field if non-nil, zero value otherwise.

### GetExchangeTradeOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetExchangeTradeOk() (*bool, bool)`

GetExchangeTradeOk returns a tuple with the ExchangeTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTrade

`func (o *GetAlertsResponseDataInnerTickerWarning) SetExchangeTrade(v bool)`

SetExchangeTrade sets ExchangeTrade field to given value.

### HasExchangeTrade

`func (o *GetAlertsResponseDataInnerTickerWarning) HasExchangeTrade() bool`

HasExchangeTrade returns a boolean if a field has been set.

### GetRegionId

`func (o *GetAlertsResponseDataInnerTickerWarning) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *GetAlertsResponseDataInnerTickerWarning) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *GetAlertsResponseDataInnerTickerWarning) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetRules

`func (o *GetAlertsResponseDataInnerTickerWarning) GetRules() []GetAlertsResponseDataInnerTickerWarningRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetRulesOk() (*[]GetAlertsResponseDataInnerTickerWarningRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GetAlertsResponseDataInnerTickerWarning) SetRules(v []GetAlertsResponseDataInnerTickerWarningRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *GetAlertsResponseDataInnerTickerWarning) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetShowCode

`func (o *GetAlertsResponseDataInnerTickerWarning) GetShowCode() string`

GetShowCode returns the ShowCode field if non-nil, zero value otherwise.

### GetShowCodeOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetShowCodeOk() (*string, bool)`

GetShowCodeOk returns a tuple with the ShowCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCode

`func (o *GetAlertsResponseDataInnerTickerWarning) SetShowCode(v string)`

SetShowCode sets ShowCode field to given value.

### HasShowCode

`func (o *GetAlertsResponseDataInnerTickerWarning) HasShowCode() bool`

HasShowCode returns a boolean if a field has been set.

### GetTickerId

`func (o *GetAlertsResponseDataInnerTickerWarning) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *GetAlertsResponseDataInnerTickerWarning) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *GetAlertsResponseDataInnerTickerWarning) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetTickerName

`func (o *GetAlertsResponseDataInnerTickerWarning) GetTickerName() string`

GetTickerName returns the TickerName field if non-nil, zero value otherwise.

### GetTickerNameOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetTickerNameOk() (*string, bool)`

GetTickerNameOk returns a tuple with the TickerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerName

`func (o *GetAlertsResponseDataInnerTickerWarning) SetTickerName(v string)`

SetTickerName sets TickerName field to given value.

### HasTickerName

`func (o *GetAlertsResponseDataInnerTickerWarning) HasTickerName() bool`

HasTickerName returns a boolean if a field has been set.

### GetTickerSymbol

`func (o *GetAlertsResponseDataInnerTickerWarning) GetTickerSymbol() string`

GetTickerSymbol returns the TickerSymbol field if non-nil, zero value otherwise.

### GetTickerSymbolOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetTickerSymbolOk() (*string, bool)`

GetTickerSymbolOk returns a tuple with the TickerSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerSymbol

`func (o *GetAlertsResponseDataInnerTickerWarning) SetTickerSymbol(v string)`

SetTickerSymbol sets TickerSymbol field to given value.

### HasTickerSymbol

`func (o *GetAlertsResponseDataInnerTickerWarning) HasTickerSymbol() bool`

HasTickerSymbol returns a boolean if a field has been set.

### GetTickerType

`func (o *GetAlertsResponseDataInnerTickerWarning) GetTickerType() int32`

GetTickerType returns the TickerType field if non-nil, zero value otherwise.

### GetTickerTypeOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetTickerTypeOk() (*int32, bool)`

GetTickerTypeOk returns a tuple with the TickerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerType

`func (o *GetAlertsResponseDataInnerTickerWarning) SetTickerType(v int32)`

SetTickerType sets TickerType field to given value.

### HasTickerType

`func (o *GetAlertsResponseDataInnerTickerWarning) HasTickerType() bool`

HasTickerType returns a boolean if a field has been set.

### GetTinyName

`func (o *GetAlertsResponseDataInnerTickerWarning) GetTinyName() string`

GetTinyName returns the TinyName field if non-nil, zero value otherwise.

### GetTinyNameOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetTinyNameOk() (*string, bool)`

GetTinyNameOk returns a tuple with the TinyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTinyName

`func (o *GetAlertsResponseDataInnerTickerWarning) SetTinyName(v string)`

SetTinyName sets TinyName field to given value.

### HasTinyName

`func (o *GetAlertsResponseDataInnerTickerWarning) HasTinyName() bool`

HasTinyName returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetAlertsResponseDataInnerTickerWarning) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetAlertsResponseDataInnerTickerWarning) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetAlertsResponseDataInnerTickerWarning) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetWarningFrequency

`func (o *GetAlertsResponseDataInnerTickerWarning) GetWarningFrequency() int32`

GetWarningFrequency returns the WarningFrequency field if non-nil, zero value otherwise.

### GetWarningFrequencyOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetWarningFrequencyOk() (*int32, bool)`

GetWarningFrequencyOk returns a tuple with the WarningFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningFrequency

`func (o *GetAlertsResponseDataInnerTickerWarning) SetWarningFrequency(v int32)`

SetWarningFrequency sets WarningFrequency field to given value.

### HasWarningFrequency

`func (o *GetAlertsResponseDataInnerTickerWarning) HasWarningFrequency() bool`

HasWarningFrequency returns a boolean if a field has been set.

### GetWarningInterval

`func (o *GetAlertsResponseDataInnerTickerWarning) GetWarningInterval() int32`

GetWarningInterval returns the WarningInterval field if non-nil, zero value otherwise.

### GetWarningIntervalOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetWarningIntervalOk() (*int32, bool)`

GetWarningIntervalOk returns a tuple with the WarningInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningInterval

`func (o *GetAlertsResponseDataInnerTickerWarning) SetWarningInterval(v int32)`

SetWarningInterval sets WarningInterval field to given value.

### HasWarningInterval

`func (o *GetAlertsResponseDataInnerTickerWarning) HasWarningInterval() bool`

HasWarningInterval returns a boolean if a field has been set.

### GetWarningMode

`func (o *GetAlertsResponseDataInnerTickerWarning) GetWarningMode() int32`

GetWarningMode returns the WarningMode field if non-nil, zero value otherwise.

### GetWarningModeOk

`func (o *GetAlertsResponseDataInnerTickerWarning) GetWarningModeOk() (*int32, bool)`

GetWarningModeOk returns a tuple with the WarningMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMode

`func (o *GetAlertsResponseDataInnerTickerWarning) SetWarningMode(v int32)`

SetWarningMode sets WarningMode field to given value.

### HasWarningMode

`func (o *GetAlertsResponseDataInnerTickerWarning) HasWarningMode() bool`

HasWarningMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


