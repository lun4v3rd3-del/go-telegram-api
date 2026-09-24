package interfaces

import "telegram-api-service/internal/entitiy"

type Handler interface {
	Handle(event *entitiy.Event, ctx *entitiy.Context)
}
