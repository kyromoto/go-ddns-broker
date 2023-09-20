package entities

import "github.com/google/uuid"

type Entity struct {
	_uuid uuid.UUID
}

func (e *Entity) GetUuid() uuid.UUID {
	return e._uuid
}

func NewEntity() Entity {
	return Entity{
		_uuid: uuid.New(),
	}
}
