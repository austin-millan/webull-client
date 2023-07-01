# PostOptionOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuxPrice** | Pointer to **float32** |  | [optional] 
**LmtPrice** | Pointer to **float32** |  | [optional] 
**OrderType** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**Orders** | Pointer to [**[]OptionOrder**](OptionOrder.md) |  | [optional] 
**SerialId** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to [**TimeInForce**](TimeInForce.md) |  | [optional] 

## Methods

### NewPostOptionOrderRequest

`func NewPostOptionOrderRequest() *PostOptionOrderRequest`

NewPostOptionOrderRequest instantiates a new PostOptionOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostOptionOrderRequestWithDefaults

`func NewPostOptionOrderRequestWithDefaults() *PostOptionOrderRequest`

NewPostOptionOrderRequestWithDefaults instantiates a new PostOptionOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuxPrice

`func (o *PostOptionOrderRequest) GetAuxPrice() float32`

GetAuxPrice returns the AuxPrice field if non-nil, zero value otherwise.

### GetAuxPriceOk

`func (o *PostOptionOrderRequest) GetAuxPriceOk() (*float32, bool)`

GetAuxPriceOk returns a tuple with the AuxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxPrice

`func (o *PostOptionOrderRequest) SetAuxPrice(v float32)`

SetAuxPrice sets AuxPrice field to given value.

### HasAuxPrice

`func (o *PostOptionOrderRequest) HasAuxPrice() bool`

HasAuxPrice returns a boolean if a field has been set.

### GetLmtPrice

`func (o *PostOptionOrderRequest) GetLmtPrice() float32`

GetLmtPrice returns the LmtPrice field if non-nil, zero value otherwise.

### GetLmtPriceOk

`func (o *PostOptionOrderRequest) GetLmtPriceOk() (*float32, bool)`

GetLmtPriceOk returns a tuple with the LmtPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmtPrice

`func (o *PostOptionOrderRequest) SetLmtPrice(v float32)`

SetLmtPrice sets LmtPrice field to given value.

### HasLmtPrice

`func (o *PostOptionOrderRequest) HasLmtPrice() bool`

HasLmtPrice returns a boolean if a field has been set.

### GetOrderType

`func (o *PostOptionOrderRequest) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *PostOptionOrderRequest) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *PostOptionOrderRequest) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *PostOptionOrderRequest) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetOrders

`func (o *PostOptionOrderRequest) GetOrders() []OptionOrder`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *PostOptionOrderRequest) GetOrdersOk() (*[]OptionOrder, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *PostOptionOrderRequest) SetOrders(v []OptionOrder)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *PostOptionOrderRequest) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSerialId

`func (o *PostOptionOrderRequest) GetSerialId() string`

GetSerialId returns the SerialId field if non-nil, zero value otherwise.

### GetSerialIdOk

`func (o *PostOptionOrderRequest) GetSerialIdOk() (*string, bool)`

GetSerialIdOk returns a tuple with the SerialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialId

`func (o *PostOptionOrderRequest) SetSerialId(v string)`

SetSerialId sets SerialId field to given value.

### HasSerialId

`func (o *PostOptionOrderRequest) HasSerialId() bool`

HasSerialId returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PostOptionOrderRequest) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PostOptionOrderRequest) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PostOptionOrderRequest) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PostOptionOrderRequest) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


