package muxserver

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ssoroka/slice"
	"io"
	"log"
	"mediacore/core"
	"mediacore/models/dam"
	"mediacore/pkg/config"
	"mediacore/pkg/persist"
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
			ctx := req.Context()
			r := recover()
			if r != nil {

				var (
					errDetails  interface{}
					errHttpCode = http.StatusInternalServerError
				)

				switch t := r.(type) {
				case *pq.Error:
					captureSentryErr(req, payload, t)
					pqError := t
					switch pqError.Code {
					case "23505":
						errDetails = dam.UniqueValueError(ctx)
						errHttpCode = http.StatusConflict
					case "P0001":
						errDetails = dam.AccountLimitError(ctx)
						errHttpCode = http.StatusPaymentRequired
					default:
						errDetails = fmt.Sprintf("%s : %s", pqError.Code, pqError.Message)
						errHttpCode = http.StatusBadRequest
					}
				case *core.DamError:
					if !slice.Contains(persist.SentryIgnoreStatusCodes, t.StatusCode) {
						captureSentryErr(req, payload, t)
					}
					l := t.DamError()
					errDetails = l.(core.DamError).Details
					errHttpCode = l.(core.DamError).StatusCode
					if errHttpCode == 0 {
						errHttpCode = http.StatusInternalServerError
					}
					if errHttpCode > http.StatusInternalServerError {
						log.Println(string(debug.Stack()))
						log.Println(t.Error())
					}
				case string:
					captureSentryErr(req, payload, errors.New(t))
					errDetails = t
				case core.ValidateError:
					captureSentryErr(req, payload, t)
					errHttpCode = http.StatusBadRequest
					errDetails = core.ValidateErrors{t}.DamError()
				case core.ValidateErrors:
					captureSentryErr(req, payload, t)
					errHttpCode = http.StatusBadRequest
					errDetails = t.DamError()
				case error:
					if config.AppConfig.Debug == true {
						log.Println(string(debug.Stack()))
						log.Println(t.Error())
					}

					errDetails = t.Error()

				default:
					if config.AppConfig.Debug == true {
						log.Println(string(debug.Stack()))
					}
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
