# GetAcccountRequestPositionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetType** | Pointer to **string** |  | [optional] 
**BrokerId** | Pointer to **int32** |  | [optional] 
**Cost** | Pointer to **string** |  | [optional] 
**CostPrice** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ExchangeRate** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LastOpenTime** | Pointer to **string** |  | [optional] 
**LastPrice** | Pointer to **string** |  | [optional] 
**Lock** | Pointer to **bool** |  | [optional] 
**MarketValue** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **string** |  | [optional] 
**PositionProportion** | Pointer to **string** |  | [optional] 
**Ticker** | Pointer to [**GetAcccountRequestPositionsInnerTicker**](GetAcccountRequestPositionsInnerTicker.md) |  | [optional] 
**UnrealizedProfitLoss** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLossRate** | Pointer to **string** |  | [optional] 
**UpdatePositionTimeStamp** | Pointer to **float32** |  | [optional] 

## Methods

### NewGetAcccountRequestPositionsInner

`func NewGetAcccountRequestPositionsInner() *GetAcccountRequestPositionsInner`

NewGetAcccountRequestPositionsInner instantiates a new GetAcccountRequestPositionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAcccountRequestPositionsInnerWithDefaults

`func NewGetAcccountRequestPositionsInnerWithDefaults() *GetAcccountRequestPositionsInner`

NewGetAcccountRequestPositionsInnerWithDefaults instantiates a new GetAcccountRequestPositionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetType

`func (o *GetAcccountRequestPositionsInner) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *GetAcccountRequestPositionsInner) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *GetAcccountRequestPositionsInner) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *GetAcccountRequestPositionsInner) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetBrokerId

`func (o *GetAcccountRequestPositionsInner) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *GetAcccountRequestPositionsInner) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *GetAcccountRequestPositionsInner) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *GetAcccountRequestPositionsInner) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetCost

`func (o *GetAcccountRequestPositionsInner) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetAcccountRequestPositionsInner) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetAcccountRequestPositionsInner) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetAcccountRequestPositionsInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCostPrice

`func (o *GetAcccountRequestPositionsInner) GetCostPrice() string`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *GetAcccountRequestPositionsInner) GetCostPriceOk() (*string, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *GetAcccountRequestPositionsInner) SetCostPrice(v string)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *GetAcccountRequestPositionsInner) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *GetAcccountRequestPositionsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetAcccountRequestPositionsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetAcccountRequestPositionsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetAcccountRequestPositionsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchangeRate

`func (o *GetAcccountRequestPositionsInner) GetExchangeRate() string`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *GetAcccountRequestPositionsInner) GetExchangeRateOk() (*string, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *GetAcccountRequestPositionsInner) SetExchangeRate(v string)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *GetAcccountRequestPositionsInner) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetId

`func (o *GetAcccountRequestPositionsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAcccountRequestPositionsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAcccountRequestPositionsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetAcccountRequestPositionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastOpenTime

`func (o *GetAcccountRequestPositionsInner) GetLastOpenTime() string`

GetLastOpenTime returns the LastOpenTime field if non-nil, zero value otherwise.

### GetLastOpenTimeOk

`func (o *GetAcccountRequestPositionsInner) GetLastOpenTimeOk() (*string, bool)`

GetLastOpenTimeOk returns a tuple with the LastOpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOpenTime

`func (o *GetAcccountRequestPositionsInner) SetLastOpenTime(v string)`

SetLastOpenTime sets LastOpenTime field to given value.

### HasLastOpenTime

`func (o *GetAcccountRequestPositionsInner) HasLastOpenTime() bool`

HasLastOpenTime returns a boolean if a field has been set.

### GetLastPrice

`func (o *GetAcccountRequestPositionsInner) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *GetAcccountRequestPositionsInner) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *GetAcccountRequestPositionsInner) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *GetAcccountRequestPositionsInner) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLock

`func (o *GetAcccountRequestPositionsInner) GetLock() bool`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *GetAcccountRequestPositionsInner) GetLockOk() (*bool, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *GetAcccountRequestPositionsInner) SetLock(v bool)`

SetLock sets Lock field to given value.

### HasLock

`func (o *GetAcccountRequestPositionsInner) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMarketValue

`func (o *GetAcccountRequestPositionsInner) GetMarketValue() string`

GetMarketValue returns the MarketValue field if non-nil, zero value otherwise.

### GetMarketValueOk

`func (o *GetAcccountRequestPositionsInner) GetMarketValueOk() (*string, bool)`

GetMarketValueOk returns a tuple with the MarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketValue

`func (o *GetAcccountRequestPositionsInner) SetMarketValue(v string)`

SetMarketValue sets MarketValue field to given value.

### HasMarketValue

`func (o *GetAcccountRequestPositionsInner) HasMarketValue() bool`

HasMarketValue returns a boolean if a field has been set.

### GetPosition

`func (o *GetAcccountRequestPositionsInner) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GetAcccountRequestPositionsInner) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GetAcccountRequestPositionsInner) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *GetAcccountRequestPositionsInner) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPositionProportion

