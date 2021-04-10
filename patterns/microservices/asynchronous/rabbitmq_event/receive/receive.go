package main

import (
	"log"

	rabbitmq "microservices/asynchronous/rabbitmq_event"

	"sync"

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
		false,   // durable: Queue def will survive a server restart
		false,   // delete when unused
		false,   // exclusive: Used by only one connection
		false,   // no-wait
		nil,     // arguments: Optional
	)
	rabbitmq.FailOnError(err, "Failed to delare a queue")

	// read from queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	rabbitmq.FailOnError(err, "Failed to delare a queue")

	// create wait group so programme doesn't end after following go func is created
	var wg sync.WaitGroup
	go func(msgs <-chan amqp.Delivery, wg *sync.WaitGroup) {
		defer wg.Done()
		wg.Add(1)

		for d := range msgs {
			log.Printf("Received message: %s\n", d.Body)
		}

	}(msgs, &wg)

	// wait until go func finishes
	wg.Wait()

}
