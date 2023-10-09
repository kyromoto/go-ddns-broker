package clientmanager

import (
	"github.com/google/uuid"
	"inet.af/netaddr"
)

type Service interface {
	UpdateIp(clientid uuid.UUID, ip netaddr.IP) error
	Authenticate(clientid uuid.UUID, password string) bool
}

func New(clientRepository ClientRepository, eventbus Eventbus) Service {
	return &service{
		clientRepository: clientRepository,
		eventbus:         eventbus,
	}
}

type service struct {
	clientRepository ClientRepository
	eventbus         Eventbus
}

func (s *service) UpdateIp(clientid uuid.UUID, ip netaddr.IP) error {
	return s.eventbus.PublishClientIpUpdate(clientid, ip)
}

func (s *service) Authenticate(clientid uuid.UUID, password string) bool {
	client, err := s.clientRepository.FindById(clientid)

	if err != nil {
		return false
	}

	if err := PasswordCompare(client.PasswordHash, password); err != nil {
		return false
	}

	return true
}
