# GetOrdersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetOrdersItem**](GetOrdersItem.md) |  | [optional] 

## Methods

### NewGetOrdersResponse

`func NewGetOrdersResponse() *GetOrdersResponse`

NewGetOrdersResponse instantiates a new GetOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrdersResponseWithDefaults

`func NewGetOrdersResponseWithDefaults() *GetOrdersResponse`

NewGetOrdersResponseWithDefaults instantiates a new GetOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetOrdersResponse) GetData() []GetOrdersItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetOrdersResponse) GetDataOk() (*[]GetOrdersItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetOrdersResponse) SetData(v []GetOrdersItem)`

SetData sets Data field to given value.

### HasData

`func (o *GetOrdersResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


