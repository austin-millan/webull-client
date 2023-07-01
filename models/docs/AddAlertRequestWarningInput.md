# AddAlertRequestWarningInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]SmartRule**](SmartRule.md) |  | [optional] 
**TickerId** | Pointer to **int32** |  | [optional] [default to 913243251]
**WarningFrequency** | Pointer to **int32** | 1 is once a day, 2 is once a minute | [optional] [default to 1]
**WarningInterval** | Pointer to **int32** | 1 is once, 0 is repeating | [optional] [default to 1]

## Methods

### NewAddAlertRequestWarningInput

`func NewAddAlertRequestWarningInput() *AddAlertRequestWarningInput`

NewAddAlertRequestWarningInput instantiates a new AddAlertRequestWarningInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAlertRequestWarningInputWithDefaults

`func NewAddAlertRequestWarningInputWithDefaults() *AddAlertRequestWarningInput`

NewAddAlertRequestWarningInputWithDefaults instantiates a new AddAlertRequestWarningInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *AddAlertRequestWarningInput) GetRules() []SmartRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AddAlertRequestWarningInput) GetRulesOk() (*[]SmartRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AddAlertRequestWarningInput) SetRules(v []SmartRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *AddAlertRequestWarningInput) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTickerId

`func (o *AddAlertRequestWarningInput) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *AddAlertRequestWarningInput) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *AddAlertRequestWarningInput) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *AddAlertRequestWarningInput) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetWarningFrequency

`func (o *AddAlertRequestWarningInput) GetWarningFrequency() int32`

GetWarningFrequency returns the WarningFrequency field if non-nil, zero value otherwise.

### GetWarningFrequencyOk

`func (o *AddAlertRequestWarningInput) GetWarningFrequencyOk() (*int32, bool)`

GetWarningFrequencyOk returns a tuple with the WarningFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningFrequency

`func (o *AddAlertRequestWarningInput) SetWarningFrequency(v int32)`

SetWarningFrequency sets WarningFrequency field to given value.

### HasWarningFrequency

`func (o *AddAlertRequestWarningInput) HasWarningFrequency() bool`

HasWarningFrequency returns a boolean if a field has been set.

### GetWarningInterval

`func (o *AddAlertRequestWarningInput) GetWarningInterval() int32`

GetWarningInterval returns the WarningInterval field if non-nil, zero value otherwise.

### GetWarningIntervalOk

`func (o *AddAlertRequestWarningInput) GetWarningIntervalOk() (*int32, bool)`

GetWarningIntervalOk returns a tuple with the WarningInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningInterval

`func (o *AddAlertRequestWarningInput) SetWarningInterval(v int32)`

SetWarningInterval sets WarningInterval field to given value.

### HasWarningInterval

`func (o *AddAlertRequestWarningInput) HasWarningInterval() bool`

HasWarningInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


