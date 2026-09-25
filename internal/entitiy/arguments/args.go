package arguments

import (
	"regexp"
	"telegram-api-service/internal/entitiy"
	"telegram-api-service/internal/infrastructure/interfaces"
)

type HandlerArgs struct {
	Re      *regexp.Regexp
	H       interfaces.Handler
	Cd      string
	State   *entitiy.State
	Pattern string
}
