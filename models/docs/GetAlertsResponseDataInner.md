# GetAlertsResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisExchangeCode** | Pointer to **string** |  | [optional] 
**DisSymbol** | Pointer to **string** |  | [optional] 
**EventWarning** | Pointer to [**GetAlertsResponseDataInnerEventWarning**](GetAlertsResponseDataInnerEventWarning.md) |  | [optional] 
**ExchangeCode** | Pointer to **string** |  | [optional] 
**ExchangeTrade** | Pointer to **bool** |  | [optional] 
**RegionId** | Pointer to **int32** |  | [optional] 
**ShowCode** | Pointer to **string** |  | [optional] 
**TickerId** | Pointer to **int32** |  | [optional] 
**TickerName** | Pointer to **string** |  | [optional] 
**TickerSymbol** | Pointer to **string** |  | [optional] 
**TickerType** | Pointer to **int32** |  | [optional] 
**TickerWarning** | Pointer to [**GetAlertsResponseDataInnerTickerWarning**](GetAlertsResponseDataInnerTickerWarning.md) |  | [optional] 
**TinyName** | Pointer to **string** |  | [optional] 

## Methods

### NewGetAlertsResponseDataInner

`func NewGetAlertsResponseDataInner() *GetAlertsResponseDataInner`

NewGetAlertsResponseDataInner instantiates a new GetAlertsResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlertsResponseDataInnerWithDefaults

`func NewGetAlertsResponseDataInnerWithDefaults() *GetAlertsResponseDataInner`

NewGetAlertsResponseDataInnerWithDefaults instantiates a new GetAlertsResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisExchangeCode

`func (o *GetAlertsResponseDataInner) GetDisExchangeCode() string`

GetDisExchangeCode returns the DisExchangeCode field if non-nil, zero value otherwise.

### GetDisExchangeCodeOk

`func (o *GetAlertsResponseDataInner) GetDisExchangeCodeOk() (*string, bool)`

GetDisExchangeCodeOk returns a tuple with the DisExchangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisExchangeCode

`func (o *GetAlertsResponseDataInner) SetDisExchangeCode(v string)`

SetDisExchangeCode sets DisExchangeCode field to given value.

### HasDisExchangeCode

`func (o *GetAlertsResponseDataInner) HasDisExchangeCode() bool`

HasDisExchangeCode returns a boolean if a field has been set.

### GetDisSymbol

`func (o *GetAlertsResponseDataInner) GetDisSymbol() string`

GetDisSymbol returns the DisSymbol field if non-nil, zero value otherwise.

### GetDisSymbolOk

`func (o *GetAlertsResponseDataInner) GetDisSymbolOk() (*string, bool)`

GetDisSymbolOk returns a tuple with the DisSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisSymbol

`func (o *GetAlertsResponseDataInner) SetDisSymbol(v string)`

SetDisSymbol sets DisSymbol field to given value.

### HasDisSymbol

`func (o *GetAlertsResponseDataInner) HasDisSymbol() bool`

HasDisSymbol returns a boolean if a field has been set.

### GetEventWarning

`func (o *GetAlertsResponseDataInner) GetEventWarning() GetAlertsResponseDataInnerEventWarning`

GetEventWarning returns the EventWarning field if non-nil, zero value otherwise.

### GetEventWarningOk

`func (o *GetAlertsResponseDataInner) GetEventWarningOk() (*GetAlertsResponseDataInnerEventWarning, bool)`

GetEventWarningOk returns a tuple with the EventWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventWarning

`func (o *GetAlertsResponseDataInner) SetEventWarning(v GetAlertsResponseDataInnerEventWarning)`

SetEventWarning sets EventWarning field to given value.

### HasEventWarning

`func (o *GetAlertsResponseDataInner) HasEventWarning() bool`

HasEventWarning returns a boolean if a field has been set.

### GetExchangeCode

`func (o *GetAlertsResponseDataInner) GetExchangeCode() string`

GetExchangeCode returns the ExchangeCode field if non-nil, zero value otherwise.

### GetExchangeCodeOk

`func (o *GetAlertsResponseDataInner) GetExchangeCodeOk() (*string, bool)`

GetExchangeCodeOk returns a tuple with the ExchangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeCode

`func (o *GetAlertsResponseDataInner) SetExchangeCode(v string)`

SetExchangeCode sets ExchangeCode field to given value.

### HasExchangeCode

`func (o *GetAlertsResponseDataInner) HasExchangeCode() bool`

HasExchangeCode returns a boolean if a field has been set.

### GetExchangeTrade

`func (o *GetAlertsResponseDataInner) GetExchangeTrade() bool`

GetExchangeTrade returns the ExchangeTrade field if non-nil, zero value otherwise.

### GetExchangeTradeOk

`func (o *GetAlertsResponseDataInner) GetExchangeTradeOk() (*bool, bool)`

