# AddAlertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisExchangeCode** | Pointer to **string** |  | [optional] 
**DisSymbol** | Pointer to **string** |  | [optional] 
**EventWarningInput** | Pointer to [**AddAlertRequestEventWarningInput**](AddAlertRequestEventWarningInput.md) |  | [optional] 
**ExchangeCode** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **string** |  | [optional] 
**ShowCode** | Pointer to **string** |  | [optional] 
**TickerId** | Pointer to **int32** |  | [optional] [default to 913243251]
**TickerName** | Pointer to **string** |  | [optional] 
**TickerSymbol** | Pointer to **string** |  | [optional] 
**TickerType** | Pointer to **string** |  | [optional] 
**TinyName** | Pointer to **string** |  | [optional] 
**WarningInput** | Pointer to [**AddAlertRequestWarningInput**](AddAlertRequestWarningInput.md) |  | [optional] 

## Methods

### NewAddAlertRequest

`func NewAddAlertRequest() *AddAlertRequest`

NewAddAlertRequest instantiates a new AddAlertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAlertRequestWithDefaults

`func NewAddAlertRequestWithDefaults() *AddAlertRequest`

NewAddAlertRequestWithDefaults instantiates a new AddAlertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisExchangeCode

`func (o *AddAlertRequest) GetDisExchangeCode() string`

GetDisExchangeCode returns the DisExchangeCode field if non-nil, zero value otherwise.

### GetDisExchangeCodeOk

`func (o *AddAlertRequest) GetDisExchangeCodeOk() (*string, bool)`

GetDisExchangeCodeOk returns a tuple with the DisExchangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisExchangeCode

`func (o *AddAlertRequest) SetDisExchangeCode(v string)`

SetDisExchangeCode sets DisExchangeCode field to given value.

### HasDisExchangeCode

`func (o *AddAlertRequest) HasDisExchangeCode() bool`

HasDisExchangeCode returns a boolean if a field has been set.

### GetDisSymbol

`func (o *AddAlertRequest) GetDisSymbol() string`

GetDisSymbol returns the DisSymbol field if non-nil, zero value otherwise.

### GetDisSymbolOk

`func (o *AddAlertRequest) GetDisSymbolOk() (*string, bool)`

GetDisSymbolOk returns a tuple with the DisSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisSymbol

`func (o *AddAlertRequest) SetDisSymbol(v string)`

SetDisSymbol sets DisSymbol field to given value.

### HasDisSymbol

`func (o *AddAlertRequest) HasDisSymbol() bool`

HasDisSymbol returns a boolean if a field has been set.

### GetEventWarningInput

`func (o *AddAlertRequest) GetEventWarningInput() AddAlertRequestEventWarningInput`

GetEventWarningInput returns the EventWarningInput field if non-nil, zero value otherwise.

### GetEventWarningInputOk

`func (o *AddAlertRequest) GetEventWarningInputOk() (*AddAlertRequestEventWarningInput, bool)`

GetEventWarningInputOk returns a tuple with the EventWarningInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventWarningInput

`func (o *AddAlertRequest) SetEventWarningInput(v AddAlertRequestEventWarningInput)`

SetEventWarningInput sets EventWarningInput field to given value.

### HasEventWarningInput

`func (o *AddAlertRequest) HasEventWarningInput() bool`

HasEventWarningInput returns a boolean if a field has been set.

### GetExchangeCode

`func (o *AddAlertRequest) GetExchangeCode() string`

GetExchangeCode returns the ExchangeCode field if non-nil, zero value otherwise.

### GetExchangeCodeOk

`func (o *AddAlertRequest) GetExchangeCodeOk() (*string, bool)`

GetExchangeCodeOk returns a tuple with the ExchangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeCode

`func (o *AddAlertRequest) SetExchangeCode(v string)`

SetExchangeCode sets ExchangeCode field to given value.

### HasExchangeCode

`func (o *AddAlertRequest) HasExchangeCode() bool`

