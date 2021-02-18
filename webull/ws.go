package webull
// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"context"
	"crypto/tls"
	"encoding/json"
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

type wsMessage struct {
	Message interface{}
	Topic Topic
}

var (
	bufferedMessages = make(chan wsMessage, 1000)
)

// parse incoming MQTT message and put in channel for processing
func pushChannel(client MQTT.Client, msg MQTT.Message) {
	var (
		outMsg = wsMessage{}
		topic = msg.Topic()
		payload = msg.Payload()
	)
	if err := json.Unmarshal([]byte(topic), &outMsg.Topic); err != nil {
		fmt.Printf("Got error during unmarshal: %s", err.Error())
	} else {
		fmt.Printf("No error")
	}
	switch outMsg.Topic.Type {
	case 101:
		outMsg.Message = Type101Message{}
	case 102:
		outMsg.Message = Type102Message{}
	case 103:
		outMsg.Message = Type103Message{}
	case 104:
		outMsg.Message = Type104Message{}
	case 105:
		outMsg.Message = Type105Message{}
	case 106:
		outMsg.Message = Type106Message{}
	case 107:
		outMsg.Message = Type107Message{}
	case 108:
		outMsg.Message = Type108Message{}
	default:
		// unhandled topic
		return
	}
	if err := json.Unmarshal(payload, &outMsg.Message); err != nil {
		fmt.Printf("Got error during unmarshal: %s", err.Error())
	} else {
		fmt.Printf("No error")
	}
	bufferedMessages <- outMsg
}

// Topic should be used in tandem with *Message structs in user defined callbacks
type Topic struct {
	Type     int      `json:"type"`
	TickerID int      `json:"tickerId"`
	Modules  []string `json:"modules"`
	Flag     string   `json:"flag"`
}

// Type101Message is a message type returned from Webull MQTT quote streams
type Type101Message struct {
	TransID     int    `json:"transId"`
	Change      string `json:"change"`
	MarketValue string `json:"marketValue"`
	ChangeRatio string `json:"changeRatio"`
	PubID       int    `json:"pubId"`
	Topic       int    `json:"topic"`
	TickerID    int    `json:"tickerId"`
	TradeStamp  int64  `json:"tradeStamp"`
	Close       string `json:"close"`
	TrdSeq      int    `json:"trdSeq"`
	Status      string `json:"status"`
}

// Type102Message is a message type returned from Webull MQTT quote streams
type Type102Message struct {
	TransID        int    `json:"transId"`
	High           string `json:"high"`
	Low            string `json:"low"`
	ChangeRatio    string `json:"changeRatio"`
	PubID          int    `json:"pubId"`
	TradeStamp     int64  `json:"tradeStamp"`
	Close          string `json:"close"`
	Change         string `json:"change"`
	MarketValue    string `json:"marketValue"`
	VibrateRatio   string `json:"vibrateRatio"`
	Volume         string `json:"volume"`
	TradeTime      string `json:"tradeTime"`
	NegMarketValue string `json:"negMarketValue"`
	TurnoverRate   string `json:"turnoverRate"`
	Topic          int    `json:"topic"`
	TickerID       int    `json:"tickerId"`
	TrdSeq         int    `json:"trdSeq"`
	Open           string `json:"open"`
	Status         string `json:"status"`
}

// Type103Message is a message type returned from Webull MQTT quote streams
type Type103Message struct {
	Deal struct {
		TrdBs     string `json:"trdBs"`
		Volume    string `json:"volume"`
		TradeTime string `json:"tradeTime"`
		Price     string `json:"price"`
	} `json:"deal"`
	TransID    int    `json:"transId"`
	PubID      int    `json:"pubId"`
	Topic      int    `json:"topic"`
	TickerID   int    `json:"tickerId"`
	TradeStamp int64  `json:"tradeStamp"`
	TrdSeq     int    `json:"trdSeq"`
	Status     string `json:"status"`
}

// Type104Message is a message type returned from Webull MQTT quote streams
type Type104Message struct {
	TransID int `json:"transId"`
	BidList []struct {
		Price  string `json:"price,omitempty"`
		Volume string `json:"volume,omitempty"`
	} `json:"bidList"`
	AskList []struct {
		Price  string `json:"price,omitempty"`
		Volume string `json:"volume,omitempty"`
	} `json:"askList"`
	PubID    int    `json:"pubId"`
	Topic    int    `json:"topic"`
	TickerID int    `json:"tickerId"`
	TrdSeq   int    `json:"trdSeq"`
	Status   string `json:"status"`
}

// Type105Message is a message type returned from Webull MQTT quote streams
type Type105Message struct {
	Deal struct {
		TrdBs     string `json:"trdBs"`
		Volume    string `json:"volume"`
		TradeTime string `json:"tradeTime"`
		Price     string `json:"price"`
	} `json:"deal"`
	TransID        int    `json:"transId"`
	High           string `json:"high"`
	Low            string `json:"low"`
	ChangeRatio    string `json:"changeRatio"`
	PubID          int    `json:"pubId"`
	TradeStamp     int64  `json:"tradeStamp"`
	Close          string `json:"close"`
	Change         string `json:"change"`
	MarketValue    string `json:"marketValue"`
	VibrateRatio   string `json:"vibrateRatio"`
	Volume         string `json:"volume"`
	TradeTime      string `json:"tradeTime"`
	NegMarketValue string `json:"negMarketValue"`
	TurnoverRate   string `json:"turnoverRate"`
	Topic          int    `json:"topic"`
	TickerID       int    `json:"tickerId"`
	TrdSeq         int    `json:"trdSeq"`
	Open           string `json:"open"`
	Status         string `json:"status"`
}

