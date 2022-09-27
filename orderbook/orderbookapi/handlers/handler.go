// Package classification Spot OrderBook API
//
// Documentation for Spot OrderBook API
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//  Consumes:
//  - application/json
//  Produces:
//  - application/json
//  swagger:meta
package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	ob "github.com/muzykantov/orderbook"
	"log"
	"net/http"
	"spotob/orderbook/client/grpc"
	"spotob/orderbook/env"
	"spotob/orderbook/models"
	rc "spotob/orderbook/redis"
)

type App struct {
	OrderBook   *ob.OrderBook
	Router      *mux.Router
	RedisClient *rc.OpsClient
	Client      *grpc.OrderBookgRPCClient
	ctx         context.Context
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/processLimitOrder", a.ProcessLimitOrder).Methods("POST")
	a.Router.HandleFunc("/processMarketOrder", a.ProcessMarketOrder).Methods("POST")
	a.Router.HandleFunc("/cancelOrder/{id}", a.CancelOrder)
	a.Router.HandleFunc("/depth", a.Depth)
	a.Router.HandleFunc("/health", a.Health)
	/// mux.CORSMethodMiddleware(a.Router)
	// documentation for developers
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)

	// documentation for share
	// opts := middleware.RedocOpts{SpecURL: "./swagger.yaml"}
	// sh := middleware.Redoc(opts, nil)

	swaggerPath := "./swagger/"
	if env.IsDev() {
		swaggerPath = "./orderbook/orderbookapi/swagger/"
	}

	a.Router.Handle("/docs", sh)
	a.Router.Handle("/swagger.yaml", http.FileServer(http.Dir(swaggerPath)))
}

func (a *App) Initialize(ctx context.Context, orderBook *ob.OrderBook, rc *rc.OpsClient, serverPort string) {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
	a.OrderBook = orderBook
	a.RedisClient = rc
	a.ctx = ctx
	client, err := grpc.NewOrderBookClient(":"+serverPort, a.OrderBook)
	if err != nil {
		log.Fatal(err)
	}

	a.Client = client
}

//  swagger:route POST /processLimitOrder processLimitOrder
//  Create a new limit order
//  Responses:
//    401: ErrorResponse
//    404: ErrorResponse
//    500: ErrorResponse
//    200: body:LimitOrderResponse
//   Parameters:
//     + name: LimitOrder
//       in: body
//       required: true
//       type: LimitOrder

// ProcessLimitOrder create a limit order for processing
func (a *App) ProcessLimitOrder(w http.ResponseWriter, r *http.Request) {
	var limitOrder models.LimitOrder
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&limitOrder); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload"+err.Error())
		return
	}
	defer r.Body.Close()

	_, err := validateLimitOrderRequest(&limitOrder)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	res, err := a.Client.ProcessLimitOrder(r.Context(), limitOrder)
	// done, partial, partialQuantityProcessed, err := obs.ProcessLimitOrder(a.OrderBook, lr.Side, lr.OrderId, lr.Quantity, lr.Price)
	if err != nil {
		// log this msg
		respondWithError(w, http.StatusBadRequest, "error processing limit order")
		return
	}
	respondWithJSON(w, http.StatusOK, res)

}

func (a *App) Health(w http.ResponseWriter, r *http.Request) {
	type data struct{ msg string }
	respondWithJSON(w, http.StatusOK, data{msg: "health"})
}

//  swagger:route POST /processMarketOrder processMarketOrder
//  Create a new market order for processing
//  Responses:
//    401: ErrorResponse
//    200: body:MarketOrderResponse
//   Parameters:
//     + name: MarketOrderRequest
//       in: body
//       required: true
//       type: MarketOrderRequest

// ProcessMarketOrder create a market order for processing
func (a *App) ProcessMarketOrder(w http.ResponseWriter, r *http.Request) {
	var req models.MarketOrderRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()
	if err := validateMarketOrderRequest(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	mor, err := a.Client.ProcessMarketOrder(r.Context(), req)
	if err != nil {
		// log this msg
		respondWithError(w, http.StatusBadRequest, "Error processing market order")
		return
	}
	respondWithJSON(w, http.StatusOK, mor)

}

// swagger:route GET /cancelOrder/{id} cancelOrder
//   Responses:
//     200: body:Order
//     401: ErrorResponse
//   Parameters:
//     + name: order id
//       in: path
//       description: id of the order to cancel
//       required: true
//       type: string
//       format: string

// CancelOrder cancel a specific order by id
func (a *App) CancelOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["id"]
	if len(orderId) == 0 {
		respondWithError(w, http.StatusBadRequest, "id cannot be empty")
		return
	}

	order, err := a.Client.CancelOrder(r.Context(), orderId)
	if err != nil {
		// log this msg
		respondWithError(w, http.StatusBadRequest, "Error cancelling order")
		return
	}

	respondWithJSON(w, http.StatusOK, order)
}

// swagger:route GET /depth depth
// Responses:
//   200: OrderBookDepth
//   401: ErrorResponse
func (a *App) Depth(w http.ResponseWriter, r *http.Request) {
	depthRes, err := a.Client.Depth(r.Context())
	if err != nil {
		// log this msg
		respondWithError(w, http.StatusBadRequest, "Error getting depth")
		return
	}
	respondWithJSON(w, http.StatusOK, depthRes)
}

func validateLimitOrderRequest(lo *models.LimitOrder) (*models.LimitOrderType, error) {
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

func validateMarketOrderRequest(req *models.MarketOrderRequest) error {
	if req.Side < 0 || req.Side > 1 {
		return errors.New("invalid side. sell(0), sell(1)")
	}

	if req.Quantity.IsZero() {
		return errors.New("invalid quantity, quantity should be not zero")
	}
	return nil
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	er := models.ErrorResponse{Code: code, Errors: map[string]string{"error": message}}
	respondWithJSON(w, er.Code, er)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
