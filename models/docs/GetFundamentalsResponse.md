# GetFundamentalsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analysis** | Pointer to [**GetFundamentalsResponseAnalysis**](GetFundamentalsResponseAnalysis.md) |  | [optional] 
**Forecast** | Pointer to [**[]GetFundamentalsResponseForecastInner**](GetFundamentalsResponseForecastInner.md) |  | [optional] 
**Remind** | Pointer to [**GetFundamentalsResponseRemind**](GetFundamentalsResponseRemind.md) |  | [optional] 
**SimpleStatement** | Pointer to [**[]GetFundamentalsResponseSimpleStatementInner**](GetFundamentalsResponseSimpleStatementInner.md) |  | [optional] 

## Methods

### NewGetFundamentalsResponse

`func NewGetFundamentalsResponse() *GetFundamentalsResponse`

NewGetFundamentalsResponse instantiates a new GetFundamentalsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFundamentalsResponseWithDefaults

`func NewGetFundamentalsResponseWithDefaults() *GetFundamentalsResponse`

NewGetFundamentalsResponseWithDefaults instantiates a new GetFundamentalsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysis

`func (o *GetFundamentalsResponse) GetAnalysis() GetFundamentalsResponseAnalysis`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *GetFundamentalsResponse) GetAnalysisOk() (*GetFundamentalsResponseAnalysis, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *GetFundamentalsResponse) SetAnalysis(v GetFundamentalsResponseAnalysis)`

SetAnalysis sets Analysis field to given value.

### HasAnalysis

`func (o *GetFundamentalsResponse) HasAnalysis() bool`

HasAnalysis returns a boolean if a field has been set.

### GetForecast

`func (o *GetFundamentalsResponse) GetForecast() []GetFundamentalsResponseForecastInner`

GetForecast returns the Forecast field if non-nil, zero value otherwise.

### GetForecastOk

`func (o *GetFundamentalsResponse) GetForecastOk() (*[]GetFundamentalsResponseForecastInner, bool)`

GetForecastOk returns a tuple with the Forecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecast

`func (o *GetFundamentalsResponse) SetForecast(v []GetFundamentalsResponseForecastInner)`

SetForecast sets Forecast field to given value.

### HasForecast

`func (o *GetFundamentalsResponse) HasForecast() bool`

HasForecast returns a boolean if a field has been set.

### GetRemind

`func (o *GetFundamentalsResponse) GetRemind() GetFundamentalsResponseRemind`

GetRemind returns the Remind field if non-nil, zero value otherwise.

### GetRemindOk

`func (o *GetFundamentalsResponse) GetRemindOk() (*GetFundamentalsResponseRemind, bool)`

GetRemindOk returns a tuple with the Remind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemind

`func (o *GetFundamentalsResponse) SetRemind(v GetFundamentalsResponseRemind)`

SetRemind sets Remind field to given value.

### HasRemind

`func (o *GetFundamentalsResponse) HasRemind() bool`

HasRemind returns a boolean if a field has been set.

### GetSimpleStatement

`func (o *GetFundamentalsResponse) GetSimpleStatement() []GetFundamentalsResponseSimpleStatementInner`

GetSimpleStatement returns the SimpleStatement field if non-nil, zero value otherwise.

### GetSimpleStatementOk

`func (o *GetFundamentalsResponse) GetSimpleStatementOk() (*[]GetFundamentalsResponseSimpleStatementInner, bool)`

GetSimpleStatementOk returns a tuple with the SimpleStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleStatement

`func (o *GetFundamentalsResponse) SetSimpleStatement(v []GetFundamentalsResponseSimpleStatementInner)`

SetSimpleStatement sets SimpleStatement field to given value.

### HasSimpleStatement

`func (o *GetFundamentalsResponse) HasSimpleStatement() bool`

HasSimpleStatement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


