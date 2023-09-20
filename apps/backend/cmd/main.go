package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/kyromoto/go-ddns-broker/ClientManager/controllers"
	"github.com/kyromoto/go-ddns-broker/ClientManager/infrastructures"
	"github.com/kyromoto/go-ddns-broker/ClientManager/usecases"
	"github.com/mustafaturan/bus/v3"
	"github.com/mustafaturan/monoton"
	"github.com/mustafaturan/monoton/sequencer"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewBus() *bus.Bus {
	node := uint64(1)
	initTime := uint64(0)

	m, err := monoton.New(sequencer.NewMillisecond(), node, initTime)

	if err != nil {
		log.Fatal().Err(err)
	}

	var idGenerator bus.Next = m.Next

	bus, err := bus.NewBus(idGenerator)

	if err != nil {
		log.Fatal().Err(err)
	}

	return bus
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	app := fiber.New()
	bus := NewBus()

	var messageService usecases.MessageService = infrastructures.NewMessageService(bus)
	var clientRepository usecases.ClientRepository = infrastructures.NewClientRepository()

	app.Use(helmet.New())
	app.Use(requestid.New(requestid.Config{ContextKey: "correlationID"}))

	app.Post("/api/client/update", controllers.AuthorizeClientWithBasicAuth(clientRepository), controllers.UpdateClientIp(clientRepository, messageService))

	err := app.Listen(":4000")

	if err != nil {
		log.Fatal().Err(err)
	}

}
