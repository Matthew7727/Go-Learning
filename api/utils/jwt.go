package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64 ) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email" : email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (error, int64) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token)(interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return  nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("couldn't parse token"), 0
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return errors.New("invalid token"), 0
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return errors.New("invalid token claims"), 0
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return nil, userId
}