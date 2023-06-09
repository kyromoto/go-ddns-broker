package httpcontroller

import (
	"net/netip"

	"github.com/gofiber/fiber/v2"
)

type UpdatePublisher interface {
	PublishUpdate(username string, ip netip.Addr) error
}

func PostUpdate(updatePublisher UpdatePublisher) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ipv4 := ctx.Query("ipv4")
		ipv6 := ctx.Query("ipv6")
		username := "k-rm-28.ddns.kyro.space"

		var ipStrArray []string
		var ipAddrArray []netip.Addr

		if len(ipv4) > 0 {
			ipStrArray = append(ipStrArray, ipv4)
		}

		if len(ipv6) > 0 {
			ipStrArray = append(ipStrArray, ipv6)
		}

		if len(ipStrArray) == 0 {
			ctx.SendStatus(fiber.ErrBadRequest.Code)
			ctx.SendString("fatal")
			return nil
		}

		for _, ip := range ipStrArray {
			addr, err := netip.ParseAddr(ip)

			if err == nil {
				ipAddrArray = append(ipAddrArray, addr)
			}
		}

		if len(ipStrArray) != len(ipAddrArray) {
			ctx.SendStatus(fiber.ErrBadRequest.Code)
			ctx.SendString("fatal")
			return nil
		}

		for _, addr := range ipAddrArray {
			err := updatePublisher.PublishUpdate(username, addr)

			if err != nil {
				ctx.SendStatus(fiber.StatusInternalServerError)
				ctx.SendString("fatal")
				return err
			}
		}

		return ctx.SendString("ok")
	}

}
