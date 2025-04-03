package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const key = "ssdfgfgsdfsdgfd"

func GenerateToken(email string, userId int64) (string, error) {
	fmt.Println("params----", email, userId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(key))
}

func VerifyToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(key), nil
	})

	if err != nil {
		return 0, errors.New("could not parse the token")
	}

	isValidToken := token.Valid

	if !isValidToken {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token")
	}
	fmt.Println("claims===", claims)
	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
