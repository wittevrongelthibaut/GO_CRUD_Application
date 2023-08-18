package rabbitmq

import (
	"log"
	"main/src/util"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitiateConnection(url string) (*amqp.Connection) {
	conn, err := amqp.Dial(url)
	util.FailOnError(err, "Failed to connect to RabbitMQ")

	log.Printf("Connection to RabbitMQ established")

	return conn
}

func InitiateChannel(conn *amqp.Connection) (*amqp.Channel) {
	channel, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")

	log.Printf("Successfully opened a channel")

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	util.FailOnError(err, "Failed to set QoS")

	return channel
}