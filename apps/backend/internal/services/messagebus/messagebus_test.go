package messagebus_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/kyromoto/go-ddns/internal/services/messagebus"
)

func TestMessageBus(t *testing.T) {
	topic := "client.update-ip"
	msg := "Hello, World"

	bus := messagebus.NewMessageBus()
	q1 := messagebus.NewQueue()

	c1 := messagebus.NewConsumer(func(ctx context.Context, topic string, data interface{}) {
		fmt.Printf("c1 @ %v: %v", topic, data)

		if data != msg {
			t.Errorf("c1: %v != %v", data, msg)
		}
	})

	c2 := messagebus.NewConsumer(func(ctx context.Context, topic string, data interface{}) {
		fmt.Printf("c1 @ %v: %v", topic, data)

		if data != msg {
			t.Errorf("c1: %v != %v", data, msg)
		}
	})

	if err := bus.RegisterTopic(topic); err != nil {
		t.Error(err)
	}

	if err := bus.Subscribe(topic, &q1); err != nil {
		t.Error(err)
	}

	if err := q1.Subscribe(&c1); err != nil {
		t.Error(err)
	}

	if err := q1.Subscribe(&c2); err != nil {
		t.Error(err)
	}

	if err := bus.Publish(context.Background(), topic, msg); err != nil {
		t.Error(err)
	}
}
