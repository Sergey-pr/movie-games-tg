package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Sergey-pr/movie-games-tg/config"
	"github.com/Sergey-pr/movie-games-tg/models"
	"github.com/Sergey-pr/movie-games-tg/muxserver/forms"
	"github.com/Sergey-pr/movie-games-tg/muxserver/serializers"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// Login user by email and password
func Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var form forms.LoginForm
	OrPanic(ValidateForm(r, &form))
	query := ObjOrPanic(url.ParseQuery(form.InitData))

	print(query)

	if !checkHash(query) {
		Resp(w, nil, http.StatusForbidden)
	}

	var userForm forms.UserForm
	OrPanic(json.Unmarshal([]byte(query.Get("user")), &userForm))

	user, err := models.LoginUser(ctx, userForm.TelegramId)
	if err != nil {
		user = &models.User{
			TelegramId: userForm.TelegramId,
			Name:       userForm.Name,
			UserName:   userForm.UserName,
			Language:   userForm.Language,
		}
		OrPanic(user.Save(ctx))
	}

	token, expirationTime, err := user.GetJwtToken()
	OrPanic(err)

	Resp(w, serializers.JwtToken{
		Token:   token,
		ExpTime: expirationTime.Unix(),
	})
}

func checkHash(query url.Values) bool {
	hash := query.Get("hash")
	checkData := make([]string, 0)
	for key, value := range query {
		switch key {
		case "hash":
			continue
		default:
			checkData = append(checkData, fmt.Sprintf("%s=%s", key, value[0]))
		}
	}
	sort.Strings(checkData)
	dataCheckString := strings.Join(checkData, "\n")

	secretHmac := hmac.New(sha256.New, []byte("WebAppData"))
	secretHmac.Write([]byte(config.AppConfig.TelegramBotToken))

	h := hmac.New(sha256.New, secretHmac.Sum(nil))
	h.Write([]byte(dataCheckString))
	dataHash := hex.EncodeToString(h.Sum(nil))
	if dataHash == hash {
		return true
	}
	return false
}
