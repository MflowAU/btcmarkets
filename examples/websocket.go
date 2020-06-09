package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MflowAU/btcmarkets/pkg/btcmarkets"
)

func websocketExample() {
	hc := &http.Client{
		Timeout: time.Second * 10,
	}

	conf := btcmarkets.ClientConfig{
		BaseURL:     nil,
		WsURL:       nil,
		APIKey:      "25d55ef7-f33e-49e8",
		APISecret:   "TXlTdXBlclNlY3JldEtleQ==",
		Httpclient:  hc,
		RateLimiter: nil,
	}

	c, err := btcmarkets.NewBTCMClient(conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	subm := btcmarkets.WSSubscribeMessage{
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
			"orderChange",
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
	fmt.Println("Done...")
}
