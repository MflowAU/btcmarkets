package btcmarkets

import (
	"net/http"
	"path"
)

// MarketService is an interface for interfacing with
// the BTCMarkets Markets API
type MarketService interface {
	AllMarkets() ([]Market, error)
	GetMarketTicker(marketID string) (*Ticker, error)
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

//AllMarkets Returns all available markets on the exchange
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
