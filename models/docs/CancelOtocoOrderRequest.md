# CancelOtocoOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialId** | Pointer to **string** |  | [optional] 
**CancelOrders** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCancelOtocoOrderRequest

`func NewCancelOtocoOrderRequest() *CancelOtocoOrderRequest`

NewCancelOtocoOrderRequest instantiates a new CancelOtocoOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelOtocoOrderRequestWithDefaults

`func NewCancelOtocoOrderRequestWithDefaults() *CancelOtocoOrderRequest`

NewCancelOtocoOrderRequestWithDefaults instantiates a new CancelOtocoOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialId

`func (o *CancelOtocoOrderRequest) GetSerialId() string`

GetSerialId returns the SerialId field if non-nil, zero value otherwise.

### GetSerialIdOk

`func (o *CancelOtocoOrderRequest) GetSerialIdOk() (*string, bool)`

GetSerialIdOk returns a tuple with the SerialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialId

`func (o *CancelOtocoOrderRequest) SetSerialId(v string)`

SetSerialId sets SerialId field to given value.

### HasSerialId

`func (o *CancelOtocoOrderRequest) HasSerialId() bool`

HasSerialId returns a boolean if a field has been set.

### GetCancelOrders

`func (o *CancelOtocoOrderRequest) GetCancelOrders() []string`

GetCancelOrders returns the CancelOrders field if non-nil, zero value otherwise.

### GetCancelOrdersOk

`func (o *CancelOtocoOrderRequest) GetCancelOrdersOk() (*[]string, bool)`

GetCancelOrdersOk returns a tuple with the CancelOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOrders

`func (o *CancelOtocoOrderRequest) SetCancelOrders(v []string)`

SetCancelOrders sets CancelOrders field to given value.

### HasCancelOrders

`func (o *CancelOtocoOrderRequest) HasCancelOrders() bool`

HasCancelOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


