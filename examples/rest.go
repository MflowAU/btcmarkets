package main

import (
	"log"
	"net/http"
	"time"

	"github.com/MflowAU/btcmarkets/pkg/btcmarkets"
	"golang.org/x/time/rate"
)

func restExample() {
	hc := &http.Client{
		Timeout: time.Second * 10,
	}

	rl := rate.NewLimiter(rate.Every(10*time.Second), 50)

	conf := btcmarkets.ClientConfig{
		BaseURL:     nil,
		WsURL:       nil,
		APIKey:      "25d55ef7-f33e-49e8",
		APISecret:   "TXlTdXBlclNlY3JldEtleQ==",
		Httpclient:  hc,
		RateLimiter: rl,
	}

	c, err := btcmarkets.NewBTCMClient(conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	markets := make([]string, 0, 5)
	markets = append(markets, "BTC-AUD")

	o, err := c.Order.ListOrders("BTC-AUD", "close", 0, 0, 10)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("%+v \n\n\n", o)

	cpo, err := c.Order.CancelOpenOrdersByPairs(markets)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("%+v \n\n\n", cpo)

	cao, err := c.Order.CancelAllOpenOrders()
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("%+v \n\n\n", cao)

	no, err := c.Order.PlaceNewOrder("BTC-AUD", 0.01, 1, "Limit", "Bid", 0.0, 0.0, "", false, "", "")
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("%+v \n\n\n", no)

	da, err := c.FundManagement.GetDepositeAddress("ltc", 78234876, 0, 10)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("%+v \n\n\n", da)

	wf, err := c.FundManagement.GetWithdrawalFees()
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("%+v \n\n\n", wf)

	la, err := c.FundManagement.ListAssets()
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("%+v \n\n\n", la)

	tf, err := c.Account.GetTradingFees()
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("%+v \n\n\n", tf)
}
