package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	go client()
	go server()

	var a string
	fmt.Scanln(&a)
}

func client() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil)

	failOnError(err, "Client queue failed")

	for msg := range msgs {
		log.Printf("Received messages from body: %s", msg.Body)
	}
}

func server() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello RabbitMQ"),
	}
	for {
		ch.Publish("", q.Name, false, false, msg)
	}

}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	//establishing connection with RabbitMQ
	conn, err := amqp.Dial("amqp://guest@localhost:5672")
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
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
