# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventWarning** | Pointer to [**AlertEventWarning**](AlertEventWarning.md) |  | [optional] 
**TickerWarning** | Pointer to [**AlertTickerWarning**](AlertTickerWarning.md) |  | [optional] 
**WarningInput** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAlert

`func NewAlert() *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventWarning

`func (o *Alert) GetEventWarning() AlertEventWarning`

GetEventWarning returns the EventWarning field if non-nil, zero value otherwise.

### GetEventWarningOk

`func (o *Alert) GetEventWarningOk() (*AlertEventWarning, bool)`

GetEventWarningOk returns a tuple with the EventWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventWarning

`func (o *Alert) SetEventWarning(v AlertEventWarning)`

SetEventWarning sets EventWarning field to given value.

### HasEventWarning

`func (o *Alert) HasEventWarning() bool`

HasEventWarning returns a boolean if a field has been set.

### GetTickerWarning

`func (o *Alert) GetTickerWarning() AlertTickerWarning`

GetTickerWarning returns the TickerWarning field if non-nil, zero value otherwise.

### GetTickerWarningOk

`func (o *Alert) GetTickerWarningOk() (*AlertTickerWarning, bool)`

GetTickerWarningOk returns a tuple with the TickerWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerWarning

`func (o *Alert) SetTickerWarning(v AlertTickerWarning)`

SetTickerWarning sets TickerWarning field to given value.

### HasTickerWarning

`func (o *Alert) HasTickerWarning() bool`

HasTickerWarning returns a boolean if a field has been set.

### GetWarningInput

`func (o *Alert) GetWarningInput() map[string]interface{}`

GetWarningInput returns the WarningInput field if non-nil, zero value otherwise.

### GetWarningInputOk

`func (o *Alert) GetWarningInputOk() (*map[string]interface{}, bool)`

GetWarningInputOk returns a tuple with the WarningInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningInput

`func (o *Alert) SetWarningInput(v map[string]interface{})`

SetWarningInput sets WarningInput field to given value.

### HasWarningInput

`func (o *Alert) HasWarningInput() bool`

HasWarningInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


