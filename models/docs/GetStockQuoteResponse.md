# GetStockQuoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgVol10D** | Pointer to **string** |  | [optional] 
**AvgVol3M** | Pointer to **string** |  | [optional] 
**Beta3Y** | Pointer to **string** |  | [optional] 
**Bps** | Pointer to **string** |  | [optional] 
**Change** | Pointer to **string** |  | [optional] 
**ChangeRatio** | Pointer to **string** |  | [optional] 
**Close** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**DerivativeSupport** | Pointer to **int32** |  | [optional] 
**Dividend** | Pointer to **string** |  | [optional] 
**Eps** | Pointer to **string** |  | [optional] 
**EpsTtm** | Pointer to **string** |  | [optional] 
**FiftyTwoWkHigh** | Pointer to **string** |  | [optional] 
**FiftyTwoWkLow** | Pointer to **string** |  | [optional] 
**High** | Pointer to **string** |  | [optional] 
**InceptionDate** | Pointer to **string** |  | [optional] 
**LatestDividendDate** | Pointer to **string** |  | [optional] 
**LimitDown** | Pointer to **string** |  | [optional] 
**LimitUp** | Pointer to **string** |  | [optional] 
**LotSize** | Pointer to **string** |  | [optional] 
**Low** | Pointer to **string** |  | [optional] 
**NetAsset** | Pointer to **string** |  | [optional] 
**NetExpenseRatio** | Pointer to **string** |  | [optional] 
**NetValue** | Pointer to **string** |  | [optional] 
**NtvSize** | Pointer to **int32** |  | [optional] 
**Open** | Pointer to **string** |  | [optional] 
**PChRatio** | Pointer to **string** |  | [optional] 
**PChange** | Pointer to **string** |  | [optional] 
**PPrice** | Pointer to **string** |  | [optional] 
**Pe** | Pointer to **string** |  | [optional] 
**PeTtm** | Pointer to **string** |  | [optional] 
**PreClose** | Pointer to **string** |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**ReturnThisYear** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TickerId** | Pointer to **int32** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**TotalAsset** | Pointer to **string** |  | [optional] 
**TradeTime** | Pointer to **string** |  | [optional] 
**TzName** | Pointer to **string** |  | [optional] 
**VibrateRatio** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to **string** |  | [optional] 
**Yield** | Pointer to **string** |  | [optional] 
**Yield1Y** | Pointer to **string** |  | [optional] 

## Methods

### NewGetStockQuoteResponse

`func NewGetStockQuoteResponse() *GetStockQuoteResponse`

NewGetStockQuoteResponse instantiates a new GetStockQuoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStockQuoteResponseWithDefaults

`func NewGetStockQuoteResponseWithDefaults() *GetStockQuoteResponse`

NewGetStockQuoteResponseWithDefaults instantiates a new GetStockQuoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgVol10D

`func (o *GetStockQuoteResponse) GetAvgVol10D() string`

GetAvgVol10D returns the AvgVol10D field if non-nil, zero value otherwise.

### GetAvgVol10DOk

`func (o *GetStockQuoteResponse) GetAvgVol10DOk() (*string, bool)`

GetAvgVol10DOk returns a tuple with the AvgVol10D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgVol10D

`func (o *GetStockQuoteResponse) SetAvgVol10D(v string)`

SetAvgVol10D sets AvgVol10D field to given value.

### HasAvgVol10D

`func (o *GetStockQuoteResponse) HasAvgVol10D() bool`

HasAvgVol10D returns a boolean if a field has been set.

### GetAvgVol3M

`func (o *GetStockQuoteResponse) GetAvgVol3M() string`

GetAvgVol3M returns the AvgVol3M field if non-nil, zero value otherwise.

### GetAvgVol3MOk

`func (o *GetStockQuoteResponse) GetAvgVol3MOk() (*string, bool)`

GetAvgVol3MOk returns a tuple with the AvgVol3M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgVol3M

`func (o *GetStockQuoteResponse) SetAvgVol3M(v string)`

SetAvgVol3M sets AvgVol3M field to given value.

### HasAvgVol3M

`func (o *GetStockQuoteResponse) HasAvgVol3M() bool`

HasAvgVol3M returns a boolean if a field has been set.

### GetBeta3Y

`func (o *GetStockQuoteResponse) GetBeta3Y() string`

GetBeta3Y returns the Beta3Y field if non-nil, zero value otherwise.

### GetBeta3YOk

