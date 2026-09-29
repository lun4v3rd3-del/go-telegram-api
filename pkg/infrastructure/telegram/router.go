package telegram

import (
	"fmt"
	"regexp"

	"github.com/lun4v3rd3-del/go-telegram-api/pkg/entitiy"
	"github.com/lun4v3rd3-del/go-telegram-api/pkg/entitiy/arguments"
	"github.com/lun4v3rd3-del/go-telegram-api/pkg/infrastructure/interfaces"
)

type HandlerHolder struct {
	State       *entitiy.State
	Re          *regexp.Regexp
	Cd          string
	HandlerFunc *interfaces.HandlerFunc
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
		Cd:          callbackData,
		Re:          args.Re,
		State:       args.State,
		HandlerFunc: &args.H,
	}

	r.HandlerPool = append(r.HandlerPool, holder)
}

func (r *Router) FindHandler(args arguments.HandlerArgs) *interfaces.HandlerFunc {
	callbackData := "none"
	if args.Cd != "" {
		callbackData = args.Cd
	}

	var hPattern *interfaces.HandlerFunc
	var hCallback *interfaces.HandlerFunc

	fmt.Println("args:", callbackData, args.Pattern)

	for _, holder := range r.HandlerPool {
		fmt.Println("holder:", holder.Cd, holder.Re)

		if holder.State != nil && holder.State == args.State {
			return holder.HandlerFunc
		}

		if callbackData != "none" && holder.Cd == callbackData {
			hCallback = holder.HandlerFunc
		}

		if holder.Re != nil && args.Pattern != "" {
			if holder.Re.MatchString(args.Pattern) {
				hPattern = holder.HandlerFunc
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
