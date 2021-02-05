package webull

import (
	"fmt"
	"net/url"
	"strconv"

	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

// GetTicker gets ticker information for a provided stock symbol
func (c *Client) GetTicker(symbol string) (*model.LookupTickerResponse, error) {
	var (
		u, _        = url.Parse(StockInfoEndpoint + "/search/tickers5")
		response    model.LookupTickerResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID

	queryParams["keys"] = symbol
	queryParams["queryNumber"] = strconv.Itoa(int(1))

	err := c.GetAndDecode(*u, &response, &headersMap, &queryParams)
	if err != nil {
		return &response, err
	}

	return &response, err
}

// GetTickerID is a helper function for getting a ticker ID from a stock symbol
func (c *Client) GetTickerID(symbol string) (string, error) {
	res, err := c.GetTicker(symbol)
	if err != nil {
		return "", err
	}
	if len(res.List) < 1 {
		return "", fmt.Errorf("No ticker found")
	}
	for _, symbolInfo := range res.List {
		return fmt.Sprintf("%d", symbolInfo.TickerId), nil
	}
	return "", nil
}

// GetRealtimeStockQuote gets real-time data for ticker `tickerID`
func (c *Client) GetRealtimeStockQuote(tickerID string) (*model.GetStockQuoteResponse, error) {
	var (
		u, _       = url.Parse(QuotesEndpoint + "/quote/tickerRealTimes/v5/" + tickerID)
		response   model.GetStockQuoteResponse
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

// GetStockFundamentals gets stock fundamentals for ticker `tickerID`
func (c *Client) GetStockFundamentals(tickerID string) (*model.GetFundamentalsResponse, error) {
	var (
		u, _       = url.Parse(QuotesEndpoint + "/securities/financial/index/" + tickerID)
		response   model.GetFundamentalsResponse
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

// GetActiveGainersLosers gets the day's active gainers or losers.
func (c *Client) GetActiveGainersLosers(direction, regionID, userRegionID string) (*[]model.ActiveGainersLosers, error) {
	var (
		u, _        = url.Parse(SecuritiesEndpoint + "/securities/market/v5/card/stockActivityPc." + direction + "/list")
		response    []model.ActiveGainersLosers
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	queryParams["regionId"] = regionID
	queryParams["userRegionId"] = userRegionID

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID

	err := c.GetAndDecode(*u, &response, &headersMap, &queryParams)
	if err != nil {
		return &response, err
	}
	return &response, err
}

// GetStockAnalysis gets Webull stock analysis for tickerID `tickerID`
func (c *Client) GetStockAnalysis(tickerID string) (*model.GetStockAnalysisResponse, error) {
	var (
		u, _        = url.Parse(StockInfoEndpoint + "/securities/ticker/v5/analysis/" + tickerID)
		response    model.GetStockAnalysisResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID

	err := c.GetAndDecode(*u, &response, &headersMap, &queryParams)
	if err != nil {
		return &response, err
	}

	return &response, err
}
