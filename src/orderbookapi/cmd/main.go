package main

import (
	"flag"
	"fmt"
	ob "github.com/muzykantov/orderbook"
	"log"
	"net/http"
)

func main() {

	port := flag.String("port", "8080", "the port for the server")
	flag.Parse()
	o := ob.NewOrderBook()
	app := App{}
	app.Initialize(o)
	log.Fatal(http.ListenAndServe(":"+*port, app.Router))
	fmt.Printf("Listening on port %s\n", *port)
}
