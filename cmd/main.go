package main

import (
	"log"

	"github.com/GJLuCkY/todo-app"
	"github.com/GJLuCkY/todo-app/pkg/handler"
	"github.com/GJLuCkY/todo-app/pkg/repository"
	"github.com/GJLuCkY/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
