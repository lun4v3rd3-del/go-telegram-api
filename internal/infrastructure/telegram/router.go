package telegram

import (
	"maps"
	"regexp"
	"telegram-api-service/internal/infrastructure/interfaces"
)

type Registrator interface {
	RegistrateHandler(re regexp.Regexp, h interfaces.Handler)
}

type Router struct {
	HandlerPool map[string]map[string]interfaces.Handler
}

func NewRouter() *Router {
	return &Router{
		HandlerPool: make(map[string]map[string]interfaces.Handler),
	}
}

func (r *Router) RegisterHandler(re *regexp.Regexp, h interfaces.Handler, qd ...string) {
	callbackData := "none"
	if len(qd) > 0 && qd[0] != "" {
		callbackData = qd[0]
	}

	pattern := re.String()

	if r.HandlerPool[pattern] == nil {
		r.HandlerPool[pattern] = make(map[string]interfaces.Handler)
	}
	r.HandlerPool[pattern][callbackData] = h
}

func (r *Router) FindHandler(re string, qd ...string) *interfaces.Handler {
	var callbackData = "none"
	if len(qd) > 0 && qd[0] != "" {
		callbackData = qd[0]
	}

	for key := range maps.Keys(r.HandlerPool) {
		reg, err := regexp.Compile(key)
		if err != nil {
			continue
		}
		if reg.MatchString(re) {
			handlerFamily := r.HandlerPool[key]
			handler, ok := handlerFamily[callbackData]
			if ok {
				return &handler
			}
			break
		}
	}
	return nil
}
