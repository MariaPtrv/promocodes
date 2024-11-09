package main

import (
	handler "admin/pkg/handlers"
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

	e := echo.New()

	admin := e.Group("/admin")

	admin.GET("/", func(c echo.Context) error {
		return c.String(200, "Admin here")
	})
	promocode := admin.Group("/promocode")

	promocode.POST("/new", handler.NewPromocode)

	promocode.DELETE("/:id", handler.DeletePromocode)

	promocode.PUT("/:id", handler.UpdatePromocode)

	reward := admin.Group("/reward")

	reward.POST("/new", handler.NewReward)

	reward.DELETE("/:id", func(c echo.Context) error {
		return c.String(200, "Delete reward :id")
	})

	reward.PUT("/:id", func(c echo.Context) error {
		return c.String(200, "Update reward :id")
	})

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
