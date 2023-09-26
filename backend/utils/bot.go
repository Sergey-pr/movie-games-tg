package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Sergey-pr/movie-games-tg/config"
	"net/http"
)

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

func SendBotMessage(chatId int, text string) error {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.AppConfig.TelegramBotToken)
	body := BotMessage{
		ChatId: chatId,
		Text:   text,
	}

	jsonValue, _ := json.Marshal(body)
	_, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	return nil
}
