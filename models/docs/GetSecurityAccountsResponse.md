# GetSecurityAccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetSecurityAccountsResponseDataInner**](GetSecurityAccountsResponseDataInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetSecurityAccountsResponse

`func NewGetSecurityAccountsResponse() *GetSecurityAccountsResponse`

NewGetSecurityAccountsResponse instantiates a new GetSecurityAccountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecurityAccountsResponseWithDefaults

`func NewGetSecurityAccountsResponseWithDefaults() *GetSecurityAccountsResponse`

NewGetSecurityAccountsResponseWithDefaults instantiates a new GetSecurityAccountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSecurityAccountsResponse) GetData() []GetSecurityAccountsResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSecurityAccountsResponse) GetDataOk() (*[]GetSecurityAccountsResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSecurityAccountsResponse) SetData(v []GetSecurityAccountsResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetSecurityAccountsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSuccess

`func (o *GetSecurityAccountsResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetSecurityAccountsResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetSecurityAccountsResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetSecurityAccountsResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


