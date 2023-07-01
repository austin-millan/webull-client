# GetStockOptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Change** | Pointer to **string** |  | [optional] 
**ChangeRatio** | Pointer to **string** |  | [optional] 
**Close** | Pointer to **string** |  | [optional] 
**DisExchangeCode** | Pointer to **string** |  | [optional] 
**DisSymbol** | Pointer to **string** |  | [optional] 
**ExpireDate** | Pointer to **string** |  | [optional] 
**ExpireDateList** | Pointer to [**[]GetStockOptionsResponseExpireDateListInner**](GetStockOptionsResponseExpireDateListInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TickerId** | Pointer to **int32** |  | [optional] 
**UnSymbol** | Pointer to **string** |  | [optional] 

## Methods

### NewGetStockOptionsResponse

`func NewGetStockOptionsResponse() *GetStockOptionsResponse`

NewGetStockOptionsResponse instantiates a new GetStockOptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStockOptionsResponseWithDefaults

`func NewGetStockOptionsResponseWithDefaults() *GetStockOptionsResponse`

NewGetStockOptionsResponseWithDefaults instantiates a new GetStockOptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChange

`func (o *GetStockOptionsResponse) GetChange() string`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *GetStockOptionsResponse) GetChangeOk() (*string, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *GetStockOptionsResponse) SetChange(v string)`

SetChange sets Change field to given value.

### HasChange

`func (o *GetStockOptionsResponse) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetChangeRatio

`func (o *GetStockOptionsResponse) GetChangeRatio() string`

GetChangeRatio returns the ChangeRatio field if non-nil, zero value otherwise.

### GetChangeRatioOk

`func (o *GetStockOptionsResponse) GetChangeRatioOk() (*string, bool)`

GetChangeRatioOk returns a tuple with the ChangeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRatio

`func (o *GetStockOptionsResponse) SetChangeRatio(v string)`

SetChangeRatio sets ChangeRatio field to given value.

### HasChangeRatio

`func (o *GetStockOptionsResponse) HasChangeRatio() bool`

HasChangeRatio returns a boolean if a field has been set.

### GetClose

`func (o *GetStockOptionsResponse) GetClose() string`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *GetStockOptionsResponse) GetCloseOk() (*string, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *GetStockOptionsResponse) SetClose(v string)`

SetClose sets Close field to given value.

### HasClose

`func (o *GetStockOptionsResponse) HasClose() bool`

HasClose returns a boolean if a field has been set.

### GetDisExchangeCode

`func (o *GetStockOptionsResponse) GetDisExchangeCode() string`

GetDisExchangeCode returns the DisExchangeCode field if non-nil, zero value otherwise.

### GetDisExchangeCodeOk

`func (o *GetStockOptionsResponse) GetDisExchangeCodeOk() (*string, bool)`

GetDisExchangeCodeOk returns a tuple with the DisExchangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisExchangeCode

`func (o *GetStockOptionsResponse) SetDisExchangeCode(v string)`

SetDisExchangeCode sets DisExchangeCode field to given value.

### HasDisExchangeCode

`func (o *GetStockOptionsResponse) HasDisExchangeCode() bool`

HasDisExchangeCode returns a boolean if a field has been set.

### GetDisSymbol

`func (o *GetStockOptionsResponse) GetDisSymbol() string`

GetDisSymbol returns the DisSymbol field if non-nil, zero value otherwise.

### GetDisSymbolOk

`func (o *GetStockOptionsResponse) GetDisSymbolOk() (*string, bool)`

GetDisSymbolOk returns a tuple with the DisSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisSymbol

`func (o *GetStockOptionsResponse) SetDisSymbol(v string)`

SetDisSymbol sets DisSymbol field to given value.

### HasDisSymbol

`func (o *GetStockOptionsResponse) HasDisSymbol() bool`

HasDisSymbol returns a boolean if a field has been set.

### GetExpireDate

`func (o *GetStockOptionsResponse) GetExpireDate() string`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *GetStockOptionsResponse) GetExpireDateOk() (*string, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *GetStockOptionsResponse) SetExpireDate(v string)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *GetStockOptionsResponse) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetExpireDateList

`func (o *GetStockOptionsResponse) GetExpireDateList() []GetStockOptionsResponseExpireDateListInner`

GetExpireDateList returns the ExpireDateList field if non-nil, zero value otherwise.

### GetExpireDateListOk

`func (o *GetStockOptionsResponse) GetExpireDateListOk() (*[]GetStockOptionsResponseExpireDateListInner, bool)`

GetExpireDateListOk returns a tuple with the ExpireDateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDateList

`func (o *GetStockOptionsResponse) SetExpireDateList(v []GetStockOptionsResponseExpireDateListInner)`

SetExpireDateList sets ExpireDateList field to given value.

### HasExpireDateList

`func (o *GetStockOptionsResponse) HasExpireDateList() bool`

HasExpireDateList returns a boolean if a field has been set.

### GetName

`func (o *GetStockOptionsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetStockOptionsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetStockOptionsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetStockOptionsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTickerId

`func (o *GetStockOptionsResponse) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *GetStockOptionsResponse) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *GetStockOptionsResponse) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *GetStockOptionsResponse) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetUnSymbol

`func (o *GetStockOptionsResponse) GetUnSymbol() string`

GetUnSymbol returns the UnSymbol field if non-nil, zero value otherwise.

### GetUnSymbolOk

`func (o *GetStockOptionsResponse) GetUnSymbolOk() (*string, bool)`

GetUnSymbolOk returns a tuple with the UnSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnSymbol

`func (o *GetStockOptionsResponse) SetUnSymbol(v string)`

SetUnSymbol sets UnSymbol field to given value.

### HasUnSymbol

`func (o *GetStockOptionsResponse) HasUnSymbol() bool`

HasUnSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


