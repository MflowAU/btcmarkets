package btcmarkets

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/net/context"
)

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

// WSSubscribeMessage Subscribe message to initiate WebSocket Connection
type WSSubscribeMessage struct {
	Channels    []string `json:"channels"`
	MarketIds   []string `json:"marketIds"`
	MessageType string   `json:"messageType"`
	Timestamp   string   `json:"timestamp"`
	Key         string   `json:"key"`
	Signature   string   `json:"signature"`
}

// WebSocketServiceOp WebSocket feed provides real-time market data covering
//  orderbook updates, order life cycle and trades
type WebSocketServiceOp struct {
	client *BTCMClient
}

// Subscribe returns a channel of bytes with messages from the websocket.
// The consumer of this method will need to handle the Implicit type Conversion
// of the bytes returned on the channel.
// This method needs to be called with a ContextWithCancel as first parameter to be able
// close the websocket and a SubscribeMessage to start receiving events for the
// specified channels and marketIds
func (ws *WebSocketServiceOp) Subscribe(ctx context.Context, m WSSubscribeMessage) (chan []byte, error) {
	wsmessages := make(chan []byte)

	c, _, err := websocket.DefaultDialer.Dial(ws.client.WSURL.String(), nil)

	if err != nil {
		fmt.Println("Error Dialing WebSocket Connection: ", err.Error())
		return nil, err
	}

	if len(ws.client.apiKey) > 0 {
		m.Key = ws.client.apiKey
	}

	if len(ws.client.privateKey) > 0 {
		t := strconv.FormatInt(time.Now().UTC().UnixNano()/1000000, 10)
		m.Timestamp = t
		strToSign := "/users/self/subscribe" + "\n" + t
		m.Signature = ws.client.signMessage(strToSign)
	}
	m.MessageType = "subscribe"

	err = c.WriteJSON(m)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	go func() {
		defer c.Close()
		defer close(wsmessages)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_, payload, err := c.ReadMessage()
				if err != nil {
					fmt.Println(err.Error())
					wsmessages <- []byte(err.Error())
					// return
					// TODO: check if this code block should return on error
				}
				wsmessages <- payload
			}
		}
	}()

	return wsmessages, nil
}
