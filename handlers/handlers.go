package handlers

import "github.com/gofiber/fiber/v2"

func createPastaBin(c *fiber.Ctx) error {
	return c.BodyParser(c.Body())

}
