package orderbookpb

import (
	ob "github.com/muzykantov/orderbook"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertOrderToOrderpb(order *ob.Order) *Order {

	if order == nil {
		return &Order{}
	}

	time := timestamppb.New(order.Time())
	return &Order{
		Side:                 int32(order.Side()),
		Id:                   order.ID(),
		Timestamp:            time,
		Quantity:             order.Quantity().String(),
		Price:                order.Price().String(),
	}
}

func ConvertOrdersToOrderpb(orders []*ob.Order) []*Order {
	var orderList []*Order
	for _, order := range orders {
		orderList = append(orderList, ConvertOrderToOrderpb(order))
	}

	return orderList
}

func ToPriceLevelpb(pl *ob.PriceLevel) *PriceLevel {
	return &PriceLevel{
		Price: pl.Price.String(),
		Quantity: pl.Quantity.String(),
	}
}

func ToPriceLevelpbs(pls []*ob.PriceLevel) []*PriceLevel {
	var priceLevels []*PriceLevel
	for _,pl := range pls {
		priceLevels = append(priceLevels, ToPriceLevelpb(pl))
	}
	return priceLevels
}