HasExchangeCode returns a boolean if a field has been set.

### GetRegionId

`func (o *AddAlertRequest) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AddAlertRequest) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AddAlertRequest) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *AddAlertRequest) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetShowCode

`func (o *AddAlertRequest) GetShowCode() string`

GetShowCode returns the ShowCode field if non-nil, zero value otherwise.

### GetShowCodeOk

`func (o *AddAlertRequest) GetShowCodeOk() (*string, bool)`

GetShowCodeOk returns a tuple with the ShowCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCode

`func (o *AddAlertRequest) SetShowCode(v string)`

SetShowCode sets ShowCode field to given value.

### HasShowCode

`func (o *AddAlertRequest) HasShowCode() bool`

HasShowCode returns a boolean if a field has been set.

### GetTickerId

`func (o *AddAlertRequest) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *AddAlertRequest) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *AddAlertRequest) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *AddAlertRequest) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetTickerName

`func (o *AddAlertRequest) GetTickerName() string`

GetTickerName returns the TickerName field if non-nil, zero value otherwise.

### GetTickerNameOk

`func (o *AddAlertRequest) GetTickerNameOk() (*string, bool)`

GetTickerNameOk returns a tuple with the TickerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerName

`func (o *AddAlertRequest) SetTickerName(v string)`

SetTickerName sets TickerName field to given value.

### HasTickerName

`func (o *AddAlertRequest) HasTickerName() bool`

HasTickerName returns a boolean if a field has been set.

### GetTickerSymbol

`func (o *AddAlertRequest) GetTickerSymbol() string`

GetTickerSymbol returns the TickerSymbol field if non-nil, zero value otherwise.

### GetTickerSymbolOk

`func (o *AddAlertRequest) GetTickerSymbolOk() (*string, bool)`

GetTickerSymbolOk returns a tuple with the TickerSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerSymbol

`func (o *AddAlertRequest) SetTickerSymbol(v string)`

SetTickerSymbol sets TickerSymbol field to given value.

### HasTickerSymbol

`func (o *AddAlertRequest) HasTickerSymbol() bool`

HasTickerSymbol returns a boolean if a field has been set.

### GetTickerType

`func (o *AddAlertRequest) GetTickerType() string`

GetTickerType returns the TickerType field if non-nil, zero value otherwise.

### GetTickerTypeOk

`func (o *AddAlertRequest) GetTickerTypeOk() (*string, bool)`

GetTickerTypeOk returns a tuple with the TickerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerType

`func (o *AddAlertRequest) SetTickerType(v string)`

SetTickerType sets TickerType field to given value.

### HasTickerType

`func (o *AddAlertRequest) HasTickerType() bool`

HasTickerType returns a boolean if a field has been set.

### GetTinyName

`func (o *AddAlertRequest) GetTinyName() string`

GetTinyName returns the TinyName field if non-nil, zero value otherwise.

### GetTinyNameOk

`func (o *AddAlertRequest) GetTinyNameOk() (*string, bool)`

GetTinyNameOk returns a tuple with the TinyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTinyName

`func (o *AddAlertRequest) SetTinyName(v string)`

SetTinyName sets TinyName field to given value.

### HasTinyName

`func (o *AddAlertRequest) HasTinyName() bool`

HasTinyName returns a boolean if a field has been set.

### GetWarningInput

`func (o *AddAlertRequest) GetWarningInput() AddAlertRequestWarningInput`

GetWarningInput returns the WarningInput field if non-nil, zero value otherwise.

### GetWarningInputOk

`func (o *AddAlertRequest) GetWarningInputOk() (*AddAlertRequestWarningInput, bool)`

GetWarningInputOk returns a tuple with the WarningInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningInput

`func (o *AddAlertRequest) SetWarningInput(v AddAlertRequestWarningInput)`

SetWarningInput sets WarningInput field to given value.

### HasWarningInput

`func (o *AddAlertRequest) HasWarningInput() bool`

HasWarningInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


