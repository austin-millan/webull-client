package webull

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

func TestGetAccounts(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	c, err := NewClient(nil)
	asrt := assert.New(t)
	err = c.Login(Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
	})
	asrt.Empty(err)
	res, err := c.GetAccounts()
	asrt.Empty(err)
	asrt.True(res.Success)
}

func TestGetAccount(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(nil)
	err = c.Login(Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
	})
	asrt.Empty(err)
	accs, err := c.GetAccounts()
	asrt.Empty(err)
	asrt.True(accs.Success)
	if len(accs.Data) < 1 {
		t.Errorf("No accounts returned")
		t.FailNow()
	}
	acc, err := c.GetAccount(int(accs.Data[0].SecAccountId))
	asrt.Empty(err)
	asrt.NotNil(acc)
}

func TestGetAccountID(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(nil)
	err = c.Login(Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
	})
	asrt.Empty(err)
	res, err := c.GetAccountID()
	asrt.Empty(err)
	asrt.NotEmpty(res)
}
