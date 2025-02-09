package main

import (
	"gofr.dev/pkg/gofr"
	"threelayer/handler"
	"threelayer/models"
	"threelayer/services"
	"threelayer/store"
)

func main() {
	router := gofr.New()

	log := make(map[string]models.Logs)

	store := store.New(log)

	service := services.New(&store)

	handlers := handler.New(&service)

	router.GET("/log/{id}", handlers.GetByID)
	router.POST("/log", handlers.Create)

	router.Run()
}
