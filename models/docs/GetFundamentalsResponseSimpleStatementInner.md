# GetFundamentalsResponseSimpleStatementInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyName** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to [**[]GetFundamentalsResponseSimpleStatementInnerLabelsInner**](GetFundamentalsResponseSimpleStatementInnerLabelsInner.md) |  | [optional] 
**List** | Pointer to [**[]GetFundamentalsResponseSimpleStatementInnerListInner**](GetFundamentalsResponseSimpleStatementInnerListInner.md) |  | [optional] 
**ReportType** | Pointer to **int32** |  | [optional] 
**Single** | Pointer to [**GetFundamentalsResponseSimpleStatementInnerSingle**](GetFundamentalsResponseSimpleStatementInnerSingle.md) |  | [optional] 
**StatementType** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewGetFundamentalsResponseSimpleStatementInner

`func NewGetFundamentalsResponseSimpleStatementInner() *GetFundamentalsResponseSimpleStatementInner`

NewGetFundamentalsResponseSimpleStatementInner instantiates a new GetFundamentalsResponseSimpleStatementInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFundamentalsResponseSimpleStatementInnerWithDefaults

`func NewGetFundamentalsResponseSimpleStatementInnerWithDefaults() *GetFundamentalsResponseSimpleStatementInner`

NewGetFundamentalsResponseSimpleStatementInnerWithDefaults instantiates a new GetFundamentalsResponseSimpleStatementInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyName

`func (o *GetFundamentalsResponseSimpleStatementInner) GetCurrencyName() string`

GetCurrencyName returns the CurrencyName field if non-nil, zero value otherwise.

### GetCurrencyNameOk

`func (o *GetFundamentalsResponseSimpleStatementInner) GetCurrencyNameOk() (*string, bool)`

GetCurrencyNameOk returns a tuple with the CurrencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyName

`func (o *GetFundamentalsResponseSimpleStatementInner) SetCurrencyName(v string)`

SetCurrencyName sets CurrencyName field to given value.

### HasCurrencyName

`func (o *GetFundamentalsResponseSimpleStatementInner) HasCurrencyName() bool`

HasCurrencyName returns a boolean if a field has been set.

### GetLabels

`func (o *GetFundamentalsResponseSimpleStatementInner) GetLabels() []GetFundamentalsResponseSimpleStatementInnerLabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetFundamentalsResponseSimpleStatementInner) GetLabelsOk() (*[]GetFundamentalsResponseSimpleStatementInnerLabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetFundamentalsResponseSimpleStatementInner) SetLabels(v []GetFundamentalsResponseSimpleStatementInnerLabelsInner)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetFundamentalsResponseSimpleStatementInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetList

`func (o *GetFundamentalsResponseSimpleStatementInner) GetList() []GetFundamentalsResponseSimpleStatementInnerListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *GetFundamentalsResponseSimpleStatementInner) GetListOk() (*[]GetFundamentalsResponseSimpleStatementInnerListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *GetFundamentalsResponseSimpleStatementInner) SetList(v []GetFundamentalsResponseSimpleStatementInnerListInner)`

SetList sets List field to given value.

### HasList

`func (o *GetFundamentalsResponseSimpleStatementInner) HasList() bool`

HasList returns a boolean if a field has been set.

### GetReportType

`func (o *GetFundamentalsResponseSimpleStatementInner) GetReportType() int32`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *GetFundamentalsResponseSimpleStatementInner) GetReportTypeOk() (*int32, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *GetFundamentalsResponseSimpleStatementInner) SetReportType(v int32)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *GetFundamentalsResponseSimpleStatementInner) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetSingle

`func (o *GetFundamentalsResponseSimpleStatementInner) GetSingle() GetFundamentalsResponseSimpleStatementInnerSingle`

GetSingle returns the Single field if non-nil, zero value otherwise.

### GetSingleOk

`func (o *GetFundamentalsResponseSimpleStatementInner) GetSingleOk() (*GetFundamentalsResponseSimpleStatementInnerSingle, bool)`

GetSingleOk returns a tuple with the Single field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingle

`func (o *GetFundamentalsResponseSimpleStatementInner) SetSingle(v GetFundamentalsResponseSimpleStatementInnerSingle)`

SetSingle sets Single field to given value.

### HasSingle

`func (o *GetFundamentalsResponseSimpleStatementInner) HasSingle() bool`

HasSingle returns a boolean if a field has been set.

### GetStatementType

`func (o *GetFundamentalsResponseSimpleStatementInner) GetStatementType() int32`

GetStatementType returns the StatementType field if non-nil, zero value otherwise.

### GetStatementTypeOk

`func (o *GetFundamentalsResponseSimpleStatementInner) GetStatementTypeOk() (*int32, bool)`

GetStatementTypeOk returns a tuple with the StatementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementType

`func (o *GetFundamentalsResponseSimpleStatementInner) SetStatementType(v int32)`

SetStatementType sets StatementType field to given value.

### HasStatementType

`func (o *GetFundamentalsResponseSimpleStatementInner) HasStatementType() bool`

HasStatementType returns a boolean if a field has been set.

### GetTitle

`func (o *GetFundamentalsResponseSimpleStatementInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetFundamentalsResponseSimpleStatementInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetFundamentalsResponseSimpleStatementInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetFundamentalsResponseSimpleStatementInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


