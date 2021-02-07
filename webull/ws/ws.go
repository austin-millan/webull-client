package ws
// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const (
	// Topics
	streamingQuotesAddr = "wss://wspush.webullbroker.com:443/mqtt"
	ordersAddr = "wss://platpush.webullfintech.com:443/mqtt"
)

var choke = make(chan [2]string)

func msgCallback(client MQTT.Client, msg MQTT.Message) {
	log.Printf("TOPIC: %s\n", msg.Topic())
	log.Printf("	MSG: %s\n", msg.Payload())
	choke <- [2]string{msg.Topic(), string(msg.Payload())}
}

// ConnectStreamingQuotes is a utility function for connecting to WS streaming API
func ConnectStreamingQuotes(username, password, deviceID, accessToken string, tickerIDs []string) error {
	var (
		addr = streamingQuotesAddr
		qos = 0
		num = 5
		opts = MQTT.NewClientOptions()
		receiveCount = 0
	)
	MQTT.DEBUG = log.New(os.Stdout, "", 0)
	MQTT.ERROR = log.New(os.Stdout, "", 0)
	subscriptionReq := fmt.Sprintf(`{"tickerIds": [%v],"type": "105"}`, strings.Join(tickerIDs, `,`))
	opts.SetClientID(deviceID)
	opts.AddBroker(addr)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(6 * time.Second)
	//opts.SetHTTPHeaders(req.Header)
	opts.SetTLSConfig(&tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert})
	opts.SetCleanSession(true)
	opts.SetDefaultPublishHandler(msgCallback)
	log.Printf("Connecting to: %s", addr)
	client := MQTT.NewClient(opts)
	client.AddRoute("#", msgCallback) // wildcard
	defer client.Disconnect(250)

	if token := client.Connect(); token.WaitTimeout(time.Second*60) && token.Error() != nil {
		log.Printf("Got error: %s", token.Error().Error())
		return token.Error()
	}
	helloMsg := fmt.Sprintf(
		`{"header":{"access_token":"%s","did":"%s","hl":"en","os":"web","osv":"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0","ver":"3.22.20","app":"global","platform":"web","device-type":"Web"}}`,
		accessToken,
		deviceID,
	)
	if token := client.Subscribe(helloMsg, byte(qos), msgCallback); token.WaitTimeout(time.Second*60) && token.Error() != nil {
		return token.Error()
	}

	//subData := `{"tickerIds": [913256135],"type": "105"}` // TODO: paramaterize func to create subscription request
	if token := client.Subscribe(subscriptionReq, byte(qos), msgCallback); token.WaitTimeout(time.Second*60) && token.Error() != nil {
		return token.Error()
	}
	for receiveCount < num {
		incoming := <-choke
		log.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
		receiveCount++
	}

	fmt.Println("Subscriber disconnecting now")
	return nil
}
