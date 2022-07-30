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

// swagger:model
// Limit order request
type LimitOrder struct {
	// enum 0 for sell and 1 for buy
	// Side
	// required: true
	// example: 0
	Side int32 `json:"side"`

	// OrderId
	// required: true
	// example: order-1
	OrderId string `json:"order_id"`

	// quantity
	// required: true
	// example: 4.0
	Quantity Decimal `json:"quantity"`

	// price
	// required: true
	// example: 4.5
	Price Decimal `json:"price"`
}

/*

// swagger:model
type ReqLimitOrderBody struct {
	//  in: body
	LimitOrder *LimitOrder
}

*/

// An Error response
// swagger:response
type ErrorResponse struct {
	// The error message
	// name code
	Code int `json:"code"`

	// errors
	Errors map[string]string `json:"errors"`
}

// swagger: response ErrorResponse
type ValidationError struct {
	// - name: error
	//  in: body
	//  description: request body for limit order
	Body *ErrorResponse
}

// swagger:model
// represents an order whether a buy or sell
type Order struct {
	// enum 0 for sell and 1 for buy
	// Side
	// required: true
	// example: 1
	Side int32 `json:"side"`

	// order id
	// required: true
	// example: order-1
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

// swagger:model
// in:body
type LimitOrderResponse struct {
	// Orders that are done
	Done []*ob.Order `json:"done"`

	// Partially done orders
	Partial *ob.Order `json:"partial"`

	// Quantity of orders that have been partially completed
	PartialQuantityProcessed Decimal `json:"partialQuantityProcessed"`
}

// swagger:model
// in:body
type MarketOrderResponse struct {
	// The Market Order response
	// Orders that are done
	Done []*ob.Order `json:"done"`

	// Partially done orders
	Partial *ob.Order `json:"partial"`

	// Quantity of orders that have been partially completed
	PartialQuantityProcessed Decimal `json:"partialQuantityProcessed"`

	// Quantity of orders left
	QuantityLeft Decimal `json:"quantityLeft"`
}

// swagger:model
// Market order request
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

// swagger: model Order
type OrderResponse struct {
	// Order response
	// in: body
	Order *Order
}

type LimitOrderType struct {
	Side     ob.Side
	OrderId  string
	Quantity Decimal
	Price    Decimal
}

// swagger:model
// in:body
// depth response
type OrderBookDepth struct {
	// bids
	// swagger: model
	Bids []*ob.PriceLevel `json:"bids"`
	// asks
	// swagger: model
	Asks []*ob.PriceLevel `json:"asks"`
}

type RedisOrderMessageType struct {
	Type MessageType `json:"message_type"`
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
