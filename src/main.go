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
	// Post User endpoints
	app.Post("/api/users", routes.CreateUser)
	// Get all users
	app.Get("/api/users", routes.Getusers)
	// Get user by id
	app.Get("/api/users/:id", routes.GetUser)
	// Update User
	app.Put("/api/users/:id", routes.UpdateUser)
	// Delete User
	app.Delete("/api/users/:id", routes.DeleteUser)

}

func main() {
	database.ConnectDb()

	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
