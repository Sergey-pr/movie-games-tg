package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Sergey-pr/movie-games-tg/config"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type BotMessage struct {
	ChatId      int          `json:"chat_id"`
	Text        string       `json:"text"`
	ReplyMarkup *replyMarkup `json:"reply_markup"`
	ParseMode   string       `json:"parse_mode"`
}

type replyMarkup struct {
	InlineKeyboard [][]*keyboardButton `json:"inline_keyboard"`
}

type keyboardButton struct {
	Text   string      `json:"text"`
	WebApp *webAppInfo `json:"web_app"`
}

type webAppInfo struct {
	Url string `json:"url"`
}

func SendStartBotMessage(chatId int, text string, buttonText string) error {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.AppConfig.TelegramBotToken)
	body := BotMessage{
		ChatId: chatId,
		Text:   text,
		ReplyMarkup: &replyMarkup{InlineKeyboard: [][]*keyboardButton{{
			&keyboardButton{
				Text:   buttonText,
				WebApp: &webAppInfo{Url: config.AppConfig.FrontendHostname},
			},
		}}},
		ParseMode: "MarkdownV2",
	}

	jsonValue, _ := json.Marshal(body)
	_, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	return nil
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

type GetFileResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		FileId       string `json:"file_id"`
		FileUniqueId string `json:"file_unique_id"`
		FileSize     int    `json:"file_size"`
		FilePath     string `json:"file_path"`
	} `json:"result"`
}

func DownloadBotImage(imageId string) error {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s", config.AppConfig.TelegramBotToken, imageId)

	r, err := http.Get(endpoint)
	if err != nil {
		return err
	}

	var fileResponse GetFileResponse
	err = json.NewDecoder(r.Body).Decode(&fileResponse)
	if err != nil {
		return err
	}

	newPath := filepath.Join(".", "card_files")
	err = os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		return err
	}
	// Create the file
	out, err := os.Create(filepath.Join("card_files", imageId))
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", config.AppConfig.TelegramBotToken, fileResponse.Result.FilePath))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func RegisterCallback() error {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook", config.AppConfig.TelegramBotToken)
	body := map[string]string{
		"url": fmt.Sprintf("%s/api/public/bot-updates/", config.AppConfig.BackendHostname),
	}

	jsonValue, _ := json.Marshal(body)
	_, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	return nil
}
