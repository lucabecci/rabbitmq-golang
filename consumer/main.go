package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Consumer connected to rbmq")

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	chMessages, err := ch.Consume("go", "", true, false, false, false, nil)

	if err != nil {
		log.Fatal(err)
	}

	noStop := make(chan bool)

	go func() {
		for messages := range chMessages {
			fmt.Println("message:" + string(messages.Body))
		}
	}()

	<-noStop
}
