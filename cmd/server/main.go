package main

import (
	"fmt"
	"os"
	"os/signal"
    amqp "github.com/rabbitmq/amqp091-go"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
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

	ch, err := con.Channel()
	if err != nil {
		fmt.Println("Error creating channel")
	}
	defer ch.Close()

	// Publish a message
	val, _ := json.Marshal(routing.PlayingState{IsPaused: true})
	err = pubsub.PublishJSON(ch, routing.ExchangePerilDirect, routing.PauseKey, val)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	fmt.Println("Blocking, press ctrl+c to continue...")
	<-done  // Will block here until user hits ctrl+c
	fmt.Println("Exiting...")
	

}
