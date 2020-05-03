package main

import (
	"context"
	"fmt"
	"time"

	"github.com/MflowAU/btcmarkets/pkg/btcmarkets"
)

func main() {
	c, err := btcmarkets.NewBTCMClient("XXXXXXXXXXXXXXXXXXXXXXXX", "YYYYYYYYYYYYYYYYYYYYYYYY")

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
