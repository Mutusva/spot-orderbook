package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	ob "github.com/muzykantov/orderbook"
	"log"
	"net/http"
	"os"
	"os/signal"
	"spotob/src/orderbookapi/handlers"
	rc "spotob/src/redis"
	"time"
)

func main() {

	port := flag.String("server_port", "8080", "the port for the server")
	redisHost := flag.String("redis_host", "localhost:6379", "redis host")
	channel := flag.String("redis_channel", "orderbook", "redis order book channel")

	flag.Parse()
	rdb := redis.NewClient(&redis.Options{
		Addr:     *redisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	redisClient := rc.NewOpsClient(rdb, *channel)
	o := ob.NewOrderBook()
	app := handlers.App{}
	app.Initialize(ctx, o, redisClient)
	fmt.Printf("Listening on port %s\n", *port)

	s := &http.Server{
		Addr:         ":" + *port,
		Handler:      app.Router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Second,
	}

	// start redis connection
	go app.ProcessOrders(ctx)

	// graceful shut down - allows cleaning up of resources.
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
