package handlers

import (
	"encoding/json"
	"github.com/Sergey-pr/movie-games-tg/utils"
	"net/http"
)

func OrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ObjOrPanic[T any](o T, err error) T {
	if err != nil {
		panic(err)
	}
	return o
}
func DecodeRequest(r *http.Request, i interface{}) error {
	return json.NewDecoder(r.Body).Decode(i)
}

func ValidateForm(r *http.Request, i interface{}) error {
	if err := DecodeRequest(r, i); err != nil {
		return err
	}

	v := utils.NewValidator()
	errList := utils.ValidateStruct(v, i)
	if errList != nil {
		return errList
	}
	return nil
}

func Resp(w http.ResponseWriter, i interface{}, statusCode ...int) {
	var status = http.StatusOK
	if len(statusCode) > 0 {
		status = statusCode[0]
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if i != nil {
		_ = json.NewEncoder(w).Encode(i)
	}
}
