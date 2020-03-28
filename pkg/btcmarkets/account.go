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

// TradingFeeData stores trading fee data
type TradingFeeData struct {
	MakerFeeRate float64 `json:"makerFeeRate,string"`
	TakerFeeRate float64 `json:"takerFeeRate,string"`
	MarketID     string  `json:"marketId"`
}

// TradingFeeResponse stores trading fee data
type TradingFeeResponse struct {
	MonthlyVolume float64          `json:"volume30Day,string"`
	FeeByMarkets  []TradingFeeData `json:"FeeByMarkets"`
}

// AccountBalance struct to represent Account balance JSON response
type AccountBalance struct {
	AssetName string `json:"assetName"`
	Available string `json:"available"`
	Balance   string `json:"balance"`
	Locked    string `json:"locked"`
}

// TransactionData stores data from past transactions
type TransactionData struct {
	ID           string    `json:"id"`
	CreationTime time.Time `json:"creationTime"`
	Description  string    `json:"description"`
	AssetName    string    `json:"assetName"`
	Amount       float64   `json:"amount,string"`
	Balance      float64   `json:"balance,string"`
	FeeType      string    `json:"type"`
	RecordType   string    `json:"recordType"`
	ReferrenceID string    `json:"referrenceId"`
}

// AccountServiceOp perform Account API actions
type AccountServiceOp struct {
	client *BTCMClient
}

// GetTradingFees Returns 30 day trading fee volume plus trading fee per market covering marker and taker.
func (a *AccountServiceOp) GetTradingFees() (TradingFeeResponse, error) {
	var tf TradingFeeResponse

	req, err := a.client.NewRequest(http.MethodGet, btcMarketsTradingFees, nil)
	if err != nil {
		return tf, err
	}

	_, err = a.client.DoAuthenticated(req, nil, &tf)
	if err != nil {
		return tf, err
	}

	return tf, nil
}

// GetWithdrawalLimits Gets Withdrawal limit
// func (a *AccountServiceOp) GetWithdrawalLimits() ()

// GetBalances Returns list of assets covering balance,
// available, and locked amount for each asset due to open orders or
// pending withdrawals. This formula represents the relationship
// between those three elements: balance = available + locked
func (a *AccountServiceOp) GetBalances() ([]AccountBalance, error) {
	var gb []AccountBalance

	req, err := a.client.NewRequest(http.MethodGet, btcMarketsAccountBalance, nil)
	if err != nil {
		return gb, err
	}

	_, err = a.client.DoAuthenticated(req, nil, &gb)
	if err != nil {
		return gb, err
	}

	return gb, nil
}

// ListTransactions Returns detail ledger recoerds for underlying wallets. This API supports pagination.
func (a *AccountServiceOp) ListTransactions(assetName string, before, after int64, limit int32) ([]TransactionData, error) {
	var lt []TransactionData

	if (before > 0) && (after > 0) {
		return lt, errors.New("BTCMarkets only supports either before or after, not both")
	}

	params := url.Values{}
	if assetName != "" {
		assetName = strings.ToUpper(assetName)
		params.Set("assetName", assetName)
	}
	if before > 0 {
		params.Set("before", strconv.FormatInt(before, 10))
	}
	if after > 0 {
		params.Set("after", strconv.FormatInt(after, 10))
	}
	if limit > 0 {
		params.Add("limit", strconv.Itoa(int(limit)))
	}

	req, err := a.client.NewRequest(http.MethodGet, path.Join(btcMarketsTransactions, "?"+params.Encode()), nil)
	if err != nil {
		return lt, err
	}

	_, err = a.client.DoAuthenticated(req, nil, &lt)
	if err != nil {
		return lt, err
	}

	return lt, nil
}
