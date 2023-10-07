package jwt

import (
	"errors"
	"github.com/Sergey-pr/movie-games-tg/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenKey = "JWT"

// InvalidTokenError is an error shortcut
var InvalidTokenError = errors.New("token invalid")

// tokenSecretBytes is a jwt secret in bytes
var tokenSecretBytes = []byte(config.AppConfig.JwtSecret)

// UserClaims is a user object to use in JWT claims
type UserClaims struct {
	Id         int    `json:"id"`
	TelegramId int    `json:"tg_id"`
	Name       string `json:"name"`
}

// Claims is a standard JWT claims with user object claims
type Claims struct {
	User UserClaims `json:"user"`
	jwt.StandardClaims
}

// GetJwtToken returns JWT token by claims
func GetJwtToken(claims *Claims) (string, error) {
	// Set expiration of token
	expirationAt := time.Now().Add(time.Hour)
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expirationAt.Unix(),
	}
	// Create token string
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(tokenSecretBytes)
	return tokenString, err
}

// RenewJwtToken renews expired JWT token
func RenewJwtToken(tokenString string) (string, *time.Time, error) {
	claims := &Claims{}
	// Parse token claims
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return tokenSecretBytes, nil
	})
	if err != nil {
		return "", nil, err
	}
	if !tkn.Valid {
		return "", nil, InvalidTokenError
	}
	// If token is not expired for a minute or more don't create a new token
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > time.Minute {
		expirationAt := time.Unix(claims.ExpiresAt, 0)
		return tokenString, &expirationAt, nil
	}
	// Create a 5-minute JWT token
	expirationAt := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationAt.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(tokenSecretBytes)
	if err != nil {
		return "", nil, err
	}
	return tokenString, &expirationAt, nil
}

// ParseJwtToken parses and validates JWT token string, returns claims
func ParseJwtToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
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
