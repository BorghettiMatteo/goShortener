package main

import (
	"encoding/json"
	"main/models"
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"

	amqp "github.com/rabbitmq/amqp091-go"
)

var globalChannelMQ *amqp.Channel

func createMessage(c *fiber.Ctx) error {
	response := fiber.Get("https://poetrydb.org/author/shakespeare")
	_, double, _ := response.Bytes()
	// unmarshaling
	var dump []models.Poem
	err := json.Unmarshal(double, &dump)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	//find a random poem and send it back

	currentPoem := dump[rand.Intn(len(dump))]
	var desiredOutput string
	for _, line := range currentPoem.Lines {
		desiredOutput += line
	}

	//add to mq server
	error := globalChannelMQ.Publish("", "poem", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(desiredOutput),
	})
	if error != nil {
		return c.SendString(error.Error())
	}
	return c.SendStatus(http.StatusCreated)

}

func initServer() {
	//init rabbit mq server
	globalConnMQ, err := amqp.Dial("amqp://guest:guest@0.0.0.0:5672/")
	if err != nil {
		panic(err)
	}
	defer globalConnMQ.Close()
	globalChannelMQ, err = globalConnMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer globalChannelMQ.Close()
	//declare queue

	_, _ = globalChannelMQ.QueueDeclare("poem", false, false, false, false, nil)
	//webserver
	app := fiber.New()
	app.Get("/", createMessage)
	app.Listen(":5000")

}

func main() {
	println("asdassd")
	initServer()
}
