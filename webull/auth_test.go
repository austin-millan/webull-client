package webull

import (
	"os"
	"testing"

	model "gitlab.com/brokerage-api/webull-openapi/openapi"
	"github.com/stretchr/testify/assert"
)

var testCacher = false

func TestLogin(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(nil)
	asrt.Empty(err)
	err = c.Login(Credentials{
		Username: os.Getenv("WEBULL_USERNAME"),
		Password: os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
		DeviceName: deviceName(),
	})
	asrt.Empty(err)
	asrt.NotEmpty(c.AccessToken)
	res, err := c.Token()
	asrt.Empty(err)
	asrt.NotNil(res)
}

func TestTradeToken(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	if os.Getenv("WEBULL_PIN") == "" {
		t.Skip("Trade PIN not set. PIN required to retrieve trade token.")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(nil)
	asrt.Empty(err)
	// Must get access token first
	err = c.Login(Credentials{
		Username: os.Getenv("WEBULL_USERNAME"),
		Password: os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
		DeviceName: deviceName(),
	})
	asrt.NoError(err)

	// Finally get trake token
	err = c.TradeLogin(Credentials{
		Username: os.Getenv("WEBULL_USERNAME"),
		TradePIN: os.Getenv("WEBULL_PIN"),
		DeviceName: deviceName(),
	})
	asrt.Empty(err)
	asrt.NotEmpty(c.TradeToken)
}
