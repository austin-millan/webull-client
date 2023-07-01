# AlertEventWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remove** | Pointer to **bool** |  | [optional] 
**Rules** | Pointer to [**[]SmartRule**](SmartRule.md) |  | [optional] 

## Methods

### NewAlertEventWarning

`func NewAlertEventWarning() *AlertEventWarning`

NewAlertEventWarning instantiates a new AlertEventWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertEventWarningWithDefaults

`func NewAlertEventWarningWithDefaults() *AlertEventWarning`

NewAlertEventWarningWithDefaults instantiates a new AlertEventWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemove

`func (o *AlertEventWarning) GetRemove() bool`

GetRemove returns the Remove field if non-nil, zero value otherwise.

### GetRemoveOk

`func (o *AlertEventWarning) GetRemoveOk() (*bool, bool)`

GetRemoveOk returns a tuple with the Remove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemove

`func (o *AlertEventWarning) SetRemove(v bool)`

SetRemove sets Remove field to given value.

### HasRemove

`func (o *AlertEventWarning) HasRemove() bool`

HasRemove returns a boolean if a field has been set.

### GetRules

`func (o *AlertEventWarning) GetRules() []SmartRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AlertEventWarning) GetRulesOk() (*[]SmartRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AlertEventWarning) SetRules(v []SmartRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *AlertEventWarning) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


