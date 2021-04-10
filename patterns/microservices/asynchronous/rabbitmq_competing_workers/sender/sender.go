package main

import (
	"fmt"
	rabbitmq "microservices/asynchronous/rabbitmq_event"
	"strconv"

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
		true,    // durable: Queue def will survive a server restart
		false,   // delete when unused
		false,   // exclusive: Used by only one connection
		false,   // no-wait
		nil,     // arguments: Optional
	)
	rabbitmq.FailOnError(err, "Failed to delare a queue")

	for {

		// get user input
		fmt.Println("Enter amount of jobs: ")
		var input string
		_, err = fmt.Scanln(&input)
		rabbitmq.FailOnError(err, "Failed on user input")

		// convert input to integer
		numberOfJobs, err := strconv.ParseInt(input, 10, 32)
		rabbitmq.FailOnError(err, "Failed on job string convert")

		// send that amount of jobs
		for i := 0; i < int(numberOfJobs); i++ {
			// publish
			err = ch.Publish(
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					DeliveryMode: amqp.Persistent, // messages wont be lost is queue gets restarted
					ContentType:  "text/plain",
					Body:         []byte(fmt.Sprintf("Job number %v\n", i)),
				},
			)
			rabbitmq.FailOnError(err, "Failed to publish message")
		}

	}

}
