package main

import (
	"context"
	"encoding/json"
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
	orderBkLoc := flag.String("ob_fpath", "orderbook.json", "order book persistance location")

	flag.Parse()
	rdb := redis.NewClient(&redis.Options{
		Addr:     *redisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	redisClient := rc.NewOpsClient(rdb, *channel)

	odb := readFile(*orderBkLoc)

	// o := ob.NewOrderBook()
	app := handlers.App{}
	app.Initialize(ctx, odb, redisClient)
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
	saveAppState(odb, *orderBkLoc)
	s.Shutdown(tc)
}

func readFile(s string) *ob.OrderBook {
	b, err := os.ReadFile(s)
	if err != nil {
		log.Println("Error reading file or file does not exits")
	}

	if len(b) == 0 {
		return ob.NewOrderBook()
	}
	var o ob.OrderBook
	err = json.Unmarshal(b, &o)
	if err != nil {
		log.Println("Error creating order book from file")
		return ob.NewOrderBook()
	}

	return &o
}

func saveAppState(orderBook *ob.OrderBook, loc string) {
	obStr, err := orderBook.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	f, err := os.Create(loc)
	if err != nil {
		log.Println(err)
	}
	_, err = f.Write(obStr)
	if err != nil {
		log.Println(err)
	}
}
