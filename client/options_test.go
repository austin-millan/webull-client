package webull

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	model "gitlab.com/brokerage-api/webull-client/models"
)

func TestGetStockOptions(t *testing.T) {
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
	acc, err := c.GetStockOptions("913243251", "12/16/2022", "calls", -1, 1, 0)
	asrt.NotEmpty(acc)
	asrt.Empty(err)
}

func TestGetOptionsQuotes(t *testing.T) {
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
	acc, err := c.GetOptionsQuotes("913256135", "913256135")
	asrt.NotEmpty(acc)
	asrt.Empty(err)
}