GetExchangeTradeOk returns a tuple with the ExchangeTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTrade

`func (o *GetAlertsResponseDataInner) SetExchangeTrade(v bool)`

SetExchangeTrade sets ExchangeTrade field to given value.

### HasExchangeTrade

`func (o *GetAlertsResponseDataInner) HasExchangeTrade() bool`

HasExchangeTrade returns a boolean if a field has been set.

### GetRegionId

`func (o *GetAlertsResponseDataInner) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *GetAlertsResponseDataInner) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *GetAlertsResponseDataInner) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *GetAlertsResponseDataInner) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetShowCode

`func (o *GetAlertsResponseDataInner) GetShowCode() string`

GetShowCode returns the ShowCode field if non-nil, zero value otherwise.

### GetShowCodeOk

`func (o *GetAlertsResponseDataInner) GetShowCodeOk() (*string, bool)`

GetShowCodeOk returns a tuple with the ShowCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCode

`func (o *GetAlertsResponseDataInner) SetShowCode(v string)`

SetShowCode sets ShowCode field to given value.

### HasShowCode

`func (o *GetAlertsResponseDataInner) HasShowCode() bool`

HasShowCode returns a boolean if a field has been set.

### GetTickerId

`func (o *GetAlertsResponseDataInner) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *GetAlertsResponseDataInner) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *GetAlertsResponseDataInner) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *GetAlertsResponseDataInner) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetTickerName

`func (o *GetAlertsResponseDataInner) GetTickerName() string`

GetTickerName returns the TickerName field if non-nil, zero value otherwise.

### GetTickerNameOk

`func (o *GetAlertsResponseDataInner) GetTickerNameOk() (*string, bool)`

GetTickerNameOk returns a tuple with the TickerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerName

`func (o *GetAlertsResponseDataInner) SetTickerName(v string)`

SetTickerName sets TickerName field to given value.

### HasTickerName

`func (o *GetAlertsResponseDataInner) HasTickerName() bool`

HasTickerName returns a boolean if a field has been set.

### GetTickerSymbol

`func (o *GetAlertsResponseDataInner) GetTickerSymbol() string`

GetTickerSymbol returns the TickerSymbol field if non-nil, zero value otherwise.

### GetTickerSymbolOk

`func (o *GetAlertsResponseDataInner) GetTickerSymbolOk() (*string, bool)`

GetTickerSymbolOk returns a tuple with the TickerSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerSymbol

`func (o *GetAlertsResponseDataInner) SetTickerSymbol(v string)`

SetTickerSymbol sets TickerSymbol field to given value.

### HasTickerSymbol

`func (o *GetAlertsResponseDataInner) HasTickerSymbol() bool`

HasTickerSymbol returns a boolean if a field has been set.

### GetTickerType

`func (o *GetAlertsResponseDataInner) GetTickerType() int32`

GetTickerType returns the TickerType field if non-nil, zero value otherwise.

### GetTickerTypeOk

`func (o *GetAlertsResponseDataInner) GetTickerTypeOk() (*int32, bool)`

GetTickerTypeOk returns a tuple with the TickerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerType

`func (o *GetAlertsResponseDataInner) SetTickerType(v int32)`

SetTickerType sets TickerType field to given value.

### HasTickerType

`func (o *GetAlertsResponseDataInner) HasTickerType() bool`

HasTickerType returns a boolean if a field has been set.

### GetTickerWarning

`func (o *GetAlertsResponseDataInner) GetTickerWarning() GetAlertsResponseDataInnerTickerWarning`

GetTickerWarning returns the TickerWarning field if non-nil, zero value otherwise.

### GetTickerWarningOk

`func (o *GetAlertsResponseDataInner) GetTickerWarningOk() (*GetAlertsResponseDataInnerTickerWarning, bool)`

GetTickerWarningOk returns a tuple with the TickerWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerWarning

`func (o *GetAlertsResponseDataInner) SetTickerWarning(v GetAlertsResponseDataInnerTickerWarning)`

SetTickerWarning sets TickerWarning field to given value.

### HasTickerWarning

`func (o *GetAlertsResponseDataInner) HasTickerWarning() bool`

HasTickerWarning returns a boolean if a field has been set.

### GetTinyName

`func (o *GetAlertsResponseDataInner) GetTinyName() string`

GetTinyName returns the TinyName field if non-nil, zero value otherwise.

### GetTinyNameOk

`func (o *GetAlertsResponseDataInner) GetTinyNameOk() (*string, bool)`

GetTinyNameOk returns a tuple with the TinyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTinyName

`func (o *GetAlertsResponseDataInner) SetTinyName(v string)`

SetTinyName sets TinyName field to given value.

### HasTinyName

`func (o *GetAlertsResponseDataInner) HasTinyName() bool`

HasTinyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


