# GetAccountResponsePositionsInner

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
**UpdatePositionTimeStamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetAccountResponsePositionsInner

`func NewGetAccountResponsePositionsInner() *GetAccountResponsePositionsInner`

NewGetAccountResponsePositionsInner instantiates a new GetAccountResponsePositionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountResponsePositionsInnerWithDefaults

`func NewGetAccountResponsePositionsInnerWithDefaults() *GetAccountResponsePositionsInner`

NewGetAccountResponsePositionsInnerWithDefaults instantiates a new GetAccountResponsePositionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetType

`func (o *GetAccountResponsePositionsInner) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *GetAccountResponsePositionsInner) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *GetAccountResponsePositionsInner) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *GetAccountResponsePositionsInner) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetBrokerId

`func (o *GetAccountResponsePositionsInner) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *GetAccountResponsePositionsInner) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *GetAccountResponsePositionsInner) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *GetAccountResponsePositionsInner) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetCost

`func (o *GetAccountResponsePositionsInner) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetAccountResponsePositionsInner) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetAccountResponsePositionsInner) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetAccountResponsePositionsInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCostPrice

`func (o *GetAccountResponsePositionsInner) GetCostPrice() string`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *GetAccountResponsePositionsInner) GetCostPriceOk() (*string, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *GetAccountResponsePositionsInner) SetCostPrice(v string)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *GetAccountResponsePositionsInner) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *GetAccountResponsePositionsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetAccountResponsePositionsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetAccountResponsePositionsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetAccountResponsePositionsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchangeRate

`func (o *GetAccountResponsePositionsInner) GetExchangeRate() string`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *GetAccountResponsePositionsInner) GetExchangeRateOk() (*string, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *GetAccountResponsePositionsInner) SetExchangeRate(v string)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *GetAccountResponsePositionsInner) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetId

`func (o *GetAccountResponsePositionsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAccountResponsePositionsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAccountResponsePositionsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetAccountResponsePositionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastOpenTime

`func (o *GetAccountResponsePositionsInner) GetLastOpenTime() string`

GetLastOpenTime returns the LastOpenTime field if non-nil, zero value otherwise.

### GetLastOpenTimeOk

`func (o *GetAccountResponsePositionsInner) GetLastOpenTimeOk() (*string, bool)`

GetLastOpenTimeOk returns a tuple with the LastOpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOpenTime

`func (o *GetAccountResponsePositionsInner) SetLastOpenTime(v string)`

SetLastOpenTime sets LastOpenTime field to given value.

### HasLastOpenTime

`func (o *GetAccountResponsePositionsInner) HasLastOpenTime() bool`

HasLastOpenTime returns a boolean if a field has been set.

### GetLastPrice

`func (o *GetAccountResponsePositionsInner) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *GetAccountResponsePositionsInner) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *GetAccountResponsePositionsInner) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *GetAccountResponsePositionsInner) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLock

`func (o *GetAccountResponsePositionsInner) GetLock() bool`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *GetAccountResponsePositionsInner) GetLockOk() (*bool, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *GetAccountResponsePositionsInner) SetLock(v bool)`

SetLock sets Lock field to given value.

### HasLock

`func (o *GetAccountResponsePositionsInner) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMarketValue

`func (o *GetAccountResponsePositionsInner) GetMarketValue() string`

GetMarketValue returns the MarketValue field if non-nil, zero value otherwise.

### GetMarketValueOk

`func (o *GetAccountResponsePositionsInner) GetMarketValueOk() (*string, bool)`

GetMarketValueOk returns a tuple with the MarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketValue

`func (o *GetAccountResponsePositionsInner) SetMarketValue(v string)`

SetMarketValue sets MarketValue field to given value.

### HasMarketValue

`func (o *GetAccountResponsePositionsInner) HasMarketValue() bool`

HasMarketValue returns a boolean if a field has been set.

### GetPosition

`func (o *GetAccountResponsePositionsInner) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GetAccountResponsePositionsInner) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GetAccountResponsePositionsInner) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *GetAccountResponsePositionsInner) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPositionProportion

