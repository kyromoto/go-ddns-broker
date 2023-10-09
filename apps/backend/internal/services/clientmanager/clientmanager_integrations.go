package clientmanager

import (
	"github.com/google/uuid"
	"inet.af/netaddr"
)

type SubscriberClientIpUpdate interface {
	OnClientIpUpdate(clientipd uuid.UUID, ip netaddr.IP) error
}

type Eventbus interface {
	PublishClientIpUpdate(clientid uuid.UUID, ip netaddr.IP) error
	SubscribeClientIpUpdate(subscriber SubscriberClientIpUpdate) error
}
