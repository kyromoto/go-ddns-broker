package clientmanager

import "github.com/google/uuid"

type ClientRepository interface {
	FindById(clientid uuid.UUID) (Client, error)
}
