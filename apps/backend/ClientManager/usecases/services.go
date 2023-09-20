package usecases

import (
	"context"

	"github.com/google/uuid"
	"inet.af/netaddr"
)

type MessageService interface {
	ClientIpUpdated(ctx *context.Context, clientuuid uuid.UUID, ip netaddr.IP) (ok bool)
}
