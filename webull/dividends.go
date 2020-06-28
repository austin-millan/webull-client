package webull

import (
	"net/url"

	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

// GetAccountDividends gets account `accountID` total dividends.
func (c *Client) GetAccountDividends(accountID string) (*model.GetDividendsResponse, error) {
	var (
		u, _        = url.Parse(TradeEndpoint + "/v2/account/" + accountID + "/dividends")
		response    model.GetDividendsResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID

	queryParams["direct"] = "in"

	err := c.GetAndDecode(*u, &response, &headersMap, &queryParams)
	if err != nil {
		return &response, err
	}

	return &response, err
}
