package btcmarkets

import "net/http"

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