`func (o *GetStockQuoteResponse) GetBeta3YOk() (*string, bool)`

GetBeta3YOk returns a tuple with the Beta3Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta3Y

`func (o *GetStockQuoteResponse) SetBeta3Y(v string)`

SetBeta3Y sets Beta3Y field to given value.

### HasBeta3Y

`func (o *GetStockQuoteResponse) HasBeta3Y() bool`

HasBeta3Y returns a boolean if a field has been set.

### GetBps

`func (o *GetStockQuoteResponse) GetBps() string`

GetBps returns the Bps field if non-nil, zero value otherwise.

### GetBpsOk

`func (o *GetStockQuoteResponse) GetBpsOk() (*string, bool)`

GetBpsOk returns a tuple with the Bps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBps

`func (o *GetStockQuoteResponse) SetBps(v string)`

SetBps sets Bps field to given value.

### HasBps

`func (o *GetStockQuoteResponse) HasBps() bool`

HasBps returns a boolean if a field has been set.

### GetChange

`func (o *GetStockQuoteResponse) GetChange() string`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *GetStockQuoteResponse) GetChangeOk() (*string, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *GetStockQuoteResponse) SetChange(v string)`

SetChange sets Change field to given value.

### HasChange

`func (o *GetStockQuoteResponse) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetChangeRatio

`func (o *GetStockQuoteResponse) GetChangeRatio() string`

GetChangeRatio returns the ChangeRatio field if non-nil, zero value otherwise.

### GetChangeRatioOk

`func (o *GetStockQuoteResponse) GetChangeRatioOk() (*string, bool)`

GetChangeRatioOk returns a tuple with the ChangeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRatio

`func (o *GetStockQuoteResponse) SetChangeRatio(v string)`

SetChangeRatio sets ChangeRatio field to given value.

### HasChangeRatio

`func (o *GetStockQuoteResponse) HasChangeRatio() bool`

HasChangeRatio returns a boolean if a field has been set.

### GetClose

`func (o *GetStockQuoteResponse) GetClose() string`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *GetStockQuoteResponse) GetCloseOk() (*string, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *GetStockQuoteResponse) SetClose(v string)`

SetClose sets Close field to given value.

### HasClose

`func (o *GetStockQuoteResponse) HasClose() bool`

HasClose returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GetStockQuoteResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GetStockQuoteResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GetStockQuoteResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GetStockQuoteResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDerivativeSupport

`func (o *GetStockQuoteResponse) GetDerivativeSupport() int32`

GetDerivativeSupport returns the DerivativeSupport field if non-nil, zero value otherwise.

### GetDerivativeSupportOk

`func (o *GetStockQuoteResponse) GetDerivativeSupportOk() (*int32, bool)`

GetDerivativeSupportOk returns a tuple with the DerivativeSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivativeSupport

`func (o *GetStockQuoteResponse) SetDerivativeSupport(v int32)`

SetDerivativeSupport sets DerivativeSupport field to given value.

### HasDerivativeSupport

`func (o *GetStockQuoteResponse) HasDerivativeSupport() bool`

HasDerivativeSupport returns a boolean if a field has been set.

### GetDividend

`func (o *GetStockQuoteResponse) GetDividend() string`

GetDividend returns the Dividend field if non-nil, zero value otherwise.

### GetDividendOk

`func (o *GetStockQuoteResponse) GetDividendOk() (*string, bool)`

GetDividendOk returns a tuple with the Dividend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividend

`func (o *GetStockQuoteResponse) SetDividend(v string)`

SetDividend sets Dividend field to given value.

### HasDividend

`func (o *GetStockQuoteResponse) HasDividend() bool`

HasDividend returns a boolean if a field has been set.

### GetEps

`func (o *GetStockQuoteResponse) GetEps() string`

GetEps returns the Eps field if non-nil, zero value otherwise.

### GetEpsOk

`func (o *GetStockQuoteResponse) GetEpsOk() (*string, bool)`

GetEpsOk returns a tuple with the Eps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEps

`func (o *GetStockQuoteResponse) SetEps(v string)`

SetEps sets Eps field to given value.

### HasEps

`func (o *GetStockQuoteResponse) HasEps() bool`

HasEps returns a boolean if a field has been set.

### GetEpsTtm

`func (o *GetStockQuoteResponse) GetEpsTtm() string`

GetEpsTtm returns the EpsTtm field if non-nil, zero value otherwise.

### GetEpsTtmOk

`func (o *GetStockQuoteResponse) GetEpsTtmOk() (*string, bool)`

GetEpsTtmOk returns a tuple with the EpsTtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsTtm

`func (o *GetStockQuoteResponse) SetEpsTtm(v string)`

SetEpsTtm sets EpsTtm field to given value.

### HasEpsTtm

`func (o *GetStockQuoteResponse) HasEpsTtm() bool`

HasEpsTtm returns a boolean if a field has been set.

### GetFiftyTwoWkHigh

`func (o *GetStockQuoteResponse) GetFiftyTwoWkHigh() string`

GetFiftyTwoWkHigh returns the FiftyTwoWkHigh field if non-nil, zero value otherwise.

### GetFiftyTwoWkHighOk

`func (o *GetStockQuoteResponse) GetFiftyTwoWkHighOk() (*string, bool)`

GetFiftyTwoWkHighOk returns a tuple with the FiftyTwoWkHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWkHigh

`func (o *GetStockQuoteResponse) SetFiftyTwoWkHigh(v string)`

SetFiftyTwoWkHigh sets FiftyTwoWkHigh field to given value.

### HasFiftyTwoWkHigh

`func (o *GetStockQuoteResponse) HasFiftyTwoWkHigh() bool`

HasFiftyTwoWkHigh returns a boolean if a field has been set.

### GetFiftyTwoWkLow

`func (o *GetStockQuoteResponse) GetFiftyTwoWkLow() string`

GetFiftyTwoWkLow returns the FiftyTwoWkLow field if non-nil, zero value otherwise.

### GetFiftyTwoWkLowOk

`func (o *GetStockQuoteResponse) GetFiftyTwoWkLowOk() (*string, bool)`

GetFiftyTwoWkLowOk returns a tuple with the FiftyTwoWkLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWkLow

`func (o *GetStockQuoteResponse) SetFiftyTwoWkLow(v string)`

SetFiftyTwoWkLow sets FiftyTwoWkLow field to given value.

### HasFiftyTwoWkLow

`func (o *GetStockQuoteResponse) HasFiftyTwoWkLow() bool`

HasFiftyTwoWkLow returns a boolean if a field has been set.

### GetHigh

`func (o *GetStockQuoteResponse) GetHigh() string`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *GetStockQuoteResponse) GetHighOk() (*string, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *GetStockQuoteResponse) SetHigh(v string)`

