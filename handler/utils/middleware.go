package utils

import (
	"AppServer/types"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("x0x45nhkyotoka")

type SessionInfo struct {
	UserID string
	Role   string
}

type CustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func AuthMiddleware(r *http.Request) (SessionInfo, error) {
	cookie, err := r.Cookie("sid")
	if err != nil || cookie.Value == "" {
		return SessionInfo{}, errors.New("cooike not found")
	}

	val := func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &CustomClaims{}, val)
	if err != nil {
		return SessionInfo{}, errors.New("invalid session id")
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return SessionInfo{}, errors.New("invalid session id")
	}

	return SessionInfo{
		UserID: claims.UserID,
		Role:   claims.Role,
	}, nil
}

func NewCookie(user *types.User) (*http.Cookie, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &CustomClaims{
		UserID: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	session_id, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:    "sid",
		Value:   session_id,
		Expires: expirationTime,
	}, nil
}

func NewEmptyCookie() *http.Cookie {
	return &http.Cookie{
		Name:    "sid",
		Value:   "",
		Expires: time.Now(),
	}
}
