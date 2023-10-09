package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kyromoto/go-ddns/internal/environment"
	"github.com/kyromoto/go-ddns/internal/httphandlers"
	"github.com/kyromoto/go-ddns/internal/integrations/eventbus"
	"github.com/kyromoto/go-ddns/internal/repositories/clientrepository"
	"github.com/kyromoto/go-ddns/internal/services/clientmanager"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open(environment.DbSqliteFilename("./db.sqlite")))

	if err != nil {
		panic(err)
	}

	eventbus := eventbus.NewInMemory()

	clientRepository := clientrepository.New(db)
	clientmanager := clientmanager.New(clientRepository, eventbus)

	app := fiber.New(fiber.Config{})
	// app.Use(helmet.New())

	client := app.Group("/client")
	// client.Use(httphandlers.ClientAuthenticate(clientmanager))
	client.Post("/update", httphandlers.ClientUpdateIp(clientmanager))

	if err := app.Listen(fmt.Sprintf(":%v", environment.HttpPort(3333))); err != nil {
		panic(err)
	}
}
