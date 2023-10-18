package client

import (
	"fmt"

	"github.com/google/uuid"
)

type Client struct {
	id          uuid.UUID
	description string
	password    string
}

func (c *Client) ID() uuid.UUID {
	return c.id
}

func (c *Client) Description() string {
	return c.description
}

func (c *Client) VerifyPassword(password string) bool {
	return c.password == password
}

func (c *Client) ChangePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must have min 8 characters")
	}

	c.password = password

	return nil
}

func NewClient(description string, password string) (Client, error) {
	var client Client

	client.id = uuid.New()
	client.description = description

	if err := client.ChangePassword(password); err != nil {
		return Client{}, err
	}

	return client, nil
}
