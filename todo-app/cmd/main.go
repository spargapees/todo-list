package main

import (
	"log"

	todoapp "github.com/spargapees/todo-app/todo-app"
	"github.com/spargapees/todo-app/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todoapp.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