`func (o *GetAccountResponsePositionsInner) GetPositionProportion() string`

GetPositionProportion returns the PositionProportion field if non-nil, zero value otherwise.

### GetPositionProportionOk

`func (o *GetAccountResponsePositionsInner) GetPositionProportionOk() (*string, bool)`

GetPositionProportionOk returns a tuple with the PositionProportion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionProportion

`func (o *GetAccountResponsePositionsInner) SetPositionProportion(v string)`

SetPositionProportion sets PositionProportion field to given value.

### HasPositionProportion

`func (o *GetAccountResponsePositionsInner) HasPositionProportion() bool`

HasPositionProportion returns a boolean if a field has been set.

### GetTicker

`func (o *GetAccountResponsePositionsInner) GetTicker() GetAcccountRequestPositionsInnerTicker`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *GetAccountResponsePositionsInner) GetTickerOk() (*GetAcccountRequestPositionsInnerTicker, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *GetAccountResponsePositionsInner) SetTicker(v GetAcccountRequestPositionsInnerTicker)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *GetAccountResponsePositionsInner) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetUnrealizedProfitLoss

`func (o *GetAccountResponsePositionsInner) GetUnrealizedProfitLoss() string`

GetUnrealizedProfitLoss returns the UnrealizedProfitLoss field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossOk

`func (o *GetAccountResponsePositionsInner) GetUnrealizedProfitLossOk() (*string, bool)`

GetUnrealizedProfitLossOk returns a tuple with the UnrealizedProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLoss

`func (o *GetAccountResponsePositionsInner) SetUnrealizedProfitLoss(v string)`

SetUnrealizedProfitLoss sets UnrealizedProfitLoss field to given value.

### HasUnrealizedProfitLoss

`func (o *GetAccountResponsePositionsInner) HasUnrealizedProfitLoss() bool`

HasUnrealizedProfitLoss returns a boolean if a field has been set.

### GetUnrealizedProfitLossRate

`func (o *GetAccountResponsePositionsInner) GetUnrealizedProfitLossRate() string`

GetUnrealizedProfitLossRate returns the UnrealizedProfitLossRate field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossRateOk

`func (o *GetAccountResponsePositionsInner) GetUnrealizedProfitLossRateOk() (*string, bool)`

GetUnrealizedProfitLossRateOk returns a tuple with the UnrealizedProfitLossRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLossRate

`func (o *GetAccountResponsePositionsInner) SetUnrealizedProfitLossRate(v string)`

SetUnrealizedProfitLossRate sets UnrealizedProfitLossRate field to given value.

### HasUnrealizedProfitLossRate

`func (o *GetAccountResponsePositionsInner) HasUnrealizedProfitLossRate() bool`

HasUnrealizedProfitLossRate returns a boolean if a field has been set.

### GetUpdatePositionTimeStamp

`func (o *GetAccountResponsePositionsInner) GetUpdatePositionTimeStamp() int64`

GetUpdatePositionTimeStamp returns the UpdatePositionTimeStamp field if non-nil, zero value otherwise.

### GetUpdatePositionTimeStampOk

`func (o *GetAccountResponsePositionsInner) GetUpdatePositionTimeStampOk() (*int64, bool)`

GetUpdatePositionTimeStampOk returns a tuple with the UpdatePositionTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePositionTimeStamp

`func (o *GetAccountResponsePositionsInner) SetUpdatePositionTimeStamp(v int64)`

SetUpdatePositionTimeStamp sets UpdatePositionTimeStamp field to given value.

### HasUpdatePositionTimeStamp

`func (o *GetAccountResponsePositionsInner) HasUpdatePositionTimeStamp() bool`

HasUpdatePositionTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


