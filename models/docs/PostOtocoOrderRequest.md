# PostOtocoOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewOrders** | Pointer to [**[]PostStockOrderRequest**](PostStockOrderRequest.md) |  | [optional] 

## Methods

### NewPostOtocoOrderRequest

`func NewPostOtocoOrderRequest() *PostOtocoOrderRequest`

NewPostOtocoOrderRequest instantiates a new PostOtocoOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostOtocoOrderRequestWithDefaults

`func NewPostOtocoOrderRequestWithDefaults() *PostOtocoOrderRequest`

NewPostOtocoOrderRequestWithDefaults instantiates a new PostOtocoOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewOrders

`func (o *PostOtocoOrderRequest) GetNewOrders() []PostStockOrderRequest`

GetNewOrders returns the NewOrders field if non-nil, zero value otherwise.

### GetNewOrdersOk

`func (o *PostOtocoOrderRequest) GetNewOrdersOk() (*[]PostStockOrderRequest, bool)`

GetNewOrdersOk returns a tuple with the NewOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrders

`func (o *PostOtocoOrderRequest) SetNewOrders(v []PostStockOrderRequest)`

SetNewOrders sets NewOrders field to given value.

### HasNewOrders

`func (o *PostOtocoOrderRequest) HasNewOrders() bool`

HasNewOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


