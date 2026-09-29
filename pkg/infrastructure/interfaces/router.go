package interfaces

import "regexp"

type Router interface {
	RegisterHandler(re *regexp.Regexp, h Handler)
}
