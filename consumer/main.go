package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	globalConnMQ, err := amqp.Dial("amqp://guest:guest@0.0.0.0:5672/")
	if err != nil {
		panic(err)
	}
	defer globalConnMQ.Close()
	globalChannelMQ, err := globalConnMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer globalChannelMQ.Close()
	_, _ = globalChannelMQ.QueueDeclare("poem", false, false, false, false, nil)

	var enternalLoop chan struct{}

	messages, _ := globalChannelMQ.Consume(
		"poem", // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	go func() {
		for mex := range messages {
			println("this is the messagge : " + string(mex.Body))
		}
	}() <-enternalLoop
}
