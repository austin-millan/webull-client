package webull

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	model "gitlab.com/brokerage-api/webull-client/models"
)

func TestGetAccountDividends(t *testing.T) {
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
	res, err := c.GetAccountDividends(acc)
	asrt.Empty(err)
	asrt.NotEmpty(res)
}
