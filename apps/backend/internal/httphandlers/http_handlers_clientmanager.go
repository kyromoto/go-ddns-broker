package httphandlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns/internal/services/clientmanager"
	"inet.af/netaddr"
)

func ClientUpdateIp(clientmanager clientmanager.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ipsAsStr := []string{}
		ips := []netaddr.IP{}

		clientIdAsStr := c.Query("clientid")
		ipv4AsStr := c.Query("ipv4")
		ipv6AsStr := c.Query("ipv6")

		clientid, err := uuid.Parse(clientIdAsStr)

		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("fatal")
		}

		if len(ipv4AsStr) != 0 {
			ipsAsStr = append(ipsAsStr, ipv4AsStr)
		}

		if len(ipv6AsStr) != 0 {
			ipsAsStr = append(ipsAsStr, ipv6AsStr)
		}

		if len(ipsAsStr) == 0 {
			return c.Status(http.StatusBadRequest).SendString("fatal")
		}

		for _, ipAsStr := range ipsAsStr {
			ip, err := netaddr.ParseIP(ipAsStr)

			if err != nil {
				return c.Status(http.StatusBadRequest).SendString("fatal")
			}

			ips = append(ips, ip)
		}

		for _, ip := range ips {
			if err := clientmanager.UpdateIp(clientid, ip); err != nil {
				return c.Status(http.StatusInternalServerError).SendString("fatal")
			}
		}

		return c.Status(http.StatusOK).SendString("good")
	}
}

func ClientAuthenticate(clientmanager clientmanager.Service) fiber.Handler {
	return basicauth.New(basicauth.Config{
		Authorizer: func(clientidAsStr, password string) bool {
			clientid, err := uuid.Parse(clientidAsStr)

			if err != nil {
				return false
			}

			return clientmanager.Authenticate(clientid, password)
		},
	})
}
