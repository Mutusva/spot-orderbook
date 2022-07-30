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
	"spotob/src/env"
	"spotob/src/orderbookapi/handlers"
	rc "spotob/src/redis"
	"time"
)

var obKey = "orderbook"

func main() {

	port := flag.String("server_port", "8080", "the port for the server")
	environment := flag.String("env", "local", "redis host")
	channel := flag.String("redis_channel", "orderbook", "redis order book channel")
	flag.Parse()

	appConfig := env.GetRedisConfig(*environment)
	rdb := redis.NewClient(&redis.Options{
		Addr:     appConfig.RedisHost + ":" + appConfig.RedisPort,
		Password: appConfig.RedisPassword,
		DB:       0, // use default DB
	})

	ctx := context.Background()
	redisClient := rc.NewOpsClient(rdb, *channel)

	odb := getSavedOrderBook(ctx, redisClient, obKey)

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
	saveAppState(ctx, odb, redisClient, obKey)
	s.Shutdown(tc)
}

func getSavedOrderBook(ctx context.Context, rdb *rc.OpsClient, key string) *ob.OrderBook {
	orderBk, err := rdb.GetSavedOrderBook(ctx, key)
	if err != nil {
		log.Println("cannot create order book from redis " + err.Error())
		return ob.NewOrderBook()
	}

	if len(orderBk) == 0 {
		return ob.NewOrderBook()
	}
	var o ob.OrderBook
	err = json.Unmarshal([]byte(orderBk), &o)
	if err != nil {
		log.Println("Error creating order book from file")
		return ob.NewOrderBook()
	}
	return &o
}

func saveAppState(ctx context.Context, orderBook *ob.OrderBook, rdb *rc.OpsClient, key string) {
	obStr, err := orderBook.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	err = rdb.SaveOrderBook(ctx, key, string(obStr))
	if err != nil {
		log.Println(err)
	}
}
