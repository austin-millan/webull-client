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
// Fundamental struct for Fundamental
type Fundamental struct {
	Cusip string `json:"cusip,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	Description string `json:"description,omitempty"`
	Exchange string `json:"exchange,omitempty"`
	AssetType AssetType `json:"assetType,omitempty"`
	Fundamental FundamentalFundamental `json:"fundamental,omitempty"`
}
