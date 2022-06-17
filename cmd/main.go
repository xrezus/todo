package main

import (
	"github.com/xrezus/todo"
	"github.com/xrezus/todo/pkg/handler"
	"github.com/xrezus/todo/pkg/repository"
	"github.com/xrezus/todo/pkg/service"
	"log"
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