// Type106Message is a message type returned from Webull MQTT quote streams
type Type106Message struct {
	Depth struct {
		NtvAggAskList []struct {
			Price  string `json:"price"`
			Volume string `json:"volume"`
		} `json:"ntvAggAskList"`
		NtvAggBidList []struct {
			Price  string `json:"price"`
			Volume string `json:"volume"`
		} `json:"ntvAggBidList"`
	} `json:"depth"`
	PubID    int    `json:"pubId"`
	Topic    int    `json:"topic"`
	TickerID int    `json:"tickerId"`
	Status   string `json:"status"`
}

// Type107Message is a message type returned from Webull MQTT quote streams
type Type107Message struct {
	Deal struct {
		TrdBs     string `json:"trdBs"`
		Volume    string `json:"volume"`
		TradeTime string `json:"tradeTime"`
		Price     string `json:"price"`
	} `json:"deal"`
	TransID        int    `json:"transId"`
	High           string `json:"high"`
	Low            string `json:"low"`
	ChangeRatio    string `json:"changeRatio"`
	PubID          int    `json:"pubId"`
	TradeStamp     int64  `json:"tradeStamp"`
	Close          string `json:"close"`
	Change         string `json:"change"`
	MarketValue    string `json:"marketValue"`
	VibrateRatio   string `json:"vibrateRatio"`
	Volume         string `json:"volume"`
	TradeTime      string `json:"tradeTime"`
	NegMarketValue string `json:"negMarketValue"`
	TurnoverRate   string `json:"turnoverRate"`
	Topic          int    `json:"topic"`
	TickerID       int    `json:"tickerId"`
	TrdSeq         int    `json:"trdSeq"`
	Open           string `json:"open"`
	Status         string `json:"status"`
}

// Type108Message is a message type returned from Webull MQTT quote streams
type Type108Message struct {
	Deal struct {
		TrdBs     string `json:"trdBs"`
		Volume    string `json:"volume"`
		TradeTime string `json:"tradeTime"`
		Price     string `json:"price"`
	} `json:"deal"`
	TransID        int    `json:"transId"`
	ChangeRatio    string `json:"changeRatio"`
	PubID          int    `json:"pubId"`
	TradeStamp     int64  `json:"tradeStamp"`
	Close          string `json:"close"`
	Change         string `json:"change"`
	MarketValue    string `json:"marketValue"`
	Volume         string `json:"volume"`
	TradeTime      string `json:"tradeTime"`
	NegMarketValue string `json:"negMarketValue"`
	Topic          int    `json:"topic"`
	TickerID       int    `json:"tickerId"`
	TrdSeq         int    `json:"trdSeq"`
	Status         string `json:"status"`
}

// ConnectStreamingQuotes is a utility function for connecting to WS streaming API
func (c *Client) ConnectStreamingQuotes(ctx context.Context, username, password, deviceID, accessToken string, messageTypes, tickerIDs []string) error {
	var (
		addr = streamingQuotesAddr
		qos = 1
		opts = MQTT.NewClientOptions()
	)
	MQTT.DEBUG = log.New(os.Stdout, "", 0)
	MQTT.ERROR = log.New(os.Stdout, "", 0)
	opts.SetClientID(deviceID)
	opts.AddBroker(addr)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(6 * time.Second)
	opts.SetTLSConfig(&tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert})
	opts.SetCleanSession(true)
	opts.SetDefaultPublishHandler(pushChannel)
	log.Printf("Connecting to: %s", addr)
	client := MQTT.NewClient(opts)
	//client.AddRoute("#", pushChannel) // wildcard
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
	if token := client.Subscribe(helloMsg, byte(qos), pushChannel); token.WaitTimeout(time.Second*60) && token.Error() != nil {
		return token.Error()
	}
	for _, messageType := range messageTypes {
		subscriptionReq := fmt.Sprintf(`{"tickerIds": [%v],"type": "%s"}`, strings.Join(tickerIDs, `,`), messageType)
		if token := client.Subscribe(subscriptionReq, byte(qos), pushChannel); token.WaitTimeout(time.Second*20) && token.Error() != nil {
			return token.Error()
		}
	}
	for {
		select {
		case <-ctx.Done():
			switch ctx.Err() {
			case context.DeadlineExceeded:
				fmt.Println("context timeout exceeded")
				return nil
			case context.Canceled:
				fmt.Println("context cancelled by force. whole process is complete")
				return nil
			}
		default:
			// keep going
			var incoming wsMessage
			select {
			case incoming = <-bufferedMessages:
				if callback, ok := c.WebsocketCallbacks[fmt.Sprintf("%d", incoming.Topic.Type)]; ok {
					if err := callback(ctx, incoming.Topic, incoming.Message); err != nil {
						log.Print("Error running user callback: ", err.Error())
					}
				} else { /* Skip, no user callback assigned yet (can change run-time) */}
			default:
				fmt.Printf("Nothing matched")
			}
		}
		time.Sleep(time.Millisecond*250)
	}
}
