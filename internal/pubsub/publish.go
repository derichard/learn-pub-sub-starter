package pubsub

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"encoding/json"
	"fmt"
)


func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	ctx := context.Background()
	body, _ := json.Marshal(val)
	msg := amqp.Publishing{
		ContentType: "application/json",
		Body: body,
	}
	err := ch.PublishWithContext(ctx, exchange, key, false, false, msg)
	if err != nil {
		fmt.Println("Error publishing message")
	}
	return err
}

