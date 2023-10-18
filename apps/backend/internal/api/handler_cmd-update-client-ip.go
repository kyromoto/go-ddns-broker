package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns/internal/lib"
	"github.com/kyromoto/go-ddns/internal/services/client"
	"inet.af/netaddr"
)

func HandleCmdClientUpdateIp(clientUpdateIpService client.UpdateIpService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := lib.CreateContextWithCorrelationIDFromRequest(c)
		logger := lib.LoggerWithCorrelationID(ctx)

		logger.Debug().Msg("got client update ip request")

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
			props := client.UpdateIpServiceProps{
				ID: clientid,
				IP: ip,
			}

			if err := clientUpdateIpService(ctx, props); err != nil {
				return c.Status(http.StatusInternalServerError).SendString("fatal")
			}
		}

		return c.Status(http.StatusOK).SendString("good")
	}
}
