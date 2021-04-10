package main

import (
	"log"
	"time"

	rabbitmq "microservices/asynchronous/rabbitmq_event"

	"github.com/streadway/amqp"
)

func main() {

	// open connection
	con, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	rabbitmq.FailOnError(err, "Failed to connect to RabbitMQ")
	defer con.Close()

	// open channel
	ch, err := con.Channel()
	rabbitmq.FailOnError(err, "Failed to open channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	rabbitmq.FailOnError(err, "Failed to declare an exchange")

	// delare a queue to receive from as we might start the consume before the send
	// builds it (queue's are idempotent)
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable: Queue def will survive a server restart
		false, // delete when unused
		true,  // exclusive: Used by only one connection - when connection closes, the queue will be deleted
		false, // no-wait
		nil,   // arguments: Optional
	)
	rabbitmq.FailOnError(err, "Failed to delare a queue")

	// queue bind: This is to tell the exchange where to send messages to
	err = ch.QueueBind(
		q.Name,
		"",
		"logs",
		false,
		nil,
	)
	rabbitmq.FailOnError(err, "Failed binding to queue")

	// read from queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack: Setting to false, will mean an acknowledgement of done will need to be sent.
		//			 This ensures that if a process fails, the message will be re queued
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	rabbitmq.FailOnError(err, "Failed to delare a queue")

	forever := make(chan bool)

	go func(msgs <-chan amqp.Delivery) {

		for d := range msgs {
			log.Printf("Received message: %s\n", d.Body)
			time.Sleep(1 * time.Second)
			d.Ack(false) // send ack that message has been processed correctly
		}

	}(msgs)

	<-forever
}
