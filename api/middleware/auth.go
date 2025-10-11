package middleware

import (
	"mall/api/httputils"
	"net/http"
	"strings"

	tokenService "mall/internal/service/token"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func Auth(c *gin.Context) {
	bearToken := c.GetHeader("Authorization")
	if bearToken == "" {
		c.AbortWithStatusJSON(http.StatusOK, httputils.Error(httputils.AuthError))
		return
	}
	parts := strings.Fields(bearToken)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		c.AbortWithStatusJSON(http.StatusOK, httputils.Error(httputils.AuthError))
		return
	}
	token := parts[1]

	userMap, err := tokenService.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, httputils.Error(httputils.AuthError))
		return
	}
	userId := cast.ToInt(userMap["user_id"])

	c.Set("user_id", userId)
	c.Next()
}
