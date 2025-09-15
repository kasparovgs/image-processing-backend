package main

import (
	appConfig "image_processor/cmd/app/config"
	rabbitMQ "image_processor/repository/rabbit_mq"
	"log"
	"pkg/config"
)

func main() {
	appFlags := appConfig.ParseFlags()
	var cfg appConfig.AppConfig
	config.MustLoad(appFlags.ConfigPath, &cfg)

	taskConsumer, err := rabbitMQ.NewRabbitMQConsumer(cfg.RabbitMQ)
	if err != nil {
		log.Fatalf("failed creating rabbitMQ: %s", err.Error())
	}

	taskConsumer.Consume()
}
