package infrastructures

import (
	"context"

	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns-broker/ClientManager/usecases"
	"github.com/mustafaturan/bus/v3"
	"inet.af/netaddr"
)

const TopicClientIpUpdated = "client.ip-updated"

type messageServiceImpl struct {
	bus *bus.Bus
}

type MessageClientIpUpdated struct {
	clientuuid uuid.UUID
	ip         netaddr.IP
}

func (ms *messageServiceImpl) ClientIpUpdated(ctx *context.Context, clientuuid uuid.UUID, ip netaddr.IP) (ok bool) {
	ms.bus.RegisterTopics(TopicClientIpUpdated)

	err := ms.bus.Emit(*ctx, TopicClientIpUpdated, MessageClientIpUpdated{clientuuid: clientuuid, ip: ip})

	return err == nil
}

func NewMessageService(bus *bus.Bus) usecases.MessageService {
	return &messageServiceImpl{
		bus: bus,
	}
}
