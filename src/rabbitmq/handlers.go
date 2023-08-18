package rabbitmq

import (
	"context"
	"encoding/json"
	"log"
	"main/src/domain"
	"main/src/rabbitmq/handlers"
	"main/src/util"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitiateHandlers(ch *amqp.Channel) {
	studentmsgs, err := ch.Consume(util.GetEnvVar("STUDENT_REQUEST_QUEUE"), util.GetEnvVar("STUDENT_REQUEST_CONSUMER_TAG"), false, false, false, false, nil)
	util.FailOnError(err, "Failed to register Student consumer")
	companymsgs, err := ch.Consume(util.GetEnvVar("COMPANY_REQUEST_QUEUE"), util.GetEnvVar("COMPANY_REQUEST_CONSUMER_TAG"), false, false, false, false, nil)
	util.FailOnError(err, "Failed to register Company consumer")
	internshipmsgs, err := ch.Consume(util.GetEnvVar("INTERNSHIP_REQUEST_QUEUE"), util.GetEnvVar("INTERNSHIP_REQUEST_CONSUMER_TAG"), false, false, false, false, nil)
	util.FailOnError(err, "Failed to register Internship consumer")
	applicationmsgs, err := ch.Consume(util.GetEnvVar("APPLICATION_REQUEST_QUEUE"), util.GetEnvVar("APPLICATION_REQUEST_CONSUMER_TAG"), false, false, false, false, nil)
	util.FailOnError(err, "Failed to register Application consumer")
	teachermsgs, err := ch.Consume(util.GetEnvVar("TEACHER_REQUEST_QUEUE"), util.GetEnvVar("TEACHER_REQUEST_CONSUMER_TAG"), false, false, false, false, nil)
	util.FailOnError(err, "Failed to register Teacher consumer")
	managermsgs, err := ch.Consume(util.GetEnvVar("MANAGER_REQUEST_QUEUE"), util.GetEnvVar("MANAGER_REQUEST_CONSUMER_TAG"), false, false, false, false, nil)
	util.FailOnError(err, "Failed to register Manager consumer")
	chatmessagemsgs, err := ch.Consume(util.GetEnvVar("CHATMESSAGE_REQUEST_QUEUE"), util.GetEnvVar("CHATMESSAGE_REQUEST_CONSUMER_TAG"), false, false, false, false, nil)
	util.FailOnError(err, "Failed to register ChatMessage consumer")
	categorymsgs, err := ch.Consume(util.GetEnvVar("CATEGORY_REQUEST_QUEUE"), util.GetEnvVar("CATEGORY_REQUEST_CONSUMER_TAG"), false, false, false, false, nil)
	util.FailOnError(err, "Failed to register Category consumer")

	log.Printf("Successfully initialised consumers")

	go studentHandler(studentmsgs, ch)
	go companyHandler(companymsgs, ch)
	go internshipHandler(internshipmsgs, ch)
	go applicationHandler(applicationmsgs, ch)
	go teacherHandler(teachermsgs, ch)
	go managerHandler(managermsgs, ch)
	go chatMessageHandler(chatmessagemsgs, ch)
	go categoryHandler(categorymsgs, ch)
}

func studentHandler(msgs <-chan amqp.Delivery, ch *amqp.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {
			log.Println("received a student request")
			message := domain.Message{}
			json.Unmarshal(d.Body, &message)
			response, _ := json.Marshal(handlers.HandleStudentMessage(message, ch)) 
			err := ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          response,
				})
			util.FailOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
}

func companyHandler(msgs <-chan amqp.Delivery, ch *amqp.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {
			log.Println("received a company request")
			message := domain.Message{}
			json.Unmarshal(d.Body, &message)
			response, _ := json.Marshal(handlers.HandleCompanyMessage(message))
			err := ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          response,
				})
			util.FailOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
}

func internshipHandler(msgs <-chan amqp.Delivery, ch *amqp.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {
			log.Println("received an internship request")
			message := domain.Message{}
			json.Unmarshal(d.Body, &message)
			response, _ := json.Marshal(handlers.HandleInternshipMessage(message, ch)) 
			err := ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          response,
				})
			util.FailOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
}

func applicationHandler(msgs <-chan amqp.Delivery, ch *amqp.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {
			log.Println("received an application request")
			message := domain.Message{}
			json.Unmarshal(d.Body, &message)
			response, _ := json.Marshal(handlers.HandleApplicationMessage(message))
			err := ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          response,
				})
			util.FailOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
}

func teacherHandler(msgs <-chan amqp.Delivery, ch *amqp.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {
			log.Println("received a teacher request")
			message := domain.Message{}
			json.Unmarshal(d.Body, &message)
			response, _ := json.Marshal(handlers.HandleTeacherMessage(message))
			err := ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          response,
				})
			util.FailOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
}

func managerHandler(msgs <-chan amqp.Delivery, ch *amqp.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {
			log.Println("received a manager request")
			message := domain.Message{}
			json.Unmarshal(d.Body, &message)
			response, _ := json.Marshal(handlers.HandleManagerMessage(message))
			err := ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          response,
				})
			util.FailOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
}

func chatMessageHandler(msgs <-chan amqp.Delivery, ch *amqp.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {
			log.Println("received a chat message")
			message := domain.Message{}
			json.Unmarshal(d.Body, &message)
			response, _ := json.Marshal(handlers.HandleChatMessage(message))
			err := ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          response,
				})
			util.FailOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
}

func categoryHandler(msgs <-chan amqp.Delivery, ch *amqp.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {
			log.Println("received a category request")
			message := domain.Message{}
			json.Unmarshal(d.Body, &message)
			response, _ := json.Marshal(handlers.HandleCategoryMessage(message))
			err := ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          response,
				})
			util.FailOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
}