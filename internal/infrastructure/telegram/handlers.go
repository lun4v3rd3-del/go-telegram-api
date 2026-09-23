package telegram

import (
	"fmt"
	"telegram-api-service/internal/entitiy"
)

type GreetHandler struct {
	c *HttpClient
}

func NewGreetHandler(c *HttpClient) *GreetHandler {
	return &GreetHandler{c: c}
}

func (g *GreetHandler) Handle(event *entitiy.Event) {
	fmt.Print("GreetHandler handle")
	g.c.SendMessage(event.Msg.Chat.ID, event.Msg.Text)
}
