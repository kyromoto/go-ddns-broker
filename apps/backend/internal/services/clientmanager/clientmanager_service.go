package clientmanager

import (
	"context"

	"github.com/google/uuid"
	"inet.af/netaddr"
)

type Service interface {
	UpdateIp(ctx context.Context, clientid uuid.UUID, ip netaddr.IP) error
	Authenticate(ctx context.Context, clientid uuid.UUID, password string) bool
}

func New(clientRepository ClientRepository, eventbus MessageBus) Service {
	return &service{
		clientRepository: clientRepository,
		eventbus:         eventbus,
	}
}

type service struct {
	clientRepository ClientRepository
	eventbus         MessageBus
}

func (s *service) UpdateIp(ctx context.Context, clientid uuid.UUID, ip netaddr.IP) error {
	return s.eventbus.PublishClientIpUpdate(ctx, clientid, ip)
}

func (s *service) Authenticate(ctx context.Context, clientid uuid.UUID, password string) bool {
	client, err := s.clientRepository.FindById(clientid)

	if err != nil {
		return false
	}

	if err := PasswordCompare(client.PasswordHash, password); err != nil {
		return false
	}

	return true
}
