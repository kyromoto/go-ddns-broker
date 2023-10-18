package user

import (
	"fmt"

	"github.com/google/uuid"
)

type CreateService func(user User) error
type DeleteService func(id uuid.UUID) error
type AuthenticateService func(username string, password string) bool

func NewCreateService(userRepository UserRepository) CreateService {
	return func(user User) error {
		if err := userRepository.Save(user); err != nil {
			return fmt.Errorf("save user to db failed")
		}

		return nil
	}
}

func NewDeleteService(userRepository UserRepository) DeleteService {
	return func(id uuid.UUID) error {
		if err := userRepository.Delete(id); err != nil {
			return fmt.Errorf("delete user from db failed")
		}

		return nil
	}
}

func NewAuthenticateService(userRepository UserRepository) AuthenticateService {
	return func(username, password string) bool {
		user, err := userRepository.ReadByUsername(username)

		if err != nil {
			return false
		}

		return user.Password() == password
	}
}
