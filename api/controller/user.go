package controller

import (
	"mall/internal/entity"
	"mall/internal/httputils"
	"mall/internal/log"
	"mall/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetUserInfo(c *gin.Context) {
	var req = entity.SetUserInfoReq{}

	// 记录请求开始
	log.Info("SetUserInfo request started")

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to bind JSON request", zap.Error(err))
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}

	// 记录请求参数（注意不要记录敏感信息）
	log.Debug("SetUserInfo request data", zap.Any("request", req))

	if err := service.SetUserInfo(c, req); err != nil {
		log.Error("SetUserInfo service error", zap.Error(err))
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}

	log.Info("SetUserInfo completed successfully")
	c.JSON(http.StatusOK, httputils.SuccessWithData(nil))
}
