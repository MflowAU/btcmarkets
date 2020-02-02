package btcmarkets

import "net/http"

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
func (o *OrderServiceOp) ListOrders(markerID, status string, before, after int64, limit int32) ([]Order, error) {
	var orders []Order

	req, err := o.client.NewRequest(http.MethodGet, btcMarketsOrders, nil)
	if err != nil {
		return nil, err
	}

	_, err = o.client.DoAuthenticated(req, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
