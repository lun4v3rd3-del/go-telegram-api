package interfaces

import "context"

type Bot interface {
	Start(context.Context) error
	Stop()
}
