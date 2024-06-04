package main

import (
	"context"
	"main/models"
	"main/routes"

	"github.com/gofiber/fiber/v2"
)

func sampleServer(c *fiber.Ctx) error {
	return c.SendString("boiade guarda qui")
}

func setupServer() {
	app := fiber.New()
	v1 := app.Group("/v1")
	//create ShortenedUrl
	v1.Post("/shortener", routes.CreateShortener)
	//get true url from shortenedUrl
	v1.Get("/shortener/:url", routes.ComputeShortener)
	//database
	models.CreateDatabase()
	ctx := context.Background()
	if err := models.Database.Ping(ctx).Err(); err != nil {
		panic("Not able to setup redis connection, aborting")
	}

	app.Listen(":8080")
}

func main() {
	setupServer()
}
