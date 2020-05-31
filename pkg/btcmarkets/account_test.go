package btcmarkets

import (
	"net/http"
	"testing"
)

func TestAccountGetTradingFees(t *testing.T) {
	mockserver := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"volume30Day": "0.0098275",
			"feeByMarkets": [
			  {
				"makerFeeRate": "0.00849999",
				"takerFeeRate": "0.00849999",
				"marketId": "BTC-AUD"
			  },
			  {
				"makerFeeRate": "0.00849999",
				"takerFeeRate": "0.00849999",
				"marketId": "LTC-AUD"
			  },
			  {
				"makerFeeRate": "0.0022",
				"takerFeeRate": "0.0022",
				"marketId": "LTC-BTC"
			  }
			]
		  }
		`))
	}

	client, mux, _, teardown, err := setup()
	defer teardown()
	if err != nil {
		t.Fatal(err)
	}
	mux.HandleFunc("/v3/accounts/me/trading-fees", mockserver)

	_, err = client.Account.GetTradingFees()
	if err != nil {
		t.Errorf(err.Error())
	}
}
