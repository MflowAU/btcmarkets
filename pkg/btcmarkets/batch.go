package btcmarkets

import (
	"errors"
	"net/http"
	"time"
)

// BatchPlaceData stores data for placed batch orders
type BatchPlaceData struct {
	OrderID       string    `json:"orderId"`
	MarketID      string    `json:"marketId"`
	Side          string    `json:"side"`
	Type          string    `json:"type"`
	CreationTime  time.Time `json:"creationTime"`
	Price         float64   `json:"price,string"`
	Amount        float64   `json:"amount,string"`
	OpenAmount    float64   `json:"openAmount,string"`
	Status        string    `json:"status"`
	ClientOrderID string    `json:"clientOrderId"`
}

// CancelBatch stores data for batch cancel request
type CancelBatch struct {
	OrderID       string `json:"orderId,omitempty"`
	ClientOrderID string `json:"clientOrderId,omitempty"`
}

// PlaceBatch stores data for place batch request
type PlaceBatch struct {
	MarketID      string  `json:"marketId"`
	Price         float64 `json:"price"`
	Amount        float64 `json:"amount"`
	OrderType     string  `json:"type"`
	Side          string  `json:"side"`
	TriggerPrice  float64 `json:"triggerPrice,omitempty"`
	TriggerAmount float64 `json:"triggerAmount,omitempty"`
	TimeInForce   string  `json:"timeInForce,omitempty"`
	PostOnly      bool    `json:"postOnly,omitempty"`
	SelfTrade     string  `json:"selfTrade,omitempty"`
	ClientOrderID string  `json:"clientOrderId,omitempty"`
}

// BatchPlaceCancelResponse stores place and cancel batch data
type BatchPlaceCancelResponse struct {
	PlacedOrders      []BatchPlaceData       `json:"placeOrders"`
	CancelledOrders   []CancelOrderResp      `json:"cancelOrders"`
	UnprocessedOrders []UnprocessedBatchResp `json:"unprocessedRequests"`
}

// CancelOrderMethod stores data for Cancel request
type CancelOrderMethod struct {
	CancelOrder CancelBatch `json:"cancelOrder,omitempty"`
}

// BatchOrderServiceOp performs BatchOrder placements
// Operations on BTCMarkets
type BatchOrderServiceOp struct {
	client *BTCMClient
}

// BatchPlaceCancelOrders Use this API to place multiple new
// orders or cancel existing ones via a single request.
// The request for batch processing is an array of items,
// and each item contains instructions for an order placement
// and a cancellation. There are restrictions on the number
// of items in a batch (currently set to 10) so a batch can c
// ontain up to 4 items in any form that is needed
func (b *BatchOrderServiceOp) BatchPlaceCancelOrders(cancelOrders []CancelBatch, placeOrders []PlaceBatch) (BatchPlaceCancelResponse, error) {
	var resp BatchPlaceCancelResponse
	var orderRequests []interface{}

	if len(cancelOrders)+len(placeOrders) > 4 {
		return resp, errors.New("Cannot supply more than 4 orders at a time")
	}

	for x := range cancelOrders {
		orderRequests = append(orderRequests, CancelOrderMethod{CancelOrder: cancelOrders[x]})
	}
	for y := range placeOrders {
		if placeOrders[y].ClientOrderID == "" {
			return resp, errors.New("placeorders must have clientorderids filled")
		}
		orderRequests = append(orderRequests, PlaceOrderMethod{PlaceOrder: placeOrders[y]})
	}

	req, err := b.client.NewRequest(http.MethodPost, btcMarketsBatchOrders, orderRequests)
	if err != nil {
		return resp, err
	}

	_, err = b.client.DoAuthenticated(req, orderRequests, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
