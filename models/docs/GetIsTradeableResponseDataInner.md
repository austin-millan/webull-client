# GetIsTradeableResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrokerId** | Pointer to **int32** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**CurrencyId** | Pointer to **int32** |  | [optional] 
**DisExchangeCode** | Pointer to **string** |  | [optional] 
**DisSymbol** | Pointer to **string** |  | [optional] 
**ExchangeCode** | Pointer to **string** |  | [optional] 
**ExchangeId** | Pointer to **int32** |  | [optional] 
**ExchangeTrade** | Pointer to **bool** |  | [optional] 
**ExtType** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ListStatus** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **int32** |  | [optional] 
**RegionIsoCode** | Pointer to **string** |  | [optional] 
**RegionName** | Pointer to **string** |  | [optional] 
**SecType** | Pointer to **[]int32** |  | [optional] 
**ShowCode** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TickerId** | Pointer to **int32** |  | [optional] 
**TickerName** | Pointer to **string** |  | [optional] 
**TickerType** | Pointer to **int32** |  | [optional] 
**TinyName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**Types** | Pointer to **string** |  | [optional] 

## Methods

### NewGetIsTradeableResponseDataInner

`func NewGetIsTradeableResponseDataInner() *GetIsTradeableResponseDataInner`

NewGetIsTradeableResponseDataInner instantiates a new GetIsTradeableResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIsTradeableResponseDataInnerWithDefaults

`func NewGetIsTradeableResponseDataInnerWithDefaults() *GetIsTradeableResponseDataInner`

NewGetIsTradeableResponseDataInnerWithDefaults instantiates a new GetIsTradeableResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokerId

`func (o *GetIsTradeableResponseDataInner) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *GetIsTradeableResponseDataInner) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *GetIsTradeableResponseDataInner) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *GetIsTradeableResponseDataInner) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GetIsTradeableResponseDataInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GetIsTradeableResponseDataInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GetIsTradeableResponseDataInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GetIsTradeableResponseDataInner) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencyId

`func (o *GetIsTradeableResponseDataInner) GetCurrencyId() int32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *GetIsTradeableResponseDataInner) GetCurrencyIdOk() (*int32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *GetIsTradeableResponseDataInner) SetCurrencyId(v int32)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *GetIsTradeableResponseDataInner) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetDisExchangeCode

`func (o *GetIsTradeableResponseDataInner) GetDisExchangeCode() string`

GetDisExchangeCode returns the DisExchangeCode field if non-nil, zero value otherwise.

### GetDisExchangeCodeOk

`func (o *GetIsTradeableResponseDataInner) GetDisExchangeCodeOk() (*string, bool)`

GetDisExchangeCodeOk returns a tuple with the DisExchangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisExchangeCode

`func (o *GetIsTradeableResponseDataInner) SetDisExchangeCode(v string)`

SetDisExchangeCode sets DisExchangeCode field to given value.

### HasDisExchangeCode

`func (o *GetIsTradeableResponseDataInner) HasDisExchangeCode() bool`

HasDisExchangeCode returns a boolean if a field has been set.

### GetDisSymbol

`func (o *GetIsTradeableResponseDataInner) GetDisSymbol() string`

GetDisSymbol returns the DisSymbol field if non-nil, zero value otherwise.

### GetDisSymbolOk

`func (o *GetIsTradeableResponseDataInner) GetDisSymbolOk() (*string, bool)`

GetDisSymbolOk returns a tuple with the DisSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisSymbol

`func (o *GetIsTradeableResponseDataInner) SetDisSymbol(v string)`

SetDisSymbol sets DisSymbol field to given value.

### HasDisSymbol

`func (o *GetIsTradeableResponseDataInner) HasDisSymbol() bool`

HasDisSymbol returns a boolean if a field has been set.

### GetExchangeCode

`func (o *GetIsTradeableResponseDataInner) GetExchangeCode() string`

GetExchangeCode returns the ExchangeCode field if non-nil, zero value otherwise.

### GetExchangeCodeOk

