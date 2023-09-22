package usecases

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

type AuthenticateUserRequestDTO struct {
	username string
	password string
}

func NewAuthenticateUserUC(userRepository UserRepository) func(ctx context.Context, dto AuthenticateUserRequestDTO) (ok bool) {
	return func(ctx context.Context, dto AuthenticateUserRequestDTO) (ok bool) {
		client, err := userRepository.FindByUsername(dto.username)

		if err != nil {
			log.Error().Err(err).Ctx(ctx).Send()
			fmt.Errorf("User not found")
		}

		return client.AssertPassword(dto.password)
	}
}
