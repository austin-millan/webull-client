package webull

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

func TestGetOrders(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
		DeviceName:  deviceName(),
	})
	asrt.Empty(err)
	asrt.NotNil(c)
	acc, err := c.GetAccountID()
	asrt.NotEmpty(acc)
	asrt.Empty(err)
	orders, err := c.GetOrders(acc, model.FILLED, 200)
	asrt.Empty(err)
	asrt.NotEmpty(orders)
}

func TestIsTradeable(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
		DeviceName:  deviceName(),
	})
	asrt.Empty(err)
	asrt.NotNil(c)
	orders, err := c.IsTradeable("913243251")
	asrt.Empty(err)
	asrt.NotEmpty(orders)
}

func TestPlaceTrade(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
		DeviceName:  deviceName(),
	})
	asrt.Empty(err)
	asrt.NotNil(c)

	accID, err := c.GetAccountID()
	asrt.Empty(err)
	asrt.NotEmpty(accID)

	err = c.TradeLogin(Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		TradePIN:    os.Getenv("WEBULL_PIN"),
		AccountType: model.AccountType(2),
		DeviceName:  deviceName(),
	})

	res, err := c.PlaceOrder(accID, model.PostStockOrderRequest{
		Action:                    model.BUY.Ptr(),
		ComboType:                 String("NORMAL"),
		LmtPrice:                  Float32(0.05),
		OrderType:                 model.LMT.Ptr(),
		OutsideRegularTradingHour: Bool(true),
		Quantity:                  Int32(1),
		SerialId:                  String(c.UUID),
		TickerId:                  Int32(913243251),
		TimeInForce:               model.DAY.Ptr(),
	})
	asrt.Empty(err)
	asrt.NotEmpty(res)
}

// func TestPlaceOtocoTrade(t *testing.T) {
// 	if os.Getenv("WEBULL_USERNAME") == "" {
// 		t.Skip("No username set")
// 		return
// 	}
// 	asrt := assert.New(t)
// 	c, err := NewClient(&Credentials{
// 		Username:    os.Getenv("WEBULL_USERNAME"),
// 		Password:    os.Getenv("WEBULL_PASSWORD"),
// 		AccountType: model.AccountType(2),
// 	})
// 	asrt.Empty(err)
// 	asrt.NotNil(c)

// 	accID, err := c.GetAccountID()
// 	asrt.Empty(err)
// 	asrt.NotEmpty(accID)

// 	err = c.TradeLogin(Credentials{
// 		Username:    os.Getenv("WEBULL_USERNAME"),
// 		Password:    os.Getenv("WEBULL_PASSWORD"),
// 		TradePIN:    os.Getenv("WEBULL_PIN"),
// 		AccountType: model.AccountType(2),
// 	})

// 	res, err := c.PlaceOtocoOrder(accID, model.PostStockOrderRequest{
// 		Action:                    model.BUY,
// 		ComboType:                 "NORMAL",
// 		LmtPrice:                  4.69,
// 		OrderType:                 model.MKT,
// 		OutsideRegularTradingHour: true,
// 		Quantity:                  1,
// 		SerialId:                  c.UUID,
// 		TickerId:                  913243251,
// 		TimeInForce:               model.DAY,
// 	})
// 	asrt.Empty(err)
// 	asrt.NotEmpty(res)
// }

func TestCancelTrade(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
		DeviceName:  deviceName(),
	})
	asrt.Empty(err)
	asrt.NotNil(c)

	accID, err := c.GetAccountID()
	asrt.Empty(err)
	asrt.NotEmpty(accID)

	err = c.TradeLogin(Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		TradePIN:    os.Getenv("WEBULL_PIN"),
		AccountType: model.AccountType(2),
		DeviceName:  deviceName(),
	})

	// Place Trade
	placed, err := c.PlaceOrder(accID, model.PostStockOrderRequest{
		Action:                    model.BUY.Ptr(),
		ComboType:                 String("NORMAL"),
		LmtPrice:                  Float32(4.69),
		OrderType:                 model.MKT.Ptr(),
		OutsideRegularTradingHour: Bool(true),
		Quantity:                  Int32(1),
		SerialId:                  String(c.UUID),
		TickerId:                  Int32(913243251),
		TimeInForce:               model.DAY.Ptr(),
	})
	asrt.Empty(err)
	asrt.NotEmpty(placed)

	// Cancel Trade
	cancelled, err := c.CancelPaperOrder(accID, fmt.Sprintf("%d", Int32Value(placed.OrderId)))
	asrt.Empty(err)
	asrt.NotEmpty(cancelled)
}

