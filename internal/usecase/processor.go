package usecase

import (
	"log"
	"maps"
	"regexp"
	"telegram-api-service/internal/entitiy"
)

type Processor interface {
	Process(event entitiy.Event)
}

func (e *EventProcessor) Process(event entitiy.Event) {
	re := event.Msg.Text

	for key := range maps.Keys(e.router.HandlerPool) {
		ok, err := regexp.MatchString(key, re)

		if err != nil {
			log.Fatal(err)
		}

		if ok {
			e.router.HandlerPool[key].Handle(&event)
		}
	}
}
