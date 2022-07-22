package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type OpsClient struct {
	Rc *redis.Client
	Ch string
}

func NewOpsClient(cl *redis.Client, channel string) *OpsClient {
	return &OpsClient{
		Rc: cl,
		Ch: channel,
	}
}

func (c *OpsClient) PublishMessage(ctx context.Context, msg string) error {
	err := c.Rc.Publish(ctx, c.Ch, msg).Err()
	if err != nil {
		return errors.New("could not publish the message")
	}
	return nil
}

func (c *OpsClient) SaveOrderBook(ctx context.Context, key string, orderBook string) error {
	err := c.Rc.Set(ctx, key, orderBook, time.Minute*24).Err()
	if err != nil {
		log.Println(err)
		return nil
	}
	return nil
}

func (c *OpsClient) GetSavedOrderBook(ctx context.Context, key string) (string, error) {
	val, err := c.Rc.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
