package main

import (
	"fmt"
	"os"
	"os/signal"
    amqp "github.com/rabbitmq/amqp091-go"
	"github.com/derichard/learn-pubsub/internal/routing"
	"github.com/derichard/learn-pubsub/internal/pubsub"
	"encoding/json"
)

func main() {
	fmt.Println("Starting Peril server...")
	con_str := "amqp://guest:guest@localhost:5672/"
	con, err := amqp.Dial(con_str)
	if err != nil {
		fmt.Println("Error connecting to RabbitMQ")
	}
	defer con.Close()
	fmt.Println("Connected to RabbitMQ")

	chan, err := con.Channel()
	if err != nil {
		fmt.Println("Error creating channel")
	}
	defer chan.Close()

	// Publish a message
	val := json.Marshal(models.PlayingState{IsPaused: true})
	err = publish.PublishJSON(chan, routing.ExchangePerilDirect, routing.PauseKey, val)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	fmt.Println("Blocking, press ctrl+c to continue...")
	<-done  // Will block here until user hits ctrl+c
	fmt.Println("Exiting...")
	

}
