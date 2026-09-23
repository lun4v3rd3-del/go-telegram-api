package usecase

import (
	"bytes"
	"encoding/json"
	"telegram-api-service/internal/entitiy"
)

type Fetcher interface {
	Fetch() []entitiy.Event
}

func (e *EventProcessor) Fetch() []entitiy.Event {
	raw_events := bytes.TrimRight(e.client.Updates(), "\x00")

	var response entitiy.Response

	if err := json.Unmarshal(raw_events, &response); err != nil {
		panic(err)
	}

	return response.Events
}
