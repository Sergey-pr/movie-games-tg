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

// Login user or create new user by telegram init data
func Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Fill form struct with data from request and validate it
	var form forms.LoginForm
	OrPanic(json.NewDecoder(r.Body).Decode(&form))
	// Parse initData which is formatted as query params
	initData := ObjOrPanic(url.ParseQuery(form.InitData))
	// Check hash to validate telegram data
	if !checkHash(initData) {
		JsonResponse(w, nil, http.StatusForbidden)
	}
	// Unmarshal user data from initData to form
	var userForm forms.UserForm
	OrPanic(json.Unmarshal([]byte(initData.Get("user")), &userForm))
	// Check if user exist
	user, err := models.LoginUser(ctx, userForm.TelegramId)
	if err != nil {
		// Create new user
		user = &models.User{
			TelegramId: userForm.TelegramId,
			Name:       userForm.Name,
			UserName:   userForm.UserName,
			Language:   userForm.Language,
		}
		if userForm.LastName != nil {
			user.LastName = userForm.LastName
		}
		OrPanic(user.Save(ctx))
	}
	if user.LastName == nil && userForm.LastName != nil {
		user.LastName = userForm.LastName
		OrPanic(user.Save(ctx))
	}
	// Return jwt token
	JsonResponse(w, serializers.JwtToken{Token: ObjOrPanic(user.GetJwtToken())})
}

// Checking hash by telegram hash algorithm
func checkHash(initData url.Values) bool {
	hash := initData.Get("hash")
	checkData := make([]string, 0)
	// Remove hash from init data and format as a key=value strings slice
	for key, value := range initData {
		if key == "hash" {
			continue
		}
		checkData = append(checkData, fmt.Sprintf("%s=%s", key, value[0]))
	}
	// Sort strings alphabetically
	sort.Strings(checkData)
	// Make secret string from bot token
	secretHmac := hmac.New(sha256.New, []byte("WebAppData"))
	secretHmac.Write([]byte(config.AppConfig.TelegramBotToken))
	// Create dataCheckString by joining all alphabetically sorted key=value with \n
	dataCheckString := strings.Join(checkData, "\n")
	// Getting string hash and comparing it with initData hash
	h := hmac.New(sha256.New, secretHmac.Sum(nil))
	h.Write([]byte(dataCheckString))
	dataHash := hex.EncodeToString(h.Sum(nil))
	if dataHash == hash {
		return true
	}
	return false
}