SetHigh sets High field to given value.

### HasHigh

`func (o *GetStockQuoteResponse) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetInceptionDate

`func (o *GetStockQuoteResponse) GetInceptionDate() string`

GetInceptionDate returns the InceptionDate field if non-nil, zero value otherwise.

### GetInceptionDateOk

`func (o *GetStockQuoteResponse) GetInceptionDateOk() (*string, bool)`

GetInceptionDateOk returns a tuple with the InceptionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInceptionDate

`func (o *GetStockQuoteResponse) SetInceptionDate(v string)`

SetInceptionDate sets InceptionDate field to given value.

### HasInceptionDate

`func (o *GetStockQuoteResponse) HasInceptionDate() bool`

HasInceptionDate returns a boolean if a field has been set.

### GetLatestDividendDate

`func (o *GetStockQuoteResponse) GetLatestDividendDate() string`

GetLatestDividendDate returns the LatestDividendDate field if non-nil, zero value otherwise.

### GetLatestDividendDateOk

`func (o *GetStockQuoteResponse) GetLatestDividendDateOk() (*string, bool)`

GetLatestDividendDateOk returns a tuple with the LatestDividendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDividendDate

`func (o *GetStockQuoteResponse) SetLatestDividendDate(v string)`

SetLatestDividendDate sets LatestDividendDate field to given value.

### HasLatestDividendDate

`func (o *GetStockQuoteResponse) HasLatestDividendDate() bool`

HasLatestDividendDate returns a boolean if a field has been set.

### GetLimitDown

`func (o *GetStockQuoteResponse) GetLimitDown() string`

GetLimitDown returns the LimitDown field if non-nil, zero value otherwise.

### GetLimitDownOk

`func (o *GetStockQuoteResponse) GetLimitDownOk() (*string, bool)`

GetLimitDownOk returns a tuple with the LimitDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDown

`func (o *GetStockQuoteResponse) SetLimitDown(v string)`

SetLimitDown sets LimitDown field to given value.

### HasLimitDown

`func (o *GetStockQuoteResponse) HasLimitDown() bool`

HasLimitDown returns a boolean if a field has been set.

### GetLimitUp

`func (o *GetStockQuoteResponse) GetLimitUp() string`

GetLimitUp returns the LimitUp field if non-nil, zero value otherwise.

### GetLimitUpOk

`func (o *GetStockQuoteResponse) GetLimitUpOk() (*string, bool)`

GetLimitUpOk returns a tuple with the LimitUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitUp

`func (o *GetStockQuoteResponse) SetLimitUp(v string)`

SetLimitUp sets LimitUp field to given value.

### HasLimitUp

`func (o *GetStockQuoteResponse) HasLimitUp() bool`

HasLimitUp returns a boolean if a field has been set.

### GetLotSize

`func (o *GetStockQuoteResponse) GetLotSize() string`

GetLotSize returns the LotSize field if non-nil, zero value otherwise.

### GetLotSizeOk

`func (o *GetStockQuoteResponse) GetLotSizeOk() (*string, bool)`

GetLotSizeOk returns a tuple with the LotSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotSize

`func (o *GetStockQuoteResponse) SetLotSize(v string)`

SetLotSize sets LotSize field to given value.

### HasLotSize

`func (o *GetStockQuoteResponse) HasLotSize() bool`

HasLotSize returns a boolean if a field has been set.

### GetLow

`func (o *GetStockQuoteResponse) GetLow() string`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *GetStockQuoteResponse) GetLowOk() (*string, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *GetStockQuoteResponse) SetLow(v string)`

