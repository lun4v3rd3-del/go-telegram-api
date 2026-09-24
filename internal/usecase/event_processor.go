package usecase

import (
	"telegram-api-service/internal/entitiy"
	"telegram-api-service/internal/infrastructure/interfaces"
	"telegram-api-service/internal/infrastructure/telegram"
)

type EventProcessor struct {
	client         interfaces.Client
	router         *telegram.Router
	handlerContext entitiy.Context
}

func NewEventProcessor(client interfaces.Client, router *telegram.Router) *EventProcessor {
	return &EventProcessor{client: client, router: router}
}
