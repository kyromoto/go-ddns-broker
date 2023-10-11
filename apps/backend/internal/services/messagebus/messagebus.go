package messagebus

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type message struct {
	id    uuid.UUID
	ctx   context.Context
	topic string
	data  interface{}
}

type MessageBus struct {
	topics map[string]map[uuid.UUID]subscriber
}

func NewMessageBus() *MessageBus {
	return &MessageBus{
		topics: make(map[string]map[uuid.UUID]subscriber),
	}
}

func (b *MessageBus) RegisterTopic(topic string) error {
	_, exists := b.topics[topic]

	if exists {
		return nil
	}

	b.topics[topic] = make(map[uuid.UUID]subscriber)

	return nil
}

func (b *MessageBus) Subscribe(topic string, subscriber subscriber) error {
	subscribers, exists := b.topics[topic]

	if !exists {
		return fmt.Errorf("topic %v not registered", topic)
	}

	_, exists = subscribers[subscriber.GetID()]

	if exists {
		return fmt.Errorf("subscriber %v already registered", subscriber.GetID().String())
	}

	subscribers[subscriber.GetID()] = subscriber

	return nil

}

func (b *MessageBus) Publish(ctx context.Context, topic string, data interface{}) error {
	msg := message{
		id:    uuid.New(),
		ctx:   ctx,
		topic: topic,
		data:  data,
	}

	subscribers, exists := b.topics[msg.topic]

	if !exists {
		return fmt.Errorf("topic %v not registered", msg.topic)
	}

	if len(subscribers) == 0 {
		return fmt.Errorf("no subscriptions for %v", msg.topic)
	}

	for _, subscriber := range subscribers {
		subscriber.HandleMessage(msg)
	}

	return nil
}
