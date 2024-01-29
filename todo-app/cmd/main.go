package main

import (
	"log"

	todoapp "github.com/spargapees/todo-app/todo-app"
	"github.com/spargapees/todo-app/todo-app/pkg/handler"
	"github.com/spargapees/todo-app/todo-app/pkg/repository"
	"github.com/spargapees/todo-app/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error intializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todoapp.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("todo-app/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
