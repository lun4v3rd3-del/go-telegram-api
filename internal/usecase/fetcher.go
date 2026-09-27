package usecase

import (
	"bytes"
	"encoding/json"
	"telegram-api-service/internal/entitiy"
)

type Fetcher interface {
	Fetch() ([]entitiy.Event, error)
}

func (e *EventProcessor) Fetch() ([]entitiy.Event, error) {
	raw_events := bytes.TrimRight(e.client.Updates(), "\x00")

	var response entitiy.Response

	if err := json.Unmarshal(raw_events, &response); err != nil {
		return nil, err
	}

	var maxOffset int64 = -1
	for _, event := range response.Events {
		if maxOffset < event.ID {
			maxOffset = event.ID
		}
	}
	e.client.OffsetUpdate(maxOffset)

	return response.Events, nil
}
