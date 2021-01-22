package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Producer connected to rbmq")

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare("go", false, false, false, false, nil)

	if err != nil {
		log.Fatal(err)
	}
	var counter int = 0
	for {
		err := ch.Publish("", q.Name, false, false,
			amqp.Publishing{
				Headers:     nil,
				ContentType: "text/plain",
				Body:        []byte("sent at" + time.Now().String()),
			})
		if err != nil {
			break
		}

		time.Sleep(2 * time.Second)
		counter = counter + 1

		if counter > 10 {
			fmt.Println("Finish service")
			break
		}
	}
}
