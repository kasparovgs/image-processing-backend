package main

import (
	appConfig "image_processor/cmd/app/config"
	rabbitMQ "image_processor/repository/rabbit_mq"
	"log"
	"net/http"
	"pkg/config"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		log.Println("Metrics endpoint is listening on :9090/metrics")
		if err := http.ListenAndServe(":9090", nil); err != nil {
			log.Fatalf("metrics HTTP server failed: %v", err)
		}
	}()
}

func main() {
	appFlags := appConfig.ParseFlags()
	var cfg appConfig.AppConfig
	config.MustLoad(appFlags.ConfigPath, &cfg)

	taskConsumer, err := rabbitMQ.NewRabbitMQConsumer(cfg.RabbitMQ)
	if err != nil {
		log.Fatalf("failed creating rabbitMQ: %s", err.Error())
	}

	http.Handle("/metrics", promhttp.Handler())
	recordMetrics()

	taskConsumer.Consume()
}
