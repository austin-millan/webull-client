# AddAlertRequestEventWarningInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Del** | Pointer to **bool** |  | [optional] 
**Remove** | Pointer to **bool** |  | [optional] 
**Rules** | Pointer to [**[]SmartRule**](SmartRule.md) |  | [optional] 
**TickerId** | Pointer to **int32** |  | [optional] [default to 913243251]

## Methods

### NewAddAlertRequestEventWarningInput

`func NewAddAlertRequestEventWarningInput() *AddAlertRequestEventWarningInput`

NewAddAlertRequestEventWarningInput instantiates a new AddAlertRequestEventWarningInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAlertRequestEventWarningInputWithDefaults

`func NewAddAlertRequestEventWarningInputWithDefaults() *AddAlertRequestEventWarningInput`

NewAddAlertRequestEventWarningInputWithDefaults instantiates a new AddAlertRequestEventWarningInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDel

`func (o *AddAlertRequestEventWarningInput) GetDel() bool`

GetDel returns the Del field if non-nil, zero value otherwise.

### GetDelOk

`func (o *AddAlertRequestEventWarningInput) GetDelOk() (*bool, bool)`

GetDelOk returns a tuple with the Del field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDel

`func (o *AddAlertRequestEventWarningInput) SetDel(v bool)`

SetDel sets Del field to given value.

### HasDel

`func (o *AddAlertRequestEventWarningInput) HasDel() bool`

HasDel returns a boolean if a field has been set.

### GetRemove

`func (o *AddAlertRequestEventWarningInput) GetRemove() bool`

GetRemove returns the Remove field if non-nil, zero value otherwise.

### GetRemoveOk

`func (o *AddAlertRequestEventWarningInput) GetRemoveOk() (*bool, bool)`

GetRemoveOk returns a tuple with the Remove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemove

`func (o *AddAlertRequestEventWarningInput) SetRemove(v bool)`

SetRemove sets Remove field to given value.

### HasRemove

`func (o *AddAlertRequestEventWarningInput) HasRemove() bool`

HasRemove returns a boolean if a field has been set.

### GetRules

`func (o *AddAlertRequestEventWarningInput) GetRules() []SmartRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AddAlertRequestEventWarningInput) GetRulesOk() (*[]SmartRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AddAlertRequestEventWarningInput) SetRules(v []SmartRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *AddAlertRequestEventWarningInput) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTickerId

`func (o *AddAlertRequestEventWarningInput) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *AddAlertRequestEventWarningInput) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *AddAlertRequestEventWarningInput) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *AddAlertRequestEventWarningInput) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


