package webull

import (
	"net/url"

	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

// GetUser returns user details of your account (e.g. name, email/phone, region, ...).
func (c *Client) GetUser() (*model.GetUserDetailsResponse, error) {
	var (
		u, _        = url.Parse(UserEndpoint + "/user")
		response    model.GetUserDetailsResponse
		headersMap  = make(map[string]string)
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID


	err := c.GetAndDecode(*u, &response, &headersMap, nil)
	if err != nil {
		return &response, err
	}

	return &response, err
}
