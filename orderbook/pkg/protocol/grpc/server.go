package grpc

import (
	"context"
	ob "github.com/muzykantov/orderbook"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"spotob/orderbook/orderbookpb"
	"spotob/orderbook/server"
)

func NewOrderbookRPCServer(orderBook *ob.OrderBook) *grpc.Server {
	gsrv := grpc.NewServer()
	srv := server.OrderbookServer{
		Orderbook: orderBook,
	}
	orderbookpb.RegisterOrderBookgRPCServiceServer(gsrv, &srv)
	return gsrv
}

func RunServer(ctx context.Context, port string, orderBook *ob.OrderBook) error {
	log.Println("Starting listening on port " + port)
	listen, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		return err
	}

	log.Printf("Listening on %s", port)
	// register service
	srv := NewOrderbookRPCServer(orderBook)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			srv.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	return srv.Serve(listen)
}
