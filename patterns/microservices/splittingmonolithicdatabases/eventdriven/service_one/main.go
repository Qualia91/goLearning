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

	// get/generate exchange for data_has_changed
	err = ch.ExchangeDeclare(
		"data_has_changed", // name
		"fanout",           // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	rabbitmq.FailOnError(err, "Failed to declare an exchange for data_has_changed")

	// get/generate exchange for other_data_has_changed
	err = ch.ExchangeDeclare(
		"other_data_has_changed", // name
		"fanout",                 // type
		true,                     // durable
		false,                    // auto-deleted
		false,                    // internal
		false,                    // no-wait
		nil,                      // arguments
	)
	rabbitmq.FailOnError(err, "Failed to declare an exchange for other_data_has_changed")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable: Queue def will survive a server restart
		false, // delete when unused
		true,  // exclusive: Used by only one connection
		false, // no-wait
		nil,   // arguments: Optional
	)
	rabbitmq.FailOnError(err, "Failed to delare a queue")

	// bind to queue data_has_changed
	err = ch.QueueBind(
		q.Name,             // queue name
		"",                 // routing key
		"data_has_changed", // exchange
		false,
		nil,
	)
	rabbitmq.FailOnError(err, "Failed to bind a queue data_has_changed")

	// bind to queue data_has_changed
	err = ch.QueueBind(
		q.Name,                   // queue name
		"",                       // routing key
		"other_data_has_changed", // exchange
		false,
		nil,
	)
	rabbitmq.FailOnError(err, "Failed to bind a queue other_data_has_changed")

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
