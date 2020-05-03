package main

import (
	"fmt"

	"github.com/MflowAU/btcmarkets/pkg/btcmarkets"
)

func mainb() {
	c, err := btcmarkets.NewBTCMClient("XXXXXXXXXXXXXXXXXXXXXXXX", "YYYYYYYYYYYYYYYYYYYYYYYY")
	markets := make([]string, 0, 5)
	markets = append(markets, "BTC-AUD")

	o, err := c.Order.ListOrders("BTC-AUD", "close", 0, 0, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", o)

	cno, err := c.Order.CancelOrder("5251295067")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", cno)

	cpo, err := c.Order.CancelOpenOrdersByPairs(markets)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", cpo)

	cao, err := c.Order.CancelAllOpenOrders()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", cao)

	geto, err := c.Order.GetOrder("5251457662")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", geto)

	no, err := c.Order.PlaceNewOrder("BTC-AUD", 0.01, 1, "Limit", "Bid", 0.0, 0.0, "", false, "", "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n\n\n", no)
}
