package webull

import (
	"encoding/json"
	"fmt"
	// "fmt"
	"net/url"

	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

// GetOrders returns orders.
func (c *Client) GetOrders(accountID string, status model.OrderStatus, count int32) (*[]model.GetOrdersItem, error) {
	var (
		u, _        = url.Parse(TradeEndpoint + "/v2/option/list")
		response    []model.GetOrdersItem
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID
	headersMap[HeaderKeyTradeToken] = c.TradeToken
	headersMap[HeaderKeyTradeTime] = getTimeSeconds()

	queryParams["secAccountId"] = accountID
	queryParams["startTime"] = "1970-0-1"
	queryParams["dateType"] = "ORDER"
	queryParams["pageSize"] = fmt.Sprintf("%d", count)
	queryParams["status"] = string(status)

	err := c.GetAndDecode(*u, &response, &headersMap, &queryParams)
	if err != nil {
		return &response, err
	}

	return &response, err
}

// IsTradeable returns information on where a specific ticker is traded
func (c *Client) IsTradeable(tickerID string) (*model.GetIsTradeableResponse, error) {
	var (
		u, _        = url.Parse(TradeEndpoint + "/ticker/broker/permissionV2")
		response    model.GetIsTradeableResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	queryParams[QueryKeyTickerID] = tickerID

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID
	headersMap[HeaderKeyTradeTime] = getTimeSeconds()

	err := c.GetAndDecode(*u, &response, &headersMap, &queryParams)
	if err != nil {
		return &response, err
	}
	return &response, err
}

// PlaceOrder places trade (TODO)
func (c *Client) PlaceOrder(accountID string, input model.PostStockOrderRequest) (*model.PostOrderResponse, error) {
	var (
		u, _       = url.Parse(TradeEndpoint + "/order/" + accountID + "/placeStockOrder")
		headersMap = make(map[string]string)
		response   model.PostOrderResponse
	)

	if input.SerialId == "" {
		input.SerialId = c.UUID
	}

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID
	headersMap[HeaderKeyTradeToken] = c.TradeToken
	headersMap[HeaderKeyTradeTime] = getTimeSeconds()
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	err = c.PostAndDecode(*u, &response, &headersMap, nil, payload)
	if err != nil {
		return &response, err
	}
	if response.OrderId == 0 {
		err = fmt.Errorf("OrderId should not be 0")
	}
	return &response, err
}

// CheckOtocoOrder checks OTOCO order (TODO)
func (c *Client) CheckOtocoOrder(accountID string, input model.PostOtocoOrderRequest) (*interface{}, error) {
	var (
		u, _       = url.Parse(TradeEndpoint + "/v2/corder/stock/place/" + accountID)
		headersMap = make(map[string]string)
		response   interface{}
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID
	headersMap[HeaderKeyTradeToken] = c.TradeToken
	headersMap[HeaderKeyTradeTime] = getTimeSeconds()
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	err = c.PostAndDecode(*u, &response, &headersMap, nil, payload)
	if err != nil {
		return &response, err
	}
	return &response, err
}

// PlaceOtocoOrder places OTOCO trade (TODO)
func (c *Client) PlaceOtocoOrder(accountID string, input model.PostOtocoOrderRequest) (*interface{}, error) {
	var (
		u, _       = url.Parse(TradeEndpoint + "/v2/corder/stock/place/" + accountID)
		headersMap = make(map[string]string)
		response   interface{}
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID
	headersMap[HeaderKeyTradeToken] = c.TradeToken
	headersMap[HeaderKeyTradeTime] = getTimeSeconds()
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	err = c.PostAndDecode(*u, &response, &headersMap, nil, payload)
	if err != nil {
		return &response, err
	}
	return &response, err
}

// CancelOrder cancels trade
func (c *Client) CancelOrder(accountID, orderID string) (*interface{}, error) {
	var (
		u, _       = url.Parse(TradeEndpoint + "/paper/1/acc/" + accountID + "/orderop/cancel/" + orderID)
		headersMap = make(map[string]string)
	)
	var response interface{}

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID
	headersMap[HeaderKeyTradeToken] = c.TradeToken
	headersMap[HeaderKeyTradeTime] = getTimeSeconds()

	err := c.PostAndDecode(*u, &response, &headersMap, nil, nil)
	if err != nil {
		return &response, err
	}
	return &response, err
}

// ModifyOrder modifies trade (TODO)
func (c *Client) ModifyOrder(accountID string, orderID string, input model.PostStockOrderRequest) (*interface{}, error) {
	var (
		u, _       = url.Parse(TradeEndpoint + "/order/" + accountID + "/modifyStockOrder/" + orderID)
		headersMap = make(map[string]string)
	)
	var response interface{}

	if input.SerialId == "" {
		input.SerialId = c.UUID
	}

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID
	headersMap[HeaderKeyTradeToken] = c.TradeToken
	headersMap[HeaderKeyTradeTime] = getTimeSeconds()
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	err = c.PostAndDecode(*u, &response, &headersMap, nil, payload)
	if err != nil {
		return &response, err
	}
	return &response, err
}
