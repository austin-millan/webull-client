# PaperOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**OrderSide**](OrderSide.md) |  | [optional] 
**CanCancel** | Pointer to **bool** |  | [optional] 
**CanModify** | Pointer to **bool** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**CreateTime0** | Pointer to **int64** |  | [optional] 
**FilledQuantity** | Pointer to **string** |  | [optional] 
**LmtPrice** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int32** |  | [optional] 
**OrderType** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**OutsideRegularTradingHour** | Pointer to **bool** |  | [optional] 
**PaperId** | Pointer to **int32** |  | [optional] 
**PlacedTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusStr** | Pointer to **string** |  | [optional] 
**Ticker** | Pointer to [**GetOrdersItemOrdersInnerTicker**](GetOrdersItemOrdersInnerTicker.md) |  | [optional] 
**TimeInForce** | Pointer to [**TimeInForce**](TimeInForce.md) |  | [optional] 
**TotalQuantity** | Pointer to **string** |  | [optional] 

## Methods

### NewPaperOrder

`func NewPaperOrder() *PaperOrder`

NewPaperOrder instantiates a new PaperOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaperOrderWithDefaults

`func NewPaperOrderWithDefaults() *PaperOrder`

NewPaperOrderWithDefaults instantiates a new PaperOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PaperOrder) GetAction() OrderSide`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PaperOrder) GetActionOk() (*OrderSide, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PaperOrder) SetAction(v OrderSide)`

SetAction sets Action field to given value.

### HasAction

`func (o *PaperOrder) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCanCancel

`func (o *PaperOrder) GetCanCancel() bool`

GetCanCancel returns the CanCancel field if non-nil, zero value otherwise.

### GetCanCancelOk

`func (o *PaperOrder) GetCanCancelOk() (*bool, bool)`

GetCanCancelOk returns a tuple with the CanCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCancel

`func (o *PaperOrder) SetCanCancel(v bool)`

SetCanCancel sets CanCancel field to given value.

### HasCanCancel

`func (o *PaperOrder) HasCanCancel() bool`

HasCanCancel returns a boolean if a field has been set.

### GetCanModify

`func (o *PaperOrder) GetCanModify() bool`

GetCanModify returns the CanModify field if non-nil, zero value otherwise.

### GetCanModifyOk

`func (o *PaperOrder) GetCanModifyOk() (*bool, bool)`

GetCanModifyOk returns a tuple with the CanModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanModify

`func (o *PaperOrder) SetCanModify(v bool)`

SetCanModify sets CanModify field to given value.

### HasCanModify

`func (o *PaperOrder) HasCanModify() bool`

HasCanModify returns a boolean if a field has been set.

### GetCreateTime

`func (o *PaperOrder) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *PaperOrder) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *PaperOrder) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *PaperOrder) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreateTime0

`func (o *PaperOrder) GetCreateTime0() int64`

GetCreateTime0 returns the CreateTime0 field if non-nil, zero value otherwise.

### GetCreateTime0Ok

`func (o *PaperOrder) GetCreateTime0Ok() (*int64, bool)`

GetCreateTime0Ok returns a tuple with the CreateTime0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime0

`func (o *PaperOrder) SetCreateTime0(v int64)`

SetCreateTime0 sets CreateTime0 field to given value.

### HasCreateTime0

`func (o *PaperOrder) HasCreateTime0() bool`

HasCreateTime0 returns a boolean if a field has been set.

### GetFilledQuantity

`func (o *PaperOrder) GetFilledQuantity() string`

GetFilledQuantity returns the FilledQuantity field if non-nil, zero value otherwise.

### GetFilledQuantityOk

`func (o *PaperOrder) GetFilledQuantityOk() (*string, bool)`

GetFilledQuantityOk returns a tuple with the FilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQuantity

`func (o *PaperOrder) SetFilledQuantity(v string)`

SetFilledQuantity sets FilledQuantity field to given value.

### HasFilledQuantity

`func (o *PaperOrder) HasFilledQuantity() bool`

HasFilledQuantity returns a boolean if a field has been set.

### GetLmtPrice

`func (o *PaperOrder) GetLmtPrice() string`

GetLmtPrice returns the LmtPrice field if non-nil, zero value otherwise.

### GetLmtPriceOk

