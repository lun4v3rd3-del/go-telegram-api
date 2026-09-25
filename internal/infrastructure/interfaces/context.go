package interfaces

import "telegram-api-service/internal/entitiy"

type Context interface {
	GetState() *entitiy.State
	SetState(*entitiy.State)
	ClearState()
}

type StatesGroup interface {
	Init()
}
