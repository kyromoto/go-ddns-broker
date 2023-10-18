package adapter

import (
	"context"

	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns/internal/services/client"
	"github.com/kyromoto/go-ddns/internal/services/messagebus"
	"inet.af/netaddr"
)

func NewClientUpdateIpMessageBusAdapter(bus *messagebus.MessageBus, topic string) client.Messagebus {

	bus.RegisterTopic(topic)

	return &clientUpdateIpMessageBusAdapter{
		bus:   bus,
		topic: topic,
	}

}

type clientUpdateIpMessageBusAdapter struct {
	bus   *messagebus.MessageBus
	topic string
}

func (b *clientUpdateIpMessageBusAdapter) SendMessage(ctx context.Context, message client.IpUpdatedMessage) error {
	return b.bus.Publish(ctx, b.topic, struct {
		ClientID uuid.UUID
		ClientIP netaddr.IP
	}{
		ClientID: message.ClientID,
		ClientIP: message.ClientIP,
	})
}
