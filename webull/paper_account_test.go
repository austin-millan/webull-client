package webull

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

func TestGetPaperTradeAccounts(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
		DeviceName: deviceName(),
	})
	asrt.Empty(err)
	asrt.NotNil(c)
	res, err := c.GetPaperTradeAccounts()
	asrt.Empty(err)
	asrt.NotEmpty(res)
}

func TestGetPaperTradeAccountID(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
		DeviceName: deviceName(),
	})
	asrt.Empty(err)
	asrt.NotNil(c)
	paperAccID, err := c.GetPaperTradeAccountID()
	asrt.Empty(err)
	asrt.NotEmpty(paperAccID)
}
