package webull

import (
	"encoding/json"
	"fmt"
	"net/url"

	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

func (c *Client) CancelAllPaperOrders(accountID string) ([]int32, error) {
	if paperOrders, err := c.GetPaperOrders(accountID, "", "", model.WORKING); err != nil {
		return nil, err
	} else if paperOrders == nil {
		return nil, fmt.Errorf("no orders returned")
	} else {
		cancelledOrders := make([]int32, 0)
		for _, order := range *paperOrders {
			cancellation, err := c.CancelPaperOrder(accountID, fmt.Sprintf("%d", Int32Value(order.OrderId)))
			if err != nil {
				fmt.Printf("TODO: fix marshalling error\n")
				cancelledOrders = append(cancelledOrders, Int32Value(order.OrderId))
				//return cancelledOrders, err
			} else {
				cancelledOrders = append(cancelledOrders, Int32Value(order.OrderId))
			}
			fmt.Printf("cancellation: %v", cancellation)
		}
		return cancelledOrders, nil
	}

}

// PlacePaperOrder places paper trade
func (c *Client) PlacePaperOrder(accountID string, input model.PostStockOrderRequest) (*model.PostPaperOrderResponse, error) {
	var (
		u, _       = url.Parse(PaperTradeEndpoint + "/paper/1/acc/" + accountID + "/orderop/place/" + fmt.Sprintf("%d", Int32Value(input.TickerId)))
		headersMap = make(map[string]string)
		response   model.PostPaperOrderResponse
	)

	if StringValue(input.SerialId) == "" {
		input.SerialId = String(c.UUID)
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

	if StringValue(input.SerialId) == "" {
		input.SerialId = String(c.UUID)
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
