package webull

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	model "gitlab.com/brokerage-api/webull-openapi/openapi"
)

const (
	// DefaultClientID if none is supplied by user.
	DefaultClientID = "2292c4714f144eb08ed3edec7f7ce289"
	// PasswordSalt is used for salting your password
	PasswordSalt = "wl_app-a&b@!423^"
	// DefaultDeviceName is a device name
	DefaultDeviceName = "test"
	// DefaultTokenExpiryFormat is used to parse the custom datetime returned by Webull
	DefaultTokenExpiryFormat = "2006-01-02T15:04:05.000+0000"
)

// Credentials implements oauth2 using the webull implementation
type Credentials struct {
	Username string
	Password string
	ClientID string
	TradePIN string
	MFA      string
	AccountType model.AccountType
}


// Login implements TokenSource
func (c *Client) Login(creds Credentials) (err error) {
	var (
		// Login URL
		u, _ = url.Parse(UserBrokerEndpoint + "/passport/login/v3/account")
		hasher = md5.New()
		response   model.PostLoginResponse
	)

	// Client ID
	c.DeviceID = creds.ClientID
	if c.DeviceID == "" {
		c.DeviceID= DefaultClientID
	}
	// UTF-8 encoded salted password
	hasher.Write([]byte(PasswordSalt + creds.Password))
	pwd := hex.EncodeToString(hasher.Sum(nil))

	// Login request body
	request := model.PostLoginParametersRequest{
		Account:     creds.Username,
		AccountType: creds.AccountType,
		DeviceId:    DefaultClientID,
		DeviceName:  DefaultDeviceName,
		Grade:       0.0,
		Pwd:         pwd,
		RegionId:    1,
	}
	request.ExtInfo.VerificationCode = creds.MFA
	requestBody, _ := json.Marshal(request)
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(requestBody))
	req.Header.Add(HeaderKeyDeviceID, c.DeviceID)
	req.Header.Add(HeaderKeyAccessToken, c.AccessToken)
	if err != nil {
		return errors.Wrap(err, "could not create request")
	}
	err = c.DoAndDecode(req, &response)
	if err != nil {
		return err
	}
	c.AccessToken = response.AccessToken
	c.AccessTokenExpiration, err = time.Parse(DefaultTokenExpiryFormat, response.TokenExpireTime)
	c.RefreshToken = response.RefreshToken
	return
}

// TradeLogin implements TokenSource
func (c *Client) TradeLogin(creds Credentials) (err error) {
	var (
		// Login URL
		u, _ = url.Parse(TradeEndpoint + "/login")
		response   model.PostTradeTokenResponse
		hasher = md5.New()
	)

	// Client ID
	c.DeviceID = creds.ClientID
	if c.DeviceID == "" {
		c.DeviceID = DefaultClientID
	}
	// UTF-8 encoded salted password
	hasher.Write([]byte(PasswordSalt + creds.TradePIN))
	pwd := hex.EncodeToString(hasher.Sum(nil))

	// Login request body
	request := model.PostLoginParametersRequest{
		Account:     creds.Username,
		AccountType: creds.AccountType,
		DeviceId:    DefaultClientID,
		DeviceName:  DefaultDeviceName,
		Grade:       0.0,
		Pwd:         pwd,
		RegionId:    6,
	}
	request.ExtInfo.VerificationCode = creds.MFA
	requestBody, _ := json.Marshal(request)
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(requestBody))
	req.Header.Add(HeaderKeyDeviceID, c.DeviceID)
	req.Header.Add(HeaderKeyAccessToken, c.AccessToken)
	if err != nil {
		return errors.Wrap(err, "could not create request")
	}
	// Send and parse request
	err = c.DoAndDecode(req, &response)
	if err != nil {
		return err
	}
	if response.Success {
		c.TradeToken = response.Data.TradeToken
		c.TradeTokenExpiration = time.Now().Add(time.Duration(response.Data.TradeTokenExpireIn) * time.Millisecond) // Assuming ms?
	}	
	return nil
}