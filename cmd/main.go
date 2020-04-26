package main

import (
	"context"
	"fmt"
	"time"

	"github.com/MflowAU/btcmarkets/pkg/btcmarkets"
)

func main() {
	c, err := btcmarkets.NewBTCMClient(adminSecret())
	// c, err := btcmarkets.NewBTCMClient(secret())
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(c)
	// markets, err := c.Market.AllMarkets()
	// fmt.Printf("%+v \n\n\n", markets)

	// ticker, err := c.Market.GetMarketTicker("BTC-AUD")
	// fmt.Printf("%+v \n\n\n", ticker

	// t, err := c.Market.GetMarketTicker("BTC-AUD")
	// fmt.Printf("%+v \n\n\n", t)

	// trades, err := c.Market.GetMarketTrades("BTC-AUD", 0, 0, 10)
	// fmt.Printf("%+v \n\n\n", trades)

	// orders, err := c.Market.GetMarketOrderbook("BTC-AUD", 2)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// candles, err := c.Market.GetMarketCandles("BTC-AUD", "1d", nil, nil, 0, 0, 10)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", candles)

	markets := make([]string, 0, 5)
	markets = append(markets, "BTC-AUD")
	// tickers, err := c.Market.GetMultipleTickers(markets)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", tickers)

	// orders, err := c.Market.GetMultipleOrderbooks(markets, 1)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", orders)

	// o, err := c.Order.ListOrders("BTC-AUD", "close", 0, 0, 10)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", o)

	// cno, err := c.Order.CancelOrder("5251295067")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", cno)

	// cpo, err := c.Order.CancelOpenOrdersByPairs(markets)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", cpo)

	// cao, err := c.Order.CancelAllOpenOrders()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", cao)

	// geto, err := c.Order.GetOrder("5251457662")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", geto)

	// no, err := c.Order.PlaceNewOrder("BTC-AUD", 0.01, 1, "Limit", "Bid", 0.0, 0.0, "", false, "", "")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", no)

	// wd, err := c.FundManagement.ListWithdrawls(0, 78234876, 10)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", wd)

	// lwd, err := c.FundManagement.ListTransfers(0, 78234876, 10)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", lwd)

	// gt, err := c.FundManagement.GetTransfers("931809710")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", gt)

	// da, err := c.FundManagement.GetDepositeAddress("ltc", 78234876, 0, 10)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", da)

	// wf, err := c.FundManagement.GetWithdrawalFees()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", wf)

	// la, err := c.FundManagement.ListAssets()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", la)

	// tf, err := c.Account.GetTradingFees()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", tf)

	// gb, err := c.Account.GetBalances()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v \n\n\n", gb)

	// lt, err := c.Account.ListTransactions("btc", 0, 1, 10)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("List Transactions...")
	// fmt.Printf("%+v \n\n\n", lt)

	subm := btcmarkets.WSSubscribeMessageAuth{
		MessageType: "subscribe",
		MarketIds: []string{
			"BTC-AUD",
			"LTC-AUD",
			"XRP-AUD",
			"ETH-AUD",
			"ETC-AUD",
			"OMG-AUD",
			"GNT-AUD",
			"XLM-AUD",
		},
		Channels: []string{
			"trade",
			"tick",
			// "orderbook",
			// "orderbookUpdate",
			// "heartbeat",
		},
	}
	ctx, ctxCancel := context.WithCancel(context.Background())

	ch, err := c.WebSocket.Subscribe(ctx, subm)
	if err != nil {
		fmt.Println(err.Error())
	}

	go func() {
		time.Sleep(time.Second * 10)
		ctxCancel()
	}()

	for i := 0; i < 50; i++ {
		fmt.Println(string(<-ch))
	}

}
