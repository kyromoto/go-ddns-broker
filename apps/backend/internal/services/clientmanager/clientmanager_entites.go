package clientmanager

import "github.com/google/uuid"

type Client struct {
	Id uuid.UUID

	PasswordHash string
}
