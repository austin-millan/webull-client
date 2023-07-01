# GetAccountsResponseV5Positions2InnerPositionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PositionId** | Pointer to **float32** |  | [optional] 
**BrokerPosId** | Pointer to **string** |  | [optional] 
**BrokerId** | Pointer to **float32** |  | [optional] 
**AssetType** | Pointer to **string** |  | [optional] 
**TickerType** | Pointer to **string** |  | [optional] 
**Ticker** | Pointer to [**GetAccountsResponseV5Positions2InnerPositionsInnerTicker**](GetAccountsResponseV5Positions2InnerPositionsInnerTicker.md) |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **string** |  | [optional] 
**TickerId** | Pointer to **float32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**CostPrice** | Pointer to **string** |  | [optional] 
**Cost** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**MarketValue** | Pointer to **string** |  | [optional] 
**LastPrice** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLoss** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLossRate** | Pointer to **string** |  | [optional] 
**PositionProportion** | Pointer to **string** |  | [optional] 
**UpdatePositionTime** | Pointer to **float32** |  | [optional] 
**OptionCanExercise** | Pointer to **float32** |  | [optional] 
**RecentlyExpireFlag** | Pointer to **bool** |  | [optional] 
**RemainDay** | Pointer to **float32** |  | [optional] 
**IsLending** | Pointer to **string** |  | [optional] 

## Methods

### NewGetAccountsResponseV5Positions2InnerPositionsInner

`func NewGetAccountsResponseV5Positions2InnerPositionsInner() *GetAccountsResponseV5Positions2InnerPositionsInner`

NewGetAccountsResponseV5Positions2InnerPositionsInner instantiates a new GetAccountsResponseV5Positions2InnerPositionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountsResponseV5Positions2InnerPositionsInnerWithDefaults

`func NewGetAccountsResponseV5Positions2InnerPositionsInnerWithDefaults() *GetAccountsResponseV5Positions2InnerPositionsInner`

NewGetAccountsResponseV5Positions2InnerPositionsInnerWithDefaults instantiates a new GetAccountsResponseV5Positions2InnerPositionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionId

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetPositionId() float32`

GetPositionId returns the PositionId field if non-nil, zero value otherwise.

### GetPositionIdOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetPositionIdOk() (*float32, bool)`

GetPositionIdOk returns a tuple with the PositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionId

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetPositionId(v float32)`

SetPositionId sets PositionId field to given value.

### HasPositionId

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasPositionId() bool`

HasPositionId returns a boolean if a field has been set.

### GetBrokerPosId

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetBrokerPosId() string`

GetBrokerPosId returns the BrokerPosId field if non-nil, zero value otherwise.

### GetBrokerPosIdOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetBrokerPosIdOk() (*string, bool)`

GetBrokerPosIdOk returns a tuple with the BrokerPosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerPosId

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetBrokerPosId(v string)`

SetBrokerPosId sets BrokerPosId field to given value.

### HasBrokerPosId

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasBrokerPosId() bool`

HasBrokerPosId returns a boolean if a field has been set.

