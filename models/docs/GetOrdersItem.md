# GetOrdersItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComboId** | Pointer to **string** |  | [optional] 
**ComboTickerType** | Pointer to **string** |  | [optional] 
**ComboType** | Pointer to **string** |  | [optional] 
**Orders** | Pointer to [**[]GetOrdersItemOrdersInner**](GetOrdersItemOrdersInner.md) |  | [optional] 
**OutsideRegularTradingHour** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetOrdersItem

`func NewGetOrdersItem() *GetOrdersItem`

NewGetOrdersItem instantiates a new GetOrdersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrdersItemWithDefaults

`func NewGetOrdersItemWithDefaults() *GetOrdersItem`

NewGetOrdersItemWithDefaults instantiates a new GetOrdersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComboId

`func (o *GetOrdersItem) GetComboId() string`

GetComboId returns the ComboId field if non-nil, zero value otherwise.

### GetComboIdOk

`func (o *GetOrdersItem) GetComboIdOk() (*string, bool)`

GetComboIdOk returns a tuple with the ComboId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComboId

`func (o *GetOrdersItem) SetComboId(v string)`

SetComboId sets ComboId field to given value.

### HasComboId

`func (o *GetOrdersItem) HasComboId() bool`

HasComboId returns a boolean if a field has been set.

### GetComboTickerType

`func (o *GetOrdersItem) GetComboTickerType() string`

GetComboTickerType returns the ComboTickerType field if non-nil, zero value otherwise.

### GetComboTickerTypeOk

`func (o *GetOrdersItem) GetComboTickerTypeOk() (*string, bool)`

GetComboTickerTypeOk returns a tuple with the ComboTickerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComboTickerType

`func (o *GetOrdersItem) SetComboTickerType(v string)`

SetComboTickerType sets ComboTickerType field to given value.

### HasComboTickerType

`func (o *GetOrdersItem) HasComboTickerType() bool`

HasComboTickerType returns a boolean if a field has been set.

### GetComboType

`func (o *GetOrdersItem) GetComboType() string`

GetComboType returns the ComboType field if non-nil, zero value otherwise.

### GetComboTypeOk

`func (o *GetOrdersItem) GetComboTypeOk() (*string, bool)`

GetComboTypeOk returns a tuple with the ComboType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComboType

`func (o *GetOrdersItem) SetComboType(v string)`

SetComboType sets ComboType field to given value.

### HasComboType

`func (o *GetOrdersItem) HasComboType() bool`

HasComboType returns a boolean if a field has been set.

### GetOrders

`func (o *GetOrdersItem) GetOrders() []GetOrdersItemOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *GetOrdersItem) GetOrdersOk() (*[]GetOrdersItemOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *GetOrdersItem) SetOrders(v []GetOrdersItemOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *GetOrdersItem) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetOutsideRegularTradingHour

`func (o *GetOrdersItem) GetOutsideRegularTradingHour() bool`

GetOutsideRegularTradingHour returns the OutsideRegularTradingHour field if non-nil, zero value otherwise.

### GetOutsideRegularTradingHourOk

`func (o *GetOrdersItem) GetOutsideRegularTradingHourOk() (*bool, bool)`

GetOutsideRegularTradingHourOk returns a tuple with the OutsideRegularTradingHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideRegularTradingHour

`func (o *GetOrdersItem) SetOutsideRegularTradingHour(v bool)`

SetOutsideRegularTradingHour sets OutsideRegularTradingHour field to given value.

### HasOutsideRegularTradingHour

`func (o *GetOrdersItem) HasOutsideRegularTradingHour() bool`

HasOutsideRegularTradingHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


