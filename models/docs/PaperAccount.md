# PaperAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**CurrencyId** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**OrderTypes** | Pointer to [**[]OrderType**](OrderType.md) |  | [optional] 
**PaperId** | Pointer to **int32** |  | [optional] 
**PaperName** | Pointer to **string** |  | [optional] 
**PaperTickerPoolCode** | Pointer to **string** |  | [optional] 
**PaperType** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**SupportOutsideRth** | Pointer to **bool** |  | [optional] 
**TimeInForces** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPaperAccount

`func NewPaperAccount() *PaperAccount`

NewPaperAccount instantiates a new PaperAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaperAccountWithDefaults

`func NewPaperAccountWithDefaults() *PaperAccount`

NewPaperAccountWithDefaults instantiates a new PaperAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PaperAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaperAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaperAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaperAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrencyId

`func (o *PaperAccount) GetCurrencyId() int32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *PaperAccount) GetCurrencyIdOk() (*int32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *PaperAccount) SetCurrencyId(v int32)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *PaperAccount) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetId

`func (o *PaperAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaperAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaperAccount) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PaperAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderTypes

`func (o *PaperAccount) GetOrderTypes() []OrderType`

GetOrderTypes returns the OrderTypes field if non-nil, zero value otherwise.

### GetOrderTypesOk

`func (o *PaperAccount) GetOrderTypesOk() (*[]OrderType, bool)`

GetOrderTypesOk returns a tuple with the OrderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypes

`func (o *PaperAccount) SetOrderTypes(v []OrderType)`

SetOrderTypes sets OrderTypes field to given value.

### HasOrderTypes

`func (o *PaperAccount) HasOrderTypes() bool`

HasOrderTypes returns a boolean if a field has been set.

### GetPaperId

`func (o *PaperAccount) GetPaperId() int32`

GetPaperId returns the PaperId field if non-nil, zero value otherwise.

### GetPaperIdOk

`func (o *PaperAccount) GetPaperIdOk() (*int32, bool)`

GetPaperIdOk returns a tuple with the PaperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaperId

`func (o *PaperAccount) SetPaperId(v int32)`

SetPaperId sets PaperId field to given value.

### HasPaperId

`func (o *PaperAccount) HasPaperId() bool`

HasPaperId returns a boolean if a field has been set.

### GetPaperName

`func (o *PaperAccount) GetPaperName() string`

GetPaperName returns the PaperName field if non-nil, zero value otherwise.

### GetPaperNameOk

`func (o *PaperAccount) GetPaperNameOk() (*string, bool)`

GetPaperNameOk returns a tuple with the PaperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaperName

`func (o *PaperAccount) SetPaperName(v string)`

SetPaperName sets PaperName field to given value.

### HasPaperName

`func (o *PaperAccount) HasPaperName() bool`

HasPaperName returns a boolean if a field has been set.

### GetPaperTickerPoolCode

`func (o *PaperAccount) GetPaperTickerPoolCode() string`

GetPaperTickerPoolCode returns the PaperTickerPoolCode field if non-nil, zero value otherwise.

### GetPaperTickerPoolCodeOk

`func (o *PaperAccount) GetPaperTickerPoolCodeOk() (*string, bool)`

GetPaperTickerPoolCodeOk returns a tuple with the PaperTickerPoolCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaperTickerPoolCode

`func (o *PaperAccount) SetPaperTickerPoolCode(v string)`

SetPaperTickerPoolCode sets PaperTickerPoolCode field to given value.

### HasPaperTickerPoolCode

`func (o *PaperAccount) HasPaperTickerPoolCode() bool`

HasPaperTickerPoolCode returns a boolean if a field has been set.

### GetPaperType

`func (o *PaperAccount) GetPaperType() int32`

GetPaperType returns the PaperType field if non-nil, zero value otherwise.

### GetPaperTypeOk

`func (o *PaperAccount) GetPaperTypeOk() (*int32, bool)`

GetPaperTypeOk returns a tuple with the PaperType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaperType

`func (o *PaperAccount) SetPaperType(v int32)`

SetPaperType sets PaperType field to given value.

### HasPaperType

`func (o *PaperAccount) HasPaperType() bool`

HasPaperType returns a boolean if a field has been set.

### GetStatus

`func (o *PaperAccount) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaperAccount) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaperAccount) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaperAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportOutsideRth

`func (o *PaperAccount) GetSupportOutsideRth() bool`

GetSupportOutsideRth returns the SupportOutsideRth field if non-nil, zero value otherwise.

### GetSupportOutsideRthOk

`func (o *PaperAccount) GetSupportOutsideRthOk() (*bool, bool)`

GetSupportOutsideRthOk returns a tuple with the SupportOutsideRth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportOutsideRth

`func (o *PaperAccount) SetSupportOutsideRth(v bool)`

SetSupportOutsideRth sets SupportOutsideRth field to given value.

### HasSupportOutsideRth

`func (o *PaperAccount) HasSupportOutsideRth() bool`

HasSupportOutsideRth returns a boolean if a field has been set.

### GetTimeInForces

`func (o *PaperAccount) GetTimeInForces() []string`

GetTimeInForces returns the TimeInForces field if non-nil, zero value otherwise.

### GetTimeInForcesOk

`func (o *PaperAccount) GetTimeInForcesOk() (*[]string, bool)`

GetTimeInForcesOk returns a tuple with the TimeInForces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForces

`func (o *PaperAccount) SetTimeInForces(v []string)`

SetTimeInForces sets TimeInForces field to given value.

### HasTimeInForces

`func (o *PaperAccount) HasTimeInForces() bool`

HasTimeInForces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


