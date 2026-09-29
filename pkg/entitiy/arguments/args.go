package arguments

import (
	"regexp"

	"github.com/lun4v3rd3-del/go-telegram-api/pkg/entitiy"
	"github.com/lun4v3rd3-del/go-telegram-api/pkg/infrastructure/interfaces"
)

type HandlerArgs struct {
	Re      *regexp.Regexp
	H       interfaces.HandlerFunc
	Cd      string
	State   *entitiy.State
	Pattern string
}
