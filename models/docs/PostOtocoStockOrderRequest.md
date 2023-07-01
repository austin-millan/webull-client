# PostOtocoStockOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**OrderSide**](OrderSide.md) |  | [optional] 
**LmtPrice** | Pointer to **float32** | limit price | [optional] 
**OrderType** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**OutsideRegularTradingHour** | Pointer to **bool** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**SerialId** | Pointer to **string** |  | [optional] 
**TickerId** | Pointer to **int32** |  | [optional] 
**TimeInForce** | Pointer to [**TimeInForce**](TimeInForce.md) |  | [optional] 

## Methods

### NewPostOtocoStockOrderRequest

`func NewPostOtocoStockOrderRequest() *PostOtocoStockOrderRequest`

NewPostOtocoStockOrderRequest instantiates a new PostOtocoStockOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostOtocoStockOrderRequestWithDefaults

`func NewPostOtocoStockOrderRequestWithDefaults() *PostOtocoStockOrderRequest`

NewPostOtocoStockOrderRequestWithDefaults instantiates a new PostOtocoStockOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PostOtocoStockOrderRequest) GetAction() OrderSide`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PostOtocoStockOrderRequest) GetActionOk() (*OrderSide, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PostOtocoStockOrderRequest) SetAction(v OrderSide)`

SetAction sets Action field to given value.

### HasAction

`func (o *PostOtocoStockOrderRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetLmtPrice

`func (o *PostOtocoStockOrderRequest) GetLmtPrice() float32`

GetLmtPrice returns the LmtPrice field if non-nil, zero value otherwise.

### GetLmtPriceOk

`func (o *PostOtocoStockOrderRequest) GetLmtPriceOk() (*float32, bool)`

GetLmtPriceOk returns a tuple with the LmtPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmtPrice

`func (o *PostOtocoStockOrderRequest) SetLmtPrice(v float32)`

SetLmtPrice sets LmtPrice field to given value.

### HasLmtPrice

`func (o *PostOtocoStockOrderRequest) HasLmtPrice() bool`

HasLmtPrice returns a boolean if a field has been set.

### GetOrderType

`func (o *PostOtocoStockOrderRequest) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *PostOtocoStockOrderRequest) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *PostOtocoStockOrderRequest) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *PostOtocoStockOrderRequest) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetOutsideRegularTradingHour

`func (o *PostOtocoStockOrderRequest) GetOutsideRegularTradingHour() bool`

GetOutsideRegularTradingHour returns the OutsideRegularTradingHour field if non-nil, zero value otherwise.

### GetOutsideRegularTradingHourOk

`func (o *PostOtocoStockOrderRequest) GetOutsideRegularTradingHourOk() (*bool, bool)`

GetOutsideRegularTradingHourOk returns a tuple with the OutsideRegularTradingHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideRegularTradingHour

`func (o *PostOtocoStockOrderRequest) SetOutsideRegularTradingHour(v bool)`

SetOutsideRegularTradingHour sets OutsideRegularTradingHour field to given value.

### HasOutsideRegularTradingHour

`func (o *PostOtocoStockOrderRequest) HasOutsideRegularTradingHour() bool`

HasOutsideRegularTradingHour returns a boolean if a field has been set.

### GetQuantity

`func (o *PostOtocoStockOrderRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PostOtocoStockOrderRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PostOtocoStockOrderRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PostOtocoStockOrderRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSerialId

`func (o *PostOtocoStockOrderRequest) GetSerialId() string`

GetSerialId returns the SerialId field if non-nil, zero value otherwise.

### GetSerialIdOk

`func (o *PostOtocoStockOrderRequest) GetSerialIdOk() (*string, bool)`

GetSerialIdOk returns a tuple with the SerialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialId

`func (o *PostOtocoStockOrderRequest) SetSerialId(v string)`

SetSerialId sets SerialId field to given value.

### HasSerialId

`func (o *PostOtocoStockOrderRequest) HasSerialId() bool`

HasSerialId returns a boolean if a field has been set.

### GetTickerId

`func (o *PostOtocoStockOrderRequest) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *PostOtocoStockOrderRequest) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *PostOtocoStockOrderRequest) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *PostOtocoStockOrderRequest) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PostOtocoStockOrderRequest) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PostOtocoStockOrderRequest) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PostOtocoStockOrderRequest) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PostOtocoStockOrderRequest) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


