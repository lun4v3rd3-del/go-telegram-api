package usecase

import (
	"telegram-api-service/internal/infrastructure/interfaces"
)

type EventProcessor struct {
	client interfaces.Client
}

func NewEventProcessor(client interfaces.Client) *EventProcessor {
	return &EventProcessor{client: client}
}
