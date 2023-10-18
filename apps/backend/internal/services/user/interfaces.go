package user

import "github.com/google/uuid"

type UserRepository interface {
	Save(user User) error
	Delete(id uuid.UUID) error

	Read() ([]User, error)
	ReadById(id uuid.UUID) (User, error)
	ReadByUsername(username string) (User, error)
}
