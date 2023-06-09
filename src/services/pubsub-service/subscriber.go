package pubsubservice

import (
	"context"
	"encoding/json"
	"log"
	"net/netip"

	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/mempubsub"
)

type Subscriber struct {
	ctx          context.Context
	subscription *pubsub.Subscription
}

func NewSubscriber(topicUrl string) *Subscriber {
	ctx := context.Background()
	subscription, err := pubsub.OpenSubscription(ctx, topicUrl)

	if err != nil {
		log.Fatal(err)
	}

	return &Subscriber{
		ctx:          ctx,
		subscription: subscription,
	}
}

func (s *Subscriber) Close() {
	defer s.subscription.Shutdown(s.ctx)
}

func (s *Subscriber) Listen(handle func(username string, ip netip.Addr) error) {
	for {
		msg, err := s.subscription.Receive(s.ctx)

		if err != nil {
			log.Fatal(err)
		}

		payload := updateMessagePayload{}

		err = json.Unmarshal(msg.Body, &payload)

		if err != nil {
			if msg.Nackable() {
				msg.Nack()
			}

			log.Fatal(err)
		}

		err = handle(payload.Username, payload.Ip)

		if err != nil {
			if msg.Nackable() {
				msg.Nack()
			}

			log.Fatal(err)
		}

		msg.Ack()
	}
}
