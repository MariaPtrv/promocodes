package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	handler "promocodes/pkg/handlers"
	"promocodes/pkg/repository"
	"promocodes/pkg/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func init() {
	err := godotenv.Load("./configs/.env")
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

	ctxBg, cancel := context.WithCancel(context.Background())
	defer cancel()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} method: ${method} uri: ${uri} status: ${status} ${error}\n",
	}))

	e.Use(middleware.CORS())

	log.Printf("DB %s connected %s:%s", viper.GetString("db.dbname"), viper.GetString("db.host"), viper.GetString("db.port"))

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes(e)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()

	if err := e.Shutdown(ctxBg); err != nil {
		e.Logger.Fatal(err)
	}

	if err := db.Close(); err != nil {
		e.Logger.Fatal("error occured while db shutting down:", err)
	} else {
		e.Logger.Info("db shut down")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
