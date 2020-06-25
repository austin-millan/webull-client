package webull

import (
	"fmt"
	"net/url"

	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)


// GetAccounts returns all the accounts associated with a login/client.
func (c *Client) GetAccounts() (*model.GetSecurityAccountsResponse, error) {
	var (
		path       = TradeEndpoint + "/account/getSecAccountList/v4"
		response   model.GetSecurityAccountsResponse
		headersMap = make(map[string]string)
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID
	err := c.GetAndDecode(path, &response, headersMap)
	if err != nil {
		return &response, err
	}

	return &response, err
}

// GetAccount returns an account
func (c *Client) GetAccount(accountID int) (*model.GetAccountResponse, error) {
	var (
		path       = fmt.Sprintf(TradeEndpoint+"/v2/home/%d", accountID)
		headersMap = make(map[string]string)
		response   model.GetAccountResponse
		u, _       = url.Parse(path)
		q          = u.Query()
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID

	u.RawQuery = q.Encode()

	err := c.GetAndDecode(u.String(), &response, headersMap)
	if err != nil {
		return &response, err
	}
	return &response, err
}
