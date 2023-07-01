# GetDividendsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DividendTotal** | Pointer to **string** |  | [optional] 
**HasTax** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetDividendsResponse

`func NewGetDividendsResponse() *GetDividendsResponse`

NewGetDividendsResponse instantiates a new GetDividendsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDividendsResponseWithDefaults

`func NewGetDividendsResponseWithDefaults() *GetDividendsResponse`

NewGetDividendsResponseWithDefaults instantiates a new GetDividendsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDividendTotal

`func (o *GetDividendsResponse) GetDividendTotal() string`

GetDividendTotal returns the DividendTotal field if non-nil, zero value otherwise.

### GetDividendTotalOk

`func (o *GetDividendsResponse) GetDividendTotalOk() (*string, bool)`

GetDividendTotalOk returns a tuple with the DividendTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendTotal

`func (o *GetDividendsResponse) SetDividendTotal(v string)`

SetDividendTotal sets DividendTotal field to given value.

### HasDividendTotal

`func (o *GetDividendsResponse) HasDividendTotal() bool`

HasDividendTotal returns a boolean if a field has been set.

### GetHasTax

`func (o *GetDividendsResponse) GetHasTax() bool`

GetHasTax returns the HasTax field if non-nil, zero value otherwise.

### GetHasTaxOk

`func (o *GetDividendsResponse) GetHasTaxOk() (*bool, bool)`

GetHasTaxOk returns a tuple with the HasTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTax

`func (o *GetDividendsResponse) SetHasTax(v bool)`

SetHasTax sets HasTax field to given value.

### HasHasTax

`func (o *GetDividendsResponse) HasHasTax() bool`

HasHasTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


