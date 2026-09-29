package usecase

import (
	"github.com/lun4v3rd3-del/go-telegram-api/pkg/entitiy"
	"github.com/lun4v3rd3-del/go-telegram-api/pkg/entitiy/arguments"
	"log"
)

type Processor interface {
	Process(event entitiy.Event) error
}

func (e *EventProcessor) Process(event entitiy.Event) error {
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
		(*h).Handle(&event, &e.context, e.client)
	} else {
		log.Println("No handler for event:", event.ID)
	}
	return nil
}
