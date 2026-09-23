package usecase

import (
	"fmt"
	"log"
	"maps"
	"regexp"
	"telegram-api-service/internal/entitiy"
	"telegram-api-service/internal/infrastructure/interfaces"
)

type Processor interface {
	Process(event entitiy.Event, handlerPool map[string]interfaces.Handler)
}

func (e *EventProcessor) Process(event entitiy.Event, handlerPool map[string]interfaces.Handler) {
	re := event.Msg.Text

	for key := range maps.Keys(handlerPool) {
		ok, err := regexp.MatchString(key, re)
		fmt.Printf("matching: %s %s %f\n", re, key, ok)

		if err != nil {
			log.Fatal(err)
		}

		if ok {
			handlerPool[key].Handle(&event)
		}
	}
}
