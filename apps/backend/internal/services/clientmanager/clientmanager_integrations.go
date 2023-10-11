package clientmanager

import (
	"context"

	"github.com/google/uuid"
	"inet.af/netaddr"
)

type MessageBus interface {
	PublishClientIpUpdate(ctx context.Context, clientid uuid.UUID, ip netaddr.IP) error
}
