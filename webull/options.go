package webull

import (
	// "fmt"

	"fmt"
	"net/url"
	// "strconv"

	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

// GetStockOptions returns options quotes.
func (c *Client) GetStockOptions(tickerID, expireDate, direction string, count, includeWeekly, queryAll int32) (*model.GetStockOptionsResponse, error) {
	var (
		u, _        = url.Parse(BrokerQuotesEndpoint + "/quote/option/" + tickerID + "/list")
		response    model.GetStockOptionsResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID

	if direction == "" {
		direction = "all"
	}

	queryParams["count"] = fmt.Sprintf("%d", count)
	queryParams["includeWeekly"] = fmt.Sprintf("%d", includeWeekly)
	queryParams["direction"] = direction
	queryParams["expireDate"] = expireDate
	queryParams["queryAll"] = fmt.Sprintf("%d", queryAll)

	err := c.GetAndDecode(*u, &response, &headersMap, &queryParams)
	if err != nil {
		return &response, err
	}

	return &response, err
}

// GetOptionsQuotes returns options quotes.
func (c *Client) GetOptionsQuotes(tickerID, derivativeIds string) (*model.GetOptionsQuotesResponse, error) {
	var (
		u, _        = url.Parse(BrokerQuotesGWEndpoint + "/quote/option/query/list")
		response    model.GetOptionsQuotesResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID

	queryParams["tickerID"] = tickerID
	queryParams["derivativeIds"] = derivativeIds

	err := c.GetAndDecode(*u, &response, &headersMap, &queryParams)
	if err != nil {
		return &response, err
	}

	return &response, err
}