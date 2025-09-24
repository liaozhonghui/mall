package test

import (
	"mall/internal/core"
	"mall/internal/entity"
	"mall/internal/logger"
	"mall/internal/service"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestUserLogin(t *testing.T) {
	_ = core.InitConfig("../configs/config.yaml")
	_ = logger.InitLogger()
	login := entity.LoginReq{
		Account:  "test002",
		Password: "123456",
	}
	var user1, user2 entity.LoginResp
	var err error

	gin.SetMode(gin.TestMode)
	c1, _ := gin.CreateTestContext(nil)
	c2, _ := gin.CreateTestContext(nil)

	go func() {
		user1, err = service.APILogin(c1, login)
		if err != nil {
			t.Error(err)
		}
	}()
	go func() {
		user2, err = service.APILogin(c2, login)
		if err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(time.Second * 1)
	if user1.UserId != user2.UserId {
		t.Errorf("user1 and user2 not equal, user1: %v, user2: %v", user1, user2)
	}
}
