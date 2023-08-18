package internship

import (
	"context"
	"fmt"
	"log"
	"main/src/util"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateContracts(ch *amqp.Channel, internshipId int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := fmt.Sprintf("{\"id\": %d}", internshipId)
	fmt.Println(request)
	err := ch.PublishWithContext(ctx,
		"microservice.fileserver",        // exchange
		"contracts.generate", // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			Body: []byte(request),
		})
	util.FailOnError(err, "Failed to publish a message")
	log.Println("published a contract request")
}