package usecase

import (
	"log"
	"telegram-api-service/internal/entitiy"
)

type Processor interface {
	Process(event entitiy.Event)
}

func (e *EventProcessor) Process(event entitiy.Event) {
	var re string
	var qd string
	if event.Message != nil {
		re = event.Message.Text
	} else {
		re = event.CallbackQuery.Message.Text
		qd = event.CallbackQuery.Data
	}

	h := e.router.FindHandler(re, qd)
	if h != nil {
		(*h).Handle(&event, &e.handlerContext)
	} else {
		log.Println("No handler for event:", event.ID)
	}
}