SetLow sets Low field to given value.

### HasLow

`func (o *GetStockQuoteResponse) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetNetAsset

`func (o *GetStockQuoteResponse) GetNetAsset() string`

GetNetAsset returns the NetAsset field if non-nil, zero value otherwise.

### GetNetAssetOk

`func (o *GetStockQuoteResponse) GetNetAssetOk() (*string, bool)`

GetNetAssetOk returns a tuple with the NetAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAsset

`func (o *GetStockQuoteResponse) SetNetAsset(v string)`

SetNetAsset sets NetAsset field to given value.

### HasNetAsset

`func (o *GetStockQuoteResponse) HasNetAsset() bool`

HasNetAsset returns a boolean if a field has been set.

### GetNetExpenseRatio

`func (o *GetStockQuoteResponse) GetNetExpenseRatio() string`

GetNetExpenseRatio returns the NetExpenseRatio field if non-nil, zero value otherwise.

### GetNetExpenseRatioOk

`func (o *GetStockQuoteResponse) GetNetExpenseRatioOk() (*string, bool)`

GetNetExpenseRatioOk returns a tuple with the NetExpenseRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetExpenseRatio

`func (o *GetStockQuoteResponse) SetNetExpenseRatio(v string)`

SetNetExpenseRatio sets NetExpenseRatio field to given value.

### HasNetExpenseRatio

`func (o *GetStockQuoteResponse) HasNetExpenseRatio() bool`

HasNetExpenseRatio returns a boolean if a field has been set.

### GetNetValue

`func (o *GetStockQuoteResponse) GetNetValue() string`

GetNetValue returns the NetValue field if non-nil, zero value otherwise.

### GetNetValueOk

`func (o *GetStockQuoteResponse) GetNetValueOk() (*string, bool)`

GetNetValueOk returns a tuple with the NetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetValue

`func (o *GetStockQuoteResponse) SetNetValue(v string)`

SetNetValue sets NetValue field to given value.

### HasNetValue

`func (o *GetStockQuoteResponse) HasNetValue() bool`

HasNetValue returns a boolean if a field has been set.

### GetNtvSize

`func (o *GetStockQuoteResponse) GetNtvSize() int32`

GetNtvSize returns the NtvSize field if non-nil, zero value otherwise.

### GetNtvSizeOk

`func (o *GetStockQuoteResponse) GetNtvSizeOk() (*int32, bool)`

GetNtvSizeOk returns a tuple with the NtvSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtvSize

`func (o *GetStockQuoteResponse) SetNtvSize(v int32)`

SetNtvSize sets NtvSize field to given value.

### HasNtvSize

`func (o *GetStockQuoteResponse) HasNtvSize() bool`

HasNtvSize returns a boolean if a field has been set.

### GetOpen

`func (o *GetStockQuoteResponse) GetOpen() string`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *GetStockQuoteResponse) GetOpenOk() (*string, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *GetStockQuoteResponse) SetOpen(v string)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *GetStockQuoteResponse) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetPChRatio

`func (o *GetStockQuoteResponse) GetPChRatio() string`

GetPChRatio returns the PChRatio field if non-nil, zero value otherwise.

### GetPChRatioOk

`func (o *GetStockQuoteResponse) GetPChRatioOk() (*string, bool)`

GetPChRatioOk returns a tuple with the PChRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPChRatio

`func (o *GetStockQuoteResponse) SetPChRatio(v string)`

SetPChRatio sets PChRatio field to given value.

### HasPChRatio

`func (o *GetStockQuoteResponse) HasPChRatio() bool`

HasPChRatio returns a boolean if a field has been set.

### GetPChange

`func (o *GetStockQuoteResponse) GetPChange() string`

GetPChange returns the PChange field if non-nil, zero value otherwise.

### GetPChangeOk

`func (o *GetStockQuoteResponse) GetPChangeOk() (*string, bool)`

GetPChangeOk returns a tuple with the PChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPChange

`func (o *GetStockQuoteResponse) SetPChange(v string)`

SetPChange sets PChange field to given value.

### HasPChange

`func (o *GetStockQuoteResponse) HasPChange() bool`

HasPChange returns a boolean if a field has been set.

### GetPPrice

`func (o *GetStockQuoteResponse) GetPPrice() string`

GetPPrice returns the PPrice field if non-nil, zero value otherwise.

### GetPPriceOk

`func (o *GetStockQuoteResponse) GetPPriceOk() (*string, bool)`

GetPPriceOk returns a tuple with the PPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPPrice

`func (o *GetStockQuoteResponse) SetPPrice(v string)`

SetPPrice sets PPrice field to given value.

### HasPPrice

`func (o *GetStockQuoteResponse) HasPPrice() bool`

HasPPrice returns a boolean if a field has been set.

### GetPe

`func (o *GetStockQuoteResponse) GetPe() string`

GetPe returns the Pe field if non-nil, zero value otherwise.

### GetPeOk

`func (o *GetStockQuoteResponse) GetPeOk() (*string, bool)`

GetPeOk returns a tuple with the Pe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPe

`func (o *GetStockQuoteResponse) SetPe(v string)`

SetPe sets Pe field to given value.

### HasPe

`func (o *GetStockQuoteResponse) HasPe() bool`

HasPe returns a boolean if a field has been set.

### GetPeTtm

`func (o *GetStockQuoteResponse) GetPeTtm() string`

GetPeTtm returns the PeTtm field if non-nil, zero value otherwise.

### GetPeTtmOk

`func (o *GetStockQuoteResponse) GetPeTtmOk() (*string, bool)`

GetPeTtmOk returns a tuple with the PeTtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeTtm

`func (o *GetStockQuoteResponse) SetPeTtm(v string)`

SetPeTtm sets PeTtm field to given value.

### HasPeTtm

`func (o *GetStockQuoteResponse) HasPeTtm() bool`

HasPeTtm returns a boolean if a field has been set.

### GetPreClose

`func (o *GetStockQuoteResponse) GetPreClose() string`

GetPreClose returns the PreClose field if non-nil, zero value otherwise.

### GetPreCloseOk

`func (o *GetStockQuoteResponse) GetPreCloseOk() (*string, bool)`

GetPreCloseOk returns a tuple with the PreClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreClose

`func (o *GetStockQuoteResponse) SetPreClose(v string)`

SetPreClose sets PreClose field to given value.

### HasPreClose

`func (o *GetStockQuoteResponse) HasPreClose() bool`

HasPreClose returns a boolean if a field has been set.

### GetRating

`func (o *GetStockQuoteResponse) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *GetStockQuoteResponse) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *GetStockQuoteResponse) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *GetStockQuoteResponse) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetReturnThisYear

