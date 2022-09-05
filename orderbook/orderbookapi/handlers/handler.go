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
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	ob "github.com/muzykantov/orderbook"
	"log"
	"net/http"
	"spotob/orderbook/env"
	"spotob/orderbook/models"
	obs "spotob/orderbook/orderbookservice"
	rc "spotob/orderbook/redis"
)

type App struct {
	OrderBook   *ob.OrderBook
	Router      *mux.Router
	RedisClient *rc.OpsClient
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
		swaggerPath = "./src/orderbookapi/swagger/"
	}

	a.Router.Handle("/docs", sh)
	a.Router.Handle("/swagger.yaml", http.FileServer(http.Dir(swaggerPath)))
}

func (a *App) Initialize(ctx context.Context, orderBook *ob.OrderBook, rc *rc.OpsClient) {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
	a.OrderBook = orderBook
	a.RedisClient = rc
	a.ctx = ctx
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

	lr, err := validateLimitOrderRequest(&limitOrder)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	done, partial, partialQuantityProcessed, err := obs.ProcessLimitOrder(a.OrderBook, lr.Side, lr.OrderId, lr.Quantity, lr.Price)
	if err != nil {
		// log this msg
		respondWithError(w, http.StatusBadRequest, "error processing limit order")
		return
	}
	respondWithJSON(w, http.StatusOK, models.LimitOrderResponse{
		Done:                     done,
		Partial:                  partial,
		PartialQuantityProcessed: partialQuantityProcessed,
	})

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

	done, partial, partialQuantityProcessed, quantityLeft, err := obs.ProcessMarketOrder(a.OrderBook, ob.Side(req.Side), req.Quantity)
	if err != nil {
		// log this msg
		respondWithError(w, http.StatusBadRequest, "Error processing market order")
		return
	}
	respondWithJSON(w, http.StatusOK, models.MarketOrderResponse{
		Done:                     done,
		Partial:                  partial,
		PartialQuantityProcessed: partialQuantityProcessed,
		QuantityLeft:             quantityLeft,
	})

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
//       type: integer
//       format: int

// CancelOrder cancel a specific order by id
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

// swagger:route GET /depth depth
// Responses:
//   200: OrderBookDepth
//   401: ErrorResponse
func (a *App) Depth(w http.ResponseWriter, _ *http.Request) {
	asks, bids := obs.Depth(a.OrderBook)
	depth := models.OrderBookDepth{
		Asks: asks,
		Bids: bids,
	}
	respondWithJSON(w, http.StatusOK, depth)
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

func (a *App) ProcessOrders(ctx context.Context) {

	defer a.RedisClient.Rc.Close()
	pubSub := a.RedisClient.Rc.Subscribe(ctx, a.RedisClient.Ch)

	for {
		msg, err := pubSub.ReceiveMessage(ctx)
		if err != nil {
			log.Println(err)
		}

		data := make(map[string]interface{})
		err = json.Unmarshal([]byte(msg.Payload), &data)
		if err != nil {
			log.Println(err)
			continue
		}

		mt, ok := data["message_type"].(float64)
		if !ok {
			//log.Println(err)
			continue
		}
		// log this to kibana or something
		fmt.Println(msg.Channel, msg.Payload, mt)
		ob.NewOrderBook()

		processMessageType(a.OrderBook, models.MessageType(int32(mt)), msg.Payload)
	}
}

func processMessageType(o *ob.OrderBook, mt models.MessageType, payload string) {
	switch mt {
	case models.LimitOrderMessageType:
		var orderMsg models.LimitOrderMessage
		err := json.Unmarshal([]byte(payload), &orderMsg)
		if err != nil {
			log.Println(err)
		}

		_, _, _, err = obs.ProcessLimitOrder(o, ob.Side(orderMsg.Message.Side), orderMsg.Message.OrderId, orderMsg.Message.Quantity, orderMsg.Message.Price)
		if err != nil {
			log.Println(err)
		}
		break

	case models.MarketOrderMessageType:
		var orderMsg models.MarketOrderMessage
		err := json.Unmarshal([]byte(payload), &orderMsg)
		if err != nil {
			log.Println(err)
		}

		_, _, _, _, err = obs.ProcessMarketOrder(o, ob.Side(orderMsg.Message.Side), orderMsg.Message.Quantity)
		if err != nil {
			log.Println(err)
		}
		break

	case models.CancelOrderMessageType:
		var orderMsg models.CancelOrderMessage
		err := json.Unmarshal([]byte(payload), &orderMsg)
		if err != nil {
			log.Println(err)
		}
		_ = obs.CancelOrder(o, orderMsg.Message)
		break
	}

}
