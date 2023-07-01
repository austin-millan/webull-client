# SmartRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **string** |  | [optional] 
**AlertRuleKey** | Pointer to **string** |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**Remark** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Example: &#39;earnPre&#39;, &#39;fastUp&#39;, &#39;fastDown&#39;, &#39;week52Up&#39;, &#39;week52Down&#39;, &#39;day5Down&#39; | [optional] 

## Methods

### NewSmartRule

`func NewSmartRule() *SmartRule`

NewSmartRule instantiates a new SmartRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartRuleWithDefaults

`func NewSmartRuleWithDefaults() *SmartRule`

NewSmartRuleWithDefaults instantiates a new SmartRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *SmartRule) GetActive() string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SmartRule) GetActiveOk() (*string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SmartRule) SetActive(v string)`

SetActive sets Active field to given value.

### HasActive

`func (o *SmartRule) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAlertRuleKey

`func (o *SmartRule) GetAlertRuleKey() string`

GetAlertRuleKey returns the AlertRuleKey field if non-nil, zero value otherwise.

### GetAlertRuleKeyOk

`func (o *SmartRule) GetAlertRuleKeyOk() (*string, bool)`

GetAlertRuleKeyOk returns a tuple with the AlertRuleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRuleKey

`func (o *SmartRule) SetAlertRuleKey(v string)`

SetAlertRuleKey sets AlertRuleKey field to given value.

### HasAlertRuleKey

`func (o *SmartRule) HasAlertRuleKey() bool`

HasAlertRuleKey returns a boolean if a field has been set.

### GetField

`func (o *SmartRule) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SmartRule) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SmartRule) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *SmartRule) HasField() bool`

HasField returns a boolean if a field has been set.

### GetRemark

`func (o *SmartRule) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *SmartRule) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *SmartRule) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *SmartRule) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### GetType

`func (o *SmartRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SmartRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SmartRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SmartRule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


