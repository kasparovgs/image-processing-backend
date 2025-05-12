package rabbitMQ

import (
	"encoding/json"
	"fmt"
	"user_backend/domain"

	"image_processor/cmd/app/config"
	"image_processor/usecases/service"

	"github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queueName  string
}

func NewRabbitMQConsumer(cfg config.RabbitMQ) (*RabbitMQConsumer, error) {
	url := fmt.Sprintf("amqp://guest:guest@%s:%d", cfg.Host, cfg.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("connecting to rabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		cfg.QueueName, // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQConsumer{
		connection: conn,
		channel:    ch,
		queueName:  cfg.QueueName,
	}, nil
}

func (r *RabbitMQConsumer) Consume() error {
	msgs, err := r.channel.Consume(
		r.queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			var task domain.Task
			if err := json.Unmarshal(msg.Body, &task); err != nil {
				continue
			}
			if err := service.ProcessTask(&task); err != nil {
				fmt.Println("Incorrect image")
			}
			if err = service.CommitTask(&task); err != nil {
				fmt.Println("Error while commiting")
			}
		}
	}()
	<-forever

	return nil
}

func (r *RabbitMQConsumer) Close() {
	r.channel.Close()
	r.connection.Close()
}
