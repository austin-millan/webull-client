package webull

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

func TestPlacePaperTrade(t *testing.T) {
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

	paperAccID, err := c.GetPaperTradeAccountID()
	asrt.Empty(err)
	asrt.NotEmpty(paperAccID)

	tickerID, err := c.GetTickerID("SPY")
	asrt.Empty(err)
	asrt.NotEmpty(tickerID)

	tickerIDNumber, err := strconv.Atoi(tickerID)

	err = c.TradeLogin(Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		TradePIN:    os.Getenv("WEBULL_PIN"),
		AccountType: model.AccountType(2),
	})

	res, err := c.PlacePaperTrade(paperAccID, model.PostStockOrderRequest{
		Action:                    model.BUY,
		ComboType:                 "NORMAL",
		LmtPrice:                  200,
		OrderType:                 model.MKT,
		OutsideRegularTradingHour: false,
		Quantity:                  1,
		SerialId:                  c.UUID,
		TickerId:                  int32(tickerIDNumber),
		TimeInForce:               model.DAY,
	})
	asrt.Empty(err)
	asrt.NotEmpty(res)
}

func TestGetPaperOrders(t *testing.T) {
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
	paperTradeAccID, err := c.GetPaperTradeAccountID()
	asrt.Empty(err)
	asrt.NotEmpty(paperTradeAccID)
	paperTradeOrders, err := c.GetPaperTradeOrders(paperTradeAccID, "", "ORDER", model.WORKING)
	err = c.TradeLogin(Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		TradePIN:    os.Getenv("WEBULL_PIN"),
		AccountType: model.AccountType(2),
	})
	asrt.Empty(err)
	asrt.NotNil(paperTradeOrders)
}

func TestCancelPaperTrade(t *testing.T) {
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

	paperAccID, err := c.GetPaperTradeAccountID()
	asrt.Empty(err)
	asrt.NotEmpty(paperAccID)

	tickerID, err := c.GetTickerID("SPY")
	asrt.Empty(err)
	asrt.NotEmpty(tickerID)

	tickerIDNumber, err := strconv.Atoi(tickerID)

	err = c.TradeLogin(Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		TradePIN:    os.Getenv("WEBULL_PIN"),
		AccountType: model.AccountType(2),
	})

	// Place Trade
	placed, err := c.PlacePaperTrade(paperAccID, model.PostStockOrderRequest{
		Action:                    model.BUY,
		ComboType:                 "NORMAL",
		LmtPrice:                  200,
		OrderType:                 model.MKT,
		OutsideRegularTradingHour: false,
		Quantity:                  1,
		SerialId:                  c.UUID,
		TickerId:                  int32(tickerIDNumber),
		TimeInForce:               model.DAY,
	})
	asrt.Empty(err)
	asrt.NotEmpty(placed)

	// Cancel Trade
	cancelled, err := c.CancelPaperTrade(paperAccID, fmt.Sprintf("%d", placed.OrderId))
	asrt.Empty(err)
	asrt.NotEmpty(cancelled)
}

func TestModifyPaperTrade(t *testing.T) {
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

	paperAccID, err := c.GetPaperTradeAccountID()
	asrt.Empty(err)
	asrt.NotEmpty(paperAccID)

	tickerID, err := c.GetTickerID("SPY")
	asrt.Empty(err)
	asrt.NotEmpty(tickerID)

	tickerIDNumber, err := strconv.Atoi(tickerID)

	err = c.TradeLogin(Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		TradePIN:    os.Getenv("WEBULL_PIN"),
		AccountType: model.AccountType(2),
	})

	// Place Trade
	placed, err := c.PlacePaperTrade(paperAccID, model.PostStockOrderRequest{
		Action:                    model.BUY,
		ComboType:                 "NORMAL",
		LmtPrice:                  200,
		OrderType:                 model.MKT,
		OutsideRegularTradingHour: false,
		Quantity:                  1,
		SerialId:                  c.UUID,
		TickerId:                  int32(tickerIDNumber),
		TimeInForce:               model.DAY,
	})
	asrt.Empty(err)
	asrt.NotEmpty(placed)

	// Cancel Trade
	_, err = c.ModifyPaperTrade(paperAccID, fmt.Sprintf("%d", placed.OrderId), model.PostStockOrderRequest{
		Action:                    model.BUY,
		ComboType:                 "NORMAL",
		LmtPrice:                  200,
		OrderType:                 model.MKT,
		OutsideRegularTradingHour: false,
		Quantity:                  1,
		SerialId:                  c.UUID,
		TickerId:                  int32(tickerIDNumber),
		TimeInForce:               model.DAY,
	})
	asrt.Empty(err)
}
