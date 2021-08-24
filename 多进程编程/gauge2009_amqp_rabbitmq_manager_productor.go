package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	// Connects opens an AMQP connection from the credentials in the URL.
	addr := "amqp://gauge:sparksubmit666@192.168.52.128:5672/"
	conn, err := amqp.Dial(addr)

	if err != nil {
		log.Fatalf("connection.open: %s", err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("channel.open: %s", err)
	}

	// We declare our topology on both the publisher and consumer to ensure they
	// are the same.  This is part of AMQP being a programmable messaging model.
	//
	// See the Channel.Publish example for the complimentary declare.
	err = channel.ExchangeDeclare("logs", "topic", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("exchange.declare: %s", err)
	}

	// Establish our queue topologies that we are responsible for
	type bind struct {
		queue string
		key   string
	}

	bindings := []bind{
		{"page", "alert"},
		{"email", "info"},
		{"firehose", "#"},
	}

	for _, b := range bindings {
		_, err = channel.QueueDeclare(b.queue, true, false, false, false, nil)
		if err != nil {
			log.Fatalf("queue.declare: %v", err)
		}

		err = channel.QueueBind(b.queue, b.key, "logs", false, nil)
		if err != nil {
			log.Fatalf("queue.bind: %v", err)
		}
	}

	// Prepare this message to be persistent.  Your publishing requirements may
	// be different.
	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         []byte("Go Go AMQP!"),
	}
	// ■ ■ ■ ■ ■ ■ ■ ■ ■ ■ ■
	// This is not a mandatory delivery, so it will be dropped if there are no
	// queues bound to the logs exchange.
	//err = channel.Publish("logs", "info", false, false, msg)
	for i := 0; i <= 100000; i++ {
		err = channel.Publish("logs", "alert", false, false, msg)
		if err != nil {
			// Since publish is asynchronous this can happen if the network connection
			// is reset or if the server has run out of resources.
			log.Fatalf("basic.publish: %v", err)
		}
	}

	time.Sleep(10 * time.Second)

}
