package grpc

import (
	"context"
	"github.com/muzykantov/orderbook"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"log"
	"spotob/orderbook/models"
	pb "spotob/orderbook/orderbookpb"
)

type OrderBookgRPCClient struct {
	Ob       *orderbook.OrderBook
	RpcConn  *grpc.ClientConn
	ObClient pb.OrderBookgRPCServiceClient
}

func (c *OrderBookgRPCClient) ProcessLimitOrder(ctx context.Context, req models.LimitOrder) (*models.LimitOrderResponse, error) {
	//Todo : Add some validation
	res, err := c.ObClient.ProcessLimitOrder(ctx, &pb.LimitOrderRequest{
		Side:     req.Side,
		OrderId:  req.OrderId,
		Quantity: req.Quantity.String(),
		Price:    req.Price.String(),
	})

	if err != nil {
		return nil, err
	}
	orders, err := pb.OrderspbToOrders(res.Orders)
	if err != nil {
		return nil, err
	}

	partial, err := pb.OrderpbToOrder(res.Partial)
	if err != nil {
		return nil, err
	}

	pqp, err := decimal.NewFromString(res.PartialQuantityProcessed)
	if err != nil {
		return nil, err
	}

	return &models.LimitOrderResponse{
		Done:                     orders,
		Partial:                  partial,
		PartialQuantityProcessed: pqp,
	}, nil
}

func (c *OrderBookgRPCClient) ProcessMarketOrder(ctx context.Context, req models.MarketOrderRequest) (*models.MarketOrderResponse, error) {
	res, err := c.ObClient.ProcessMarketOrder(ctx, &pb.MarketOrderRequest{
		Side:     req.Side,
		Quantity: req.Quantity.String(),
	})
	if err != nil {
		return nil, err
	}

	orders, err := pb.OrderspbToOrders(res.Done)
	if err != nil {
		return nil, err
	}

	partial, err := pb.OrderpbToOrder(res.Partial)
	if err != nil {
		return nil, err
	}

	pqp, err := decimal.NewFromString(res.PartialQuantityProcessed)
	if err != nil {
		return nil, err
	}

	ql, err := decimal.NewFromString(res.QuantityLeft)
	if err != nil {
		return nil, err
	}

	return &models.MarketOrderResponse{
		Done:                     orders,
		Partial:                  partial,
		PartialQuantityProcessed: pqp,
		QuantityLeft:             ql,
	}, nil
}

func (c *OrderBookgRPCClient) CancelOrder(ctx context.Context, Id string) (*models.Order, error) {
	res, err := c.ObClient.CancelOrder(ctx, &pb.CancelOrderRequest{
		Id: Id,
	})

	order, err := pb.OrderpbToOrder(res.Order)
	if err != nil {
		return nil, err
	}

	return &models.Order{
		Side:      int32(order.Side()),
		Id:        order.ID(),
		Timestamp: order.Time(),
		Quantity:  order.Quantity(),
		Price:     order.Price(),
	}, nil
}

func (c *OrderBookgRPCClient) Depth(ctx context.Context) (*models.OrderBookDepth, error) {
	res, err := c.ObClient.Depth(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}

	bids, err := pb.FromPriceLevelpbs(res.Bids)
	if err != nil {
		return nil, err
	}
	asks, err := pb.FromPriceLevelpbs(res.Asks)
	if err != nil {
		return nil, err
	}

	return &models.OrderBookDepth{
		Bids: bids,
		Asks: asks,
	}, nil
}

func NewOrderBookClient(connStri string, ob *orderbook.OrderBook) (*OrderBookgRPCClient, error) {
	conn, err := grpc.Dial(connStri, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}

	obc := pb.NewOrderBookgRPCServiceClient(conn)
	return &OrderBookgRPCClient{
		RpcConn:  conn,
		ObClient: obc,
		Ob:       ob,
	}, nil
}
