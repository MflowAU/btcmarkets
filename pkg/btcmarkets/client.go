package btcmarkets

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strconv"
	"time"
)

// BTCMClient is the main struct type representing an interface with the API as a
// particular client user.
type BTCMClient struct {
	apiKey     string
	privateKey []byte
	BaseURL    *url.URL
	client     *http.Client
	UserAgent  string

	// Services used for communicating with the API
	// Market MarketService
	Market MarketServiceOp
	Order  OrderServiceOp
}

// ServerTime holds the BTCMarket Server time returned after making
// a call to /v3/time
type ServerTime struct {
	Timestamp string `json:"timestamp"`
}

// NewBTCMClient should return a new instance of BTCMarkets Client
func NewBTCMClient(apiKey, apiSecret string) (*BTCMClient, error) {
	if apiKey == "" || apiSecret == "" {
		return nil, errors.New("")
	}

	p, err := base64.StdEncoding.DecodeString(apiSecret)
	if err != nil {
		return nil, errors.New("Error Decoding apiSecret")
	}

	u, err := url.Parse(btcMarketsAPIURL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, btcMarketsAPIVersion)

	c := &BTCMClient{
		apiKey:     apiKey,
		privateKey: p,
		BaseURL:    u,
		client:     http.DefaultClient,
		UserAgent:  "mflow/golang-client",
	}
	// c.Market = &MarketServiceOp{client: c}
	c.Market = MarketServiceOp{client: c}
	c.Order = OrderServiceOp{client: c}

	return c, nil
}

// NewRequest creates an API request.
func (c *BTCMClient) NewRequest(method, urlPath string, body interface{}) (*http.Request, error) {
	buf := new(bytes.Buffer)
	rel, err := url.Parse(urlPath)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)
	u.Path = path.Join(c.BaseURL.Path, rel.Path)

	// fmt.Println(c.BaseURL.ResolveReference(rel).String())
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do sends and API request and return the API Response.
// The response is JSON decoded and store in the value pointed
// by v, or returns an error
func (c *BTCMClient) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if v != nil {
		// err = json.NewDecoder(body).Decode(v)
		err = json.Unmarshal(body, v)
	}
	return resp, err
}

// DoAuthenticated makes API requeßst and return the API Response
func (c *BTCMClient) DoAuthenticated(req *http.Request, v interface{}) (*http.Response, error) {
	t := strconv.FormatInt(time.Now().UTC().UnixNano()/1000000, 10)
	p := req.URL.RequestURI()
	m := req.Method + p + t + "" //TODO: change empty string with data passed for the body
	h := hmac.New(sha512.New, c.privateKey)
	h.Write([]byte(m))

	// req.Header.Add("Accept", "application/json")
	// req.Header.Add("Accept-Charset", "UTF-8")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("BM-AUTH-APIKEY", c.apiKey)
	req.Header.Add("BM-AUTH-TIMESTAMP", t)
	req.Header.Add("BM-AUTH-SIGNATURE", base64.StdEncoding.EncodeToString(h.Sum(nil)))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// err = CheckResponse(resp)
	// if err != nil {
	// 	return nil, err
	// }

	body, err := ioutil.ReadAll(resp.Body)
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%q", dump)
	if v != nil {
		// err = json.NewDecoder(body).Decode(v)
		err = json.Unmarshal(body, v)
	}

	return resp, err
}

// An ErrorResponse reports the error caused by an API request
// If response is 200 but there is marshalling error there could
// be a issue with Unmarshaling the data into the struct
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Error message - maybe the json for this is "fault"
	Message string `json:"message"`
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			return err
		}
	}
	return errors.New("Error Response")

}

// GetServerTime get the BTCMarkets server time
// This is required to build a valid authenticated request
func (c *BTCMClient) GetServerTime() (ServerTime, error) {
	t := ServerTime{}
	req, err := c.NewRequest(http.MethodGet, btcMarketsGetTime, nil)
	if err != nil {
		return t, err
	}

	_, err = c.Do(req, &t)
	if err != nil {
		return t, err
	}
	return t, nil
}
