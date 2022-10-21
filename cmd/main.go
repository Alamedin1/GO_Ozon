package main

import (
	"log"
	goozon "sample-app"
	"sample-app/pkg/handler"
	"sample-app/pkg/repository"
	"sample-app/pkg/service"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error initialize configs: ", err.Error())
	}

	var db repository.UrlList

	repo := repository.NewRepo(db)
	services := service.NewService(repo, viper.GetInt("uniquestr.len"), []rune(viper.GetString("uniquestr._rune")))
	handlers := handler.NewHandler(services)

	srv := new(goozon.Server)
	if err := srv.Run(viper.GetString("port"), handlers.Routes()); err != nil {
		log.Fatal("error occured while running http server", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
