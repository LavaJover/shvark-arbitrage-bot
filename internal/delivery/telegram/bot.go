package telegram

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/LavaJover/shvark-arbitrage-bot/internal/domain"
	"github.com/LavaJover/shvark-arbitrage-bot/internal/infrastructure/kafka"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api *tgbotapi.BotAPI
	authUC domain.AuthUsecase
	disputeChan chan domain.DisputeNotification
}

func NewBot(botToken string, authUC domain.AuthUsecase) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}
	return &Bot{
		api: api,
		authUC: authUC,
		disputeChan: make(chan domain.DisputeNotification, 100),
	}, nil
}

func (b *Bot) Start(){
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.api.GetUpdatesChan(u)

	go b.listenForNotifications()

	for update := range updates {
		if update.Message != nil {
			handleMessage(b.api, update.Message, b.authUC)
		}
	}
}

func (b *Bot) Notify(event kafka.DisputeEvent) {
	dispute := domain.DisputeNotification{
		DisputeID: event.DisputeID,
		OrderID: event.OrderID,
		TraderID: event.TraderID,
		ProofUrl: event.ProofUrl,
		Reason: event.Reason,
		Status: event.Status,
		OrderAmountFiat: event.OrderAmountFiat,
		DisputeAmountFiat: event.DisputeAmountFiat,
		BankName: event.BankName,
		Phone: event.Phone,
		CardNumber: event.CardNumber,
		Owner: event.Owner,
	}
	b.disputeChan <- dispute
}

func (b *Bot) listenForNotifications() {
	for order := range b.disputeChan {
		telegramIDs, err := b.authUC.GetTelegramIDs()
		slog.Info("in listenNotifications", "DisputeID", order.DisputeID)
		if err != nil {
			log.Printf("Failed to get telegram bindings for trader %s\n", order.TraderID)
		}
		for _, telegramID := range telegramIDs {
			text := order.String()
			msg := tgbotapi.NewMessage(telegramID, text)
			fmt.Println(msg)
			_, err := b.api.Send(msg)
			if err != nil {
				log.Printf("Failed to send TG message: %v\n", err)
			}
		}
	}
}