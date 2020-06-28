package webull

import (
	"fmt"
	"net/url"

	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

// GetPaperTradeAccounts gets information for all paper accounts.
func (c *Client) GetPaperTradeAccounts() (*[]model.PaperAccount, error) {
	var (
		u, _       = url.Parse(PaperTradeEndpoint + "/myaccounts/true")
		headersMap = make(map[string]string)
		response   []model.PaperAccount
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID

	err := c.GetAndDecode(*u, &response, &headersMap, nil)
	if err != nil {
		return &response, err
	}
	return &response, err
}

// GetPaperTradeAccountID is a a helper function for getting a paper trading account.
// Currently only retrieves the first account.
func (c *Client) GetPaperTradeAccountID() (string, error) {
	res, err := c.GetPaperTradeAccounts()
	if err != nil {
		return "", err
	}
	if res == nil {
		return "", fmt.Errorf("No paper trade account found")
	}
	for _, acc := range *res {
		return fmt.Sprintf("%d", acc.Id), nil
	}
	return "", err
}