func TestModifyTrade(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
		DeviceName:  deviceName(),
	})
	asrt.Empty(err)
	asrt.NotNil(c)

	accID, err := c.GetAccountID()
	asrt.Empty(err)
	asrt.NotEmpty(accID)

	err = c.TradeLogin(Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		TradePIN:    os.Getenv("WEBULL_PIN"),
		AccountType: model.AccountType(2),
		DeviceName:  deviceName(),
	})

	// Place Trade
	placed, err := c.PlacePaperOrder(accID, model.PostStockOrderRequest{
		Action:                    model.BUY.Ptr(),
		ComboType:                 String("NORMAL"),
		LmtPrice:                  Float32(4.69),
		OrderType:                 model.MKT.Ptr(),
		OutsideRegularTradingHour: Bool(true),
		Quantity:                  Int32(1),
		SerialId:                  String(c.UUID),
		TickerId:                  Int32(913243251),
		TimeInForce:               model.DAY.Ptr(),
	})
	asrt.Empty(err)
	asrt.NotEmpty(placed)

	// Cancel Trade
	_, err = c.ModifyPaperOrder(accID, fmt.Sprintf("%d", Int32Value(placed.OrderId)), model.PostStockOrderRequest{
		Action:                    model.BUY.Ptr(),
		ComboType:                 String("NORMAL"),
		LmtPrice:                  Float32(4.69),
		OrderType:                 model.MKT.Ptr(),
		OutsideRegularTradingHour: Bool(true),
		Quantity:                  Int32(1),
		SerialId:                  String(c.UUID),
		TickerId:                  Int32(913243251),
		TimeInForce:               model.DAY.Ptr(),
	})
	asrt.Empty(err)
}

func TestCheckOtocoOrder(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
		DeviceName:  deviceName(),
	})
	asrt.Empty(err)
	asrt.NotNil(c)

	accID, err := c.GetAccountID()
	asrt.Empty(err)
	asrt.NotEmpty(accID)

	err = c.TradeLogin(Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		TradePIN:    os.Getenv("WEBULL_PIN"),
		AccountType: model.AccountType(2),
		DeviceName:  deviceName(),
	})
	model.NewPostStockOrderRequest()
	input := model.NewPostOtocoOrderRequest()
	input.SetNewOrders(
		[]model.PostStockOrderRequest{
			{
				Action:                    model.BUY.Ptr(),
				ComboType:                 String("MASTER"),
				LmtPrice:                  Float32(4.69),
				OrderType:                 model.LMT.Ptr(),
				OutsideRegularTradingHour: Bool(true),
				Quantity:                  Int32(1),
				SerialId:                  String(c.UUID),
				TickerId:                  Int32(913243251),
				TimeInForce:               model.DAY.Ptr(),
			},
			{
				OrderType:                 model.STP.Ptr(),
				TimeInForce:               model.DAY.Ptr(),
				Quantity:                  Int32(1),
				OutsideRegularTradingHour: Bool(false),
				Action:                    model.SELL.Ptr(),
				TickerId:                  Int32(913243251),
				LmtPrice:                  Float32(30),
				ComboType:                 String("STOP_LOSS"),
			},
			{
				OrderType:                 model.LMT.Ptr(),
				TimeInForce:               model.DAY.Ptr(),
				Quantity:                  Int32(1),
				OutsideRegularTradingHour: Bool(false),
				Action:                    model.SELL.Ptr(),
				TickerId:                  Int32(913243251),
				LmtPrice:                  Float32(50),
				ComboType:                 String("STOP_PROFIT"),
			},
		})

	// Place Trade
	placed, err := c.CheckOtocoOrder(accID, *input)
	asrt.Empty(err)
	asrt.NotEmpty(placed)

	// // Cancel Trade
	// _, err = c.ModifyPaperOrder(accID, fmt.Sprintf("%d", placed.OrderId), model.PostStockOrderRequest{
	// 	Action:                    model.BUY,
	// 	ComboType:                 "NORMAL",
	// 	LmtPrice:                  200,
	// 	OrderType:                 model.MKT,
	// 	OutsideRegularTradingHour: false,
	// 	Quantity:                  1,
	// 	SerialId:                  c.UUID,
	// 	TickerId:                  913243251,
	// 	TimeInForce:               model.DAY,
	// })
	// asrt.Empty(err)
}
