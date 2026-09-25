package usecase

import (
	"log"
	"telegram-api-service/internal/entitiy"
	"telegram-api-service/internal/entitiy/arguments"
)

type Processor interface {
	Process(event entitiy.Event)
}

func (e *EventProcessor) Process(event entitiy.Event) {
	var qd, pattern string
	if event.Message != nil {
		pattern = event.Message.Text
	} else {
		qd = event.CallbackQuery.Data
	}

	h := e.router.FindHandler(arguments.HandlerArgs{
		Pattern: pattern,
		Cd:      qd,
		State:   e.context.GetState(),
	})
	if h != nil {
		(*h).Handle(&event, &e.context)
	} else {
		log.Println("No handler for event:", event.ID)
	}
}
