package main

import (
	"fmt"
	"os"

	"gitlab.com/brokerage-api/webull-client/webull"
	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

func main() {
	// Setup your client
	c, _ := webull.NewClient(&webull.Credentials{ // check your email to get MFA code
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
		DeviceName: fmt.Sprintf(os.Getenv("WEBULL_USERNAME") + "@go-client"),
	})
	accID, _ := c.GetAccountID()
	if divs, err := c.GetAccountDividends(accID); divs != nil {
		fmt.Printf("Dividends for account [%s]: %s", accID, divs.DividendTotal)
	} else {
		fmt.Errorf("%s", err.Error())
	}
}
