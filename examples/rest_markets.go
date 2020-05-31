package main

import (
	"fmt"

	"github.com/MflowAU/btcmarkets/pkg/btcmarkets"
)

func maina() {
	c, err := btcmarkets.NewBTCMClient("XXXXXXXXXXXXXXXXXXXXXXXX", "YYYYYYYYYYYYYYYYYYYYYYYY")

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(c)
	markets, err := c.Market.AllMarkets()
	fmt.Printf("%+v \n\n\n", markets)

	ticker, err := c.Market.GetMarketTicker("BTC-AUD")
	fmt.Printf("%+v \n\n\n", ticker)

	t, err := c.Market.GetMarketTicker("BTC-AUD")
	fmt.Printf("%+v \n\n\n", t)

	trades, err := c.Market.GetMarketTrades("BTC-AUD", 0, 0, 10)
	fmt.Printf("%+v \n\n\n", trades)

	orders, err := c.Market.GetMarketOrderbook("BTC-AUD", 2)
	if err != nil {
		fmt.Println(err.Error())
	}

	candles, err := c.Market.GetMarketCandles("BTC-AUD", "1d", nil, nil, 0, 0, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", candles)

	markets := make([]string, 0, 5)
	markets = append(markets, "BTC-AUD")
	tickers, err := c.Market.GetMultipleTickers(markets)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", tickers)

	orders, err := c.Market.GetMultipleOrderbooks(markets, 1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", orders)

	wd, err := c.FundManagement.ListWithdrawls(0, 78234876, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", wd)

	lwd, err := c.FundManagement.ListTransfers(0, 78234876, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", lwd)

	gt, err := c.FundManagement.GetTransfers("931809710")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", gt)

	da, err := c.FundManagement.GetDepositeAddress("ltc", 78234876, 0, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", da)

	wf, err := c.FundManagement.GetWithdrawalFees()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", wf)

	la, err := c.FundManagement.ListAssets()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", la)

	tf, err := c.Account.GetTradingFees()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", tf)

	gb, err := c.Account.GetBalances()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", gb)

	lt, err := c.Account.ListTransactions("btc", 0, 1, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("List Transactions...")
	fmt.Printf("%+v \n\n\n", lt)
}