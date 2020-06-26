package webull

import (
	// "fmt"
	"net/url"
	// "strconv"

	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

// GetAlerts returns alerts.
func (c *Client) GetAlerts() (*model.GetAlertsResponse, error) {
	var (
		u, _       = url.Parse(UserBrokerEndpoint + "/user/warning/v2/query/tickers")
		response   model.GetAlertsResponse
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
