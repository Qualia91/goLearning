package rabbitmq

import "log"

// helper function to check return value of amqp call
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
