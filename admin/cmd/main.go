package main

import (
	handler "admin/pkg/handlers"
	"admin/pkg/repository"
	"admin/pkg/service"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error occured while loading .env file: %v", err)
	}

	if err := initConfig(); err != nil {
		log.Fatalf("error occured while reading config: %s", err.Error())
	}

}

func main() {
	port := viper.GetString("port")
	if port == "" {
		log.Fatal("value of 'port' must be set in config")
	}

	db_password := os.Getenv("DB_PASSWORD")
	if db_password == "" {
		log.Fatal("value of 'DB_PASSWORD' must be set in .env")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: db_password,
		DBname:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("error occured while connecting DB: %s", err.Error())
	}

	fmt.Printf("DB %s connected %s:%s", viper.GetString("db.dbname"), viper.GetString("db.host"), viper.GetString("db.port"))

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	e := echo.New()
	e.Use(middleware.Logger())
	handlers.InitRoutes(e)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(":8000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)

	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
