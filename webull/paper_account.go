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

// GetPaperTradeAccountID is a a helper function for getting a single paper trading account ID
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

// GetPaperTradeAccountIDs is a a helper function for getting all paper trading account IDs.
func (c *Client) GetPaperTradeAccountIDs() ([]string, error) {
	if res, err := c.GetPaperTradeAccounts(); err != nil {
		return []string{}, err
	} else if res == nil {
		return []string{}, fmt.Errorf("No paper trade account found")
	} else {
		accountIDs := make([]string, len(*res))
		for i, acc := range *res {
			accountIDs[i] = fmt.Sprintf("%d", acc.Id)
		}
		return accountIDs, err
	}
}

// ResetPaperAccount gets information for all paper accounts.
/*
func (c *Client) ResetPaperAccount(newBalance int32) (*model.ResetPaperAccountResponse, error) {
	var (
		headersMap = make(map[string]string)
		response   model.ResetPaperAccountResponse
	)
	accID, err := c.GetPaperTradeAccountID()
	if err != nil {
		return nil, err
	}
	var u, _       = url.Parse(PaperTradeEndpointV + "/paper/1/acc/reset/" + accID + "/" + fmt.Sprintf("%d", newBalance))
	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID

	err = c.GetAndDecode(*u, &response, &headersMap, nil)
	if err != nil {
		return &response, err
	}
	return &response, err
}
*/