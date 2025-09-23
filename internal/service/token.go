package service

import (
	"mall/internal/core"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAPIToken(userId int, failTokenMin time.Duration) (token string, err error) {
	if failTokenMin == 0 {
		failTokenMin = core.GlobalConfig.Jwt.ExpireTime * time.Second
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(failTokenMin).Unix(),
	})
	atoken, err := at.SignedString([]byte(core.GlobalConfig.Jwt.ApiSecret))
	if err != nil {
		return "", err
	}
	return atoken, nil
}

func ParseAPIToken(token string) (jwt.MapClaims, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(core.GlobalConfig.Jwt.ApiSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return claim.Claims.(jwt.MapClaims), nil
}
