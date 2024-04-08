package handler

import (
	"github.com/gofiber/fiber/v3"
)

func Handler(c fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
