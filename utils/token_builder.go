package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/markbates/goth"
	"os"
	"time"
)

func GenerateToken(payload goth.User) (string, error) {
	claims := jwt.MapClaims{
		"email": payload.Email,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		fmt.Println("Error saat membuat token:", err)
		return "", err
	}

	return tokenString, nil
}
