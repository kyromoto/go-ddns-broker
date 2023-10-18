package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kyromoto/go-ddns/internal/services/user"
)

type CmdUserCreateReqBody struct {
	Username string
	Password string
}

type CmdUserCreateResBody struct {
}

func HandleCmdUserCreate(userCreaterService user.CreateService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody CmdUserCreateReqBody

		if err := c.BodyParser(reqBody); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResponseBody{
				error: "bad request body",
			})
		}

		user, err := user.NewUser(reqBody.Username, reqBody.Password)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(ErrorResponseBody{
				error: err.Error(),
			})
		}

		if err := userCreaterService(user); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(ErrorResponseBody{
				error: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(CmdUserCreateResBody{})
	}
}
