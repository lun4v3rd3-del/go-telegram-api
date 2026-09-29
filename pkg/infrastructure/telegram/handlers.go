package telegram

import (
	"encoding/json"

	"github.com/lun4v3rd3-del/go-telegram-api/pkg/entitiy"
	"github.com/lun4v3rd3-del/go-telegram-api/pkg/infrastructure/interfaces"
)

func TestHandler(event *entitiy.Event, ctx *interfaces.Context, client interfaces.Client) {
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
	client.SendMessage(query)
}

func CallbackHandler(event *entitiy.Event, ctx *interfaces.Context, client interfaces.Client) {
	defer client.AnswerCallback(entitiy.AnswerCallbackQuery{
		CallbackQueryID: event.CallbackQuery.ID,
	})

	query := entitiy.SendMessageQuery{
		Text:   event.CallbackQuery.Data,
		ChatID: event.CallbackQuery.Message.Chat.ID,
	}

	client.SendMessage(query)
}

func StateFirstHandler(event *entitiy.Event, ctx *interfaces.Context, client interfaces.Client) {
	query := entitiy.SendMessageQuery{
		Text:   "state_1_handler_response",
		ChatID: event.Message.Chat.ID,
	}

	client.SendMessage(query)

	(*ctx).SetState(TestGet().State1)
}

func StateSecondHandler(event *entitiy.Event, ctx *interfaces.Context, client interfaces.Client) {
	query := entitiy.SendMessageQuery{
		Text:   "state_2_handler_response",
		ChatID: event.Message.Chat.ID,
	}

	client.SendMessage(query)

	(*ctx).ClearState()
}
