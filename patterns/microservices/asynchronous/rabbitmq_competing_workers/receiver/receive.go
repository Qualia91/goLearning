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

	// delare a queue to receive from as we might start the consume before the send
	// builds it (queue's are idempotent)
	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable: Queue def will survive a server restart
		false,   // delete when unused
		false,   // exclusive: Used by only one connection
		false,   // no-wait
		nil,     // arguments: Optional
	)
	rabbitmq.FailOnError(err, "Failed to delare a queue")

	// in order to make sure work is dispatched fairly, we can set a limit of messages passed to worker
	// at any time
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	rabbitmq.FailOnError(err, "Failed to set QoS")

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
