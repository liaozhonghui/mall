package token

import (
	"mall/internal/core"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Second * time.Duration(core.GlobalConfig.Jwt.ExpireTime)).Unix(),
	})
	token, err := at.SignedString([]byte(core.GlobalConfig.Jwt.ApiSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (jwt.MapClaims, error) {
	claim, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(core.GlobalConfig.Jwt.ApiSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := claim.Claims.(jwt.MapClaims); ok && claim.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
