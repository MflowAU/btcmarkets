package btcmarkets

import (
	"errors"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

// MarketService is an interface for interfacing with
// the BTCMarkets Markets API
type MarketService interface {
	AllMarkets() ([]Market, error)
	GetMarketTicker(marketID string) (*Ticker, error)
	GetMarketTrades(marketID string) ([]Trade, error)
}

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
