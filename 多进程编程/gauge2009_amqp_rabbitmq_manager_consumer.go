package main

import (
	"fmt"
	"log"
	"runtime"
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
	// ■ ■ ■ ■ ■ ■ ■ ■ ■ ■ ■ See the Channel.Publish example for the complimentary declare.
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
	// ■ ■ ■ ■ ■ ■ ■ ■ ■ ■ ■
	// Set our quality of service.  Since we're sharing 3 consumers on the same
	// channel, we want at least 3 messages in flight.
	err = channel.Qos(1, 0, false)
	if err != nil {
		log.Fatalf("basic.qos: %v", err)
	}

	// Establish our consumers that have different responsibilities.  Our first
	// two queues do not ack the messages on the server, so require to be acked
	// on the client.
	// ■ ■ ■ ■ ■ ■ ■ ■ ■ ■ ■
	pages, err := channel.Consume("page", "pager", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("basic.consume: %v", err)
	}

	go func() {
		for log := range pages {
			// // ■ ■ ■ ■ ■ ■ ■ ■ ■ ■ ■... this consumer is responsible for sending pages per log
			log.Ack(false)
			msg := string(log.Body)
			fmt.Printf("message form queue [emails] : %s", msg)
		}
	}()

	// Notice how the concern for which messages arrive here are in the AMQP
	// topology and not in the queue.  We let the server pick a consumer tag this
	// time.
	// ■ ■ ■ ■ ■ ■ ■ ■ ■ ■ ■
	emails, err := channel.Consume("email", "", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("basic.consume: %v", err)
	}

	go func() {
		for log := range emails {
			// // ■ ■ ■ ■ ■ ■ ■ ■ ■ ■ ■... this consumer is responsible for sending emails per log
			log.Ack(false)
			msg := string(log.Body)
			fmt.Printf("message form queue [emails] : %s", msg)
		}
	}()

	// This consumer requests that every message is acknowledged as soon as it's
	// delivered.
	// ■ ■ ■ ■ ■ ■ ■ ■ ■ ■ ■
	firehose, err := channel.Consume("firehose", "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("basic.consume: %v", err)
	}

	// To show how to process the items in parallel, we'll use a work pool.
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(work <-chan amqp.Delivery) {
			for range work { // ■ ■ ■ ■ ■ ■ ■ ■ ■ ■ ■
				// ... this consumer pulls from the firehose and doesn't need to acknowledge
			}
		}(firehose)
	}

	// Wait until you're ready to finish, could be a signal handler here.
	time.Sleep(86400 * time.Second)

	// Cancelling a consumer by name will finish the range and gracefully end the
	// goroutine
	//err = channel.Cancel("pager", false)
	//if err != nil {
	//	log.Fatalf("basic.cancel: %v", err)
	//}

	// deferred closing the Connection will also finish the consumer's ranges of
	// their delivery chans.  If you need every delivery to be processed, make
	// sure to wait for all consumers goroutines to finish before exiting your
	// process.
}
