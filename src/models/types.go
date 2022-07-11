package models

import (
	ob "github.com/muzykantov/orderbook"
	. "github.com/shopspring/decimal"
	"time"
)

type LimitOrder struct {
	Side     int32   `json:"side"`
	OrderId  string  `json:"order_id"`
	Quantity Decimal `json:"quantity"`
	Price    Decimal `json:"price"`
}

type Order struct {
	Side      int32     `json:"side"`
	Id        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Quantity  Decimal   `json:"quantity"`
	Price     Decimal   `json:"price"`
}

type LimitOrderResponse struct {
	Done                     []*ob.Order `json:"done"`
	Partial                  *ob.Order   `json:"partial"`
	PartialQuantityProcessed Decimal     `json:"partialQuantityProcessed"`
}

type MarketOrderResponse struct {
	Done                     []*ob.Order `json:"done"`
	Partial                  *ob.Order   `json:"partial"`
	PartialQuantityProcessed Decimal     `json:"partialQuantityProcessed"`
	QuantityLeft             Decimal     `json:"quantityLeft"`
}

type MarketOrderRequest struct {
	Side     int32   `json:"side"`
	Quantity Decimal `json:"quantity"`
}

type LimitOrderType struct {
	Side     ob.Side
	OrderId  string
	Quantity Decimal
	Price    Decimal
}

type OrderBookDepth struct {
	Bids []*ob.PriceLevel `json:"bids"`
	Asks []*ob.PriceLevel `json:"asks"`
}
