package lib

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

const PasswordSalt = "!r4uMGxngfFr*mFUY8KyGpI=O4yd+8KTHj/WH*5A5NYwQgucTOeomQw1dlSbt+/R"
const costs = 10

func HashPassword(salt string, password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(salt+password), costs)

	if err != nil {
		log.Fatal().Err(err).Msg("hash client password failed")
	}

	return hash
}

func ComparePassword(salt string, hash []byte, password string) (ok bool) {
	return bcrypt.CompareHashAndPassword(hash, []byte(salt+password)) == nil
}
