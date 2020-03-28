package btcmarkets

import "time"

const (
	btcMarketsAPIURL     = "https://api.btcmarkets.net"
	btcMarketsAPIVersion = "/v3"

	// UnAuthenticated EPs
	btcMarketsAllMarkets         = "/markets/"
	btcMarketsGetTicker          = "/ticker/"
	btcMarketsGetTickers         = "/tickers?"
	btcMarketsGetTrades          = "/trades?"
	btcMarketOrderBooks          = "/orderbook?"
	btcMarketMultipleOrderBooks  = "/orderbooks?"
	btcMarketsCandles            = "/candles?"
	btcMarketsTickers            = "tickers?"
	btcMarketsMultipleOrderbooks = "/orderbooks?"
	btcMarketsGetTime            = "/time"
	btcMarketsWithdrawalFees     = "/withdrawal-fees"
	btcMarketsUnauthPath         = btcMarketsAPIURL + btcMarketsAPIVersion + btcMarketsAllMarkets

	// Authenticated EPs
	btcMarketsAccountBalance = "/accounts/me/balances"
	btcMarketsTradingFees    = "/accounts/me/trading-fees"
	btcMarketsTransactions   = "/accounts/me/transactions"
	btcMarketsOrders         = "/orders"
	btcMarketsTradeHistory   = "/trades"
	btcMarketsWithdrawals    = "/withdrawals"
	btcMarketsDeposits       = "/deposits"
	btcMarketsTransfers      = "/transfers"
	btcMarketsAddresses      = "/addresses"
	btcMarketsAssets         = "/assets"
	btcMarketsReports        = "/reports"
	btcMarketsBatchOrders    = "/batchorders"

	btcmarketsAuthLimit   = 3
	btcmarketsUnauthLimit = 50

	orderFailed             = "Failed"
	orderPartiallyCancelled = "Partially Cancelled"
	orderCancelled          = "Cancelled"
	orderFullyMatched       = "Fully Matched"
	orderPartiallyMatched   = "Partially Matched"
	orderPlaced             = "Placed"
	orderAccepted           = "Accepted"

	ask        = "ask"
	limit      = "Limit"
	market     = "Market"
	stopLimit  = "Stop Limit"
	stop       = "Stop"
	takeProfit = "Take Profit"

	subscribe   = "subscribe"
	fundChange  = "fundChange"
	orderChange = "orderChange"
	heartbeat   = "heartbeat"
	tick        = "tick"
	wsOB        = "orderbook"
	trade       = "trade"
)

// tempOrderbook stores orderbook data
type tempOrderbook struct {
	MarketID   string      `json:"marketId"`
	SnapshotID int64       `json:"snapshotId"`
	Asks       [][2]string `json:"asks"`
	Bids       [][2]string `json:"bids"`
}

// // OBData stores orderbook data
// type OBData struct {
// 	Price  float64
// 	Volume float64
// }

// // Orderbook holds current orderbook information returned from the exchange
// type Orderbook struct {
// 	MarketID   string
// 	SnapshotID int64
// 	Asks       []OBData
// 	Bids       []OBData
// }

// MarketCandle stores candle data for a given pair
type MarketCandle struct {
	Time   time.Time
	Open   float64
	Close  float64
	Low    float64
	High   float64
	Volume float64
}

// TimeResp stores server time
type TimeResp struct {
	Time time.Time `json:"timestamp"`
}

// TradingFee 30 day trade volume
type TradingFee struct {
	Success        bool    `json:"success"`
	ErrorCode      int     `json:"errorCode"`
	ErrorMessage   string  `json:"errorMessage"`
	TradingFeeRate float64 `json:"tradingfeerate"`
	Volume30Day    float64 `json:"volume30day"`
}

// OrderToGo holds order information to be sent to the exchange
type OrderToGo struct {
	Currency        string `json:"currency"`
	Instrument      string `json:"instrument"`
	Price           int64  `json:"price"`
	Volume          int64  `json:"volume"`
	OrderSide       string `json:"orderSide"`
	OrderType       string `json:"ordertype"`
	ClientRequestID string `json:"clientRequestId"`
}

