package webull

import (
	"fmt"
	"net/url"
	"strconv"

	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

// GetAccounts gets all associated accounts
func (c *Client) GetAccounts() (*model.GetSecurityAccountsResponse, error) {
	var (
		u, _       = url.Parse(TradeEndpoint + "/account/getSecAccountList/v4")
		response   model.GetSecurityAccountsResponse
		headersMap = make(map[string]string)
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID
	err := c.GetAndDecode(*u, &response, &headersMap, nil)
	if err != nil {
		return &response, err
	}
	return &response, err
}

// GetAccountID gets an account ID
func (c *Client) GetAccountID() (string, error) {
	res, err := c.GetAccounts()
	if err != nil {
		return "", err
	}
	if res == nil {
		return "", fmt.Errorf("No paper trade account found")
	}
	for _, acc := range res.Data {
		return fmt.Sprintf("%d", acc.SecAccountId), nil
	}
	return "", err
}

// GetAccountID gets an account ID
func (c *Client) GetAccountIDs() (accountIDs []string, err error) {
	if res, err := c.GetAccounts(); err != nil {
		return accountIDs, err
	} else if res == nil {
		return accountIDs, fmt.Errorf("No paper trade account found")
	} else {
		accountIDs = make([]string, len(res.Data))
		for i, acc := range res.Data {
			accountIDs[i] = fmt.Sprintf("%d", acc.SecAccountId)
		}
		return accountIDs, err
	}
}

// GetAccount gets account details for account `accountID`
func (c *Client) GetAccount(accountID int) (*model.GetAccountResponse, error) {
	var (
		path       = TradeEndpoint + "/v2/home/" + strconv.Itoa(int(accountID))
		u, _       = url.Parse(path)
		headersMap = make(map[string]string)
		response   model.GetAccountResponse
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID

	err := c.GetAndDecode(*u, &response, &headersMap, nil)
	if err != nil {
		return &response, err
	}
	return &response, err
}

// GetAccountV5 gets account details for account. Note: Doesn't work.
func (c *Client) GetAccountV5() (*model.GetAccountsResponseV5, error) {
	var (
		path       = TradeEndpoint + "/v5/home"
		u, _       = url.Parse(path)
		headersMap = make(map[string]string)
		response   model.GetAccountsResponseV5
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID
	headersMap[HeaderKeyTradeToken] = c.TradeToken
	headersMap[HeaderKeyTradeTime] = getTimeSeconds()

	err := c.GetAndDecode(*u, &response, &headersMap, nil)
	if err != nil {
		return &response, err
	}
	return &response, err
}
