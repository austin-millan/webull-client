# GetIsTradeableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetIsTradeableResponseDataInner**](GetIsTradeableResponseDataInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetIsTradeableResponse

`func NewGetIsTradeableResponse() *GetIsTradeableResponse`

NewGetIsTradeableResponse instantiates a new GetIsTradeableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIsTradeableResponseWithDefaults

`func NewGetIsTradeableResponseWithDefaults() *GetIsTradeableResponse`

NewGetIsTradeableResponseWithDefaults instantiates a new GetIsTradeableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetIsTradeableResponse) GetData() []GetIsTradeableResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetIsTradeableResponse) GetDataOk() (*[]GetIsTradeableResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetIsTradeableResponse) SetData(v []GetIsTradeableResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetIsTradeableResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSuccess

`func (o *GetIsTradeableResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetIsTradeableResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetIsTradeableResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetIsTradeableResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


