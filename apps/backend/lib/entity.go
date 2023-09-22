package lib

import "github.com/google/uuid"

type Entity struct {
	ID uuid.UUID
}

func (e *Entity) GetUuid() uuid.UUID {
	return e.ID
}

func NewEntity() Entity {
	return Entity{
		ID: uuid.New(),
	}
}
