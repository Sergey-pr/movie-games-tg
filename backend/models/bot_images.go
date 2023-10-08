package models

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Sergey-pr/movie-games-tg/config"
	"github.com/Sergey-pr/movie-games-tg/persist"
	"github.com/doug-martin/goqu/v9"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	BotFilesTableName = "bot_files"
)

type BotFile struct {
	Id       int    `db:"id" goqu:"skipupdate,skipinsert"`
	Filename string `db:"filename"`
	FileId   string `db:"file_id"`
}

// GetFileByName return BotImage object by its filename
func GetFileByName(ctx context.Context, filename string) (*BotFile, error) {
	var obj BotFile
	exists, err := persist.Db.From(BotFilesTableName).Where(
		goqu.Ex{"filename": filename},
	).ScanStructContext(ctx, &obj)
	if err != nil {
		return nil, err
	}
	if exists == false {
		return nil, nil
	}
	return &obj, nil
}

// Save BotImage instance in DB
func (obj *BotFile) Save(ctx context.Context) error {
	var err error
	if obj.Id == 0 {
		err = obj.create(ctx)
	} else {
		err = obj.update(ctx)
	}
	if err != nil {
		return err
	}
	return nil
}

// create private method to create new BotImage DB record
func (obj *BotFile) create(ctx context.Context) error {
	insert := persist.Db.Insert(BotFilesTableName).
		Rows(obj).
		Returning("*").Executor()
	if _, err := insert.ScanStructContext(ctx, obj); err != nil {
		return err
	}
	return nil
}

// update private method to update BotImage record in DB
func (obj *BotFile) update(ctx context.Context) error {
	update := persist.Db.From(BotFilesTableName).
		Where(goqu.C("id").Eq(obj.Id)).Update().Set(obj).
		Executor()
	_, err := update.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Delete card from DB
func (obj *BotFile) Delete(ctx context.Context) error {
	_, err := persist.Db.From(BotFilesTableName).
		Where(goqu.Ex{"id": obj.Id}).
		Delete().
		Executor().ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

// BotVideoMessage is an object with a telegram sendVideo method params
type BotVideoMessage struct {
	ChatId      int          `json:"chat_id"`
	Caption     string       `json:"caption"`
	ReplyMarkup *replyMarkup `json:"reply_markup"`
	ParseMode   string       `json:"parse_mode"`
	Video       string       `json:"video"`
}

// BotMessage is an object with a telegram sendMessage method params
type BotMessage struct {
	ChatId      int          `json:"chat_id"`
	Text        string       `json:"caption"`
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

type sendVideoResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		Video struct {
			FileId string `json:"file_id"`
		} `json:"video"`
	} `json:"result"`
}

// SendStartBotMessage sends default telegram message with a button to open web app
func SendStartBotMessage(ctx context.Context, chatId int, language string) error {
	// Create sendVideo endpoint with telegram bot token
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendVideo", config.AppConfig.TelegramBotToken)

	var (
		textMsg    string
		buttonText string
		filename   string
	)
	if language == "ru" {
		textMsg = "*КИНОИГРЫ*🍿\n\nЧтобы начать игру нажмите Старт\\!"
		buttonText = "Старт"
		filename = "video_ru.mp4"
	} else {
		textMsg = "*MOVIEGAMES*🍿\n\nPress start to begin\\!"
		buttonText = "Start"
		filename = "video_en.mp4"
	}

	fileObj, err := GetFileByName(ctx, filename)
	if err != nil {
		return err
	}

	if fileObj != nil {
		// If file is already uploaded to telegram, and we have its id
		// Fill the botMessage struct with data of our message
		body := BotVideoMessage{
			ChatId:  chatId,
			Caption: textMsg,
			Video:   fileObj.FileId,
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
		_, err = http.Post(endpoint, "application/json", bytes.NewBuffer(jsonValue))
	} else {
		// If we need to upload new file to telegram
		// Fill the replyMarkupObj struct with data of our button
		replyMarkupObj := &replyMarkup{InlineKeyboard: [][]*keyboardButton{{
			&keyboardButton{
				Text:   buttonText,
				WebApp: &webAppInfo{Url: config.AppConfig.FrontendHostname},
			},
		}}}
		// Marshal reply mark up to json string
		jsonReplyMarkup, _ := json.Marshal(replyMarkupObj)
		// prepare the reader instances to encode in multipart data
		values := map[string]io.Reader{
			"video":        mustOpen(fmt.Sprintf("bot_files/%s", filename)),
			"chat_id":      strings.NewReader(strconv.Itoa(chatId)),
			"caption":      strings.NewReader(textMsg),
			"parse_mode":   strings.NewReader("MarkdownV2"),
			"reply_markup": bytes.NewBuffer(jsonReplyMarkup),
		}
		err := UploadNewFile(ctx, endpoint, values, filename)
		if err != nil {
			return err
		}
	}
	return nil
}

// UploadNewFile uploads new file to telegram, so we could use it later by id
func UploadNewFile(ctx context.Context, url string, values map[string]io.Reader, filename string) error {
	var (
		err error
		b   bytes.Buffer
	)
	// Prepare a form that you will submit to that URL.
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if _, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, filename); err != nil {
				return err
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return err
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}

	}

	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	err = w.Close()
	if err != nil {
		return err
	}

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return err
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())
	// Submit the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	// Now parse the response to get file id
	var form sendVideoResponse
	err = json.NewDecoder(res.Body).Decode(&form)
	if err != nil {
		return err
	}
	// Create BotFile obj and save it to database
	fileObj := &BotFile{
		Filename: filename,
		FileId:   form.Result.Video.FileId,
	}
	err = fileObj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

// mustOpen is a shortcut to open local file
func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
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
