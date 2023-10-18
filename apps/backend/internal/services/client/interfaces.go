package client

import (
	"context"

	"github.com/google/uuid"
	"inet.af/netaddr"
)

type ClientRepository interface {
	Save(dto ClientRepositoryDTO) error
	Delete(id uuid.UUID) error

	Read() ([]ClientRepositoryDTO, error)
	ReadById(id uuid.UUID) (ClientRepositoryDTO, error)
}

type ClientRepositoryDTO struct {
	ID          uuid.UUID
	Description string
	Password    string
}

type Messagebus interface {
	SendMessage(ctx context.Context, message IpUpdatedMessage) error
}

type IpUpdatedMessage struct {
	ClientID uuid.UUID
	ClientIP netaddr.IP
}

func NewIpUpdatedMessage(clientID uuid.UUID, ip netaddr.IP) IpUpdatedMessage {
	return IpUpdatedMessage{
		ClientID: clientID,
		ClientIP: ip,
	}
}
