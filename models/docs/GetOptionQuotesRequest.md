# GetOptionQuotesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DerivativeIds** | Pointer to **int32** |  | [optional] 
**TickerId** | Pointer to **int32** | tickerId | [optional] 

## Methods

### NewGetOptionQuotesRequest

`func NewGetOptionQuotesRequest() *GetOptionQuotesRequest`

NewGetOptionQuotesRequest instantiates a new GetOptionQuotesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOptionQuotesRequestWithDefaults

`func NewGetOptionQuotesRequestWithDefaults() *GetOptionQuotesRequest`

NewGetOptionQuotesRequestWithDefaults instantiates a new GetOptionQuotesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDerivativeIds

`func (o *GetOptionQuotesRequest) GetDerivativeIds() int32`

GetDerivativeIds returns the DerivativeIds field if non-nil, zero value otherwise.

### GetDerivativeIdsOk

`func (o *GetOptionQuotesRequest) GetDerivativeIdsOk() (*int32, bool)`

GetDerivativeIdsOk returns a tuple with the DerivativeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivativeIds

`func (o *GetOptionQuotesRequest) SetDerivativeIds(v int32)`

SetDerivativeIds sets DerivativeIds field to given value.

### HasDerivativeIds

`func (o *GetOptionQuotesRequest) HasDerivativeIds() bool`

HasDerivativeIds returns a boolean if a field has been set.

### GetTickerId

`func (o *GetOptionQuotesRequest) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *GetOptionQuotesRequest) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *GetOptionQuotesRequest) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *GetOptionQuotesRequest) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


