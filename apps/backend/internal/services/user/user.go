package user

import "github.com/google/uuid"

type User struct {
	id       uuid.UUID
	username string
	password string
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Password() string {
	return u.password
}

func NewUser(username string, password string) (User, error) {
	var user User

	user.id = uuid.New()
	user.username = username
	user.password = password

	return user, nil
}
