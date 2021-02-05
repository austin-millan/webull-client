package webull

import (
	"encoding/json"
	"fmt"
	"net/url"

	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

// PlacePaperOrder places paper trade
func (c *Client) PlacePaperOrder(accountID string, input model.PostStockOrderRequest) (*model.PostPaperOrderResponse, error) {
	var (
		u, _       = url.Parse(PaperTradeEndpoint + "/paper/1/acc/" + accountID + "/orderop/place/" + fmt.Sprintf("%d", input.TickerId))
		headersMap = make(map[string]string)
		response   model.PostPaperOrderResponse
	)

	if input.SerialId == "" {
		input.SerialId = c.UUID
	}

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID
	headersMap[HeaderKeyTradeToken] = c.TradeToken
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

// CancelPaperOrder cancels paper trade
func (c *Client) CancelPaperOrder(accountID, orderID string) (*interface{}, error) {
	var (
		u, _       = url.Parse(PaperTradeEndpoint + "/paper/1/acc/" + accountID + "/orderop/cancel/" + orderID)
		headersMap = make(map[string]string)
	)
	var response interface{}

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID
	headersMap[HeaderKeyTradeToken] = c.TradeToken

	err := c.PostAndDecode(*u, &response, &headersMap, nil, nil)
	if err != nil {
		return &response, err
	}
	return &response, err
}

// ModifyPaperOrder modifies paper trade
func (c *Client) ModifyPaperOrder(accountID string, orderID string, input model.PostStockOrderRequest) (*interface{}, error) {
	var (
		u, _       = url.Parse(PaperTradeEndpoint + "/paper/1/acc/" + accountID + "/orderop/modify/" + orderID)
		headersMap = make(map[string]string)
	)
	var response interface{}

	if input.SerialId == "" {
		input.SerialId = c.UUID
	}

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID
	headersMap[HeaderKeyTradeToken] = c.TradeToken
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

// GetPaperOrders gets user paper trades
func (c *Client) GetPaperOrders(paperAccountID string, startTime string, dateType string, orderStatus model.OrderStatus) (*[]model.PaperOrder, error) {
	var (
		u, _       = url.Parse(PaperTradeEndpoint + "/paper/1/acc/" + paperAccountID + "/order")
		headersMap = make(map[string]string)
		urlMap     = make(map[string]string)
		response   []model.PaperOrder
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID

	if startTime == "" {
		startTime = "1970-0-1"
	}
	urlMap["startTime"] = startTime
	urlMap["dateType"] = dateType
	urlMap["pageSize"] = "256"
	urlMap["status"] = string(orderStatus)
	err := c.GetAndDecode(*u, &response, &headersMap, &urlMap)
	if err != nil {
		return &response, err
	}
	return &response, err
}
