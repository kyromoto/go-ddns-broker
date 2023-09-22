package usecases

import "github.com/kyromoto/go-ddns-broker/UserManager/entities"

type UserRepository interface {
	FindByUsername(username string) (*entities.User, error)
}
