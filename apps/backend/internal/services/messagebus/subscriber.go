package messagebus

import (
	"github.com/google/uuid"
)

type subscriber interface {
	GetID() uuid.UUID
	HandleMessage(msg message) error
}
