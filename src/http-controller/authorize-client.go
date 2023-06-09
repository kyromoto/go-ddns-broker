package httpcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func AuthorizeClient() func(*fiber.Ctx) error {
	return basicauth.New(basicauth.Config{
		Authorizer: func(username string, password string) bool {
			if username != "oliver" {
				return false
			}

			if password != "skywalker" {
				return false
			}

			return true
		},
	})
}
