package interfaces

import "telegram-api-service/internal/entitiy"

type Client interface {
	Updates() []byte
	SendMessage(query entitiy.SendMessageQuery)
	OffsetUpdate(int64)
}
