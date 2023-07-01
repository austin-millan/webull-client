# PostStockOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**OrderSide**](OrderSide.md) |  | [optional] 
**ComboType** | Pointer to **string** |  | [optional] 
**LmtPrice** | Pointer to **float32** | limit price | [optional] 
**OrderType** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**OutsideRegularTradingHour** | Pointer to **bool** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**SerialId** | Pointer to **string** |  | [optional] 
**TickerId** | Pointer to **int32** |  | [optional] 
**TimeInForce** | Pointer to [**TimeInForce**](TimeInForce.md) |  | [optional] 

## Methods

### NewPostStockOrderRequest

`func NewPostStockOrderRequest() *PostStockOrderRequest`

NewPostStockOrderRequest instantiates a new PostStockOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostStockOrderRequestWithDefaults

`func NewPostStockOrderRequestWithDefaults() *PostStockOrderRequest`

NewPostStockOrderRequestWithDefaults instantiates a new PostStockOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PostStockOrderRequest) GetAction() OrderSide`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PostStockOrderRequest) GetActionOk() (*OrderSide, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PostStockOrderRequest) SetAction(v OrderSide)`

SetAction sets Action field to given value.

### HasAction

`func (o *PostStockOrderRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetComboType

`func (o *PostStockOrderRequest) GetComboType() string`

GetComboType returns the ComboType field if non-nil, zero value otherwise.

### GetComboTypeOk

`func (o *PostStockOrderRequest) GetComboTypeOk() (*string, bool)`

GetComboTypeOk returns a tuple with the ComboType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComboType

`func (o *PostStockOrderRequest) SetComboType(v string)`

SetComboType sets ComboType field to given value.

### HasComboType

`func (o *PostStockOrderRequest) HasComboType() bool`

HasComboType returns a boolean if a field has been set.

### GetLmtPrice

`func (o *PostStockOrderRequest) GetLmtPrice() float32`

GetLmtPrice returns the LmtPrice field if non-nil, zero value otherwise.

### GetLmtPriceOk

`func (o *PostStockOrderRequest) GetLmtPriceOk() (*float32, bool)`

GetLmtPriceOk returns a tuple with the LmtPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmtPrice

`func (o *PostStockOrderRequest) SetLmtPrice(v float32)`

SetLmtPrice sets LmtPrice field to given value.

### HasLmtPrice

`func (o *PostStockOrderRequest) HasLmtPrice() bool`

HasLmtPrice returns a boolean if a field has been set.

### GetOrderType

`func (o *PostStockOrderRequest) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *PostStockOrderRequest) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *PostStockOrderRequest) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *PostStockOrderRequest) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetOutsideRegularTradingHour

`func (o *PostStockOrderRequest) GetOutsideRegularTradingHour() bool`

GetOutsideRegularTradingHour returns the OutsideRegularTradingHour field if non-nil, zero value otherwise.

### GetOutsideRegularTradingHourOk

`func (o *PostStockOrderRequest) GetOutsideRegularTradingHourOk() (*bool, bool)`

GetOutsideRegularTradingHourOk returns a tuple with the OutsideRegularTradingHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideRegularTradingHour

`func (o *PostStockOrderRequest) SetOutsideRegularTradingHour(v bool)`

SetOutsideRegularTradingHour sets OutsideRegularTradingHour field to given value.

### HasOutsideRegularTradingHour

`func (o *PostStockOrderRequest) HasOutsideRegularTradingHour() bool`

HasOutsideRegularTradingHour returns a boolean if a field has been set.

### GetQuantity

`func (o *PostStockOrderRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PostStockOrderRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PostStockOrderRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PostStockOrderRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSerialId

`func (o *PostStockOrderRequest) GetSerialId() string`

GetSerialId returns the SerialId field if non-nil, zero value otherwise.

### GetSerialIdOk

`func (o *PostStockOrderRequest) GetSerialIdOk() (*string, bool)`

GetSerialIdOk returns a tuple with the SerialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialId

`func (o *PostStockOrderRequest) SetSerialId(v string)`

SetSerialId sets SerialId field to given value.

### HasSerialId

`func (o *PostStockOrderRequest) HasSerialId() bool`

HasSerialId returns a boolean if a field has been set.

### GetTickerId

`func (o *PostStockOrderRequest) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *PostStockOrderRequest) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *PostStockOrderRequest) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *PostStockOrderRequest) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PostStockOrderRequest) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PostStockOrderRequest) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PostStockOrderRequest) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PostStockOrderRequest) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


