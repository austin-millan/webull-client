# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComboTickerType** | Pointer to **string** | Example: &#39;option&#39; | [optional] 
**Legs** | Pointer to [**[]Leg**](Leg.md) |  | [optional] 
**StatusCode** | Pointer to **string** |  | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComboTickerType

`func (o *Order) GetComboTickerType() string`

GetComboTickerType returns the ComboTickerType field if non-nil, zero value otherwise.

### GetComboTickerTypeOk

`func (o *Order) GetComboTickerTypeOk() (*string, bool)`

GetComboTickerTypeOk returns a tuple with the ComboTickerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComboTickerType

`func (o *Order) SetComboTickerType(v string)`

SetComboTickerType sets ComboTickerType field to given value.

### HasComboTickerType

`func (o *Order) HasComboTickerType() bool`

HasComboTickerType returns a boolean if a field has been set.

### GetLegs

`func (o *Order) GetLegs() []Leg`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *Order) GetLegsOk() (*[]Leg, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *Order) SetLegs(v []Leg)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *Order) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetStatusCode

`func (o *Order) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *Order) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *Order) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *Order) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


