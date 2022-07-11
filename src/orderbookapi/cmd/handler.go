package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	ob "github.com/muzykantov/orderbook"
	"net/http"
	"spotob/src/models"
	obs "spotob/src/orderbookservice"
)

type App struct {
	OrderBook *ob.OrderBook
	Router    *mux.Router
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
		respondWithError(w, http.StatusBadRequest, "Invalid request payload"+err.Error())
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
	var req models.MarketOrderRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()
	if err := validateMarketOrderRequest(req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	done, partial, partialQuantityProcessed, quantityLeft, err := obs.ProcessMarketOrder(a.OrderBook, ob.Side(req.Side), req.Quantity)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error processing request")
		return
	}

	respondWithJSON(w, http.StatusOK, models.MarketOrderResponse{
		Done:                     done,
		Partial:                  partial,
		PartialQuantityProcessed: partialQuantityProcessed,
		QuantityLeft:             quantityLeft,
	})

}

func (a *App) CancelOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["id"]
	if len(orderId) == 0 {
		respondWithError(w, http.StatusBadRequest, "id cannot be empty")
		return
	}

	order := obs.CancelOrder(a.OrderBook, orderId)
	respondWithJSON(w, http.StatusOK, order)
}

func (a *App) Depth(w http.ResponseWriter, _ *http.Request) {
	asks, bids := obs.Depth(a.OrderBook)
	depth := models.OrderBookDepth{
		Asks: asks,
		Bids: bids,
	}
	respondWithJSON(w, http.StatusOK, depth)
}

func validateLimitOrderRequest(lo models.LimitOrder) (*models.LimitOrderType, error) {
	if lo.Side < 0 || lo.Side > 1 {
		return nil, errors.New("invalid side. sell(0), sell(1)")
	}

	if lo.OrderId == "" {
		return nil, errors.New("invalid order id")
	}

	if lo.Quantity.IsZero() {
		return nil, errors.New("invalid quantity, quantity should be not zero")
	}

	if lo.Price.IsZero() {
		return nil, errors.New("price should be greater than zero")
	}

	return &models.LimitOrderType{
		Side:     ob.Side(lo.Side),
		OrderId:  lo.OrderId,
		Quantity: lo.Quantity,
		Price:    lo.Price,
	}, nil
}

func validateMarketOrderRequest(req models.MarketOrderRequest) error {
	if req.Side < 0 || req.Side > 1 {
		return errors.New("invalid side. sell(0), sell(1)")
	}

	if req.Quantity.IsZero() {
		return errors.New("invalid quantity, quantity should be not zero")
	}
	return nil
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
