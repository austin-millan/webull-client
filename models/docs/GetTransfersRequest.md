# GetTransfersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageSize** | Pointer to **float32** |  | [optional] [default to 256]
**LastRecordId** | Pointer to **string** |  | [optional] [default to "0"]

## Methods

### NewGetTransfersRequest

`func NewGetTransfersRequest() *GetTransfersRequest`

NewGetTransfersRequest instantiates a new GetTransfersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransfersRequestWithDefaults

`func NewGetTransfersRequestWithDefaults() *GetTransfersRequest`

NewGetTransfersRequestWithDefaults instantiates a new GetTransfersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageSize

`func (o *GetTransfersRequest) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetTransfersRequest) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetTransfersRequest) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetTransfersRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetLastRecordId

`func (o *GetTransfersRequest) GetLastRecordId() string`

GetLastRecordId returns the LastRecordId field if non-nil, zero value otherwise.

### GetLastRecordIdOk

`func (o *GetTransfersRequest) GetLastRecordIdOk() (*string, bool)`

GetLastRecordIdOk returns a tuple with the LastRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRecordId

`func (o *GetTransfersRequest) SetLastRecordId(v string)`

SetLastRecordId sets LastRecordId field to given value.

### HasLastRecordId

`func (o *GetTransfersRequest) HasLastRecordId() bool`

HasLastRecordId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


