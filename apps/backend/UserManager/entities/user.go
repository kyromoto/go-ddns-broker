package entities

import (
	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns-broker/lib"
)

type User struct {
	lib.Entity
	username string
	password []byte
}

func (u *User) AssertPassword(password string) (ok bool) {
	return lib.ComparePassword(lib.PasswordSalt, u.password, password)
}

func NewUser(id uuid.UUID, username string, password string) {

}
