package main

import (
	rabbitmq "microservices/asynchronous/rabbitmq_event"

	"github.com/streadway/amqp"
)

func main() {

	// connect to server
	con, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	rabbitmq.FailOnError(err, "Failed to connect to RabbitMQ")
	defer con.Close()

	// create channel
	ch, err := con.Channel()
	rabbitmq.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// delare a queue to send to
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable: Queue def will survive a server restart
		false,   // delete when unused
		false,   // exclusive: Used by only one connection
		false,   // no-wait
		nil,     // arguments: Optional
	)
	rabbitmq.FailOnError(err, "Failed to delare a queue")

	// declare message body
	body := "Hello, World"

	// publish
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	rabbitmq.FailOnError(err, "Failed to publish message")

}
