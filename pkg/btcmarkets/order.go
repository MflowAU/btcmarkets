package btcmarkets

import (
	"errors"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

// Order holds order information
type Order struct {
	Amount       string `json:"amount"`
	CreationTime string `json:"creationTime"`
	MarketID     string `json:"marketId"`
	OpenAmount   string `json:"openAmount"`
	OrderID      string `json:"orderId"`
	Price        string `json:"price"`
	Side         string `json:"side"`
	Status       string `json:"status"`
	Type         string `json:"type"`
}

// CancelOrderResp stores data for cancelled orders
type CancelOrderResp struct {
	ClientOrderID string `json:"clientOrderId"`
	OrderID       string `json:"orderId"`
}

// OrderServiceOp perform Order Placement API
// Operation on BTCMarkets via the BTCMClient Market Interface
type OrderServiceOp struct {
	client *BTCMClient
}

// ListOrders Return an array of historical orders or open orders only.
// All query string parametesr are optional so by default and when no query
// parameter is provided, this API retrieves open orders only for all markets.
// This API supports pagination only when retrieving all orders status=all,
// When sending using status=open all open orders are returned and with no pagination.
func (o *OrderServiceOp) ListOrders(marketID, status string, before, after int64, limit int32) ([]Order, error) {
	var orders []Order

	params := url.Values{}
	if marketID != "" {
		params.Add("marketId", marketID)
	}
	if status != "" && (status == "all" || status == "open") {
		params.Add("status", status)
	}
	if after > 0 {
		params.Add("after", strconv.FormatInt(after, 10))
	}
	if before > 0 {
		params.Add("before", strconv.FormatInt(before, 10))
	}
	if limit > 0 {
		params.Add("limit", strconv.Itoa(int(limit)))
	}

	req, err := o.client.NewRequest(http.MethodGet, path.Join(btcMarketsOrders, "?"+params.Encode()), nil)
	if err != nil {
		return nil, err
	}

	_, err = o.client.DoAuthenticated(req, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// CancelOpenOrdersByPairs Cancels specified trading pairs for open orders for all markets or optionally
//  for a given list of marketIds only.
func (o *OrderServiceOp) CancelOpenOrdersByPairs(marketID []string) ([]CancelOrderResp, error) {
	var c []CancelOrderResp
	params := url.Values{}

	if len(marketID) < 1 {
		return nil, errors.New("CancelOpenOrdersByPairs requires atleast 1 pair")
	}

	for i := range marketID {
		params.Add("marketId", marketID[i])
	}

	req, err := o.client.NewRequest(http.MethodDelete, path.Join(btcMarketsOrders, "?"+params.Encode()), nil)
	if err != nil {
		return nil, err
	}

	_, err = o.client.DoAuthenticated(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// CancelAllOpenOrders Cancels all open orders
func (o *OrderServiceOp) CancelAllOpenOrders() ([]CancelOrderResp, error) {
	var c []CancelOrderResp

	req, err := o.client.NewRequest(http.MethodDelete, path.Join(btcMarketsOrders), nil)
	if err != nil {
		return nil, err
	}

	_, err = o.client.DoAuthenticated(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// CancelOrder Cancels a single order. this API returns http error 400 if the order is realdy cancelled or matched, patrally matched.
func (o *OrderServiceOp) CancelOrder(orderID string) (*CancelOrderResp, error) {
	var c CancelOrderResp

	req, err := o.client.NewRequest(http.MethodDelete, path.Join(btcMarketsOrders, orderID), nil)
	if err != nil {
		return nil, err
	}

	_, err = o.client.DoAuthenticated(req, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// GetOrder Returns an order by using either the exachange `orderId` or `clientOrderId`
func (o *OrderServiceOp) GetOrder(orderID string) (*Order, error) {
	var or Order

	req, err := o.client.NewRequest(http.MethodGet, path.Join(btcMarketsOrders, orderID), nil)
	if err != nil {
		return nil, err
	}

	_, err = o.client.DoAuthenticated(req, &or)
	if err != nil {
		return nil, err
	}

	return &or, nil
}

// func (o *OrderServiceOp) PlaceNewOrder() {

// }
