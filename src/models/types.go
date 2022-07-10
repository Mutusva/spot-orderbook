package models

import (
	ob "github.com/muzykantov/orderbook"
	"github.com/shopspring/decimal"
	"time"
)

type LimitOrder struct {
	Side int32 `json: "side"`
	OrderId  string `json: "order_id"`
	Quantity float64 `json: "quantity"`
	Price float64 `json: "price"`
}

type Order struct {
	Side int32 `json: "side"`
	Id string  `json: "id"`
	Timestamp time.Time `json: "timestamp"`
	Quantity float32 `json: "quantity"`
	Price float32  `json: "price"`
}

type LimitOrderResponse struct {
	Done []*ob.Order `json: "done"`
	Partial *ob.Order `json: "partial"`
	PartialQuantityProcessed decimal.Decimal `json: "partialQuantityProcessed"`
}

type LimitOrderType struct {
	Side  ob.Side
	OrderId string
	Quantity decimal.Decimal
	Price decimal.Decimal
}