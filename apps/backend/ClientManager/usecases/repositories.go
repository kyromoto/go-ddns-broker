package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns-broker/ClientManager/entities"
)

type ClientRepository interface {
	Add(ctx *context.Context, client entities.Client) error
	Delete(ctx *context.Context, uuid uuid.UUID) error
	List(ctx *context.Context) ([]entities.Client, error)
	FindByUuid(ctx *context.Context, uuid uuid.UUID) (*entities.Client, error)
}
