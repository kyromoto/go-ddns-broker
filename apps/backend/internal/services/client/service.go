package client

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"inet.af/netaddr"
)

type CreateService func(ctx context.Context, props CreateServiceProps) (CreateServiceResult, error)
type DeleteService func(ctx context.Context, props DeleteServiceProps) (DeleteServiceResult, error)
type ChangeDescriptionService func(ctx context.Context, props ChangeDescriptionServiceProps) (ChangeDescriptionServiceResult, error)
type UpdateIpService func(ctx context.Context, props UpdateIpServiceProps) error
type AuthenticateService func(ctx context.Context, props AuthenticateServiceProps) (bool, error)

type CreateServiceProps struct {
	Description string
	Password    string
}

type CreateServiceResult struct {
	ID          uuid.UUID
	Description string
}

type DeleteServiceProps struct {
	ID uuid.UUID
}

type DeleteServiceResult struct {
}

type ChangeDescriptionServiceProps struct {
	ID          uuid.UUID
	Description string
}

type ChangeDescriptionServiceResult struct {
	ID          uuid.UUID
	Description string
}

type UpdateIpServiceProps struct {
	ID uuid.UUID
	IP netaddr.IP
}

type AuthenticateServiceProps struct {
	ID       uuid.UUID
	Password string
}

func NewCreateService(clientrepository ClientRepository) CreateService {
	return func(ctx context.Context, props CreateServiceProps) (CreateServiceResult, error) {
		client, err := NewClient(props.Description, props.Password)

		if err != nil {
			return CreateServiceResult{}, err
		}

		if err := clientrepository.Save(MapClientToClientRepositoryDTO(client)); err != nil {
			return CreateServiceResult{}, fmt.Errorf("save client to db failed")
		}

		return CreateServiceResult{
			ID:          client.id,
			Description: client.description,
		}, nil
	}
}

func NewDeleteService(clientrepository ClientRepository) DeleteService {
	return func(ctx context.Context, props DeleteServiceProps) (DeleteServiceResult, error) {
		if err := clientrepository.Delete(props.ID); err != nil {
			return DeleteServiceResult{}, fmt.Errorf("delete client from db failed")
		}

		return DeleteServiceResult{}, nil
	}
}

func NewChangeDescriptionService(clientrepository ClientRepository) ChangeDescriptionService {
	return func(ctx context.Context, props ChangeDescriptionServiceProps) (ChangeDescriptionServiceResult, error) {
		clientRepositoryDTO, err := clientrepository.ReadById(props.ID)

		if err != nil {
			return ChangeDescriptionServiceResult{}, fmt.Errorf("client not found")
		}

		clientRepositoryDTO.Description = props.Description

		if err := clientrepository.Save(clientRepositoryDTO); err != nil {
			return ChangeDescriptionServiceResult{}, fmt.Errorf("save client failed")
		}

		return ChangeDescriptionServiceResult{
			ID:          clientRepositoryDTO.ID,
			Description: clientRepositoryDTO.Description,
		}, nil
	}
}

func NewAuthenticateService(clientRepository ClientRepository) AuthenticateService {
	return func(ctx context.Context, props AuthenticateServiceProps) (bool, error) {
		clientRepositoryDTO, err := clientRepository.ReadById(props.ID)

		if err != nil {
			return false, fmt.Errorf("get client from db failed")
		}

		client := MapClientFromClientRepositoryDTO(clientRepositoryDTO)
		isAuthenticated := client.VerifyPassword(props.Password)

		return isAuthenticated, nil
	}
}

func NewUpdateIpService(clientRepository ClientRepository, messagebus Messagebus) UpdateIpService {
	return func(ctx context.Context, props UpdateIpServiceProps) error {
		if _, err := clientRepository.ReadById(props.ID); err != nil {
			return fmt.Errorf("client not found")
		}

		if err := messagebus.SendMessage(ctx, NewIpUpdatedMessage(props.ID, props.IP)); err != nil {
			return fmt.Errorf("send message via messagebus failed")
		}

		return nil
	}
}
