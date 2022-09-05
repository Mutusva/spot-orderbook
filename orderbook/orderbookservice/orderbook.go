package orderbookservice

import (
	ob "github.com/muzykantov/orderbook"
	"github.com/shopspring/decimal"
)

func ProcessLimitOrder(orderBook *ob.OrderBook, side ob.Side, orderID string, quantity, price decimal.Decimal) ([]*ob.Order, *ob.Order, decimal.Decimal, error) {
	done, partial, partialQuantityProcessed, err := orderBook.ProcessLimitOrder(side, orderID, quantity, price)
	if err != nil {
		return nil, nil, decimal.Decimal{}, err
	}

	return done, partial, partialQuantityProcessed, nil
}

func ProcessMarketOrder(orderBook *ob.OrderBook, side ob.Side, quantity decimal.Decimal) ([]*ob.Order, *ob.Order, decimal.Decimal, decimal.Decimal, error) {
	done, partial,partialQuantityProcessed, quantityLeft, err := orderBook.ProcessMarketOrder(side, quantity)
	if err != nil {
		return nil, nil, decimal.Decimal{},decimal.Decimal{}, err
	}

	return done, partial, partialQuantityProcessed, quantityLeft, nil
}

func CancelOrder(orderBook *ob.OrderBook, orderID string) *ob.Order {
	order := orderBook.CancelOrder(orderID)
	return order
}

func Depth(orderBook *ob.OrderBook) ([]*ob.PriceLevel, []*ob.PriceLevel) {
	asks, bids := orderBook.Depth()
	return asks, bids
}