`func (o *GetAcccountRequestPositionsInner) GetPositionProportion() string`

GetPositionProportion returns the PositionProportion field if non-nil, zero value otherwise.

### GetPositionProportionOk

`func (o *GetAcccountRequestPositionsInner) GetPositionProportionOk() (*string, bool)`

GetPositionProportionOk returns a tuple with the PositionProportion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionProportion

`func (o *GetAcccountRequestPositionsInner) SetPositionProportion(v string)`

SetPositionProportion sets PositionProportion field to given value.

### HasPositionProportion

`func (o *GetAcccountRequestPositionsInner) HasPositionProportion() bool`

HasPositionProportion returns a boolean if a field has been set.

### GetTicker

`func (o *GetAcccountRequestPositionsInner) GetTicker() GetAcccountRequestPositionsInnerTicker`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *GetAcccountRequestPositionsInner) GetTickerOk() (*GetAcccountRequestPositionsInnerTicker, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *GetAcccountRequestPositionsInner) SetTicker(v GetAcccountRequestPositionsInnerTicker)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *GetAcccountRequestPositionsInner) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetUnrealizedProfitLoss

`func (o *GetAcccountRequestPositionsInner) GetUnrealizedProfitLoss() string`

GetUnrealizedProfitLoss returns the UnrealizedProfitLoss field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossOk

`func (o *GetAcccountRequestPositionsInner) GetUnrealizedProfitLossOk() (*string, bool)`

GetUnrealizedProfitLossOk returns a tuple with the UnrealizedProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLoss

`func (o *GetAcccountRequestPositionsInner) SetUnrealizedProfitLoss(v string)`

SetUnrealizedProfitLoss sets UnrealizedProfitLoss field to given value.

### HasUnrealizedProfitLoss

`func (o *GetAcccountRequestPositionsInner) HasUnrealizedProfitLoss() bool`

HasUnrealizedProfitLoss returns a boolean if a field has been set.

### GetUnrealizedProfitLossRate

`func (o *GetAcccountRequestPositionsInner) GetUnrealizedProfitLossRate() string`

GetUnrealizedProfitLossRate returns the UnrealizedProfitLossRate field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossRateOk

`func (o *GetAcccountRequestPositionsInner) GetUnrealizedProfitLossRateOk() (*string, bool)`

GetUnrealizedProfitLossRateOk returns a tuple with the UnrealizedProfitLossRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLossRate

`func (o *GetAcccountRequestPositionsInner) SetUnrealizedProfitLossRate(v string)`

SetUnrealizedProfitLossRate sets UnrealizedProfitLossRate field to given value.

### HasUnrealizedProfitLossRate

`func (o *GetAcccountRequestPositionsInner) HasUnrealizedProfitLossRate() bool`

HasUnrealizedProfitLossRate returns a boolean if a field has been set.

### GetUpdatePositionTimeStamp

`func (o *GetAcccountRequestPositionsInner) GetUpdatePositionTimeStamp() float32`

GetUpdatePositionTimeStamp returns the UpdatePositionTimeStamp field if non-nil, zero value otherwise.

### GetUpdatePositionTimeStampOk

`func (o *GetAcccountRequestPositionsInner) GetUpdatePositionTimeStampOk() (*float32, bool)`

GetUpdatePositionTimeStampOk returns a tuple with the UpdatePositionTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePositionTimeStamp

`func (o *GetAcccountRequestPositionsInner) SetUpdatePositionTimeStamp(v float32)`

SetUpdatePositionTimeStamp sets UpdatePositionTimeStamp field to given value.

### HasUpdatePositionTimeStamp

`func (o *GetAcccountRequestPositionsInner) HasUpdatePositionTimeStamp() bool`

HasUpdatePositionTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


