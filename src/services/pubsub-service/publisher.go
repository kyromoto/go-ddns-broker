package pubsubservice

import (
	"context"
	"encoding/json"
	"log"
	"net/netip"

	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/mempubsub"
)

type Publisher struct {
	ctx   context.Context
	topic *pubsub.Topic
}

func NewPublisher(topicUrl string) *Publisher {
	ctx := context.Background()
	topic, err := pubsub.OpenTopic(ctx, topicUrl)

	if err != nil {
		log.Fatal(err)
	}

	publisher := &Publisher{
		ctx:   ctx,
		topic: topic,
	}

	return publisher
}

func (p *Publisher) Close() {
	defer p.topic.Shutdown(p.ctx)
}

func (p Publisher) PublishUpdate(username string, ip netip.Addr) error {
	payload := updateMessagePayload{
		Username: username,
		Ip:       ip,
	}

	bytes, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	message := &pubsub.Message{
		Body: bytes,
	}

	err = p.topic.Send(p.ctx, message)

	if err != nil {
		return err
	}

	return nil
}
