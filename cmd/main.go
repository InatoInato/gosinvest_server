package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"server/api"
	"server/pkg/config"
	"server/pkg/database"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadEnv()
	db := database.InitDb(cfg)
	app := fiber.New()

	router := api.SetupRoutes(db)
	router(app)

	go func() {
		listenAddr := ":" + cfg.Port
		log.Printf("Server will started in port: %s", listenAddr)
		if err := app.Listen(listenAddr); err != nil {
			log.Fatalf("Fiber cannot to run: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Got a signal. Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Error while works Fiber: %v", err)
	}

	log.Println("Server stopped successful")
}
