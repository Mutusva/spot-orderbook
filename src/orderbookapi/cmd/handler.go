package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	ob "github.com/muzykantov/orderbook"
	"github.com/shopspring/decimal"
	"net/http"
	"spotob/src/models"
	obs "spotob/src/orderbookservice"
)

type App struct {
	OrderBook *ob.OrderBook
	Router *mux.Router
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/processLimitOrder", a.ProcessLimitOrder).Methods("POST")
	a.Router.HandleFunc("/processMarketOrder", a.ProcessMarketOrder).Methods("POST")
	a.Router.HandleFunc("/cancelOrder/{id}", a.CancelOrder)
	a.Router.HandleFunc("/depth", a.Depth)
}

func (a *App) Initialize(orderBook *ob.OrderBook) {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
	a.OrderBook = orderBook
}

func (a *App) ProcessLimitOrder(w http.ResponseWriter, r *http.Request) {
	var limitOrder models.LimitOrder
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&limitOrder); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	lr, err := validateLimitOrderRequest(limitOrder)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	done, partial, partialQuantityProcessed, err := obs.ProcessLimitOrder(a.OrderBook, lr.Side, lr.OrderId, lr.Quantity, lr.Price)
	respondWithJSON(w, http.StatusOK, models.LimitOrderResponse{
		Done:                     done,
		Partial:                  partial,
		PartialQuantityProcessed: partialQuantityProcessed,
	})
}

func (a *App) ProcessMarketOrder(w http.ResponseWriter, r *http.Request) {

}

func (a *App) CancelOrder(w http.ResponseWriter, r *http.Request) {

}

func (a *App) Depth(w http.ResponseWriter, r *http.Request) {

}

func validateLimitOrderRequest(lo models.LimitOrder) (*models.LimitOrderType, error) {
	if lo.Side < 1 || lo.Side > 2 {
		return nil, errors.New("invalid side. sell(0), sell(1)")
	}

	if lo.OrderId == "" {
		return nil, errors.New("invalid order id")
	}

	if lo.Quantity <= 0 {
		return nil, errors.New("invalid quantity, quantity should be not zero")
	}

	if lo.Price <= 0 {
		return nil, errors.New("price should be greater than zero")
	}

	return &models.LimitOrderType{
		Side:     ob.Side(lo.Side),
		OrderId:  lo.OrderId,
		Quantity: decimal.NewFromFloat(lo.Quantity),
		Price:    decimal.NewFromFloat(lo.Price),
	}, nil
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}