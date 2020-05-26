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
// TransactionTransactionItemInstrument struct for TransactionTransactionItemInstrument
type TransactionTransactionItemInstrument struct {
	Symbol string `json:"symbol,omitempty"`
	UnderlyingSymbol string `json:"underlyingSymbol,omitempty"`
	OptionExpirationDate string `json:"optionExpirationDate,omitempty"`
	OptionStrikePrice float32 `json:"optionStrikePrice,omitempty"`
	PutCall PutOrCall `json:"putCall,omitempty"`
	Cusip string `json:"cusip,omitempty"`
	Description string `json:"description,omitempty"`
	AssetType AssetType `json:"assetType,omitempty"`
	BondMaturityDate string `json:"bondMaturityDate,omitempty"`
	BondInterestRate float32 `json:"bondInterestRate,omitempty"`
}
