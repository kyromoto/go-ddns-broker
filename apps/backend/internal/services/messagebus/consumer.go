package messagebus

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type consumer struct {
	id       uuid.UUID
	lock     bool
	callback func(ctx context.Context, topic string, data interface{})
}

func NewConsumer(callback func(ctx context.Context, topic string, data interface{})) consumer {
	return consumer{
		id:       uuid.New(),
		callback: callback,
	}
}

func (c *consumer) GetID() uuid.UUID {
	return c.id
}

func (c *consumer) HandleMessage(msg message) error {
	if c.lock {
		return fmt.Errorf("consumer is locked")
	}

	c.lock = true

	c.callback(msg.ctx, msg.topic, msg.data)
	c.lock = false

	return nil
}
