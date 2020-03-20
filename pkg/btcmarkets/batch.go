package btcmarkets

import "time"

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
