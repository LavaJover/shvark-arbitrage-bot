package main

import (
	"fmt"
	"log"

	"github.com/LavaJover/shvark-arbitrage-bot/internal/config"
	"github.com/LavaJover/shvark-arbitrage-bot/internal/delivery/telegram"
	"github.com/LavaJover/shvark-arbitrage-bot/internal/grpcapi"
	"github.com/LavaJover/shvark-arbitrage-bot/internal/infrastructure/kafka"
	"github.com/LavaJover/shvark-arbitrage-bot/internal/infrastructure/postgres"
	"github.com/LavaJover/shvark-arbitrage-bot/internal/usecase"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load .env\n")
	}
	// read config
	cfg := config.MustLoad()

	ssoAddr :=  fmt.Sprintf("%s:%s", cfg.SSOService.Host, cfg.SSOService.Port)
	ssoClient, err := grpcapi.NewSSOClient(ssoAddr)
	if err != nil {
		log.Fatalf("failed to connect SSO-client")
	}

	authzAddr := fmt.Sprintf("%s:%s", cfg.AuthzService.Host, cfg.AuthzService.Port)
	authzClient, err := grpcapi.NewAuthzClient(authzAddr)
	if err != nil {
		log.Fatalf("failed to connect Authz-client")
	}

	db := postgres.InitDB(cfg)
	authRepo := postgres.NewDefaultAuthRepository(db)
	authUC := usecase.NewAuthUsecase(authRepo, ssoClient, authzClient)

	bot, err := telegram.NewBot(cfg.BotToken, authUC)
	if err != nil {
		log.Fatalf("failed to init bot " + err.Error())
	}

	go kafka.ListenToOrderEvents([]string{fmt.Sprintf("%s:%s", cfg.KafkaService.Host, cfg.KafkaService.Port)}, "dispute-events", bot.Notify)

	bot.Start()
}