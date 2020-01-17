package btcmarkets

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
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

	u, err := url.Parse(path.Join(btcMarketsAPIURL, btcMarketsAPIVersion))
	if err != nil {
		return nil, err
	}

	c := &BTCMClient{
		apiKey:     apiKey,
		privateKey: p,
		BaseURL:    u,
		client:     http.DefaultClient,
		UserAgent:  "",
	}
	c.Market = &MarketServiceOp{client: c}

	return c, nil
}

// NewRequest creates an API request.
func (c *BTCMClient) NewRequest(method, urlString string, body interface{}) (*http.Request, error) {
	buf := new(bytes.Buffer)
	rel, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)
	u.Path = path.Join(c.BaseURL.Path, rel.Path)

	req, err := http.NewRequest(method, c.BaseURL.ResolveReference(rel).String(), buf)
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
	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}
