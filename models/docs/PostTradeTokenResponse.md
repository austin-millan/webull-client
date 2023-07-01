# PostTradeTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**PostTradeTokenResponseData**](PostTradeTokenResponseData.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewPostTradeTokenResponse

`func NewPostTradeTokenResponse() *PostTradeTokenResponse`

NewPostTradeTokenResponse instantiates a new PostTradeTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTradeTokenResponseWithDefaults

`func NewPostTradeTokenResponseWithDefaults() *PostTradeTokenResponse`

NewPostTradeTokenResponseWithDefaults instantiates a new PostTradeTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PostTradeTokenResponse) GetData() PostTradeTokenResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PostTradeTokenResponse) GetDataOk() (*PostTradeTokenResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PostTradeTokenResponse) SetData(v PostTradeTokenResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *PostTradeTokenResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSuccess

`func (o *PostTradeTokenResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PostTradeTokenResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PostTradeTokenResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *PostTradeTokenResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


