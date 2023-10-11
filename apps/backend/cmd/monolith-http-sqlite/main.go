package main

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns/internal/adapter"
	"github.com/kyromoto/go-ddns/internal/apihandlers"
	"github.com/kyromoto/go-ddns/internal/environment"
	"github.com/kyromoto/go-ddns/internal/lib"
	"github.com/kyromoto/go-ddns/internal/services/clientmanager"
	"github.com/kyromoto/go-ddns/internal/services/messagebus"
	"github.com/kyromoto/go-ddns/internal/storage/dbsqlite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"inet.af/netaddr"
)

func main() {
	topicClientUpdateIp := "client.update-ip"

	bus := messagebus.NewMessageBus()

	q1 := messagebus.NewQueue()
	c1 := messagebus.NewConsumer(func(ctx context.Context, topic string, data interface{}) {
		logger := lib.LoggerWithCorrelationID(ctx)

		payload, ok := data.(struct {
			ClientID uuid.UUID
			IP       netaddr.IP
		})

		if !ok {
			log.Error().Msg("payload conversion failed")
			return
		}

		logger.Info().Str("clientid", payload.ClientID.String()).Str("ip", payload.IP.String()).Msg("client update ip received")
	})

	bus.RegisterTopic(topicClientUpdateIp)
	bus.Subscribe(topicClientUpdateIp, &c1)
	q1.Subscribe(&c1)

	db, err := gorm.Open(sqlite.Open(environment.DbSqliteFilename("./db.sqlite")))

	if err != nil {
		panic(err)
	}

	clientStore := dbsqlite.NewClientStore(db)
	clientManager := clientmanager.New(clientStore, adapter.NewClientmanagerMessageBus(bus, topicClientUpdateIp))

	app := fiber.New(fiber.Config{})
	app.Use(requestid.New(requestid.Config{
		ContextKey: "requestid",
		Generator: func() string {
			return uuid.New().String()
		},
	}))
	// app.Use(helmet.New())

	app.Get("/health", apihandlers.Health())

	clientGroup := app.Group("/client")
	// client.Use(httphandlers.ClientAuthenticate(clientmanager))
	clientGroup.Post("/update-ip", apihandlers.ClientUpdateIp(clientManager))

	if err := app.Listen(fmt.Sprintf(":%v", environment.HttpPort(3333))); err != nil {
		panic(err)
	}
}
