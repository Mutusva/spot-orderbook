package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
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

func (c *OpsClient) ReceiveMessage(ctx context.Context) string {
	return ""
}
