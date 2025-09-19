package controller

import (
	"mall/api/httputils"
	"mall/internal/entity"

	logger "mall/internal/logger"
	"mall/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUserInfo(c *gin.Context) {
	var req = entity.SetUserInfoReq{}

	// 记录请求开始
	logger.WithContext(c).Info("SetUserInfo request started")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}

	if err := service.SetUserInfo(c, req); err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}

	c.JSON(http.StatusOK, httputils.SuccessWithData(nil))
}
