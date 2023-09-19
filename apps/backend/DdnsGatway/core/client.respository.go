package core

import (
	"github.com/google/uuid"
	"inet.af/netaddr"
)

type ClientRepository interface {
	GetClientByUuid(uuid uuid.UUID) (error, Client)
	UpdateIp(uuid uuid.UUID, ip netaddr.IP) (ok bool)
}

type MessageService interface {
	ClientIpUpdated(clientuuid uuid.UUID, ip netaddr.IP) error
}