### GetBrokerId

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetBrokerId() float32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetBrokerIdOk() (*float32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetBrokerId(v float32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetAssetType

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetTickerType

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetTickerType() string`

GetTickerType returns the TickerType field if non-nil, zero value otherwise.

### GetTickerTypeOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetTickerTypeOk() (*string, bool)`

GetTickerTypeOk returns a tuple with the TickerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerType

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetTickerType(v string)`

SetTickerType sets TickerType field to given value.

### HasTickerType

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasTickerType() bool`

HasTickerType returns a boolean if a field has been set.

### GetTicker

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetTicker() GetAccountsResponseV5Positions2InnerPositionsInnerTicker`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetTickerOk() (*GetAccountsResponseV5Positions2InnerPositionsInnerTicker, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetTicker(v GetAccountsResponseV5Positions2InnerPositionsInnerTicker)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetAction

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPosition

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTickerId

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetTickerId() float32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetTickerIdOk() (*float32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetTickerId(v float32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetSymbol

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCostPrice

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetCostPrice() string`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetCostPriceOk() (*string, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetCostPrice(v string)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetCost

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCurrency

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMarketValue

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetMarketValue() string`

GetMarketValue returns the MarketValue field if non-nil, zero value otherwise.

### GetMarketValueOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetMarketValueOk() (*string, bool)`

GetMarketValueOk returns a tuple with the MarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketValue

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetMarketValue(v string)`

SetMarketValue sets MarketValue field to given value.

### HasMarketValue

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasMarketValue() bool`

HasMarketValue returns a boolean if a field has been set.

### GetLastPrice

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetUnrealizedProfitLoss

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetUnrealizedProfitLoss() string`

GetUnrealizedProfitLoss returns the UnrealizedProfitLoss field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetUnrealizedProfitLossOk() (*string, bool)`

GetUnrealizedProfitLossOk returns a tuple with the UnrealizedProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLoss

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetUnrealizedProfitLoss(v string)`

SetUnrealizedProfitLoss sets UnrealizedProfitLoss field to given value.

### HasUnrealizedProfitLoss

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasUnrealizedProfitLoss() bool`

HasUnrealizedProfitLoss returns a boolean if a field has been set.

### GetUnrealizedProfitLossRate

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetUnrealizedProfitLossRate() string`

GetUnrealizedProfitLossRate returns the UnrealizedProfitLossRate field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossRateOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetUnrealizedProfitLossRateOk() (*string, bool)`

GetUnrealizedProfitLossRateOk returns a tuple with the UnrealizedProfitLossRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLossRate

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetUnrealizedProfitLossRate(v string)`

SetUnrealizedProfitLossRate sets UnrealizedProfitLossRate field to given value.

### HasUnrealizedProfitLossRate

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasUnrealizedProfitLossRate() bool`

HasUnrealizedProfitLossRate returns a boolean if a field has been set.

### GetPositionProportion

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetPositionProportion() string`

GetPositionProportion returns the PositionProportion field if non-nil, zero value otherwise.

### GetPositionProportionOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetPositionProportionOk() (*string, bool)`

GetPositionProportionOk returns a tuple with the PositionProportion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionProportion

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetPositionProportion(v string)`

SetPositionProportion sets PositionProportion field to given value.

### HasPositionProportion

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasPositionProportion() bool`

HasPositionProportion returns a boolean if a field has been set.

### GetUpdatePositionTime

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetUpdatePositionTime() float32`

GetUpdatePositionTime returns the UpdatePositionTime field if non-nil, zero value otherwise.

### GetUpdatePositionTimeOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetUpdatePositionTimeOk() (*float32, bool)`

GetUpdatePositionTimeOk returns a tuple with the UpdatePositionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePositionTime

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetUpdatePositionTime(v float32)`

SetUpdatePositionTime sets UpdatePositionTime field to given value.

### HasUpdatePositionTime

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasUpdatePositionTime() bool`

HasUpdatePositionTime returns a boolean if a field has been set.

### GetOptionCanExercise

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetOptionCanExercise() float32`

GetOptionCanExercise returns the OptionCanExercise field if non-nil, zero value otherwise.

### GetOptionCanExerciseOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetOptionCanExerciseOk() (*float32, bool)`

GetOptionCanExerciseOk returns a tuple with the OptionCanExercise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionCanExercise

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetOptionCanExercise(v float32)`

SetOptionCanExercise sets OptionCanExercise field to given value.

### HasOptionCanExercise

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasOptionCanExercise() bool`

HasOptionCanExercise returns a boolean if a field has been set.

### GetRecentlyExpireFlag

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetRecentlyExpireFlag() bool`

GetRecentlyExpireFlag returns the RecentlyExpireFlag field if non-nil, zero value otherwise.

### GetRecentlyExpireFlagOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetRecentlyExpireFlagOk() (*bool, bool)`

GetRecentlyExpireFlagOk returns a tuple with the RecentlyExpireFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentlyExpireFlag

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetRecentlyExpireFlag(v bool)`

SetRecentlyExpireFlag sets RecentlyExpireFlag field to given value.

### HasRecentlyExpireFlag

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasRecentlyExpireFlag() bool`

HasRecentlyExpireFlag returns a boolean if a field has been set.

### GetRemainDay

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetRemainDay() float32`

GetRemainDay returns the RemainDay field if non-nil, zero value otherwise.

### GetRemainDayOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetRemainDayOk() (*float32, bool)`

GetRemainDayOk returns a tuple with the RemainDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainDay

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetRemainDay(v float32)`

SetRemainDay sets RemainDay field to given value.

### HasRemainDay

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasRemainDay() bool`

HasRemainDay returns a boolean if a field has been set.

### GetIsLending

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetIsLending() string`

GetIsLending returns the IsLending field if non-nil, zero value otherwise.

### GetIsLendingOk

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) GetIsLendingOk() (*string, bool)`

GetIsLendingOk returns a tuple with the IsLending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLending

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) SetIsLending(v string)`

SetIsLending sets IsLending field to given value.

### HasIsLending

`func (o *GetAccountsResponseV5Positions2InnerPositionsInner) HasIsLending() bool`

HasIsLending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


