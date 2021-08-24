package main

import (
	//"errors"
	"fmt"
	//"log"
	//"os"
	"time"

	//"github.com/streadway/amqp"
	"Common"
)

// This exports a Session object that wraps this library. It
// automatically reconnects when the connection fails, and
// blocks all pushes until the connection succeeds. It also
// confirms every outgoing message, so none are lost.
// It doesn't automatically ack each message, but leaves that
// to the parent process, since it is usage-dependent.
//
// Try running this in one terminal, and `rabbitmq-server` in another.
// Stop & restart RabbitMQ to see how the queue reacts.
func main() {
	name := "job_queue"
	addr := "amqp://gauge:sparksubmit666@192.168.52.128:5672/"
	queue := Common.New(name, addr)
	message := []byte("message")
	// Attempt to push a message every 2 seconds
	for {
		time.Sleep(time.Second * 3)
		if err := queue.Push(message); err != nil {
			fmt.Printf("Push failed: %s\n", err)
		} else {
			fmt.Println("Push succeeded!")
		}
	}
}
