/*
 * Webull API
 *
 * Webull API Documentation
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// GetFundamentalsResponseSingle struct for GetFundamentalsResponseSingle
type GetFundamentalsResponseSingle struct {
	CurrencyId int32 `json:"currencyId,omitempty"`
	DilutedEps string `json:"dilutedEps,omitempty"`
	GrossProfit string `json:"grossProfit,omitempty"`
	NetIncomeAfterTax string `json:"netIncomeAfterTax,omitempty"`
	OperatingIncome string `json:"operatingIncome,omitempty"`
	ParentEarning string `json:"parentEarning,omitempty"`
	ReportEndDate string `json:"reportEndDate,omitempty"`
	Revenue string `json:"revenue,omitempty"`
}
