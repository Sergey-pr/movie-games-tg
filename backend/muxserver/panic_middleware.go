package muxserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Sergey-pr/movie-games-tg/utils"
	"github.com/lib/pq"
	"io"
	"net/http"
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
					errDetails = fmt.Sprintf("%s : %s", t.Code, t.Message)
					errHttpCode = http.StatusBadRequest
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
					errDetails = "Unknown error"
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
