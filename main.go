package main

import (
	"github.com/gofiber/fiber/v3"
	handler "lasseaakjaer.com/musicbuzz/api"
)

func main() {
	app := fiber.New()

	app.Get("/", handler.Root)

	app.Listen(":3000")
}