`func (o *PaperOrder) GetLmtPriceOk() (*string, bool)`

GetLmtPriceOk returns a tuple with the LmtPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmtPrice

`func (o *PaperOrder) SetLmtPrice(v string)`

SetLmtPrice sets LmtPrice field to given value.

### HasLmtPrice

`func (o *PaperOrder) HasLmtPrice() bool`

HasLmtPrice returns a boolean if a field has been set.

### GetOrderId

`func (o *PaperOrder) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PaperOrder) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PaperOrder) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PaperOrder) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderType

`func (o *PaperOrder) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *PaperOrder) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *PaperOrder) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *PaperOrder) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetOutsideRegularTradingHour

`func (o *PaperOrder) GetOutsideRegularTradingHour() bool`

GetOutsideRegularTradingHour returns the OutsideRegularTradingHour field if non-nil, zero value otherwise.

### GetOutsideRegularTradingHourOk

`func (o *PaperOrder) GetOutsideRegularTradingHourOk() (*bool, bool)`

GetOutsideRegularTradingHourOk returns a tuple with the OutsideRegularTradingHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideRegularTradingHour

`func (o *PaperOrder) SetOutsideRegularTradingHour(v bool)`

SetOutsideRegularTradingHour sets OutsideRegularTradingHour field to given value.

### HasOutsideRegularTradingHour

`func (o *PaperOrder) HasOutsideRegularTradingHour() bool`

HasOutsideRegularTradingHour returns a boolean if a field has been set.

### GetPaperId

`func (o *PaperOrder) GetPaperId() int32`

GetPaperId returns the PaperId field if non-nil, zero value otherwise.

### GetPaperIdOk

`func (o *PaperOrder) GetPaperIdOk() (*int32, bool)`

GetPaperIdOk returns a tuple with the PaperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaperId

`func (o *PaperOrder) SetPaperId(v int32)`

SetPaperId sets PaperId field to given value.

### HasPaperId

`func (o *PaperOrder) HasPaperId() bool`

HasPaperId returns a boolean if a field has been set.

### GetPlacedTime

`func (o *PaperOrder) GetPlacedTime() string`

GetPlacedTime returns the PlacedTime field if non-nil, zero value otherwise.

### GetPlacedTimeOk

`func (o *PaperOrder) GetPlacedTimeOk() (*string, bool)`

GetPlacedTimeOk returns a tuple with the PlacedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacedTime

`func (o *PaperOrder) SetPlacedTime(v string)`

SetPlacedTime sets PlacedTime field to given value.

### HasPlacedTime

`func (o *PaperOrder) HasPlacedTime() bool`

HasPlacedTime returns a boolean if a field has been set.

### GetStatus

`func (o *PaperOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaperOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaperOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaperOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusStr

`func (o *PaperOrder) GetStatusStr() string`

GetStatusStr returns the StatusStr field if non-nil, zero value otherwise.

### GetStatusStrOk

`func (o *PaperOrder) GetStatusStrOk() (*string, bool)`

GetStatusStrOk returns a tuple with the StatusStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusStr

`func (o *PaperOrder) SetStatusStr(v string)`

SetStatusStr sets StatusStr field to given value.

### HasStatusStr

`func (o *PaperOrder) HasStatusStr() bool`

HasStatusStr returns a boolean if a field has been set.

### GetTicker

`func (o *PaperOrder) GetTicker() GetOrdersItemOrdersInnerTicker`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *PaperOrder) GetTickerOk() (*GetOrdersItemOrdersInnerTicker, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *PaperOrder) SetTicker(v GetOrdersItemOrdersInnerTicker)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *PaperOrder) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PaperOrder) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PaperOrder) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PaperOrder) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PaperOrder) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetTotalQuantity

`func (o *PaperOrder) GetTotalQuantity() string`

GetTotalQuantity returns the TotalQuantity field if non-nil, zero value otherwise.

### GetTotalQuantityOk

`func (o *PaperOrder) GetTotalQuantityOk() (*string, bool)`

GetTotalQuantityOk returns a tuple with the TotalQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantity

`func (o *PaperOrder) SetTotalQuantity(v string)`

SetTotalQuantity sets TotalQuantity field to given value.

### HasTotalQuantity

`func (o *PaperOrder) HasTotalQuantity() bool`

HasTotalQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


