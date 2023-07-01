# GetTickerChartResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]string** |  | [optional] 
**Dates** | Pointer to [**[]GetTickerChartResponseDatesInner**](GetTickerChartResponseDatesInner.md) |  | [optional] 
**ExchangeStatus** | Pointer to **bool** |  | [optional] 
**HasMore** | Pointer to **int32** |  | [optional] 
**PreClose** | Pointer to **string** |  | [optional] 
**TickerId** | Pointer to **int32** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 

## Methods

### NewGetTickerChartResponse

`func NewGetTickerChartResponse() *GetTickerChartResponse`

NewGetTickerChartResponse instantiates a new GetTickerChartResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTickerChartResponseWithDefaults

`func NewGetTickerChartResponseWithDefaults() *GetTickerChartResponse`

NewGetTickerChartResponseWithDefaults instantiates a new GetTickerChartResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTickerChartResponse) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTickerChartResponse) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTickerChartResponse) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *GetTickerChartResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDates

`func (o *GetTickerChartResponse) GetDates() []GetTickerChartResponseDatesInner`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *GetTickerChartResponse) GetDatesOk() (*[]GetTickerChartResponseDatesInner, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *GetTickerChartResponse) SetDates(v []GetTickerChartResponseDatesInner)`

SetDates sets Dates field to given value.

### HasDates

`func (o *GetTickerChartResponse) HasDates() bool`

HasDates returns a boolean if a field has been set.

### GetExchangeStatus

`func (o *GetTickerChartResponse) GetExchangeStatus() bool`

GetExchangeStatus returns the ExchangeStatus field if non-nil, zero value otherwise.

### GetExchangeStatusOk

`func (o *GetTickerChartResponse) GetExchangeStatusOk() (*bool, bool)`

GetExchangeStatusOk returns a tuple with the ExchangeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeStatus

`func (o *GetTickerChartResponse) SetExchangeStatus(v bool)`

SetExchangeStatus sets ExchangeStatus field to given value.

### HasExchangeStatus

`func (o *GetTickerChartResponse) HasExchangeStatus() bool`

HasExchangeStatus returns a boolean if a field has been set.

### GetHasMore

`func (o *GetTickerChartResponse) GetHasMore() int32`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetTickerChartResponse) GetHasMoreOk() (*int32, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetTickerChartResponse) SetHasMore(v int32)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetTickerChartResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetPreClose

`func (o *GetTickerChartResponse) GetPreClose() string`

GetPreClose returns the PreClose field if non-nil, zero value otherwise.

### GetPreCloseOk

`func (o *GetTickerChartResponse) GetPreCloseOk() (*string, bool)`

GetPreCloseOk returns a tuple with the PreClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreClose

`func (o *GetTickerChartResponse) SetPreClose(v string)`

SetPreClose sets PreClose field to given value.

### HasPreClose

`func (o *GetTickerChartResponse) HasPreClose() bool`

HasPreClose returns a boolean if a field has been set.

### GetTickerId

`func (o *GetTickerChartResponse) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *GetTickerChartResponse) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *GetTickerChartResponse) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *GetTickerChartResponse) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetTimeZone

`func (o *GetTickerChartResponse) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetTickerChartResponse) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetTickerChartResponse) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetTickerChartResponse) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


