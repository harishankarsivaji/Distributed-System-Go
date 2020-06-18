package Distributed_System_Go

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {

}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	//establishing connection with RabbitMQ
	conn, err := amqp.Dial("amqp://guest@locahost:15672")
	failOnError(err, "Failed to connect to RabbitMQ")

	// opening a Channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a Channel")

	//
	q, err := ch.QueueDeclare("Hello world",
		false,
		false,
		false,
		false,
		nil)
	failOnError(err, "Failed to declare a queue")

	return conn, ch, &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", err, msg))
	}
}
