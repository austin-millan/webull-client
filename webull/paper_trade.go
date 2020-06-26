package webull

import (
	"encoding/json"
	"fmt"
	"net/url"

	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

// PlacePaperTrade to place paper trade
func (c *Client) PlacePaperTrade(accountID string, input model.PostStockOrderRequest) (*model.PostPaperTradeResponse, error) {
	var (
		u, _       = url.Parse(PaperTradeEndpoint + "/paper/1/acc/" + accountID + "/orderop/place/" + fmt.Sprintf("%d", input.TickerId))
		headersMap = make(map[string]string)
		response   model.PostPaperTradeResponse
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

// CancelPaperTrade to cancel paper trade
func (c *Client) CancelPaperTrade(accountID, orderID string) (*interface{}, error) {
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

// ModifyPaperTrade to modify paper trade
func (c *Client) ModifyPaperTrade(accountID string, orderID string, input model.PostStockOrderRequest) (*interface{}, error) {
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

// GetPaperTradeOrders returns user paper trades
func (c *Client) GetPaperTradeOrders(paperAccountID string, startTime string, dateType string, orderStatus model.OrderStatus) (*[]model.PaperTradeOrder, error) {
	var (
		u, _       = url.Parse(PaperTradeEndpoint + "/paper/1/acc/" + paperAccountID + "/order")
		headersMap = make(map[string]string)
		urlMap     = make(map[string]string)
		response   []model.PaperTradeOrder
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
