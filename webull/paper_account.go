package webull

import (
	"fmt"
	"net/url"

	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

// GetPaperTradeAccounts returns user paper trade accounts
func (c *Client) GetPaperTradeAccounts() (*[]model.PaperTradeAccount, error) {
	var (
		u, _       = url.Parse(PaperTradeEndpoint + "/myaccounts/true")
		headersMap = make(map[string]string)
		response   []model.PaperTradeAccount
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID

	err := c.GetAndDecode(*u, &response, &headersMap, nil)
	if err != nil {
		return &response, err
	}
	return &response, err
}

// GetPaperTradeAccountID returns user paper trade accounts
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

