package adapter

import (
	"context"

	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns/internal/services/clientmanager"
	"github.com/kyromoto/go-ddns/internal/services/messagebus"
	"inet.af/netaddr"
)

func NewClientmanagerMessageBus(bus *messagebus.MessageBus, topic string) clientmanager.MessageBus {

	bus.RegisterTopic(topic)

	return &adapterClientmanagerMessageBus{
		bus:   bus,
		topic: topic,
	}

}

type adapterClientmanagerMessageBus struct {
	bus   *messagebus.MessageBus
	topic string
}

func (b *adapterClientmanagerMessageBus) PublishClientIpUpdate(ctx context.Context, clientid uuid.UUID, ip netaddr.IP) error {
	return b.bus.Publish(ctx, b.topic, struct {
		ClientID uuid.UUID
		IP       netaddr.IP
	}{
		ClientID: clientid,
		IP:       ip,
	})
}
