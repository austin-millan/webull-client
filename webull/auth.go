package webull

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"

	// "fmt"
	"io"
	"io/ioutil"

	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	model "gitlab.com/brokerage-api/webull-openapi/openapi"
	"golang.org/x/oauth2"
)

const (
	// DefaultDeviceID if none is supplied by user.
	DefaultDeviceID = "2292c4714f144eb08ed3edec7f7ce284"
	// PasswordSalt is used for salting your password
	PasswordSalt = "wl_app-a&b@!423^"
	// DefaultDeviceName is a device name
	DefaultDeviceName = "test"
	// DefaultTokenExpiryFormat is used to parse the custom datetime returned by Webull
	DefaultTokenExpiryFormat = "2006-01-02T15:04:05.000+0000"
)

// Credentials implements oauth2 using the webull implementation
type Credentials struct {
	Username    string
	Password    string
	DeviceID    string
	TradePIN    string
	MFA         string
	DeviceName  string
	AccountType model.AccountType
	Creds       oauth2.TokenSource
}

// Token implements TokenSource
func (c *Client) Token() (*oauth2.Token, error) {
	var (
		u, _       = url.Parse(UserBrokerEndpoint + "/passport/login/v3/account")
		response   model.PostLoginResponse
		httpClient = http.Client{Timeout: time.Second * 10}
		cliID      = c.DeviceID
		deviceName = c.DeviceName
	)
	// Client ID
	if cliID == "" {
		cliID = DefaultDeviceID
	}
	// Device Name
	if deviceName == "" {
		deviceName = DefaultDeviceName
	}
	// Login request body
	requestBody, err := json.Marshal(model.PostLoginParametersRequest{
		Account:     c.Username,
		AccountType: c.AccountType,
		DeviceId:    cliID,
		DeviceName:  deviceName,
		Grade:       0.0,
		Pwd:         c.HashedPassword,
		RegionId:    1,
		ExtInfo:     model.PostLoginParametersRequestExtInfo{VerificationCode: c.MFA},
	})
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(HeaderKeyDeviceID, cliID)
	if err != nil {
		return nil, errors.Wrap(err, "could not create request")
	}
	tok := oauth2.Token{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode/100 != 2 {
		b := &bytes.Buffer{}
		var e model.ErrorBody
		err = json.NewDecoder(io.TeeReader(res.Body, b)).Decode(&e)
		if err != nil {
			return nil, fmt.Errorf("got response %q and could not decode error body %q", res.Status, b.String())
		}
		return nil, fmt.Errorf(e.Msg)
	}

	if err != nil {
		return nil, fmt.Errorf("Got read error on body: %s", err.Error())
	}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("Got JSON unmarshal error on body: %s", err.Error())
	}
	if err != nil {
		return nil, err
	}
	tok.Expiry, err = time.Parse(DefaultTokenExpiryFormat, response.TokenExpireTime)
	c.AccessTokenExpiration = tok.Expiry
	tok.TokenType = "Token"
	tok.AccessToken, c.AccessToken = response.AccessToken, response.AccessToken
	tok.RefreshToken, c.RefreshToken = response.RefreshToken, response.RefreshToken
	c.UUID = response.Uuid
	return &tok, nil
}

// Login implements TokenSource
func (c *Client) Login(creds Credentials) (err error) {
	var (
		u, _     = url.Parse(UserBrokerEndpoint + "/passport/login/v3/account")
		hasher   = md5.New()
		response model.PostLoginResponse
	)

	// Client ID
	if creds.DeviceID != "" {
		c.DeviceID = creds.DeviceID
	} else {
		if c.DeviceID == "" {
			c.DeviceID = DefaultDeviceID
		}
	}

	// Client Name
	if creds.DeviceName != "" {
		c.DeviceName = creds.DeviceName
	} else {
		if c.DeviceName == "" {
			c.DeviceName = DefaultDeviceName
		}
	}

	// Client Name
	if creds.Username != "" {
		c.Username = creds.Username
	} else {
		if c.Username == "" {
			return fmt.Errorf("Username required")
		}
	}

	// Client Name
	if creds.Password != "" {
		// UTF-8 encoded salted password
		hasher.Write([]byte(PasswordSalt + creds.Password))
		c.HashedPassword = hex.EncodeToString(hasher.Sum(nil))
	} else {
		if c.HashedPassword == "" {
			return fmt.Errorf("Password has not been set")
		}
	}
	c.AccountType = creds.AccountType

	// Login request body
	request := model.PostLoginParametersRequest{
		Account:     c.Username,
		AccountType: c.AccountType,
		DeviceId:    c.DeviceID,
		DeviceName:  c.DeviceName,
		Grade:       0.0,
		Pwd:         c.HashedPassword,
		RegionId:    1,
	}
	request.ExtInfo.VerificationCode = creds.MFA
	requestBody, _ := json.Marshal(request)
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(requestBody))
	req.Header.Add(HeaderKeyDeviceID, c.DeviceID)
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
	c.UUID = response.Uuid
	return
}

// TradeLogin implements TokenSource
func (c *Client) TradeLogin(creds Credentials) (err error) {
	var (
		// Login URL
		u, _     = url.Parse(TradeEndpoint + "/login")
		response model.PostTradeTokenResponse
		hasher   = md5.New()
		pwd string
	)

	// Client ID
	if creds.DeviceID != "" {
		c.DeviceID = creds.DeviceID
	} else {
		if c.DeviceID == "" {
			c.DeviceID = DefaultDeviceID
		}
	}

	// Client Name
	if creds.DeviceName != "" {
		c.DeviceName = creds.DeviceName
	} else {
		if c.DeviceName == "" {
			c.DeviceName = DefaultDeviceName
		}
	}

	// Client Name
	if creds.Username != "" {
		c.Username = creds.Username
	} else {
		if c.Username == "" {
			return fmt.Errorf("Username required")
		}
	}

	// Client Name
	if creds.TradePIN != "" {
		// UTF-8 encoded salted password
		hasher.Write([]byte(PasswordSalt + creds.TradePIN))
		pwd = hex.EncodeToString(hasher.Sum(nil))
		c.HashedPassword = pwd
	} else {
		if c.HashedPassword == "" {
			return fmt.Errorf("Password has not been set")
		}
	}
	c.AccountType = creds.AccountType

	// Login request body
	request := model.PostLoginParametersRequest{
		Account:     creds.Username,
		AccountType: creds.AccountType,
		DeviceId:    DefaultDeviceID,
		DeviceName:  DefaultDeviceName,
		Grade:       0.0,
		Pwd:         c.HashedPassword,
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

// GetMFA requests for a 2FA code
func (c *Client) GetMFA(creds Credentials) (err error) {
	var (
		// Login URL
		u, _        = url.Parse(UserBrokerEndpoint + "/passport/verificationCode/sendCode")
		response    interface{}
		queryParams = make(map[string]string)
		headersMap  = make(map[string]string)
	)

	// Client ID
	c.DeviceID = creds.DeviceID
	if c.DeviceID == "" {
		c.DeviceID = DefaultDeviceID
	}

	headersMap["did"] = c.DeviceID

	queryParams["deviceId"] = c.DeviceID
	queryParams["accountType"] = fmt.Sprintf("%d", creds.AccountType)
	queryParams["account"] = creds.Username
	queryParams["codeType"] = "5"
	queryParams["regionCode"] = "1"

	if err != nil {
		return errors.Wrap(err, "could not create request")
	}
	// Send and parse request
	err = c.PostAndDecode(*u, response, &headersMap, &queryParams, nil)
	if err != nil {
		return err
	}
	return nil
}
