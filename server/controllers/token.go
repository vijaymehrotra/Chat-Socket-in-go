package controllers

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/vijaymehrotra/go-next-ts_chat/models"
)

type customClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserID 	 uint   `json:"user_id"`
	jwt.RegisteredClaims
}

var secret string = "secret"

func GenerateToken(user models.User) (string, error) {
	claims := customClaims{
		Username: user.Username,
		Email: user.Email,
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24))},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t,err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}
	return t, nil
}

func VerifyToken(tokenString string) (*customClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}

func RefreshToken(tokenString string) (string, error) {
	claims, err := VerifyToken(tokenString)
	if err != nil {
		return "", err
	}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func AuthenticateToken(tokenString string) (bool, error) {
	_, err := VerifyToken(tokenString)
	if err != nil {
		return false, err
	}
	return true, nil
}