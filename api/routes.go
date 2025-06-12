package api

import (
	"server/internal/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) func(*fiber.App) {
	return func(app *fiber.App) {
		repo := user.NewRepository(db)
		service := user.NewService(repo)
		handler := user.NewHandler(service)

		auth := app.Group("/auth")
		auth.Post("/register", handler.Register)
		auth.Post("/login", handler.Login)
	}
}