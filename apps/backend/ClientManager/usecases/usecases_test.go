package usecases_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns-broker/ClientManager/entities"
	"github.com/kyromoto/go-ddns-broker/ClientManager/usecases"
	"github.com/kyromoto/go-ddns-broker/lib"
	"inet.af/netaddr"
)

// --------------------------------------------

type clientRepositoryImpl struct {
	clients []entities.Client
}

func (cr *clientRepositoryImpl) Add(ctx *context.Context, client entities.Client) error {
	panic("not implemented")
}

func (cr *clientRepositoryImpl) Delete(ctx *context.Context, uuid uuid.UUID) error {
	panic("not implemented")
}

func (cr *clientRepositoryImpl) List(ctx *context.Context) ([]entities.Client, error) {
	panic("not implemented")
}

func (cr *clientRepositoryImpl) FindByUuid(ctx *context.Context, uuid uuid.UUID) (*entities.Client, error) {
	for _, client := range cr.clients {
		if client.Entity.GetUuid().String() == uuid.String() {
			return &client, nil
		}
	}

	return nil, fmt.Errorf("not found")
}

func (cr *clientRepositoryImpl) UpdateIp(ctx *context.Context, uuid uuid.UUID, ip netaddr.IP) (ok bool) {
	return true
}

// --------------------------------------------

type messageServiceImpl struct {
}

func (ms *messageServiceImpl) ClientIpUpdated(ctx *context.Context, clientuuid uuid.UUID, ip netaddr.IP) (ok bool) {
	return true
}

// --------------------------------------------

var c1 = entities.NewClient(uuid.New(), "c1", "12345")

var clientRepository = clientRepositoryImpl{
	clients: []entities.Client{c1},
}

var messageService = messageServiceImpl{}

func TestAuthenticateClientUC(t *testing.T) {
	ctx := lib.NewContextWithCorrelationId(string(uuid.New().String()))
	ok := false

	dto := usecases.AuthenticateClientDTO{
		UUID:     c1.GetUuid(),
		Password: "12345",
	}

	authenticateClient := usecases.NewAuthenticateClientUC(&clientRepository)

	ok = authenticateClient(ctx, dto)

	if !ok {
		t.Errorf("authenticate client should be ok")
	}

	dto.Password = "123"

	ok = authenticateClient(ctx, dto)

	if ok {
		t.Errorf("authenticate client should be fail")
	}

	dto.UUID = uuid.New()

	ok = authenticateClient(ctx, dto)

	if ok {
		t.Errorf("authenticate client should be fail")
	}
}

func TestHandleClientIpUpdate(t *testing.T) {
	ctx := lib.NewContextWithCorrelationId(string(uuid.New().String()))

	err := usecases.NewUpdateClientIpUC(&clientRepository, &messageService)(ctx, usecases.UpdateClientIpDTO{
		UUID: c1.GetUuid(),
		IP:   netaddr.IPv4(192, 168, 1, 1),
	})

	if err != nil {
		t.Error(err)
	}
}
