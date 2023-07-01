# Leg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**OrderSide**](OrderSide.md) |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **string** |  | [optional] 
**FilledQuantity** | Pointer to **string** |  | [optional] 
**FilledTime** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Strategy** | Pointer to **string** |  | [optional] 
**Strike** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Ticker** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewLeg

`func NewLeg() *Leg`

NewLeg instantiates a new Leg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegWithDefaults

`func NewLegWithDefaults() *Leg`

NewLegWithDefaults instantiates a new Leg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Leg) GetAction() OrderSide`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Leg) GetActionOk() (*OrderSide, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Leg) SetAction(v OrderSide)`

SetAction sets Action field to given value.

### HasAction

`func (o *Leg) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAmount

`func (o *Leg) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Leg) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Leg) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Leg) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetExpiration

`func (o *Leg) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Leg) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Leg) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *Leg) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetFilledQuantity

`func (o *Leg) GetFilledQuantity() string`

GetFilledQuantity returns the FilledQuantity field if non-nil, zero value otherwise.

### GetFilledQuantityOk

`func (o *Leg) GetFilledQuantityOk() (*string, bool)`

GetFilledQuantityOk returns a tuple with the FilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQuantity

`func (o *Leg) SetFilledQuantity(v string)`

SetFilledQuantity sets FilledQuantity field to given value.

### HasFilledQuantity

`func (o *Leg) HasFilledQuantity() bool`

HasFilledQuantity returns a boolean if a field has been set.

### GetFilledTime

`func (o *Leg) GetFilledTime() string`

GetFilledTime returns the FilledTime field if non-nil, zero value otherwise.

### GetFilledTimeOk

`func (o *Leg) GetFilledTimeOk() (*string, bool)`

GetFilledTimeOk returns a tuple with the FilledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledTime

`func (o *Leg) SetFilledTime(v string)`

SetFilledTime sets FilledTime field to given value.

### HasFilledTime

`func (o *Leg) HasFilledTime() bool`

HasFilledTime returns a boolean if a field has been set.

### GetPrice

`func (o *Leg) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Leg) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Leg) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Leg) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *Leg) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Leg) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Leg) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Leg) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSide

`func (o *Leg) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *Leg) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *Leg) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *Leg) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStrategy

`func (o *Leg) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *Leg) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *Leg) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *Leg) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetStrike

`func (o *Leg) GetStrike() string`

GetStrike returns the Strike field if non-nil, zero value otherwise.

### GetStrikeOk

`func (o *Leg) GetStrikeOk() (*string, bool)`

GetStrikeOk returns a tuple with the Strike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrike

`func (o *Leg) SetStrike(v string)`

SetStrike sets Strike field to given value.

### HasStrike

`func (o *Leg) HasStrike() bool`

HasStrike returns a boolean if a field has been set.

### GetSymbol

`func (o *Leg) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Leg) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Leg) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Leg) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTicker

`func (o *Leg) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *Leg) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *Leg) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *Leg) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetType

`func (o *Leg) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Leg) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Leg) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Leg) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


