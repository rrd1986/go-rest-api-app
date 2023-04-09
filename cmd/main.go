package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rrd1986/go-rest-api-app/database"
	"github.com/rrd1986/go-rest-api-app/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