// Order holds order information
// type Order struct {
// 	ID              int64           `json:"id"`
// 	Currency        string          `json:"currency"`
// 	Instrument      string          `json:"instrument"`
// 	OrderSide       string          `json:"orderSide"`
// 	OrderType       string          `json:"ordertype"`
// 	CreationTime    time.Time       `json:"creationTime"`
// 	Status          string          `json:"status"`
// 	ErrorMessage    string          `json:"errorMessage"`
// 	Price           float64         `json:"price"`
// 	Volume          float64         `json:"volume"`
// 	OpenVolume      float64         `json:"openVolume"`
// 	ClientRequestID string          `json:"clientRequestId"`
// 	Trades          []TradeResponse `json:"trades"`
// }

// TradeResponse holds trade information
type TradeResponse struct {
	ID           int64     `json:"id"`
	CreationTime time.Time `json:"creationTime"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	Volume       float64   `json:"volume"`
	Fee          float64   `json:"fee"`
}

// AccountData stores account data
type AccountData struct {
	AssetName string  `json:"assetName"`
	Balance   float64 `json:"balance,string"`
	Available float64 `json:"available,string"`
	Locked    float64 `json:"locked,string"`
}

// CancelOrderResp stores data for cancelled orders
// type CancelOrderResp struct {
// 	OrderID       string `json:"orderId"`
// 	ClientOrderID string `json:"clientOrderId"`
// }

// WithdrawalFeeData stores data for fees
type WithdrawalFeeData struct {
	AssetName string  `json:"assetName"`
	Fee       float64 `json:"fee,string"`
}

// AssetData stores data for given asset
type AssetData struct {
	AssetName           string  `json:"assetName"`
	MinDepositAmount    float64 `json:"minDepositAmount,string"`
	MaxDepositAmount    float64 `json:"maxDepositAmount,string"`
	DepositDecimals     float64 `json:"depositDecimals,string"`
	MinWithdrawalAmount float64 `json:"minWithdrawalAmount,string"`
	MaxWithdrawalAmount float64 `json:"maxWithdrawalAmount,string"`
	WithdrawalDecimals  float64 `json:"withdrawalDecimals,string"`
	WithdrawalFee       float64 `json:"withdrawalFee,string"`
	DepositFee          float64 `json:"depositFee,string"`
}

// TransactionData stores data from past transactions
type TransactionData struct {
	ID           string    `json:"id"`
	CreationTime time.Time `json:"creationTime"`
	Description  string    `json:"description"`
	AssetName    string    `json:"assetName"`
	Amount       float64   `json:"amount,string"`
	Balance      float64   `json:"balance,string"`
	FeeType      string    `json:"type"`
	RecordType   string    `json:"recordType"`
	ReferrenceID string    `json:"referrenceId"`
}

// CreateReportResp stores data for created report
type CreateReportResp struct {
	ReportID string `json:"reportId"`
}

// ReportData gets data for a created report
type ReportData struct {
	ID           string    `json:"id"`
	ContentURL   string    `json:"contentUrl"`
	CreationTime time.Time `json:"creationTime"`
	ReportType   string    `json:"reportType"`
	Status       string    `json:"status"`
	Format       string    `json:"format"`
}

// UnprocessedBatchResp stores data for unprocessed response
type UnprocessedBatchResp struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

// BatchTradeResponse stores the trades from batchtrades
type BatchTradeResponse struct {
	Orders              []BatchPlaceData       `json:"orders"`
	UnprocessedRequests []UnprocessedBatchResp `json:"unprocessedRequests"`
}

// BatchCancelResponse stores the cancellation details from batch cancels
type BatchCancelResponse struct {
	CancelOrders        []CancelOrderResp      `json:"cancelOrders"`
	UnprocessedRequests []UnprocessedBatchResp `json:"unprocessedRequests"`
}

// PlaceOrderMethod stores data for place request
type PlaceOrderMethod struct {
	PlaceOrder PlaceBatch `json:"placeOrder,omitempty"`
}

// TradingFeeData stores trading fee data
type TradingFeeData struct {
	MakerFeeRate float64 `json:"makerFeeRate,string"`
	TakerFeeRate float64 `json:"takerFeeRate,string"`
	MarketID     string  `json:"marketId"`
}

// TradingFeeResponse stores trading fee data
type TradingFeeResponse struct {
	MonthlyVolume float64          `json:"volume30Day,string"`
	FeeByMarkets  []TradingFeeData `json:"FeeByMarkets"`
}
