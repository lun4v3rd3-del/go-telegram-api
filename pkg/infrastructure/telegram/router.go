package telegram

import (
	"fmt"
	"regexp"
	"strings"

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
	var hPattern *interfaces.HandlerFunc

	fmt.Println("args: pattern =", args.Pattern, "| cd =", args.Cd)

	for _, holder := range r.HandlerPool {
		fmt.Println("holder: pattern =", holder.Re, "| cd =", holder.Cd)

		if holder.State != nil && args.State != nil && holder.State == args.State {

			if args.Cd != "" && args.Cd != "none" && holder.Cd != "" && strings.HasPrefix(args.Cd, holder.Cd) {
				return holder.HandlerFunc
			}

			if holder.Re != nil && args.Pattern != "" && args.Pattern != "none" && holder.Re.MatchString(args.Pattern) {
				return holder.HandlerFunc
			}

			if holder.Cd == "" && holder.Re == nil {
				if strings.HasPrefix(args.Pattern, "/") || strings.HasPrefix(args.Cd, "/") {
					continue
				}
				return holder.HandlerFunc
			}

			continue
		}

		if holder.State == nil {
			if args.Cd != "" && args.Cd != "none" && holder.Cd != "" && strings.HasPrefix(args.Cd, holder.Cd) {
				return holder.HandlerFunc
			}

			var matched bool

			userText := args.Pattern
			if userText == "" || userText == "none" {
				userText = args.Cd
			}

			if holder.Re != nil && userText != "" && userText != "none" {
				matched = holder.Re.MatchString(userText)
			}

			if !matched && holder.Cd != "" && holder.Cd != "none" && userText != "" && userText != "none" {
				if rx, err := regexp.Compile(holder.Cd); err == nil {
					matched = rx.MatchString(userText)
				} else {
					matched = (holder.Cd == userText)
				}
			}

			if matched {
				if hPattern == nil {
					hPattern = holder.HandlerFunc
				}
			}
		}
	}

	if hPattern != nil {
		return hPattern
	}

	return nil
}
