package btcmarkets

import (
	"fmt"
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

func TestGetBalance(t *testing.T) {
	mockServer := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[
			{
			  "assetName": "LTC",
			  "balance": "5",
			  "available": "5",
			  "locked": "0"
			},
			{
			  "assetName": "ETH",
			  "balance": "1.07583642",
			  "available": "1.0",
			  "locked": "0.07583642"
			}
		  ]
		`))
	}

	client, mux, _, teardown, err := setup()
	defer teardown()
	if err != nil {
		t.Fatal(err)
	}
	mux.HandleFunc("/v3/accounts/me/balances", mockServer)

	ab, err := client.Account.GetBalances()
	if err != nil {
		t.Errorf(err.Error())
	}
	for _, a := range ab {
		switch a.AssetName {
		case "ETH":
			fmt.Print("All good")
		case "LTC":
			fmt.Print("All good")
		default:
			t.Errorf("Want ETH or LTC got %v", a.AssetName)
		}
	}
}
