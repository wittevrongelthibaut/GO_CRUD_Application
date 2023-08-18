package recommendation

import (
	"context"
	"encoding/json"
	db "main/src/data"
	"main/src/util"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)


func RetrieveRecommendationFromRPC(studentId int, ch *amqp.Channel) RecommendationReadDto {

	studentMarks := db.GetMarkFromStudent(studentId)
	request, _ := json.MarshalIndent(studentMarks, "", "  ")
	callbackQueue, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		true,  // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	util.FailOnError(err, "Failed to declare a queue")

	defer ch.QueueDelete(callbackQueue.Name, false, false, false)

	msgs, err := ch.Consume(
		callbackQueue.Name, // queue
		"",                 // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	util.FailOnError(err, "Failed to register a consumer")

	corrId := util.RandomString(32)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",                     // exchange
		util.GetEnvVar("RECOMMENDATION_REQUEST_QUEUE"), // routing key
		false,                  // mandatory
		false,                  // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			ReplyTo:       callbackQueue.Name,
			Body:          request,
		})
	util.FailOnError(err, "Failed to publish a message")

	for {
		select {
		case <-ctx.Done():
			return RecommendationReadDto{}
		case d := <-msgs:
			if corrId == d.CorrelationId {
				var recommendation RecommendationReadDto
				err = json.Unmarshal(d.Body, &recommendation)
				util.FailOnError(err, "Failed to unmarshal a message")
				return recommendation
			}
		}
	}
}