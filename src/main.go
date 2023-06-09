package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	httpcontroller "github.com/kyromoto/go-ddns-broker/src/http-controller"
	"github.com/kyromoto/go-ddns-broker/src/repos"
	envservice "github.com/kyromoto/go-ddns-broker/src/services/env-service"
	inwxservice "github.com/kyromoto/go-ddns-broker/src/services/inwx-service"
	pubsubservice "github.com/kyromoto/go-ddns-broker/src/services/pubsub-service"
)

func getAccessLogFile() *os.File {
	accessLog, err := os.OpenFile("access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return accessLog
}

func main() {
	accessLog := getAccessLogFile()
	defer accessLog.Close()

	_, err := repos.OpenClientRepo()

	if err != nil {
		log.Fatal(err)
	}

	publisher := pubsubservice.NewPublisher(pubsubservice.TopicUrl)
	subscriber1 := pubsubservice.NewSubscriber(pubsubservice.TopicUrl)

	defer publisher.Close()
	defer subscriber1.Close()

	go inwxservice.HandleInwxDnsUpdates(nil)

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Output: accessLog,
	}))

	app.Post("/update", httpcontroller.AuthorizeClient(), httpcontroller.PostUpdate(publisher))
	app.Get("/status", httpcontroller.GetStatus())

	log.Fatal(app.Listen(envservice.HttpApiListen()))
}
