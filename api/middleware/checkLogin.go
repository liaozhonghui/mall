package middleware

import (
	"mall/api/httputils"
	"mall/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckLogin(c *gin.Context) {
	token := c.GetHeader("mall-auth-token")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusOK, httputils.Error(http.ErrNoCookie))
		return
	}

	userId, err := service.ParseAuthToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.Set("userId", userId) // 设置上下文userId
	c.Next()
}
