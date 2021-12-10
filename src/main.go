package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":300"))
}
