package webull

import (
	"fmt"
	"os"
	"time"
)

const (
	// HeaderKeyAccessToken variable should be used instead of hard-coding the header key for auth tokens.
	HeaderKeyAccessToken = "access_token"
	// HeaderKeyDeviceID variable should be used instead of hard-coding the header key for device IDs.
	HeaderKeyDeviceID = "did"
	// HeaderKeyTradeToken variable should be used instead of hard-coding the header key for trade auth tokens.
	HeaderKeyTradeToken = "t_token"
	// HeaderKeyTradeTime variable should be used instead of hard-coding the header key for trade time.
	HeaderKeyTradeTime = "t_time"
	// QueryKeyTickerID variable should be used instead of hard-coding the query parameter for Ticker ID.
	QueryKeyTickerID = "tickerID"
	// QueryKeyDerivativeIDs variable should be used instead of hard-coding the query parameter derivative IDs.
	QueryKeyDerivativeIDs = "derivativeIds"
)

func getTimeSeconds() string {
	now := time.Now()
	secs := now.Unix()
	return fmt.Sprintf("%d", secs)
}

func deviceName() string {
	return fmt.Sprintf(os.Getenv("WEBULL_USERNAME") + "@go-client")
}
