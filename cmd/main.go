package main

import (
	"fmt"

	"github.com/MflowAU/btcmarkets/pkg/btcmarkets"
)

func main() {
	c, err := btcmarkets.NewBTCMClient(secret())
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(c)
	// markets, err := c.Market.AllMarkets()
	// fmt.Printf("%+v", markets)
	// fmt.Printf("\n\n\n\n\n")

	// ticker, err := c.Market.GetMarketTicker("BTC-AUD")
	// fmt.Printf("%+v", ticker
	// fmt.Printf("\n\n\n\n\n")

	// t, err := c.Market.GetMarketTicker("BTC-AUD")
	// fmt.Printf("%+v", t)
	// fmt.Printf("\n\n\n\n\n")

	// trades, err := c.Market.GetMarketTrades("BTC-AUD", 0, 0, 10)
	// fmt.Printf("%+v", trades)
	// fmt.Printf("\n\n\n\n\n")

	// orders, err := c.Market.GetMarketOrderbook("BTC-AUD", 2)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v", orders)

	// candles, err := c.Market.GetMarketCandles("BTC-AUD", "1d", nil, nil, 0, 0, 10)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v", candles)

	markets := make([]string, 0, 5)
	markets = append(markets, "BTC-AUD")
	markets = append(markets, "LTC-AUD")
	markets = append(markets, "XRP-AUD")
	// tickers, err := c.Market.GetMultipleTickers(markets)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v", tickers)
	// fmt.Printf("\n\n\n\n\n")

	// orders, err := c.Market.GetMultipleOrderbooks(markets, 1)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v", orders)

	o, err := c.Order.ListOrders("BTC-AUD", "all", 0, 0, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v", o)
}