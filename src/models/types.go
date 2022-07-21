package models

import (
	ob "github.com/muzykantov/orderbook"
	. "github.com/shopspring/decimal"
	"time"
)

type MessageType int

var (
	LimitOrderMessageType  MessageType = 1
	MarketOrderMessageType MessageType = 2
	CancelOrderMessageType MessageType = 3
)

//  swagger: model
// Limit order request
type LimitOrder struct {
	// enum 0 for sell and 1 for buy
	// in: int32
	Side int32 `json:"side"`

	// order id
	// in: string
	OrderId string `json:"order_id"`

	// quantity
	// in: Decimal
	Quantity Decimal `json:"quantity"`

	// price
	// in: Decimal
	Price Decimal `json:"price"`
}

// swagger:parameters order LimitOrderRequest
type ReqLimitOrderBody struct {
	// - name: body
	//  in: body
	//  description: request body for limit order
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/ReqLimitOrderBody"
	//  required: true
	Body LimitOrder `json:"body"`
}

// An Error response
// swagger: response ErrorResponse
type ErrorResponse struct {
	// name code
	// in: int
	Code int

	// errors
	// in: map[string]string
	Errors map[string]string
}

//  swagger: model Order
// represents an order whether a buy or sell
type Order struct {
	// enum 0 for sell and 1 for buy
	// required: true
	Side int32 `json:"side"`

	// order id
	// required: true
	Id string `json:"id"`

	// time when the order was created
	// required: true
	Timestamp time.Time `json:"timestamp"`

	// Quantity of the order
	// required: true
	Quantity Decimal `json:"quantity"`

	// Price for each order
	// required: true
	Price Decimal `json:"price"`
}

//  swagger: response LimitOrderResponse
//  Limit order response
type LimitOrderResponse struct {
	// Orders that are done
	Done []*ob.Order `json:"done"`

	// Partially done orders
	Partial *ob.Order `json:"partial"`

	// Quantity of orders that have been partially completed
	PartialQuantityProcessed Decimal `json:"partialQuantityProcessed"`
}

//  swagger: response MarketOrderResponse
//  Market order response
type MarketOrderResponse struct {
	// Orders that are done
	Done []*ob.Order `json:"done"`

	// Partially done orders
	Partial *ob.Order `json:"partial"`

	// Quantity of orders that have been partially completed
	PartialQuantityProcessed Decimal `json:"partialQuantityProcessed"`

	// Quantity of orders left
	QuantityLeft Decimal `json:"quantityLeft"`
}

// swagger:parameters order MarketOrderRequest
type ReqMarketOrderBody struct {
	// - name: body
	//  in: body
	//  description: request body for market order
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/ReqMarketOrderBody"
	//  required: true
	Body MarketOrderRequest `json:"body"`
}

//  swagger: model
//  Market order request
type MarketOrderRequest struct {

	// swagger: model
	// enum 0 for sell and 1 for buy
	// required: true
	Side int32 `json:"side"`

	// swagger: model
	// quantity
	// required: true
	Quantity Decimal `json:"quantity"`
}

type LimitOrderType struct {
	Side     ob.Side
	OrderId  string
	Quantity Decimal
	Price    Decimal
}

// swagger: response OrderBookDepth
// depth response
type OrderBookDepth struct {
	// bids
	Bids []*ob.PriceLevel `json:"bids"`
	// asks
	Asks []*ob.PriceLevel `json:"asks"`
}

type LimitOrderMessage struct {
	Message LimitOrder  `json:"message"`
	Type    MessageType `json:"message_type"`
}

type MarketOrderMessage struct {
	Message MarketOrderRequest `json:"message"`
	Type    MessageType        `json:"message_type"`
}

type CancelOrderMessage struct {
	Message string      `json:"message"`
	Type    MessageType `json:"message_type"`
}
