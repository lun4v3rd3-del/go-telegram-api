package arguments

import (
	"github.com/lun4v3rd3-del/go-telegram-api/internal/entitiy"
	"github.com/lun4v3rd3-del/go-telegram-api/internal/infrastructure/interfaces"
	"regexp"
)

type HandlerArgs struct {
	Re      *regexp.Regexp
	H       interfaces.Handler
	Cd      string
	State   *entitiy.State
	Pattern string
}
