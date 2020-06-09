![Go](https://github.com/MflowAU/btcmarkets/workflows/Go/badge.svg?branch=master)

[![codecov](https://codecov.io/gh/MflowAU/btcmarkets/branch/master/graph/badge.svg)](https://codecov.io/gh/MflowAU/btcmarkets)

# btcmarkets
A GoLang API client for btcmarkets.net

## Status

| Entity Type  | Status             |
| :----------- |:--------------------
| Market Data  | Done
| Order Placement | Done
| Batch Order     | Done
| Trade           | Done
| Fund Management | Done
| Account         | Done
| Report          | Not Impelemented
| Misc            | Done
|Websocket        | Done

## Example

```go
package main

import (
    "fmt"

    "github.com/MflowAU/btcmarkets/pkg/btcmarkets"
)

func main() {
    hc :=  &http.Client{
        Timeout: time.Second * 10,
    }

    conf := ClientConfig{
        BaseURL:    nil,
        WsURL:      nil,
        APIKey:     "25d55ef7-f33e-49e8",
        APISecret:  "TXlTdXBlclNlY3JldEtleQ==",
        Httpclient: hc,
    }

    o, err := c.Order.ListOrders("BTC-AUD", "close", 0, 0, 10)
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Printf("%+v \n\n\n", o)
}

```

For more examples please check the example folder.
