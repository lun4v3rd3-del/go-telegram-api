package interfaces

import (
	"github.com/lun4v3rd3-del/go-telegram-api/pkg/entitiy"
)

type Handler interface {
	Handle(event *entitiy.Event, ctx *Context, client Client)
}
