package core

import (
	"github.com/google/uuid"
	"inet.af/netaddr"
)

func NewAuthenticateClientUC(clientRepository ClientRepository) func(clientuuid uuid.UUID, password string) (ok bool) {
	return func(clientuuid uuid.UUID, password string) bool {
		err, client := clientRepository.GetClientByUuid(clientuuid)

		if err != nil {
			return false
		}

		if !client.AssertPassword(password) {
			return false
		}

		return true
	}
}

func NewHandleClientIpUpdateUC(clientRepository ClientRepository, messageService MessageService) func(clientuuid uuid.UUID, ip netaddr.IP) (ok bool) {
	return func(clientuuid uuid.UUID, ip netaddr.IP) bool {
		err, _ := clientRepository.GetClientByUuid(clientuuid)

		if err != nil {
			return false
		}

		ok := clientRepository.UpdateIp(clientuuid, ip)

		if !ok {
			return false
		}

		err = messageService.ClientIpUpdated(clientuuid, ip)

		if err != nil {
			return false
		}

		return true
	}
}
