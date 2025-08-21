package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/x01000101x/ticket-project/config"
	"github.com/x01000101x/ticket-project/db"
	"github.com/x01000101x/ticket-project/handlers"
	"github.com/x01000101x/ticket-project/repositories"
)

func main() {

	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "Ticket Booking",
		ServerHeader: "Fiber",
	})

	//Repositories
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)

	//Routing
	server := app.Group("/api")

	//Handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)
	handlers.NewTicketHandler(server.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
