package muxserver

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Sergey-pr/movie-games-tg/config"
	"github.com/Sergey-pr/movie-games-tg/muxserver/forms"
	"github.com/Sergey-pr/movie-games-tg/utils"
	"io"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/lib/pq"
)

func panicMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		payload, err := io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewReader(payload))
		if err != nil {
			panic(err)
		}
		defer func() {
			r := recover()
			if r != nil {

				var (
					errDetails  interface{}
					errHttpCode = http.StatusOK
				)

				switch t := r.(type) {
				case *pq.Error:
					pqError := t
					switch pqError.Code {
					case "23505":
						errDetails = errors.New("value not unique")
						errHttpCode = http.StatusConflict
					case "P0001":
						errDetails = errors.New("account limit")
						errHttpCode = http.StatusPaymentRequired
					default:
						errDetails = fmt.Sprintf("%s : %s", pqError.Code, pqError.Message)
						errHttpCode = http.StatusBadRequest
					}
				case utils.ValidateError:
					errHttpCode = http.StatusBadRequest
					errDetails = utils.ValidateErrors{t}.Error()
				case utils.ValidateErrors:
					errHttpCode = http.StatusBadRequest
					errDetails = t.Error()
				case string:
					errDetails = t
				case error:
					errDetails = t.Error()

				default:
					if config.AppConfig.Debug == true {
						log.Println(string(debug.Stack()))
					}
					errDetails = "Unknown error"
				}

				var form forms.BotUpdate
				err = json.NewDecoder(io.NopCloser(bytes.NewReader(payload))).Decode(&form)
				if err == nil {
					err = utils.SendBotMessage(form.Message.Chat.Id, fmt.Sprintf("%s\n\n%s", errDetails.(string), string(debug.Stack())))
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(errHttpCode)
				_ = json.NewEncoder(w).Encode(struct {
					Message interface{} `json:"error"`
				}{errDetails})
			}
		}()
		h.ServeHTTP(w, req)
	})
}
