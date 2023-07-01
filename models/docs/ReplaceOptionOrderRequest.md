# ReplaceOptionOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuxPrice** | Pointer to **float32** |  | [optional] 
**ComboId** | Pointer to **string** |  | [optional] 
**LmtPrice** | Pointer to **float32** |  | [optional] 
**OrderType** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**Orders** | Pointer to [**[]OptionOrder**](OptionOrder.md) |  | [optional] 
**SerialId** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to [**TimeInForce**](TimeInForce.md) |  | [optional] 

## Methods

### NewReplaceOptionOrderRequest

`func NewReplaceOptionOrderRequest() *ReplaceOptionOrderRequest`

NewReplaceOptionOrderRequest instantiates a new ReplaceOptionOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceOptionOrderRequestWithDefaults

`func NewReplaceOptionOrderRequestWithDefaults() *ReplaceOptionOrderRequest`

NewReplaceOptionOrderRequestWithDefaults instantiates a new ReplaceOptionOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuxPrice

`func (o *ReplaceOptionOrderRequest) GetAuxPrice() float32`

GetAuxPrice returns the AuxPrice field if non-nil, zero value otherwise.

### GetAuxPriceOk

`func (o *ReplaceOptionOrderRequest) GetAuxPriceOk() (*float32, bool)`

GetAuxPriceOk returns a tuple with the AuxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxPrice

`func (o *ReplaceOptionOrderRequest) SetAuxPrice(v float32)`

SetAuxPrice sets AuxPrice field to given value.

### HasAuxPrice

`func (o *ReplaceOptionOrderRequest) HasAuxPrice() bool`

HasAuxPrice returns a boolean if a field has been set.

### GetComboId

`func (o *ReplaceOptionOrderRequest) GetComboId() string`

GetComboId returns the ComboId field if non-nil, zero value otherwise.

### GetComboIdOk

`func (o *ReplaceOptionOrderRequest) GetComboIdOk() (*string, bool)`

GetComboIdOk returns a tuple with the ComboId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComboId

`func (o *ReplaceOptionOrderRequest) SetComboId(v string)`

SetComboId sets ComboId field to given value.

### HasComboId

`func (o *ReplaceOptionOrderRequest) HasComboId() bool`

HasComboId returns a boolean if a field has been set.

### GetLmtPrice

`func (o *ReplaceOptionOrderRequest) GetLmtPrice() float32`

GetLmtPrice returns the LmtPrice field if non-nil, zero value otherwise.

### GetLmtPriceOk

`func (o *ReplaceOptionOrderRequest) GetLmtPriceOk() (*float32, bool)`

GetLmtPriceOk returns a tuple with the LmtPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmtPrice

`func (o *ReplaceOptionOrderRequest) SetLmtPrice(v float32)`

SetLmtPrice sets LmtPrice field to given value.

### HasLmtPrice

`func (o *ReplaceOptionOrderRequest) HasLmtPrice() bool`

HasLmtPrice returns a boolean if a field has been set.

### GetOrderType

`func (o *ReplaceOptionOrderRequest) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *ReplaceOptionOrderRequest) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *ReplaceOptionOrderRequest) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *ReplaceOptionOrderRequest) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetOrders

`func (o *ReplaceOptionOrderRequest) GetOrders() []OptionOrder`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *ReplaceOptionOrderRequest) GetOrdersOk() (*[]OptionOrder, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *ReplaceOptionOrderRequest) SetOrders(v []OptionOrder)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *ReplaceOptionOrderRequest) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSerialId

`func (o *ReplaceOptionOrderRequest) GetSerialId() string`

GetSerialId returns the SerialId field if non-nil, zero value otherwise.

### GetSerialIdOk

`func (o *ReplaceOptionOrderRequest) GetSerialIdOk() (*string, bool)`

GetSerialIdOk returns a tuple with the SerialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialId

`func (o *ReplaceOptionOrderRequest) SetSerialId(v string)`

SetSerialId sets SerialId field to given value.

### HasSerialId

`func (o *ReplaceOptionOrderRequest) HasSerialId() bool`

HasSerialId returns a boolean if a field has been set.

### GetTimeInForce

`func (o *ReplaceOptionOrderRequest) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *ReplaceOptionOrderRequest) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *ReplaceOptionOrderRequest) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *ReplaceOptionOrderRequest) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


