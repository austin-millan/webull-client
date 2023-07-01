# OptionOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**OrderSide**](OrderSide.md) |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**TickerId** | Pointer to **int32** |  | [optional] 
**TickerType** | Pointer to **string** |  | [optional] [default to "OPTION"]
**TotalQuantity** | Pointer to **int32** |  | [optional] 

## Methods

### NewOptionOrder

`func NewOptionOrder() *OptionOrder`

NewOptionOrder instantiates a new OptionOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionOrderWithDefaults

`func NewOptionOrderWithDefaults() *OptionOrder`

NewOptionOrderWithDefaults instantiates a new OptionOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *OptionOrder) GetAction() OrderSide`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OptionOrder) GetActionOk() (*OrderSide, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OptionOrder) SetAction(v OrderSide)`

SetAction sets Action field to given value.

### HasAction

`func (o *OptionOrder) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetQuantity

`func (o *OptionOrder) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OptionOrder) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OptionOrder) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OptionOrder) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTickerId

`func (o *OptionOrder) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *OptionOrder) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *OptionOrder) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *OptionOrder) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetTickerType

`func (o *OptionOrder) GetTickerType() string`

GetTickerType returns the TickerType field if non-nil, zero value otherwise.

### GetTickerTypeOk

`func (o *OptionOrder) GetTickerTypeOk() (*string, bool)`

GetTickerTypeOk returns a tuple with the TickerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerType

`func (o *OptionOrder) SetTickerType(v string)`

SetTickerType sets TickerType field to given value.

### HasTickerType

`func (o *OptionOrder) HasTickerType() bool`

HasTickerType returns a boolean if a field has been set.

### GetTotalQuantity

`func (o *OptionOrder) GetTotalQuantity() int32`

GetTotalQuantity returns the TotalQuantity field if non-nil, zero value otherwise.

### GetTotalQuantityOk

`func (o *OptionOrder) GetTotalQuantityOk() (*int32, bool)`

GetTotalQuantityOk returns a tuple with the TotalQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantity

`func (o *OptionOrder) SetTotalQuantity(v int32)`

SetTotalQuantity sets TotalQuantity field to given value.

### HasTotalQuantity

`func (o *OptionOrder) HasTotalQuantity() bool`

HasTotalQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


