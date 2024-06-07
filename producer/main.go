package main

import (
	"context"
	"encoding/json"
	"log"
	"main/models"
	"math/rand"
	"net/http"
	"reflect"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	amqp "github.com/rabbitmq/amqp091-go"
)

var globalChannelMQ *amqp.Channel
var client *mongo.Client

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
	currentPoemBytes, _ := json.Marshal(currentPoem)
	/*
		var desiredOutput string
		for _, line := range currentPoem.Lines {
			desiredOutput += line
		}
	*/

	//add to mq server
	error := globalChannelMQ.Publish("", "poem", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        currentPoemBytes,
	})
	if error != nil {
		return c.SendString(error.Error())
	}
	return c.SendStatus(http.StatusCreated)

}

func getPoems(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	coll := client.Database("db").Collection("poems")
	cur, err := coll.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		a := new(models.Poem)
		bson.Unmarshal(cur.Current, &a)
		// do something with result....
		c.JSON(a)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return c.SendStatus(http.StatusOK)
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
	app.Get("/poems", getPoems)

	mongoDbString := "mongodb://0.0.0.0:27017"
	tM := reflect.TypeOf(bson.M{})
	reg := bson.NewRegistry()
	reg.RegisterTypeMapEntry(bsontype.Type(0x03), tM)
	//.RegisterTypeMapEntry(bsontype.EmbeddedDocument, tM).Build()
	client, err = mongo.Connect(context.TODO(), options.Client().
		ApplyURI(mongoDbString).
		SetRegistry(reg),
	)
	if err != nil {
		panic("Not able to create mongodb connection: " + err.Error())
	}

	//kill mongodb connection
	//defer client.Disconnect(context.TODO())
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	app.Listen(":5000")

}

func main() {
	println("asdassd")
	initServer()
}
