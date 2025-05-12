package main

import (
	"log"
	"user_backend/api/http"
	"user_backend/cmd/app/config"
	rabbitMQ "user_backend/repository/rabbit_mq"
	"user_backend/repository/ram_storage"
	"user_backend/usecases/service"

	pkgHttp "pkg/http"
	_ "user_backend/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
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
	appFlags := config.ParseFlags()
	var cfg config.AppConfig
	config.MustLoad(appFlags.ConfigPath, &cfg)

	userRepo := ram_storage.NewUserDB()
	sessionRepo := ram_storage.NewSession()
	taskSender, err := rabbitMQ.NewRabbitMQSender(cfg.RabbitMQ)
	if err != nil {
		log.Fatalf("failed creating rabbitMQ: %s", err.Error())
	}

	taskRepo := ram_storage.NewTaskDB()
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
