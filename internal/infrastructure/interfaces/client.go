package interfaces

import "github.com/lun4v3rd3-del/go-telegram-api/internal/entitiy"

type Client interface {
	Updates() []byte
	SendMessage(query entitiy.SendMessageQuery)
	OffsetUpdate(int64)
	AnswerCallback(entitiy.AnswerCallbackQuery)
}