`func (o *GetStockQuoteResponse) GetReturnThisYear() string`

GetReturnThisYear returns the ReturnThisYear field if non-nil, zero value otherwise.

### GetReturnThisYearOk

`func (o *GetStockQuoteResponse) GetReturnThisYearOk() (*string, bool)`

GetReturnThisYearOk returns a tuple with the ReturnThisYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnThisYear

`func (o *GetStockQuoteResponse) SetReturnThisYear(v string)`

SetReturnThisYear sets ReturnThisYear field to given value.

### HasReturnThisYear

`func (o *GetStockQuoteResponse) HasReturnThisYear() bool`

HasReturnThisYear returns a boolean if a field has been set.

### GetStatus

`func (o *GetStockQuoteResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetStockQuoteResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetStockQuoteResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetStockQuoteResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTickerId

`func (o *GetStockQuoteResponse) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *GetStockQuoteResponse) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *GetStockQuoteResponse) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *GetStockQuoteResponse) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetTimeZone

`func (o *GetStockQuoteResponse) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetStockQuoteResponse) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetStockQuoteResponse) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetStockQuoteResponse) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetTotalAsset

`func (o *GetStockQuoteResponse) GetTotalAsset() string`

GetTotalAsset returns the TotalAsset field if non-nil, zero value otherwise.

### GetTotalAssetOk

`func (o *GetStockQuoteResponse) GetTotalAssetOk() (*string, bool)`

GetTotalAssetOk returns a tuple with the TotalAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAsset

`func (o *GetStockQuoteResponse) SetTotalAsset(v string)`

SetTotalAsset sets TotalAsset field to given value.

### HasTotalAsset

`func (o *GetStockQuoteResponse) HasTotalAsset() bool`

HasTotalAsset returns a boolean if a field has been set.

### GetTradeTime

`func (o *GetStockQuoteResponse) GetTradeTime() string`

GetTradeTime returns the TradeTime field if non-nil, zero value otherwise.

### GetTradeTimeOk

`func (o *GetStockQuoteResponse) GetTradeTimeOk() (*string, bool)`

GetTradeTimeOk returns a tuple with the TradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTime

`func (o *GetStockQuoteResponse) SetTradeTime(v string)`

SetTradeTime sets TradeTime field to given value.

### HasTradeTime

`func (o *GetStockQuoteResponse) HasTradeTime() bool`

HasTradeTime returns a boolean if a field has been set.

### GetTzName

`func (o *GetStockQuoteResponse) GetTzName() string`

GetTzName returns the TzName field if non-nil, zero value otherwise.

### GetTzNameOk

`func (o *GetStockQuoteResponse) GetTzNameOk() (*string, bool)`

GetTzNameOk returns a tuple with the TzName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTzName

`func (o *GetStockQuoteResponse) SetTzName(v string)`

SetTzName sets TzName field to given value.

### HasTzName

`func (o *GetStockQuoteResponse) HasTzName() bool`

HasTzName returns a boolean if a field has been set.

### GetVibrateRatio

`func (o *GetStockQuoteResponse) GetVibrateRatio() string`

GetVibrateRatio returns the VibrateRatio field if non-nil, zero value otherwise.

### GetVibrateRatioOk

`func (o *GetStockQuoteResponse) GetVibrateRatioOk() (*string, bool)`

GetVibrateRatioOk returns a tuple with the VibrateRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVibrateRatio

`func (o *GetStockQuoteResponse) SetVibrateRatio(v string)`

SetVibrateRatio sets VibrateRatio field to given value.

### HasVibrateRatio

`func (o *GetStockQuoteResponse) HasVibrateRatio() bool`

HasVibrateRatio returns a boolean if a field has been set.

### GetVolume

`func (o *GetStockQuoteResponse) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GetStockQuoteResponse) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GetStockQuoteResponse) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GetStockQuoteResponse) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetYield

`func (o *GetStockQuoteResponse) GetYield() string`

GetYield returns the Yield field if non-nil, zero value otherwise.

### GetYieldOk

`func (o *GetStockQuoteResponse) GetYieldOk() (*string, bool)`

GetYieldOk returns a tuple with the Yield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYield

`func (o *GetStockQuoteResponse) SetYield(v string)`

SetYield sets Yield field to given value.

### HasYield

`func (o *GetStockQuoteResponse) HasYield() bool`

HasYield returns a boolean if a field has been set.

### GetYield1Y

`func (o *GetStockQuoteResponse) GetYield1Y() string`

GetYield1Y returns the Yield1Y field if non-nil, zero value otherwise.

### GetYield1YOk

`func (o *GetStockQuoteResponse) GetYield1YOk() (*string, bool)`

GetYield1YOk returns a tuple with the Yield1Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYield1Y

`func (o *GetStockQuoteResponse) SetYield1Y(v string)`

SetYield1Y sets Yield1Y field to given value.

### HasYield1Y

`func (o *GetStockQuoteResponse) HasYield1Y() bool`

HasYield1Y returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


