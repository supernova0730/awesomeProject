package main

import (
	awesomeProject "app"
	"app/pkg/handler"
	"app/pkg/repository"
	"app/pkg/service"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := initConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println("configs are initialised")

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("connected to database")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	routes := handlers.InitRoutes()

	server := new(awesomeProject.Server)
	go func() {
		err = server.Run(
			viper.GetString("server.host"),
			viper.GetString("server.port"),
			routes,
		)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	fmt.Println("app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Println("app shutting down")

	err = server.Shutdown(context.Background())
	if err != nil {
		panic(err)
	}

	err = db.Close()
	if err != nil {
		panic(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
