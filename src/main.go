package main

import (
	"log"
	"main/src/rabbitmq"
	"main/src/util"
)

func main() {
	
	log.Printf("Starting Intership Service")

	util.LoadEnvVars()
	
	conn := rabbitmq.InitiateConnection(util.GetEnvVar("RABBITMQ_URL"))
	defer conn.Close()

	ch := rabbitmq.InitiateChannel(conn)
	defer ch.Close()

	log.Printf("Initialising consumers")

	rabbitmq.InitiateHandlers(ch)

	var forever chan struct{}

	log.Printf("Service up and running. Listening for messages...")

	<-forever
}
