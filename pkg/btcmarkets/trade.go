package btcmarkets

import (
	"errors"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

// TradeHistoryData stores data of past trades
type TradeHistoryData struct {
	ID            string    `json:"id"`
	MarketID      string    `json:"marketId"`
	Timestamp     time.Time `json:"timestamp"`
	Price         float64   `json:"price,string"`
	Amount        float64   `json:"amount,string"`
	Side          string    `json:"side"`
	Fee           float64   `json:"fee,string"`
	OrderID       string    `json:"orderId"`
	LiquidityType string    `json:"liquidityType"`
}

// TradeHistoryServiceOp performs Trade API
// Operations on BTCMarkets
type TradeHistoryServiceOp struct {
	client *BTCMClient
}

// ListTrades returns trade history
func (th *TradeHistoryServiceOp) ListTrades(marketID, orderID string, before, after, limit int64) ([]TradeHistoryData, error) {
	var resp []TradeHistoryData

	if (before > 0) && (after >= 0) {
		return resp, errors.New("BTCMarkets only supports either before or after, not both")
	}

	params := url.Values{}

	if marketID != "" {
		params.Set("marketId", marketID)
	}
	if orderID != "" {
		params.Set("orderId", orderID)
	}
	if before > 0 {
		params.Set("before", strconv.FormatInt(before, 10))
	}
	if after >= 0 {
		params.Set("after", strconv.FormatInt(after, 10))
	}
	if limit > 0 {
		params.Set("limit", strconv.FormatInt(limit, 10))
	}

	req, err := th.client.NewRequest(http.MethodGet, path.Join(btcMarketsTradeHistory, "?"+params.Encode()), nil)
	if err != nil {
		return resp, err
	}

	_, err = th.client.DoAuthenticated(req, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetTradeByID returns the singular trade of the ID given
func (th *TradeHistoryServiceOp) GetTradeByID(id string) (TradeHistoryData, error) {
	var resp TradeHistoryData

	req, err := th.client.NewRequest(http.MethodGet, path.Join(btcMarketsTradeHistory, id), nil)
	if err != nil {
		return resp, err
	}
	_, err = th.client.DoAuthenticated(req, nil, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
