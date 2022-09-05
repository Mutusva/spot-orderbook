package orderbook

import (
	"context"
	"spotob/orderbook/models"
)

type OrderBookClient interface {
	ProcessLimitOrder(ctx context.Context, req models.LimitOrder) (*models.LimitOrderResponse, error)
	ProcessMarketOrder(ctx context.Context, req models.MarketOrderRequest) (*models.MarketOrderResponse, error)
	CancelOrder(ctx context.Context, Id string) (*models.Order, error)
	Depth(ctx context.Context) (*models.OrderBookDepth, error)
}
