package main

import (
	"github.com/gofiber/fiber/v2"
	"main.go/models"
)

func main() {
	// init database
	var db models.MongoCon
	db.CreateDb()
	defer db.KillMongoDB()

	// init fiber
	app := fiber.New()
	app.Add("POST", "/pastabin", nil)
	app.Get("/", nil)
	app.Listen(":5555")
}
