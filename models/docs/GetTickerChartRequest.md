# GetTickerChartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **string** | Number of bars to return | [optional] 
**ExtendTrading** | Pointer to **int32** | Whether to include pre-market and afterhours bars. &#39;1&#39; is used for pre-market and after-hours bars. | [optional] [default to 1]
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewGetTickerChartRequest

`func NewGetTickerChartRequest() *GetTickerChartRequest`

NewGetTickerChartRequest instantiates a new GetTickerChartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTickerChartRequestWithDefaults

`func NewGetTickerChartRequestWithDefaults() *GetTickerChartRequest`

NewGetTickerChartRequestWithDefaults instantiates a new GetTickerChartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetTickerChartRequest) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetTickerChartRequest) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetTickerChartRequest) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetTickerChartRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetExtendTrading

`func (o *GetTickerChartRequest) GetExtendTrading() int32`

GetExtendTrading returns the ExtendTrading field if non-nil, zero value otherwise.

### GetExtendTradingOk

`func (o *GetTickerChartRequest) GetExtendTradingOk() (*int32, bool)`

GetExtendTradingOk returns a tuple with the ExtendTrading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendTrading

`func (o *GetTickerChartRequest) SetExtendTrading(v int32)`

SetExtendTrading sets ExtendTrading field to given value.

### HasExtendTrading

`func (o *GetTickerChartRequest) HasExtendTrading() bool`

HasExtendTrading returns a boolean if a field has been set.

### GetType

`func (o *GetTickerChartRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTickerChartRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTickerChartRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetTickerChartRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


