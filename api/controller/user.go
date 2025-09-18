package controller

import (
	"mall/internal/entity"
	"mall/internal/httputils"
	"mall/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUserInfo(c *gin.Context) {
	var req = entity.SetUserInfoReq{}
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
