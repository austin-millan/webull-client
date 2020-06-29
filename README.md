# Webull Client

Master Pipeline Status: [![pipeline status](https://gitlab.com/brokerage-api/webull-client/badges/master/pipeline.svg)](https://gitlab.com/brokerage-api/webull-client/commits/master)

Dev Pipeline Status: [![pipeline status](https://gitlab.com/brokerage-api/webull-client/badges/dev/pipeline.svg)](https://gitlab.com/brokerage-api/webull-client/commits/dev)

## About

A Golang library for the [Webull](https://www.webull.com/) brokerage.

## Example Usage

### Get Authorization

First set some environment variables:

```bash
export WEBULL_USERNAME='your.email@domain.com'
export WEBULL_PASSWORD='SuperStrongPasswerrrdd12345'
export WEBULL_PIN='123456'
```

Webull requires a MFA email when authenticating new devices.

```go

func main() {
	// Authenticate your device
	c, _ := webull.NewClient(nil)
	c.GetMFA(webull.Credentials{ // check your email to get MFA code
			Username:    os.Getenv("WEBULL_USERNAME"),
			AccountType: model.AccountType(2),
			DeviceName: fmt.Sprintf(os.Getenv("WEBULL_USERNAME") + "@go-client"),
    })
    // ... check your email for Webull `MFA` code
	c, _ = webull.NewClient(&webull.Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
        AccountType: model.AccountType(2),
		MFA: 123456,
    })
}

```

### Get Dividends

Once your device has been authorized, you should be able to see device name in your Webull app under "Settings" -> "Two-Factor Authentication".

Now your subsequent API calls can simply use your device ID and token. Your token is retrieved from using `Client.Login(...)`
or using the `NewClient` constructor:

```go

func main() {

	// Setup your client
	c, _ := webull.NewClient(&webull.Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
		DeviceName: fmt.Sprintf(os.Getenv("WEBULL_USERNAME") + "@go-client"),
    })
    // Your account ID is used for many operations
	accID, _ := c.GetAccountID()
	if divs, err := c.GetAccountDividends(accID); divs != nil {
		fmt.Printf("Dividends for account [%s]: %s", accID, divs.DividendTotal)
	} else {
		fmt.Errorf("%s", err.Error())
	}
}

```

## Disclaimer

Use at your own risk.

## Issues / Contributions

If you run into any problems feel free to create an issue on Gitlab or submit a pull request.
