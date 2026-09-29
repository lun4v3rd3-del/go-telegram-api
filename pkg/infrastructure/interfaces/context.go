package interfaces

import "github.com/lun4v3rd3-del/go-telegram-api/pkg/entitiy"

type Context interface {
	GetState() *entitiy.State
	SetState(*entitiy.State)
	ClearState()
}

type StatesGroup interface {
	Init()
}
