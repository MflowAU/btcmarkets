package btcmarkets

import (
	"errors"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

// Unexported helper function to check if string is found in array of strings
// TODO: move into different file ex: utils.go or something similar
func stringInArray(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

// MarketService is an interface for interfacing with
// the BTCMarkets Markets API
// type MarketService interface {
// 	AllMarkets() ([]Market, error)
// 	GetMarketTicker(marketID string) (*Ticker, error)
// 	GetMarketTrades(marketID string) ([]Trade, error)
// }

// MarketServiceOp struct is used to perform Market API
// Operation on BTCMarkets via the BTCMClient Market Interface
type MarketServiceOp struct {
	client *BTCMClient
}

// var _ MarketService = &MarketServiceOp{}

// Market holds a tradable market instrument
type Market struct {
	MarketID       string  `json:"marketId"`
	BaseAsset      string  `json:"baseAssetName"`
	QuoteAsset     string  `json:"quoteAssetName"`
	MinOrderAmount float64 `json:"minOrderAmount,string"`
	MaxOrderAmount float64 `json:"maxOrderAmount,string"`
	AmountDecimals int64   `json:"amountDecimals,string"`
	PriceDecimals  int64   `json:"priceDecimals,string"`
}

// Ticker holds ticker information
type Ticker struct {
	MarketID  string    `json:"marketId"`
	BestBID   float64   `json:"bestBid,string"`
	BestAsk   float64   `json:"bestAsk,string"`
	LastPrice float64   `json:"lastPrice,string"`
	Volume    float64   `json:"volume24h,string"`
	Change24h float64   `json:"price24h,string"`
	Low24h    float64   `json:"low24h,string"`
	High24h   float64   `json:"high24h,string"`
	Timestamp time.Time `json:"timestamp"`
}

// Trade holds trade information
type Trade struct {
	TradeID   string    `json:"id"`
	Amount    float64   `json:"amount,string"`
	Price     float64   `json:"price,string"`
	Timestamp time.Time `json:"timestamp"`
	Side      string    `json:"side"`
}

// OrderBook holds current orderbook information returned from the exchange
type OrderBook struct {
	MarketID   string     `json:"marketId"`
	SnapshotID int        `json:"snapshotId"`
	Asks       [][]string `json:"asks"`
	Bids       [][]string `json:"bids"`
}

// Candle holds time, open, high, low, close & volume information for a given trading pair
type Candle struct {
	Time   time.Time
	Open   float64
	Close  float64
	Low    float64
	High   float64
	Volume float64
}

//AllMarkets Retrieves list of active markets including configuration for each market
func (s *MarketServiceOp) AllMarkets() ([]Market, error) {
	var markets []Market

	req, err := s.client.NewRequest(http.MethodGet, path.Join(btcMarketsAllMarkets), nil)
	if err != nil {
		return nil, err
	}

	_, err = s.client.Do(req, &markets)

	if err != nil {
		return nil, err
	}

	return markets, nil
}

//GetMarketTicker Returns ticker for the given marketId
func (s *MarketServiceOp) GetMarketTicker(marketID string) (*Ticker, error) {
	var ticker Ticker

	req, err := s.client.NewRequest(http.MethodGet, path.Join(btcMarketsAllMarkets, marketID, btcMarketsGetTicker), nil)
	if err != nil {
		return nil, err
	}

	_, err = s.client.Do(req, &ticker)
	if err != nil {
		return nil, err
	}

	return &ticker, nil
}

//GetMarketTrades Retrieves list of most recent trades for the given market. This API supports pagination.
func (s *MarketServiceOp) GetMarketTrades(marketID string, after, before, limit int) ([]Trade, error) {
	var trades []Trade
	params := url.Values{}

	if after > 0 && before > 0 {
		return nil, errors.New("Using `before` and `after` simultaneously is not supported")
	}

	if limit > 0 {
		params.Set("limit", strconv.Itoa(limit))
	}
	if after >= 0 {
		params.Set("after", strconv.Itoa(after))
	}
	if before > 0 {
		params.Set("before", strconv.Itoa(before))
	}

	req, err := s.client.NewRequest(http.MethodGet, path.Join(btcMarketsAllMarkets, marketID, btcMarketsGetTrades+params.Encode()), nil)

	if err != nil {
		return nil, err
	}

	_, err = s.client.Do(req, &trades)
	if err != nil {
		return nil, err
	}

	return trades, nil
}

// GetMarketOrderbook Retrieves list of bids and asks for a given market. passing level=1 returns top 50 for bids and asks.
// level=2 returns full orderbook (full orderbook data is cached and usually updated every 10 seconds).
// Each market order is represented as an array of string [price, volume]. The attribute, snapshotId, is a
// uniqueue number associated to orderbook and it changes every time orderbook changes.
func (s *MarketServiceOp) GetMarketOrderbook(marketID string, level int) (*OrderBook, error) {
	var orderbooks OrderBook
	params := url.Values{}

	if level < 1 || level > 2 {
		return nil, errors.New("level can only take the value 1 or 2. Please check the API documentation for more details")
	}
	params.Set("level", strconv.Itoa(level))

	req, err := s.client.NewRequest(http.MethodGet, path.Join(btcMarketsAllMarkets, marketID, btcMarketOrderBooks+params.Encode()), nil)
	if err != nil {
		return nil, err
	}

	_, err = s.client.Do(req, &orderbooks)
	if err != nil {
		return nil, err
	}

	return &orderbooks, nil
}

// GetMarketCandles Retrieves array of candles for a given market. Each candle record is an array of string representing
// [time,open,high,low,close,volume] for the time window specified (default time window is 1 day).
// This API can be used to retrieve candles either by pagination (before, after, limit) or by specifying timestamp parameters
// (from and/or to). Pagination parameters can't be combined with timestamp parameters and default behavior is pagination when
// no query param is specified.
// When using timestamp parameters as query string, the maximum number of items that can be retrieved is 1000, and depending
// on the specified timeWindow this can be different time windows. For instance, when using timeWindow=1d then up to 1000 days
// of market candles can be retrieved.
func (s *MarketServiceOp) GetMarketCandles(marketID, timeWindow string, from, to *time.Time, before, after, limit int) ([]Candle, error) {
	var candles []Candle
	var temp [][]string // required to parse time.Time and int into Candles struct

	params := url.Values{}
	timeWindow = strings.ToLower(timeWindow)
	tw := []string{"1m", "1h", "1d"}

	if !stringInArray(timeWindow, tw) {
		return nil, errors.New("timeWindow needs to be set to either 1m, 1h or 1d. Please refer to the documentation for more details")
	}

	params.Set("timeWindow", timeWindow)

	if from != nil {
		params.Set("from", from.Format(time.RFC3339))
	}

	if to != nil {
		params.Set("to", from.Format(time.RFC3339))
	}

	if limit > 0 {
		params.Set("limit", strconv.Itoa(limit))
	}
	if after >= 0 {
		params.Set("after", strconv.Itoa(after))
	}
	if before > 0 {
		params.Set("before", strconv.Itoa(before))
	}

	req, err := s.client.NewRequest(http.MethodGet, path.Join(btcMarketsAllMarkets, marketID, btcMarketsCandles+params.Encode()), nil)
	if err != nil {
		return nil, err
	}

	_, err = s.client.Do(req, &temp)
	if err != nil {
		return nil, err
	}

	var candle Candle
	for i := range temp {
		t, err := time.Parse(time.RFC3339, temp[i][0])
		if err != nil {
			return nil, err
		}
		candle.Time = t

		candle.Open, err = strconv.ParseFloat(temp[i][1], 64)
		if err != nil {
			return nil, err
		}
		candle.High, err = strconv.ParseFloat(temp[i][2], 64)
		if err != nil {
			return nil, err
		}
		candle.Low, err = strconv.ParseFloat(temp[i][3], 64)
		if err != nil {
			return nil, err
		}
		candle.Close, err = strconv.ParseFloat(temp[i][4], 64)
		if err != nil {
			return nil, err
		}
		candle.Volume, err = strconv.ParseFloat(temp[i][5], 64)
		if err != nil {
			return nil, err
		}
		candles = append(candles, candle)
	}
	return candles, nil
}

// GetMultipleTickers This API works similar to /v3/markets/{marketId}/ticker except it retrieves tickers for a given list of marketIds
// provided via query string (e.g. ?marketId=ETH-BTC&marketId=XRP-BTC).
// To gain better performance, restrict the number of marketIds to the items needed for your trading app instead of requesting all markets.
func (s *MarketServiceOp) GetMultipleTickers(marketIDs []string) ([]Ticker, error) {
	var tickers []Ticker
	param := url.Values{}

	for _, v := range marketIDs {
		param.Add("marketId", v)
	}

	req, err := s.client.NewRequest(http.MethodGet, path.Join(btcMarketsAllMarkets, btcMarketsGetTickers+param.Encode()), nil)
	if err != nil {
		return nil, err
	}

	_, err = s.client.Do(req, &tickers)
	if err != nil {
		return nil, err
	}

	return tickers, nil
}

// GetMultipleOrderbooks This API works similar to /v3/markets/{marketId}/orderbook except it retrieves orderbooks for a given list of marketIds provided via query string (e.g. ?marketId=ETH-BTC&marketId=XRP-BTC).
// To gain better performance, restrict the number of marketIds to the items needed for your trading app instead of requesting all markets.
// Retrieving full orderbook (level=2), for multiple markets, was mainly provided for customers who are interested in capturing and keeping full orderbook history.
// Therefore, it's recommended to call this API with lower frequency as the data size can be large and also cached.
func (s *MarketServiceOp) GetMultipleOrderbooks(marketIDs []string, level int) ([]OrderBook, error) {
	var orderbooks []OrderBook
	params := url.Values{}

	if level < 1 || level > 2 {
		return nil, errors.New("level can only take the value 1 or 2. Please check the API documentation for more details")
	}
	params.Add("level", strconv.Itoa(level))

	for i := range marketIDs {
		params.Add("marketId", marketIDs[i])
	}

	req, err := s.client.NewRequest(http.MethodGet, path.Join(btcMarketsAllMarkets, btcMarketMultipleOrderBooks+params.Encode()), nil)
	if err != nil {
		return nil, err
	}

	_, err = s.client.Do(req, &orderbooks)
	if err != nil {
		return nil, err
	}

	return orderbooks, nil
}
