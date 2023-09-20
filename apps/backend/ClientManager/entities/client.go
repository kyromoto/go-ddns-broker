package entities

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

const costs = 10
const salt = "!r4uMGxngfFr*mFUY8KyGpI=O4yd+8KTHj/WH*5A5NYwQgucTOeomQw1dlSbt+/R"

type Client struct {
	Entity
	description string
	password    []byte
}

func (c *Client) AssertPassword(password string) (ok bool) {
	return bcrypt.CompareHashAndPassword(c.password, []byte(salt+password)) == nil
}

func (c *Client) GetPasswordHash() string {
	return string(c.password)
}

func NewClient(uuid uuid.UUID, description string, password string) Client {
	return Client{
		Entity:      Entity{_uuid: uuid},
		description: description,
		password:    HashPassword(password),
	}
}

func HashPassword(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(salt+password), costs)

	if err != nil {
		log.Fatal().Err(err).Msg("hash client password failed")
	}

	return hash
}
