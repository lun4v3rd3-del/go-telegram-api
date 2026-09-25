package telegram

import (
	"encoding/json"
	"telegram-api-service/internal/entitiy"
	"telegram-api-service/internal/infrastructure/interfaces"
)

type TestHandler struct {
	c *HttpClient
}

func NewTestHandler(c *HttpClient) *TestHandler {
	return &TestHandler{c: c}
}

func (g *TestHandler) Handle(event *entitiy.Event, ctx *interfaces.Context) {
	var message entitiy.Message

	if event.Message != nil {
		message = *event.Message
	} else {
		message = *event.CallbackQuery.Message
	}

	testString, _ := json.MarshalIndent(event, "", "  ")

	query := entitiy.SendMessageQuery{
		Text:   string(testString),
		ChatID: message.Chat.ID,
		ReplyMarkup: &entitiy.ReplyMarkup{
			InlineKeyboards: [][]entitiy.InlineKeyboardButton{
				{
					{
						Text:         "inline_keyboard_1",
						CallbackData: "callback_data_1",
					},
					{
						Text:         "inline_keyboard_2",
						CallbackData: "callback_data_2",
					},
				},
				{
					{
						Text:         "inline_keyboard_3",
						CallbackData: "callback_data_3",
					},
				},
			},
		},
	}
	g.c.SendMessage(query)
}

type CallbackHandler struct {
	c *HttpClient
}

func NewCallbackHandler(c *HttpClient) *CallbackHandler {
	return &CallbackHandler{c: c}
}

func (g *CallbackHandler) Handle(event *entitiy.Event, ctx *interfaces.Context) {
	defer g.c.AnswerCallback(entitiy.AnswerCallbackQuery{
		CallbackQueryID: event.CallbackQuery.ID,
	})

	query := entitiy.SendMessageQuery{
		Text:   event.CallbackQuery.Data,
		ChatID: event.CallbackQuery.Message.Chat.ID,
	}

	g.c.SendMessage(query)
}

type StateFirstHandler struct {
	c *HttpClient
}

func NewStateFirstHandler(c *HttpClient) *StateFirstHandler {
	return &StateFirstHandler{c: c}
}

func (h *StateFirstHandler) Handle(event *entitiy.Event, ctx *interfaces.Context) {
	query := entitiy.SendMessageQuery{
		Text:   "state_1_handler_response",
		ChatID: event.Message.Chat.ID,
	}

	h.c.SendMessage(query)

	(*ctx).SetState(TestGet().State1)
}

type StateSecondHandler struct {
	c *HttpClient
}

func NewStateSecondHandler(c *HttpClient) *StateSecondHandler {
	return &StateSecondHandler{c: c}
}

func (h *StateSecondHandler) Handle(event *entitiy.Event, ctx *interfaces.Context) {
	query := entitiy.SendMessageQuery{
		Text:   "state_2_handler_response",
		ChatID: event.Message.Chat.ID,
	}

	h.c.SendMessage(query)

	(*ctx).ClearState()
}
