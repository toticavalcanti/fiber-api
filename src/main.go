package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/toticavalcanti/fiber-api/database"
	"github.com/toticavalcanti/fiber-api/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to an Awesome API")
}

func setupRoutes(app *fiber.App) {
	// Welcome endpoint
	app.Get("/api", welcome)
	// User endpoints
	app.Post("/api/users", routes.CreateUser)
}

func main() {
	database.ConnectDb()

	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
