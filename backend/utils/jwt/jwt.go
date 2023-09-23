package jwt

import (
	"errors"
	"github.com/Sergey-pr/movie-games-tg/config"
	"github.com/Sergey-pr/movie-games-tg/utils"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenKey = "JWT"

var InvalidTokenError = errors.New("token invalid")
var tokenSecretBytes = []byte(config.AppConfig.JwtSecret)

type UserClaims struct {
	Id            int         `json:"id"`
	TelegramId    int         `json:"tg_id"`
	Name          string      `json:"name"`
	AnsweredCards utils.JSONB `json:"answered_cards"`
}

type Claims struct {
	User UserClaims `json:"user"`
	jwt.StandardClaims
}

func GetJwtToken(claims *Claims) (string, time.Time, error) {
	var ttl = time.Duration(config.AppConfig.SessionTTL * 24 * int64(time.Hour))
	if config.AppConfig.Debug == true {
		ttl = 365 * 24 * time.Hour
	}

	expirationTime := time.Now().Add(ttl)
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(tokenSecretBytes)

	return jwtToken, expirationTime, err
}

func RenewJwtToken(tknStr string) (string, time.Time, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return tokenSecretBytes, nil
	})
	if err != nil {
		return "", time.Time{}, err
	}
	if !tkn.Valid {
		return "", time.Time{}, InvalidTokenError
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return tknStr, time.Unix(claims.ExpiresAt, 0), nil
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(tokenSecretBytes)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}

func ParseJwtToken(tknStr string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return tokenSecretBytes, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, InvalidTokenError
	}
	return claims, nil
}
