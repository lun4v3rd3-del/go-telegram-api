package usecase

import (
	"github.com/lun4v3rd3-del/go-telegram-api/internal/infrastructure/interfaces"
	"github.com/lun4v3rd3-del/go-telegram-api/internal/infrastructure/telegram"
)

type EventProcessor struct {
	client  interfaces.Client
	router  *telegram.Router
	context interfaces.Context
}

func NewEventProcessor(client interfaces.Client, router *telegram.Router, context interfaces.Context) *EventProcessor {
	return &EventProcessor{client: client, router: router, context: context}
}
