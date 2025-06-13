// @title Task Management API
// @version 1.0
// @description This is a sample server for managing tasks.
// @host localhost:8080
// @BasePath /v1

package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "task-service/docs" // required for swag to register docs

	"task-service/internal/api"
	"task-service/internal/db"
	"task-service/internal/repository"
	"task-service/internal/service"
)

func main() {
	// Initialize DB
	database := db.Init()

	// Initialize layers
	taskRepo := repository.NewTaskRepository(database)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := api.NewTaskHandler(taskService)

	// Setup router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// Routes
	r.Route("/v1/tasks", func(r chi.Router) {
		r.Get("/", taskHandler.ListTasks)
		r.Get("/{id}", taskHandler.GetTask)
		r.Post("/", taskHandler.CreateTask)
		r.Put("/{id}", taskHandler.UpdateTask)
		r.Delete("/{id}", taskHandler.DeleteTask)
	})

	// Start server
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
