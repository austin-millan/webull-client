package webull

import (
	"os"
	"testing"
	model "gitlab.com/brokerage-api/webull-openapi/openapi"
	"github.com/stretchr/testify/assert"
)

func TestOauth(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username: os.Getenv("WEBULL_USERNAME"),
		Password: os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
	})
	asrt.Empty(err)
	asrt.NotNil(c)
	accs, err := c.GetAccounts()
	asrt.Empty(err)
	asrt.NotNil(accs)
}
