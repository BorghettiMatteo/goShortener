package routes

import (
	"github.com/gofiber/fiber/v2"
)

func CreateShortener(c *fiber.Ctx) error {
	return c.SendString("AAAAAAAAAAAa")

}
