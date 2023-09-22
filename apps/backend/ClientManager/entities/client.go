package entities

import (
	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns-broker/lib"
)

type Client struct {
	lib.Entity
	description string
	password    []byte
}

func (c *Client) AssertPassword(password string) (ok bool) {
	lib.ComparePassword(lib.PasswordSalt, c.password, password)
}

func (c *Client) GetPasswordHash() string {
	return string(c.password)
}

func NewClient(uuid uuid.UUID, description string, password string) Client {
	return Client{
		Entity:      lib.Entity{ID: uuid},
		description: description,
		password:    lib.HashPassword(lib.PasswordSalt, password),
	}
}
