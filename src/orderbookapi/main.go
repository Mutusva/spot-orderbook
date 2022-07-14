package main

import (
	"context"
	"flag"
	"fmt"
	ob "github.com/muzykantov/orderbook"
	"log"
	"net/http"
	"os"
	"os/signal"
	"spotob/src/orderbookapi/handlers"
	"time"
)

func main() {

	port := flag.String("port", "8080", "the port for the server")
	flag.Parse()
	o := ob.NewOrderBook()
	app := handlers.App{}
	app.Initialize(o)
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
	s.Shutdown(tc)
}
