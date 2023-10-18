package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns/internal/services/client"
)

func MiddlewareClientAuthenticate(clientAuthenticateService client.AuthenticateService) fiber.Handler {
	return basicauth.New(basicauth.Config{
		Authorizer: func(clientidAsStr, password string) bool {
			clientID, err := uuid.Parse(clientidAsStr)

			if err != nil {
				return false
			}

			isAuthenticated, err := clientAuthenticateService(context.Background(), client.AuthenticateServiceProps{
				ID:       clientID,
				Password: password,
			})

			if err != nil {
				return false
			}

			return isAuthenticated
		},
	})
}
