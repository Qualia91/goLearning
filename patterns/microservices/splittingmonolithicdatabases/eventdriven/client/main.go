package main

import (
	"fmt"
	rabbitmq "microservices/asynchronous/rabbitmq_event"

	"github.com/streadway/amqp"
)

func main() {

	// data and other_data are saved here
	fmt.Println("data has changed ")
	fmt.Println("other_data has changed ")

	// connect to server
	con, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	rabbitmq.FailOnError(err, "Failed to connect to RabbitMQ")
	defer con.Close()

	// create channel
	ch, err := con.Channel()
	rabbitmq.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// create exchange for data change
	err = ch.ExchangeDeclare(
		"data_has_changed", // name
		"fanout",           // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	rabbitmq.FailOnError(err, "Failed to declare an exchange for data change")

	// create exchange for other_data change
	err = ch.ExchangeDeclare(
		"other_data_has_changed", // name
		"fanout",                 // type
		true,                     // durable
		false,                    // auto-deleted
		false,                    // internal
		false,                    // no-wait
		nil,                      // arguments
	)
	rabbitmq.FailOnError(err, "Failed to declare an exchange for other data change")

	// publish data change
	err = ch.Publish(
		"data_has_changed", // exchange
		"",                 // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Data has changed"),
		},
	)
	rabbitmq.FailOnError(err, "Failed to publish message")

	// publish other_data change
	err = ch.Publish(
		"other_data_has_changed", // exchange
		"",                       // routing key
		false,                    // mandatory
		false,                    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Other Data has changed"),
		},
	)
	rabbitmq.FailOnError(err, "Failed to publish message")

}
