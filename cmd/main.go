package main

import (
	"log"
	"server/pkg/config"

	"github.com/gofiber/fiber"
)

func main() {
	cfg := config.LoadEnv()
	app := fiber.New()

	log.Fatal(app.Listen(":" + cfg.Port))
}