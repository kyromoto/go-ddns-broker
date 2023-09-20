package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns-broker/ClientManager/usecases"
	"github.com/kyromoto/go-ddns-broker/lib"
	"github.com/rs/zerolog/log"
	"inet.af/netaddr"
)

func AuthorizeClientWithBasicAuth(clientRepository usecases.ClientRepository) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctx := lib.NewContextWithCorrelationIdFromFiber(c)

		basicauthFn := basicauth.New(basicauth.Config{
			Authorizer: func(idStr string, password string) bool {
				id, err := uuid.Parse(idStr)

				if err != nil {
					log.Error().Err(err)
					return false
				}

				client, err := clientRepository.FindByUuid(ctx, id)

				if err != nil {
					log.Error().Err(err)
					return false
				}

				return client.AssertPassword(password)
			},
		})

		return basicauthFn(c)
	}

}

func UpdateClientIp(clientRepository usecases.ClientRepository, messageService usecases.MessageService) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctx := lib.NewContextWithCorrelationIdFromFiber(c)

		var idStr string = c.Params("id")
		var ipV4Str string = c.Params("ipV4")
		var ipV6Str string = c.Params("ipV6")

		id, err := uuid.Parse(idStr)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "fatal")
		}

		if ipV4Str == "" && ipV6Str == "" {
			return fiber.NewError(fiber.StatusBadRequest, "fatal")
		}

		dto := usecases.UpdateClientIpDTO{
			UUID: id,
			IP:   netaddr.IP{},
		}

		updateClientIp := usecases.NewUpdateClientIpUC(clientRepository, messageService)

		var doUpdate = func(ipStr string) error {
			ip, err := netaddr.ParseIP(ipStr)

			if err != nil {
				return err
			}

			dto.IP = ip

			if err := updateClientIp(ctx, dto); err != nil {
				log.Error().Err(err)
				return err
			}

			return nil
		}

		if ipV4Str != "" {
			err = doUpdate(ipV6Str)
		}

		if ipV6Str != "" {
			err = doUpdate(ipV6Str)
		}

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "fatal")
		}

		return c.SendString("good")
	}
}
