package webull

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

// Endpoints for the Webull API
const (
	QuotesEndpoint         = "https://quoteapi.webull.com/api"
	UserEndpoint           = "https://userapi.webull.com/api"
	BrokerQuotesEndpoint   = "https://quoteapi.webullbroker.com/api"
	BrokerQuotesGWEndpoint = "https://quoteapi.webullbroker.com/api"
	SecuritiesEndpoint     = "https://securitiesapi.webullbroker.com/api"
	UserBrokerEndpoint     = "https://userapi.webullbroker.com/api"
	PaperTradeEndpoint     = "https://act.webullbroker.com/webull-paper-center/api"
	TradeEndpoint          = "https://tradeapi.webulltrade.com/api/trade"
)

// Client is a helpful abstraction around some common metadata required for
// API operations.
type Client struct {
	AccessToken           string
	AccessTokenExpiration time.Time
	RefreshToken          string

	TradeToken           string
	TradeTokenExpiration time.Time

	DeviceID string

	httpClient *http.Client
}

// NewClient will return a new Webull client
func NewClient() *Client {
	c := &Client{
		httpClient: &http.Client{Timeout: time.Second * 10},
	}
	return c
}

// GetAndDecode retrieves from the endpoint and unmarshals resulting json into
// the provided destination interface, which must be a pointer.
func (c *Client) GetAndDecode(url string, dest interface{}, headers map[string]string) error {
	req, err := http.NewRequest("GET", url, nil)
	for key, val := range headers {
		req.Header.Add(key, val)
	}
	if err != nil {
		return err
	}
	return c.DoAndDecode(req, dest)
}

// DoAndDecode provides useful abstractions around common errors and decoding
// issues.
func (c *Client) DoAndDecode(req *http.Request, dest interface{}) (err error) {
	req.Header.Add("Content-Type", "application/json")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode/100 != 2 {
		b := &bytes.Buffer{}
		var e model.ErrorBody
		err = json.NewDecoder(io.TeeReader(res.Body, b)).Decode(&e)
		if err != nil {
			return fmt.Errorf("got response %q and could not decode error body %q", res.Status, b.String())
		}
		return fmt.Errorf(e.Msg)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Got read error on body: %s", err.Error())
	}
	// anything
	var anyBody interface{}
	if err = json.Unmarshal(body, &anyBody); err != nil {
		return fmt.Errorf("Unable to marshal body as interface")
	}

	if err = json.Unmarshal(body, &dest); err != nil {
		return fmt.Errorf("Got JSON unmarshal error on body: %s", err.Error())
	}
	return err
}