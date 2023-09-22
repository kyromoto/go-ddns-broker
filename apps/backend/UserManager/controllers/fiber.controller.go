package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyromoto/go-ddns-broker/UserManager/usecases"
	"github.com/kyromoto/go-ddns-broker/lib"
)

type ErrorResponse struct {
	error struct {
		message string
	}
}

func AuthorizeUserWithBasicAuth(userRepository usecases.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authenticateUserDto := usecases.AuthenticateUserRequestDTO{}

		if err := c.BodyParser(authenticateUserDto); err != nil {
			return fiber.ErrBadRequest
		}

		ctx := lib.NewContextWithCorrelationIdFromFiber(c)

		authenticateUser := usecases.NewAuthenticateUserUC(userRepository)

		if !authenticateUser(*ctx, authenticateUserDto) {
			return fiber.ErrUnauthorized
		}

		return c.SendString("ok")
	}
}
