package entitiy

type SendMessageQuery struct {
	ChatID      int64       `json:"chat_id"`
	Text        string      `json:"text"`
	ReplyMarkup ReplyMarkup `json:"reply_markup"`
}

type ReplyMarkup struct {
	InlineKeyboards [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string            `json:"text"`
	URL          string            `json:"url"`
	CallbackData string            `json:"callback_data"`
	CopyText     map[string]string `json:"copy_text"`
}
