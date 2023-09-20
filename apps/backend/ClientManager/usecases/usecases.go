package usecases

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns-broker/ClientManager/entities"
	"github.com/rs/zerolog/log"
	"inet.af/netaddr"
)

type AuthenticateClientDTO struct {
	UUID     uuid.UUID
	Password string
}

func NewAuthenticateClientUC(clientRepository ClientRepository) func(ctx *context.Context, dto AuthenticateClientDTO) (ok bool) {
	return func(ctx *context.Context, dto AuthenticateClientDTO) bool {

		client, err := clientRepository.FindByUuid(ctx, dto.UUID)

		if err != nil {
			return false
		}

		if !client.AssertPassword(dto.Password) {
			return false
		}

		return true
	}
}

type UpdateClientIpDTO struct {
	UUID uuid.UUID
	IP   netaddr.IP
}

func NewUpdateClientIpUC(clientRepository ClientRepository, messageService MessageService) func(ctx *context.Context, dto UpdateClientIpDTO) error {

	return func(ctx *context.Context, dto UpdateClientIpDTO) error {

		_, err := clientRepository.FindByUuid(ctx, dto.UUID)

		if err != nil {
			return err
		}

		if !messageService.ClientIpUpdated(ctx, dto.UUID, dto.IP) {
			return fmt.Errorf("send update ip message failed")
		}

		return nil
	}
}

type RegisterClientDTO struct {
	Description string
	Password    string
}

func NewRegisterClientUC(clientRepository ClientRepository) func(ctx *context.Context, dto RegisterClientDTO) (uuid.UUID, error) {
	return func(ctx *context.Context, dto RegisterClientDTO) (uuid.UUID, error) {

		id := uuid.New()
		client := entities.NewClient(id, dto.Description, dto.Password)

		if err := clientRepository.Add(ctx, client); err != nil {
			log.Error().Err(err).Ctx(*ctx).Send()
			return uuid.UUID{}, fmt.Errorf("register client failed")
		}

		return id, nil
	}
}

type UnregisterClientDTO struct {
	UUID uuid.UUID
}

func NewUnregisterClientUC(clientRepository ClientRepository) func(ctx *context.Context, dto UnregisterClientDTO) error {
	return func(ctx *context.Context, dto UnregisterClientDTO) error {
		panic("not implemented")
	}
}

type ListClientsDTO struct {
	Clients []struct {
		UUID        uuid.UUID
		Description string
	}
}

func NewListClientsUC(clientRepository ClientRepository) func(ctx *context.Context) (ListClientsDTO, error) {
	return func(ctx *context.Context) (ListClientsDTO, error) {
		panic("not implemented")
	}
}
