package main

import (
	// "fmt"
	"os"

	// "gitlab.com/brokerage-api/webull-client/webull"
)

func main() {
	if os.Getenv("WEBULL_USERNAME") == "" {
		return
	}
	if os.Getenv("WEBULL_PASSWORD") == "" {
		return
	}
	// o := &webull.CredsCacher{
	// 	Creds: &webull.OAuth{
	// 		Username: os.Getenv("WEBULL_USERNAME"),
	// 		Password: os.Getenv("WEBULL_PASSWORD"),
	// 	},
	// }
	// o :=  &webull.Auth{
	// 	Username: os.Getenv("WEBULL_USERNAME"),
	// 	Password: os.Getenv("WEBULL_PASSWORD"),
	// }
	// c, err := webull.Dial(&webull.CredsCacher{Creds: o})
	// if err != nil {
	// 	fmt.Errorf("Got error dialing Webull: %v", err)
	// }
	// res, err := c.GetAccounts()
	// if err != nil {
	// 	fmt.Errorf("Got error dialing Webull: %v", err)
	// }
	// fmt.Printf("%v", res)
	// // c.Token
	// fmt.Printf(c.AccessToken)

	// //err
	// i, err := c.GetInstrumentForSymbol("SPY")
	// if err != nil {
	// 	fmt.Errorf("Unable to get SPY instrument")
	// 	return
	// }
	// fmt.Printf("%v", i)
}
