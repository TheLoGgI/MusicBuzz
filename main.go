package main

import (
	"github.com/gofiber/fiber/v3"
	"lasseaakjaer.com/musicbuzz/api"
)

func main() {
	app := fiber.New()

	app.Get("/", api.Handler)

	app.Listen(":3000")
}
