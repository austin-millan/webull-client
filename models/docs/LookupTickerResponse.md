# LookupTickerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **int32** |  | [optional] 
**CategoryName** | Pointer to **string** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**List** | Pointer to [**[]LookupTickerResponseListInner**](LookupTickerResponseListInner.md) |  | [optional] 

## Methods

### NewLookupTickerResponse

`func NewLookupTickerResponse() *LookupTickerResponse`

NewLookupTickerResponse instantiates a new LookupTickerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupTickerResponseWithDefaults

`func NewLookupTickerResponseWithDefaults() *LookupTickerResponse`

NewLookupTickerResponseWithDefaults instantiates a new LookupTickerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *LookupTickerResponse) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *LookupTickerResponse) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *LookupTickerResponse) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *LookupTickerResponse) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetCategoryName

`func (o *LookupTickerResponse) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *LookupTickerResponse) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *LookupTickerResponse) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *LookupTickerResponse) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### GetHasMore

`func (o *LookupTickerResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *LookupTickerResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *LookupTickerResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *LookupTickerResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetList

`func (o *LookupTickerResponse) GetList() []LookupTickerResponseListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *LookupTickerResponse) GetListOk() (*[]LookupTickerResponseListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *LookupTickerResponse) SetList(v []LookupTickerResponseListInner)`

SetList sets List field to given value.

### HasList

`func (o *LookupTickerResponse) HasList() bool`

HasList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


