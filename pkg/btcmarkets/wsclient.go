package btcmarkets

import (
	"net/url"
)

// WSSubscribeMessageAuth Subscribe message to initiate WebSocket Connection
type WSSubscribeMessageAuth struct {
	Channels    []string `json:"channels"`
	Key         string   `json:"key"`
	MarketIds   []string `json:"marketIds"`
	MessageType string   `json:"messageType"`
	Signature   string   `json:"signature"`
	Timestamp   string   `json:"timestamp"`
}

// BTCMWSTickEvent The tick event is published every time lastPrice,
// bestBid or bestAsk is updated for a market which is the result of
// orderbook changes or trade matches.
type BTCMWSTickEvent struct {
	BestAsk     string `json:"bestAsk"`
	BestBid     string `json:"bestBid"`
	LastPrice   string `json:"lastPrice"`
	MarketID    string `json:"marketId"`
	MessageType string `json:"messageType"`
	Timestamp   string `json:"timestamp"`
	Volume24h   string `json:"volume24h"`
}

// BTCMWSTradeEvent In order to receive trade events please add trade to
// the list of channels when subscribing via WebSocket.
type BTCMWSTradeEvent struct {
	MarketID    string `json:"marketId"`
	MessageType string `json:"messageType"`
	Price       string `json:"price"`
	Side        string `json:"side"`
	Timestamp   string `json:"timestamp"`
	TradeID     int64  `json:"tradeId"`
	Volume      string `json:"volume"`
}

// BTCMWSOrderbookEvent In order to receive orderbook events please add orderbook to
// the list of channels when subscribing via WebSocket. The current orderbook event represents
// the latest orderbook state and maximum 50 bids and asks are included in each event.
type BTCMWSOrderbookEvent struct {
	Asks        [][]string `json:"asks"`
	Bids        [][]string `json:"bids"`
	MarketID    string     `json:"marketId"`
	MessageType string     `json:"messageType"`
	Timestamp   string     `json:"timestamp"`
}

// BTCMWSOrderbookUpdateEvent In many cases, it's more appropriate to maintain a local copy of
// the exchange orderbook by receiving only updates instead of the entire orderbook.
type BTCMWSOrderbookUpdateEvent struct {
	Asks        [][]interface{} `json:"asks"`
	Bids        [][]interface{} `json:"bids"`
	MarketID    string          `json:"marketId"`
	MessageType string          `json:"messageType"`
	Snapshot    bool            `json:"snapshot"`
	SnapshotID  int64           `json:"snapshotId"`
	Timestamp   string          `json:"timestamp"`
}

// BTCMWSTickResponse Response object res
type BTCMWSTickResponse struct {
	BestAsk     string `json:"bestAsk"`
	BestBid     string `json:"bestBid"`
	High24h     string `json:"high24h"`
	LastPrice   string `json:"lastPrice"`
	Low24h      string `json:"low24h"`
	MarketID    string `json:"marketId"`
	MessageType string `json:"messageType"`
	Price24h    string `json:"price24h"`
	SnapshotID  int64  `json:"snapshotId"`
	Timestamp   string `json:"timestamp"`
	Volume24h   string `json:"volume24h"`
}

// BTCMWSHeartbeatEvent if you subscribe to heartbeat event
// then the server will send you a heartbeat event every 5 seconds.
// Note: Once a new subscription request is confirmed, a single heartbeat
// event is published to the client in order to confirm the connection working.
// This is regardless of requesting to subscribe to heartbeat channel.
type BTCMWSHeartbeatEvent struct {
	Channels []struct {
		MarketIds []string `json:"marketIds"`
		Name      string   `json:"name"`
	} `json:"channels"`
	MessageType string `json:"messageType"`
}

// BTCMWSErrorEvent in case of errors, a message type of error is published.
// Authentication error
// Invalid input error
// Internal server error
// Throttle error
// Invalid Channel names
// Invalid MarketId
// Authenticate Error
type BTCMWSErrorEvent struct {
	Code        int64  `json:"code"`
	Message     string `json:"message"`
	MessageType string `json:"messageType"`
}

// BTCMWSClient is the WebSocket Client struct
// It embeds the BTCMClient to get access to the underlying properties
type BTCMWSClient struct {
	BTCMClient
	BaseURL *url.URL
}

// NewBTCMWSClient Factory function to generate a websocket connection to BTCMarket
func NewBTCMWSClient(sm WSSubscribeMessageAuth) {

}
