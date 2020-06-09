package btcmarkets

import (
	"net/http"
	"testing"
	"time"

	"golang.org/x/time/rate"
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

	client, mux, _, teardown, err := setup(nil)
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

	client, mux, _, teardown, err := setup(nil)
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
			t.Log("All good")
		case "LTC":
			t.Log("All good")
		default:
			t.Errorf("Want ETH or LTC got %v", a.AssetName)
		}
	}
}

func TestGetBalanceWithRateLimit(t *testing.T) {
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

	client, mux, _, teardown, err := setup(rate.NewLimiter(rate.Every(1*time.Second), 1))

	defer teardown()
	if err != nil {
		t.Fatal(err)
	}
	mux.HandleFunc("/v3/accounts/me/balances", mockServer)

	before := time.Now()
	for i := 0; i <= 1; i++ {
		ab, err := client.Account.GetBalances()
		if err != nil {
			t.Errorf(err.Error())
		}
		for _, a := range ab {
			switch a.AssetName {
			case "ETH":
				t.Log("All good")
			case "LTC":
				t.Log("All good")
			default:
				t.Errorf("Want ETH or LTC got %v", a.AssetName)
			}
		}
	}
	after := time.Now()
	diff := after.Sub(before)

	if diff.Seconds() < 1 {
		t.Errorf("Expected execution time > 3 but took %v", diff.Seconds())
	}

}
