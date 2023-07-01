package webull

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	model "gitlab.com/brokerage-api/webull-client/models"
)


func TestConnectWebsockets(t *testing.T) {
	var (
		tickerID = "913256135"
		msgTypes = []string{"101", "102", "103", "104", "105", "106", "107", "108"}
	)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*90)
	defer cancel()
	asrt.Empty(err)
	asrt.NotNil(c)
	err = c.ConnectWebsockets(ctx, msgTypes, []string{tickerID})
	asrt.Empty(err)
}

func TestClient_RegisterCallback(t *testing.T) {
	var (
		tickerID = "913256135"
		msgTypes = []string{"101", "102", "103", "104", "105", "106", "107", "108"}
	)
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
	callback := func(ctx context.Context, topic Topic, msg interface{}) error {
		t.Log("TOPIC: ", fmt.Sprintf("%v\n", topic))
		t.Log("MESSAGE:", fmt.Sprintf("%v\n", msg))
		return nil
	}
	err = c.RegisterCallback(false, callback, "101")
	asrt.Empty(err)
	err = c.RegisterCallback(true, callback, msgTypes...)
	asrt.NotEmpty(err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*900)
	defer cancel()
	asrt.NotNil(c)
	err = c.ConnectWebsockets(ctx, msgTypes, []string{tickerID})
	asrt.Empty(err)

}