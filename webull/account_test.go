package webull

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetAccounts(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	c := NewClient()
	asrt := assert.New(t)
	err := c.Login(Credentials{
		Username: os.Getenv("WEBULL_USERNAME"),
		Password: os.Getenv("WEBULL_PASSWORD"),
	})
	if err != nil {
		t.Errorf("Got error: %s", err.Error())
		t.FailNow()
	}
	res, err := c.GetAccounts()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	asrt.True(res.Success)
}

func TestGetAccount(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c := NewClient()
	err := c.Login(Credentials{
		Username: os.Getenv("WEBULL_USERNAME"),
		Password: os.Getenv("WEBULL_PASSWORD"),
	})
	if err != nil {
		t.Errorf("Got error: %s", err.Error())
		t.FailNow()
	}
	accs, err := c.GetAccounts()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	asrt.True(accs.Success)
	if len(accs.Data) < 1 {
		t.Errorf("No accounts returned")
		t.FailNow()
	}
	_, err = c.GetAccount(int(accs.Data[0].SecAccountId))
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	// asrt.True(acc.Success)
}
