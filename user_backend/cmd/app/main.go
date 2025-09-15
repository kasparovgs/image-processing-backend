package main

import (
	"log"
	"os"
	"time"
	"user_backend/api/http"
	appConfig "user_backend/cmd/app/config"
	"user_backend/repository/postgres_storage"
	rabbitMQ "user_backend/repository/rabbit_mq"
	"user_backend/repository/redis_storage"
	"user_backend/usecases/service"

	"pkg/config"
	pkgHttp "pkg/http"
	_ "user_backend/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

// @title My API
// @version 1.0
// @description This is a sample server.

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey SessionIDAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and your token.
func main() {
	appFlags := appConfig.ParseFlags()
	var cfg appConfig.AppConfig
	config.MustLoad(appFlags.ConfigPath, &cfg)

	connStr := os.Getenv("DB_CONN_STR")
	if connStr == "" {
		log.Fatal("DB_CONN_STR environment variable is required")
	}

	userRepo, err := postgres_storage.NewUserDB(connStr)
	if err != nil {
		log.Fatalf("no connection with postgres: %v", err)
	}

	ttl, err := time.ParseDuration(cfg.Redis.TTL)
	if err != nil {
		log.Fatalf("bad redis.ttl: %v", err)
	}
	sessionRepo := redis_storage.NewSession(cfg.Redis.Addr, ttl)

	taskSender, err := rabbitMQ.NewRabbitMQSender(cfg.RabbitMQ)
	if err != nil {
		log.Fatalf("failed creating rabbitMQ: %s", err.Error())
	}

	taskRepo, err := postgres_storage.NewTaskDB(connStr)
	if err != nil {
		log.Fatalf("no connection with postgres: %v", err)
	}
	taskService := service.NewTask(taskRepo, sessionRepo, taskSender)
	taskHandlers := http.NewTaskHandler(taskService)

	userSevice := service.NewUser(userRepo, sessionRepo)
	userHandlers := http.NewUserHandler(userSevice)

	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	taskHandlers.WithTaskHandlers(r)
	userHandlers.WithUserHandlers(r)

	log.Printf("Starting server on %s", cfg.Address)
	if err := pkgHttp.CreateAndRunServer(r, cfg.Address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
