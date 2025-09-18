package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
)

func ParseAuthToken(token string) (int, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return 0, err
	}
	userId := cast.ToInt(claims["userId"])
	return userId, nil
}
