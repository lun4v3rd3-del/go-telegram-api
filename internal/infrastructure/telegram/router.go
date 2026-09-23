package telegram

import (
	"regexp"
	"telegram-api-service/internal/infrastructure/interfaces"
)

type Registrator interface {
	RegistrateHandler(re regexp.Regexp, h interfaces.Handler)
}

type Router struct {
	HandlerPool map[string]interfaces.Handler
}

func NewRouter() *Router {
	return &Router{
		HandlerPool: make(map[string]interfaces.Handler),
	}
}

func (r *Router) RegisterHandler(re *regexp.Regexp, h interfaces.Handler) {
	r.HandlerPool[re.String()] = h
}
