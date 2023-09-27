package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Sergey-pr/movie-games-tg/config"
	"io"
	"net/http"
	"os"
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

	// Create the file
	out, err := os.Create(fmt.Sprintf("card_files/%s", imageId))
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
