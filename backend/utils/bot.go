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

// BotMessage is an object with a telegram sendMessage method params
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

// SendStartBotMessage sends default telegram message with a button to open web app
func SendStartBotMessage(chatId int, text string, buttonText string) error {
	// Create endpoint with telegram bot token
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.AppConfig.TelegramBotToken)
	// Fill the botMessage struct with data of our message
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
	// Marshal data and send request to answer in chat
	jsonValue, _ := json.Marshal(body)
	_, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	return nil
}

// SendBotMessage sends simple message to chat
func SendBotMessage(chatId int, text string) error {
	// Create endpoint with telegram bot token
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.AppConfig.TelegramBotToken)
	body := BotMessage{
		ChatId: chatId,
		Text:   text,
	}
	// Marshal data and send request to answer in chat
	jsonValue, _ := json.Marshal(body)
	_, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	return nil
}

// GetFileResponse is a struct of getFile telegram method response
type GetFileResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		FileId       string `json:"file_id"`
		FileUniqueId string `json:"file_unique_id"`
		FileSize     int    `json:"file_size"`
		FilePath     string `json:"file_path"`
	} `json:"result"`
}

// DownloadBotImage downloads an image from telegram by id and saves it locally
func DownloadBotImage(imageId string) error {
	// Create endpoint with telegram bot token and image id
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s", config.AppConfig.TelegramBotToken, imageId)
	// Make request
	r, err := http.Get(endpoint)
	if err != nil {
		return err
	}
	// Parse response
	var fileResponse GetFileResponse
	err = json.NewDecoder(r.Body).Decode(&fileResponse)
	if err != nil {
		return err
	}
	// Check if folder for images exists, if not we create it
	newPath := filepath.Join(".", "card_files")
	err = os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		return err
	}
	// Create a file for the image
	out, err := os.Create(filepath.Join("card_files", imageId))
	if err != nil {
		return err
	}
	defer out.Close()
	// Create endpoint with telegram bot token and telegram file filePath
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", config.AppConfig.TelegramBotToken, fileResponse.Result.FilePath))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Write the response body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// RegisterCallback register telegram bot callback
func RegisterCallback() error {
	// Create endpoint with telegram bot token and image id
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook", config.AppConfig.TelegramBotToken)
	// In body we send callback url
	body := map[string]string{
		"url": fmt.Sprintf("%s/api/public/bot-updates/", config.AppConfig.BackendHostname),
	}
	// Marshall data and send request
	jsonValue, _ := json.Marshal(body)
	_, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	return nil
}
