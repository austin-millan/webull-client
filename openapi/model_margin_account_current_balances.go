/*
 * TD Ameritrade API
 *
 * TD Ameritrade API
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MarginAccountCurrentBalances struct for MarginAccountCurrentBalances
type MarginAccountCurrentBalances struct {
	AccruedInterest float32 `json:"accruedInterest,omitempty"`
	CashBalance float32 `json:"cashBalance,omitempty"`
	CashReceipts float32 `json:"cashReceipts,omitempty"`
	LongOptionMarketValue float32 `json:"longOptionMarketValue,omitempty"`
	LiquidationValue float32 `json:"liquidationValue,omitempty"`
	LongMarketValue float32 `json:"longMarketValue,omitempty"`
	MoneyMarketFund float32 `json:"moneyMarketFund,omitempty"`
	Savings float32 `json:"savings,omitempty"`
	ShortMarketValue float32 `json:"shortMarketValue,omitempty"`
	PendingDeposits float32 `json:"pendingDeposits,omitempty"`
	AvailableFunds float32 `json:"availableFunds,omitempty"`
	AvailableFundsNonMarginableTrade float32 `json:"availableFundsNonMarginableTrade,omitempty"`
	BuyingPower float32 `json:"buyingPower,omitempty"`
	BuyingPowerNonMarginableTrade float32 `json:"buyingPowerNonMarginableTrade,omitempty"`
	DayTradingBuyingPower float32 `json:"dayTradingBuyingPower,omitempty"`
	DayTradingBuyingPowerCall float32 `json:"dayTradingBuyingPowerCall,omitempty"`
	Equity float32 `json:"equity,omitempty"`
	EquityPercentage float32 `json:"equityPercentage,omitempty"`
	LongMarginValue float32 `json:"longMarginValue,omitempty"`
	MaintenanceCall float32 `json:"maintenanceCall,omitempty"`
	MaintenanceRequirement float32 `json:"maintenanceRequirement,omitempty"`
	MarginBalance float32 `json:"marginBalance,omitempty"`
	RegTCall float32 `json:"regTCall,omitempty"`
	ShortBalance float32 `json:"shortBalance,omitempty"`
	ShortMarginValue float32 `json:"shortMarginValue,omitempty"`
	ShortOptionMarketValue float32 `json:"shortOptionMarketValue,omitempty"`
	Sma float32 `json:"sma,omitempty"`
	MutualFundValue float32 `json:"mutualFundValue,omitempty"`
	BondValue float32 `json:"bondValue,omitempty"`
	IsInCall bool `json:"isInCall,omitempty"`
	StockBuyingPower float32 `json:"stockBuyingPower,omitempty"`
	OptionBuyingPower float32 `json:"optionBuyingPower,omitempty"`
}
