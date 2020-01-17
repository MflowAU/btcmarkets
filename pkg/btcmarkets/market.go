package btcmarkets

import (
	"fmt"
)

type MarketService interface {
	AllMarkets() ([]Market, error)
}

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

	req, err := s.client.NewRequest("GET", "https://api.btcmarkets.net/v3/markets/", nil)
	if err != nil {
		return nil, err
	}

	_, err = s.client.Do(req, &markets)

	if err != nil {
		return nil, err
	}

	fmt.Println(markets)
	return markets, nil

}
