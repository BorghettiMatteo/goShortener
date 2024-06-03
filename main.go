package main

import (
	"main/routes"

	"github.com/gofiber/fiber/v2"
)

func sampleServer(c *fiber.Ctx) error {
	return c.SendString("boiade guarda qui")
}

func setupServer() {
	app := fiber.New()
	v1 := app.Group("/v1")
	v1.Get("/", sampleServer)
	v1.Get("/shortener", routes.CreateShortener)
	v1.Get("/prova1", sampleServer)

	app.Listen(":8080")
}

func main() {
	setupServer()
}
