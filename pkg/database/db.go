package database

import (
	"log"
	"os/user"
	"server/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(cfg *config.Config) *gorm.DB{
	db, err := gorm.Open(postgres.Open(cfg.DB_DSN), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}

	db.AutoMigrate(&user.User{})
	return db
}