package postgres

import (
	"log"

	"github.com/LavaJover/shvark-arbitrage-bot/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.ArbitrageBotConfig) *gorm.DB {
	dsn := cfg.ArbitrageBotDB.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to init database")
	}

	db.AutoMigrate(&TelegramBinding{})

	return db
}