`func (o *GetIsTradeableResponseDataInner) GetExchangeCodeOk() (*string, bool)`

GetExchangeCodeOk returns a tuple with the ExchangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeCode

`func (o *GetIsTradeableResponseDataInner) SetExchangeCode(v string)`

SetExchangeCode sets ExchangeCode field to given value.

### HasExchangeCode

`func (o *GetIsTradeableResponseDataInner) HasExchangeCode() bool`

HasExchangeCode returns a boolean if a field has been set.

### GetExchangeId

`func (o *GetIsTradeableResponseDataInner) GetExchangeId() int32`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *GetIsTradeableResponseDataInner) GetExchangeIdOk() (*int32, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *GetIsTradeableResponseDataInner) SetExchangeId(v int32)`

SetExchangeId sets ExchangeId field to given value.

### HasExchangeId

`func (o *GetIsTradeableResponseDataInner) HasExchangeId() bool`

HasExchangeId returns a boolean if a field has been set.

### GetExchangeTrade

`func (o *GetIsTradeableResponseDataInner) GetExchangeTrade() bool`

GetExchangeTrade returns the ExchangeTrade field if non-nil, zero value otherwise.

### GetExchangeTradeOk

`func (o *GetIsTradeableResponseDataInner) GetExchangeTradeOk() (*bool, bool)`

GetExchangeTradeOk returns a tuple with the ExchangeTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTrade

`func (o *GetIsTradeableResponseDataInner) SetExchangeTrade(v bool)`

SetExchangeTrade sets ExchangeTrade field to given value.

### HasExchangeTrade

`func (o *GetIsTradeableResponseDataInner) HasExchangeTrade() bool`

HasExchangeTrade returns a boolean if a field has been set.

### GetExtType

`func (o *GetIsTradeableResponseDataInner) GetExtType() []map[string]interface{}`

GetExtType returns the ExtType field if non-nil, zero value otherwise.

### GetExtTypeOk

`func (o *GetIsTradeableResponseDataInner) GetExtTypeOk() (*[]map[string]interface{}, bool)`

GetExtTypeOk returns a tuple with the ExtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtType

`func (o *GetIsTradeableResponseDataInner) SetExtType(v []map[string]interface{})`

SetExtType sets ExtType field to given value.

### HasExtType

`func (o *GetIsTradeableResponseDataInner) HasExtType() bool`

HasExtType returns a boolean if a field has been set.

### GetListStatus

`func (o *GetIsTradeableResponseDataInner) GetListStatus() int32`

GetListStatus returns the ListStatus field if non-nil, zero value otherwise.

### GetListStatusOk

`func (o *GetIsTradeableResponseDataInner) GetListStatusOk() (*int32, bool)`

GetListStatusOk returns a tuple with the ListStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatus

`func (o *GetIsTradeableResponseDataInner) SetListStatus(v int32)`

SetListStatus sets ListStatus field to given value.

### HasListStatus

`func (o *GetIsTradeableResponseDataInner) HasListStatus() bool`

HasListStatus returns a boolean if a field has been set.

### GetName

`func (o *GetIsTradeableResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIsTradeableResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIsTradeableResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIsTradeableResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegionId

`func (o *GetIsTradeableResponseDataInner) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *GetIsTradeableResponseDataInner) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *GetIsTradeableResponseDataInner) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *GetIsTradeableResponseDataInner) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetRegionIsoCode

`func (o *GetIsTradeableResponseDataInner) GetRegionIsoCode() string`

GetRegionIsoCode returns the RegionIsoCode field if non-nil, zero value otherwise.

### GetRegionIsoCodeOk

`func (o *GetIsTradeableResponseDataInner) GetRegionIsoCodeOk() (*string, bool)`

GetRegionIsoCodeOk returns a tuple with the RegionIsoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionIsoCode

`func (o *GetIsTradeableResponseDataInner) SetRegionIsoCode(v string)`

SetRegionIsoCode sets RegionIsoCode field to given value.

### HasRegionIsoCode

`func (o *GetIsTradeableResponseDataInner) HasRegionIsoCode() bool`

HasRegionIsoCode returns a boolean if a field has been set.

### GetRegionName

`func (o *GetIsTradeableResponseDataInner) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *GetIsTradeableResponseDataInner) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *GetIsTradeableResponseDataInner) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *GetIsTradeableResponseDataInner) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetSecType

