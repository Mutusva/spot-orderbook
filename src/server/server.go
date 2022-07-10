package server

import (
	"context"
	ob "github.com/muzykantov/orderbook"
	"github.com/shopspring/decimal"
	"spotob/src/orderbookpb"
	obs "spotob/src/orderbookservice"
)

type OrderbookServer struct {
	Orderbook *ob.OrderBook
}

func (O *OrderbookServer) ProcessLimitOrder(ctx context.Context, req *orderbookpb.LimitOrderRequest) (*orderbookpb.LimitOrderResponse, error)  {
	quantity, err := decimal.NewFromString(req.Quantity)
	if err != nil {
		return nil, err
	}

	price, err := decimal.NewFromString(req.Price)
	if err != nil {
		return nil, err
	}

	done, partial, partialQuantityProcessed, err := obs.ProcessLimitOrder(O.Orderbook, ob.Side(req.Side), req.OrderId, quantity, price)
	if err != nil {
		return nil, err
	}

	return &orderbookpb.LimitOrderResponse{
		Orders: orderbookpb.ConvertOrdersToOrderpb(done),
		Partial: orderbookpb.ConvertOrderToOrderpb(partial),
		PartialQuantityProcessed: partialQuantityProcessed.String(),
	}, nil

}

func (O *OrderbookServer) ProcessMarketOrder(ctx context.Context, req *orderbookpb.MarketOrderRequest) (*orderbookpb.MarketOrderResponse, error) {
	quantity, err := decimal.NewFromString(req.Quantity)
	if err != nil {
		return nil, err
	}

	done, partial, partialQuantityProcessed, quantityLeft , err := obs.ProcessMarketOrder(O.Orderbook, ob.Side(req.Side), quantity)

	if err != nil {
		return nil, err
	}

	return &orderbookpb.MarketOrderResponse{
		Done: orderbookpb.ConvertOrdersToOrderpb(done),
		Partial: orderbookpb.ConvertOrderToOrderpb(partial),
		PartialQuantityProcessed: partialQuantityProcessed.String(),
		QuantityLeft: quantityLeft.String(),
	}, nil
}

func (O *OrderbookServer) CancelOrder(ctx context.Context, req *orderbookpb.CancelOrderRequest) (*orderbookpb.CancelOrderResponse, error) {
	order := obs.CancelOrder(O.Orderbook, req.Id)
	return &orderbookpb.CancelOrderResponse{
		Order: orderbookpb.ConvertOrderToOrderpb(order),
	}, nil
}

func (O *OrderbookServer) Depth(context.Context, *orderbookpb.Empty) (*orderbookpb.DepthResponse, error) {
	asks, bids := obs.Depth(O.Orderbook)
	return &orderbookpb.DepthResponse{
		Asks: orderbookpb.ToPriceLevelpbs(asks),
		Bids: orderbookpb.ToPriceLevelpbs(bids),
	}, nil
}
