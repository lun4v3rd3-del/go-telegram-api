package telegram

import (
	"fmt"
	"github.com/lun4v3rd3-del/go-telegram-api/pkg/entitiy"
	"github.com/lun4v3rd3-del/go-telegram-api/pkg/entitiy/arguments"
	"github.com/lun4v3rd3-del/go-telegram-api/pkg/infrastructure/interfaces"
	"regexp"
)

type HandlerHolder struct {
	State   *entitiy.State
	Re      *regexp.Regexp
	Cd      string
	Handler *interfaces.Handler
}

type Registrator interface {
	RegistrateHandler(re regexp.Regexp, h interfaces.Handler)
}

type Router struct {
	HandlerPool []HandlerHolder
}

func NewRouter() *Router {
	return &Router{
		HandlerPool: make([]HandlerHolder, 1),
	}
}

func (r *Router) RegisterHandler(args arguments.HandlerArgs) {
	callbackData := "none"
	if args.Cd != "" {
		callbackData = args.Cd
	}

	holder := HandlerHolder{
		Cd:      callbackData,
		Re:      args.Re,
		State:   args.State,
		Handler: &args.H,
	}

	r.HandlerPool = append(r.HandlerPool, holder)
}

func (r *Router) FindHandler(args arguments.HandlerArgs) *interfaces.Handler {
	callbackData := "none"
	if args.Cd != "" {
		callbackData = args.Cd
	}

	var hPattern *interfaces.Handler
	var hCallback *interfaces.Handler

	fmt.Println("args:", callbackData, args.Pattern)

	for _, holder := range r.HandlerPool {
		fmt.Println("holder:", holder.Cd, holder.Re)

		if holder.State != nil && holder.State == args.State {
			return holder.Handler
		}

		if callbackData != "none" && holder.Cd == callbackData {
			hCallback = holder.Handler
		}

		if holder.Re != nil && args.Pattern != "" {
			if holder.Re.MatchString(args.Pattern) {
				hPattern = holder.Handler
			}
		}
	}

	if hCallback != nil {
		return hCallback
	}
	if hPattern != nil {
		return hPattern
	}

	return nil
}
