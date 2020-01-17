package btcmarkets

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
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
	Market MarketService
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
	c.Market = &MarketServiceOp{client: c}

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

// Do sends and API request and return the API Resonse.
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

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}

// An ErrorResponse reports the error caused by an API request
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
