# Transfers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**[]Transfer**](Transfer.md) |  | [optional] 

## Methods

### NewTransfers

`func NewTransfers() *Transfers`

NewTransfers instantiates a new Transfers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransfersWithDefaults

`func NewTransfersWithDefaults() *Transfers`

NewTransfersWithDefaults instantiates a new Transfers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *Transfers) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Transfers) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Transfers) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *Transfers) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *Transfers) GetData() []Transfer`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Transfers) GetDataOk() (*[]Transfer, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Transfers) SetData(v []Transfer)`

SetData sets Data field to given value.

### HasData

`func (o *Transfers) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


