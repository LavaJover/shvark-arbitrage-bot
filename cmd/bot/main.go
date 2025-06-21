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
)

func main() {
	// read config
	cfg := config.MustLoad()
	fmt.Println(cfg)

	ssoAddr := "localhost:50051"
	ssoClient, err := grpcapi.NewSSOClient(ssoAddr)
	if err != nil {
		log.Fatalf("failed to connect SSO-client")
	}

	authzAddr := "localhost:50054"
	authzClient, err := grpcapi.NewAuthzClient(authzAddr)
	if err != nil {
		log.Fatalf("failed to connect Authz-client")
	}

	db := postgres.InitDB(cfg)
	authRepo := postgres.NewDefaultAuthRepository(db)
	authUC := usecase.NewAuthUsecase(authRepo, ssoClient, authzClient)

	bot, err := telegram.NewBot("8035564137:AAFpzyygekkZ_43oM-TSj-IpMGiJATaZm50", authUC)
	if err != nil {
		log.Fatalf("failed to init bot")
	}

	go kafka.ListenToOrderEvents([]string{"localhost:9092"}, "dispute-events", bot.Notify)

	bot.Start()
}