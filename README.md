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
|Ratelimiting     | Partial - Global ratelimit implemented

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
    rl := rate.NewLimiter(rate.Every(10*time.Second), 50)

    conf := btcmarkets.ClientConfig{
        BaseURL:     nil,
        WsURL:       nil,
        APIKey:      "25d55ef7-f33e-49e8",
        APISecret:   "TXlTdXBlclNlY3JldEtleQ==",
        Httpclient:  hc,
        RateLimiter: rl,
    }

    o, err := c.Order.ListOrders("BTC-AUD", "close", 0, 0, 10)
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Printf("%+v \n\n\n", o)
}

```

For more examples please check the example folder.

## Known issues

* The code base requires better inline documentation and Examples to render properly on [pkg.go.dev](https://pkg.go.dev). Please see [Issue #15](https://github.com/MflowAU/btcmarkets/issues/15) if you'd like to contribute to fix this.

* Low Test coverage - This is something that will grow overtime. The project aims for the test coverage to be 80%.

* Websocket event description - BTCMarkets websocket returns a flat JSON message. The flat structure makes it very inefficient to cast the message onto the appropriate struct.

## Support

If you like this project and would want to support it please consider taking a look
at the issues section at:

https://github.com/MflowAU/btcmarkets/issues

or consider donating to

Bitcoin:  1B2kZETHm9tjhPKtCCEo6eWhwT5TfXWE6u
Etherium: 0x00912fdef62ab7d9c1cbee712a58105eb1dbd42f
BitCash:  1B2kZETHm9tjhPKtCCEo6eWhwT5TfXWE6u
