package service

import (
	"mall/internal/entity"
	"mall/internal/logger"
	"mall/internal/repo"
	tokenService "mall/internal/service/token"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetUserInfo(c *gin.Context, req entity.SetUserInfoReq) error {
	return nil
}

func APILogin(c *gin.Context, req entity.LoginReq) (entity.LoginResp, error) {
	var resp entity.LoginResp
	userRepo := repo.NewUserRepository()

	user, err := userRepo.GetUserByAccount(c, req.Account, req.Password)

	if err != nil {
		logger.WithContext(c).Errorf("APILogin GetUserByAccount failed, err: %v", err)
		return resp, err
	}

	if user.Id == 0 {
		user.Account = req.Account
		user.Password = req.Password
		user.NickName = "匿名用户" + strconv.Itoa(rand.Intn(10000))

		user.Id, err = userRepo.CreateUser(c, user)
		if err != nil {
			logger.WithContext(c).Errorf("APILogin CreateUser failed, err: %v", err)
			return resp, err
		}
	}

	token, _ := tokenService.GenerateToken(user.Id)

	resp.UserId = user.Id
	resp.Token = token

	return resp, nil
}
