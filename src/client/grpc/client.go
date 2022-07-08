package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "spotob/src/orderbookpb"
)

type OrderBookgRPCClient struct {
	RpcConn   *grpc.ClientConn
	ObClient pb.OrderBookgRPCServiceClient
}

func (client *OrderBookgRPCClient) ProcessLimitOrder(ctx context.Context, req *pb.LimitOrderRequest, opts ...grpc.CallOption) (*pb.LimitOrderResponse, error){
	res, err := client.ObClient.ProcessLimitOrder(ctx, req)
	if err!=nil {
		return nil, err
	}
	return res, nil
}

func (client *OrderBookgRPCClient) ProcessMarketOrder(ctx context.Context, req *pb.MarketOrderRequest, opts ...grpc.CallOption) (*pb.MarketOrderResponse, error){
	res, err := client.ObClient.ProcessMarketOrder(ctx, req)
	if err!=nil {
		return nil, err
	}
	return res, nil
}

func (client *OrderBookgRPCClient) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest, opts ...grpc.CallOption) (*pb.CancelOrderResponse, error) {
	res, err := client.ObClient.CancelOrder(ctx, req)
	if err!=nil {
		return nil, err
	}
	return res, nil
}

func (client *OrderBookgRPCClient) Depth(ctx context.Context, req *pb.Empty, opts ...grpc.CallOption) (*pb.DepthResponse, error){
	res, err := client.ObClient.Depth(ctx, req)
	if err!=nil {
		return nil, err
	}
	return res, nil
}

func NewOrderBookClient(address string) *OrderBookgRPCClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	obc := pb.NewOrderBookgRPCServiceClient(conn)
	 return &OrderBookgRPCClient{
		 RpcConn: conn,
		 ObClient: obc,
	 }
}
