package webull

import (
	"bytes"
	"encoding/json"

	// "fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

// GetTransfers returns Transfers.
func (c *Client) GetTransfers(accountID string, count uint32) (*model.Transfers, error) {
	var (
		u, _       = url.Parse(TradeEndpoint + "/asset/" + accountID + "/getWebullTransferList")
		response   *model.Transfers
		headersMap = make(map[string]string)
		// queryParams = make(map[string]string)
	)

	headersMap[HeaderKeyAccessToken] = c.AccessToken
	headersMap[HeaderKeyDeviceID] = c.DeviceID

	// Login request body
	request := model.GetTransfersRequest{
		PageSize:     float32(count),
		LastRecordId: "0",
	}
	requestBody, _ := json.Marshal(request)
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(requestBody))
	req.Header.Add(HeaderKeyDeviceID, c.DeviceID)
	req.Header.Add(HeaderKeyAccessToken, c.AccessToken)
	if err != nil {
		return nil, errors.Wrap(err, "could not create request")
	}
	// Send and parse request
	err = c.DoAndDecode(req, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
