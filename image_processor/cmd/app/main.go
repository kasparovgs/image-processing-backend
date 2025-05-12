package main

import (
	"image_processor/cmd/app/config"
	rabbitMQ "image_processor/repository/rabbit_mq"
	"log"
)

func main() {
	appFlags := config.ParseFlags()
	var cfg config.AppConfig
	config.MustLoad(appFlags.ConfigPath, &cfg)

	taskConsumer, err := rabbitMQ.NewRabbitMQConsumer(cfg.RabbitMQ)
	if err != nil {
		log.Fatalf("failed creating rabbitMQ: %s", err.Error())
	}

	taskConsumer.Consume()
}