`func (o *GetIsTradeableResponseDataInner) GetSecType() []int32`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *GetIsTradeableResponseDataInner) GetSecTypeOk() (*[]int32, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *GetIsTradeableResponseDataInner) SetSecType(v []int32)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *GetIsTradeableResponseDataInner) HasSecType() bool`

HasSecType returns a boolean if a field has been set.

### GetShowCode

`func (o *GetIsTradeableResponseDataInner) GetShowCode() string`

GetShowCode returns the ShowCode field if non-nil, zero value otherwise.

### GetShowCodeOk

`func (o *GetIsTradeableResponseDataInner) GetShowCodeOk() (*string, bool)`

GetShowCodeOk returns a tuple with the ShowCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCode

`func (o *GetIsTradeableResponseDataInner) SetShowCode(v string)`

SetShowCode sets ShowCode field to given value.

### HasShowCode

`func (o *GetIsTradeableResponseDataInner) HasShowCode() bool`

HasShowCode returns a boolean if a field has been set.

### GetSymbol

`func (o *GetIsTradeableResponseDataInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetIsTradeableResponseDataInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetIsTradeableResponseDataInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetIsTradeableResponseDataInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTickerId

`func (o *GetIsTradeableResponseDataInner) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *GetIsTradeableResponseDataInner) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *GetIsTradeableResponseDataInner) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *GetIsTradeableResponseDataInner) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetTickerName

`func (o *GetIsTradeableResponseDataInner) GetTickerName() string`

GetTickerName returns the TickerName field if non-nil, zero value otherwise.

### GetTickerNameOk

`func (o *GetIsTradeableResponseDataInner) GetTickerNameOk() (*string, bool)`

GetTickerNameOk returns a tuple with the TickerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerName

`func (o *GetIsTradeableResponseDataInner) SetTickerName(v string)`

SetTickerName sets TickerName field to given value.

### HasTickerName

`func (o *GetIsTradeableResponseDataInner) HasTickerName() bool`

HasTickerName returns a boolean if a field has been set.

### GetTickerType

`func (o *GetIsTradeableResponseDataInner) GetTickerType() int32`

GetTickerType returns the TickerType field if non-nil, zero value otherwise.

### GetTickerTypeOk

`func (o *GetIsTradeableResponseDataInner) GetTickerTypeOk() (*int32, bool)`

GetTickerTypeOk returns a tuple with the TickerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerType

`func (o *GetIsTradeableResponseDataInner) SetTickerType(v int32)`

SetTickerType sets TickerType field to given value.

### HasTickerType

`func (o *GetIsTradeableResponseDataInner) HasTickerType() bool`

HasTickerType returns a boolean if a field has been set.

### GetTinyName

`func (o *GetIsTradeableResponseDataInner) GetTinyName() string`

GetTinyName returns the TinyName field if non-nil, zero value otherwise.

### GetTinyNameOk

`func (o *GetIsTradeableResponseDataInner) GetTinyNameOk() (*string, bool)`

GetTinyNameOk returns a tuple with the TinyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTinyName

`func (o *GetIsTradeableResponseDataInner) SetTinyName(v string)`

SetTinyName sets TinyName field to given value.

### HasTinyName

`func (o *GetIsTradeableResponseDataInner) HasTinyName() bool`

HasTinyName returns a boolean if a field has been set.

### GetType

`func (o *GetIsTradeableResponseDataInner) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIsTradeableResponseDataInner) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIsTradeableResponseDataInner) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *GetIsTradeableResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypes

`func (o *GetIsTradeableResponseDataInner) GetTypes() string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *GetIsTradeableResponseDataInner) GetTypesOk() (*string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *GetIsTradeableResponseDataInner) SetTypes(v string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *GetIsTradeableResponseDataInner) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


