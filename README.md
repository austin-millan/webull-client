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

Webull requires an MFA email when authenticating new devices.

Once your device has been authorized, you should be able to see device name in your Webull app under "Settings" -> "Two-Factor Authentication".

Now your subsequent API calls can simply use your device ID and token. Your token is retrieved from using `Client.Login(...)`
or using the `NewClient` constructor:

```go
package main

import (
	"fmt"
	webull "gitlab.com/brokerage-api/webull-client/client"
	model "gitlab.com/brokerage-api/webull-client/models"
	"os"
)

func main() {
	// Authenticate your device
	creds := webull.Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
		DeviceName:  fmt.Sprintf(os.Getenv("WEBULL_USERNAME") + "@go-client"),
	}
	if c, err := webull.NewClient(&creds); err != nil {
		panic(err)
	} else if c == nil {
		panic("no client returned")
	} else {
		if err = c.GetMFA(creds); err != nil {
			panic(err)
		}
		if acc, err := c.GetAccountID(); err != nil {
			panic(err)
		} else {
			res, _ := c.GetAccountDividends(acc)
			fmt.Printf("Dividends: %v", res)
		}
	}
}
```

### MQTT Connection (Quotes Streaming)

You can connect and stream various data using [MQTT](https://en.wikipedia.org/wiki/MQTT).

This example will process (I think all) streamable data for a particular ticker for 90 seconds.
Note: not tested for concurrency.

```go
package main

import (
	"context"
	"fmt"
	webull "gitlab.com/brokerage-api/webull-client/webull"
	model "gitlab.com/brokerage-api/webull-client/models"
	"os"
	"time"
)

func main() {
	// Authenticate your device
	creds := webull.Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
		DeviceName:  fmt.Sprintf(os.Getenv("WEBULL_USERNAME") + "@go-client"),
	}
	if c, err := webull.NewClient(&creds); err != nil {
		panic(err)
	} else if c == nil {
		panic("no client returned")
	} else {
		var (
			tickerID = "913256135"
			msgTypes = []string{"101", "102", "103", "104", "105", "106", "107", "108"}
		)
		yourCallback := func(ctx context.Context, topic webull.Topic, msg interface{}) error {
			fmt.Printf("TOPIC: %v\nMESSAGE: %v", topic, msg)
			return nil
		}
		c.RegisterCallback(false, yourCallback, "101")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*90)
		defer cancel()
		if err = c.ConnectWebsockets(ctx, msgTypes, []string{tickerID}); err != nil {
			panic(err)
		}
	}
}
```

## Disclaimer

Use at your own risk.

## Issues / Contributions

If you run into any problems feel free to create an issue on Gitlab or submit a pull request.

## Repository Views

[![HitCount](http://hits.dwyl.com/austin-millan/webull-client.svg)](http://hits.dwyl.com/austin-millan/webull-client)
