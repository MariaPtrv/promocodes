package main

import (
	handler "admin/pkg/handlers"
	"admin/pkg/repository"
	"admin/pkg/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("error occured while loading .env file: %v", err)
	}

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Env var 'PORT' must be set")
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	e := echo.New()
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
