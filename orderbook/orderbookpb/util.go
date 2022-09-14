package orderbookpb

import (
	"errors"
	ob "github.com/muzykantov/orderbook"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertOrderToOrderpb(order *ob.Order) *Order {

	if order == nil {
		return &Order{}
	}

	time := timestamppb.New(order.Time())
	return &Order{
		Side:      int32(order.Side()),
		Id:        order.ID(),
		Timestamp: time,
		Quantity:  order.Quantity().String(),
		Price:     order.Price().String(),
	}
}

func ConvertOrdersToOrderpb(orders []*ob.Order) []*Order {
	var orderList []*Order
	for _, order := range orders {
		orderList = append(orderList, ConvertOrderToOrderpb(order))
	}

	return orderList
}

func OrderpbToOrder(o *Order) (*ob.Order, error) {
	if o == nil || o.Id == "" {
		return nil, nil
	}
	q, err := decimal.NewFromString(o.Quantity)
	if err != nil {
		return nil, err
	}

	p, err := decimal.NewFromString(o.Price)
	if err != nil {
		return nil, err
	}

	ts := o.Timestamp.AsTime()
	side := ob.Side(int(o.Side))
	return ob.NewOrder(o.Id, side, q, p, ts), nil
}

func OrderspbToOrders(ods []*Order) ([]*ob.Order, error) {
	var orders []*ob.Order
	for _, od := range ods {
		order, err := OrderpbToOrder(od)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func ToPriceLevelpb(pl *ob.PriceLevel) *PriceLevel {
	return &PriceLevel{
		Price:    pl.Price.String(),
		Quantity: pl.Quantity.String(),
	}
}

func ToPriceLevelpbs(pls []*ob.PriceLevel) []*PriceLevel {
	var priceLevels []*PriceLevel
	for _, pl := range pls {
		priceLevels = append(priceLevels, ToPriceLevelpb(pl))
	}
	return priceLevels
}

func FromPriceLevelpb(pl *PriceLevel) (*ob.PriceLevel, error) {
	if pl == nil {
		return nil, errors.New("cannot convert an empty price level")
	}

	p, err := decimal.NewFromString(pl.Price)
	if err != nil {
		return nil, err
	}

	q, err := decimal.NewFromString(pl.Quantity)
	if err != nil {
		return nil, err
	}

	return &ob.PriceLevel{
		Price:    p,
		Quantity: q,
	}, nil
}

func FromPriceLevelpbs(pls []*PriceLevel) ([]*ob.PriceLevel, error) {
	var priceLevels []*ob.PriceLevel
	for _, pl := range pls {

		obpl, err := FromPriceLevelpb(pl)
		if err != nil {
			return nil, err
		}
		priceLevels = append(priceLevels, obpl)
	}
	return priceLevels, nil
}
