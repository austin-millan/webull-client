# GetAlertsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]GetAlertsResponseDataInner**](GetAlertsResponseDataInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetAlertsResponse

`func NewGetAlertsResponse() *GetAlertsResponse`

NewGetAlertsResponse instantiates a new GetAlertsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlertsResponseWithDefaults

`func NewGetAlertsResponseWithDefaults() *GetAlertsResponse`

NewGetAlertsResponseWithDefaults instantiates a new GetAlertsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetAlertsResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetAlertsResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetAlertsResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetAlertsResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetAlertsResponse) GetData() []GetAlertsResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAlertsResponse) GetDataOk() (*[]GetAlertsResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAlertsResponse) SetData(v []GetAlertsResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetAlertsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSuccess

`func (o *GetAlertsResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetAlertsResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetAlertsResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetAlertsResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


