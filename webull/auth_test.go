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
	c := NewClient()
	err := c.Login(Credentials{
		Username: os.Getenv("WEBULL_USERNAME"),
		Password: os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
	})
	asrt.NoError(err)
	asrt.NotEmpty(c.AccessToken)
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
	c := NewClient()

	// Must get access token first
	err := c.Login(Credentials{
		Username: os.Getenv("WEBULL_USERNAME"),
		Password: os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
	})

	// Finally get trake token
	err = c.TradeLogin(Credentials{
		Username: os.Getenv("WEBULL_USERNAME"),
		TradePIN: os.Getenv("WEBULL_PIN"),
	})
	asrt.NoError(err)
	asrt.NotEmpty(c.TradeToken)
}
