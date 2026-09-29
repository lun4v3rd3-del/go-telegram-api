package usecase

import (
	"telegram-api-service/internal/infrastructure/interfaces"
	"telegram-api-service/internal/infrastructure/telegram"
)

type EventProcessor struct {
	client  interfaces.Client
	router  *telegram.Router
	context interfaces.Context
}

func NewEventProcessor(client interfaces.Client, router *telegram.Router, context interfaces.Context) *EventProcessor {
	return &EventProcessor{client: client, router: router, context: context}
}
