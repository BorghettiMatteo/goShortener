package main

import (
	"main/handlers"

	"main/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// init database
	var db models.MongoCon
	db.CreateDb()
	//defer db.KillMongoDB()

	// init fiber
	app := fiber.New()
	app.Post("/pastabin", handlers.CreatePastaBin)
	app.Get("/:pastabinid", handlers.GetPastaBin)
	app.Listen(":5555")
}
