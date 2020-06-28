package webull

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

func TestGetTicker(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
	})
	asrt.Empty(err)
	asrt.NotNil(c)
	ticker, err := c.GetTicker("AAPL")
	asrt.Empty(err)
	asrt.NotEmpty(ticker)
}

func TestGetTickerID(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
	})
	asrt.Empty(err)
	asrt.NotNil(c)
	tickerID, err := c.GetTickerID("AAPL")
	asrt.Empty(err)
	asrt.NotEmpty(tickerID)
}

func TestGetRealtimeStockQuote(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
	})
	asrt.Empty(err)
	asrt.NotNil(c)
	res, err := c.GetRealtimeStockQuote("913243251")
	asrt.Empty(err)
	asrt.NotEmpty(res)
}

func TestGetStockFundamentals(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
	})
	asrt.Empty(err)
	asrt.NotNil(c)
	res, err := c.GetStockFundamentals("913256135")
	asrt.Empty(err)
	asrt.NotEmpty(res)
}

func TestGetActiveGainersLosers(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
	})
	asrt.Empty(err)
	asrt.NotNil(c)
	res, err := c.GetActiveGainersLosers("advanced", "6", "6")
	asrt.Empty(err)
	asrt.NotEmpty(res)
}

func TestGetStockAnalysis(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
	})
	asrt.Empty(err)
	asrt.NotNil(c)
	res, err := c.GetStockAnalysis("913256135")
	asrt.Empty(err)
	asrt.NotEmpty(res)
}
