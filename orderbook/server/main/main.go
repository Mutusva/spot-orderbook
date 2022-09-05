/*

Delete this file not need. use pkg/protocol/grpc/server.go to initialise the server.
*/

package main

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
	listen, err := net.Listen("tcp", ":"+port)
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

func main() {
	log.Println("Starting listening on port 8080")
	port := ":8080"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", port)
	orderbook := ob.NewOrderBook()
	srv := NewOrderbookRPCServer(orderbook)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
