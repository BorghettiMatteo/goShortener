package main

import (
	"context"
	"encoding/json"
	"main/models"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var globalChannelMQ *amqp.Channel
var client *mongo.Client

func main() {
	var err error
	mongoDbString := "mongodb://0.0.0.0:27017"
	client, err = mongo.Connect(context.TODO(), options.Client().
		ApplyURI(
			mongoDbString,
		),
	)
	if err != nil {
		panic("Not able to create mongodb connection: " + err.Error())
	}

	//kill mongodb connection
	defer client.Disconnect(context.TODO())

	//setup rabbitmq
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
	_, _ = globalChannelMQ.QueueDeclare("poem", false, false, false, false, nil)

	var enternalLoop chan struct{}

	messages, err := globalChannelMQ.Consume(
		"poem", // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		print(err.Error())
	}

	//define collection to query on

	coll := client.Database("db").Collection("poems")

	go func() {
		for mex := range messages {
			println("this is the messagge : " + string(mex.Body))
			insert(coll, mex.Body)
		}
	}()
	<-enternalLoop
}

func insert(coll *mongo.Collection, mex []byte) {
	//insert in mongodb
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	oneMessage := new(models.Poem)
	defer cancel()
	//insert in mongodb
	_ = json.Unmarshal(mex, oneMessage)
	res, err := coll.InsertOne(ct, oneMessage)
	if err != nil {
		panic(err)
	}
	println(res.InsertedID)
}
