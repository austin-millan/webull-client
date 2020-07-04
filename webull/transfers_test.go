package webull

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

func TestGetTransfers(t *testing.T) {
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
	accountID, err := c.GetAccountID()
	asrt.Empty(err)
	asrt.NotNil(c)
	ticker, err := c.GetTransfers(accountID, 10)
	asrt.Empty(err)
	asrt.NotEmpty(ticker)
}