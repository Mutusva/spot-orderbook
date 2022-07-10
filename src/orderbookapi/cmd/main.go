package main

import (
	"github.com/gorilla/mux"
	ob "github.com/muzykantov/orderbook"
	"log"
	"net/http"
)



func main() {

	o := ob.NewOrderBook()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to spot orderbook"))
	})


	app := NewSpotApp()
	app.Initialize(o)

	log.Fatal(http.ListenAndServe(":8080",app.Router))
}
