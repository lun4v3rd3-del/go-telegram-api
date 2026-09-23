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
	query := entitiy.SendMessageQuery{
		Text:   event.Msg.Text,
		ChatID: event.Msg.Chat.ID,
		ReplyMarkup: entitiy.ReplyMarkup{
			InlineKeyboards: [][]entitiy.InlineKeyboardButton{
				{
					{
						Text:         "inline_keyboard_1",
						CallbackData: "callback_data_1",
					},
					{
						Text:         "inline_keyboard_1",
						CallbackData: "callback_data_1",
					},
				},
				{
					{
						Text:         "inline_keyboard_1",
						CallbackData: "callback_data_1",
					},
				},
			},
		},
	}
	fmt.Print("GreetHandler handle")
	g.c.SendMessage(query)
}
