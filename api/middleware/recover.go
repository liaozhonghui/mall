package middleware

import (
	"mall/internal/httputils"
	"mall/internal/log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const size = 64 << 10 // 64KB

func RecoverMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]

			// 使用 zap 记录 panic 信息
			log.Error("Panic recovered",
				zap.Any("panic", r),
				zap.String("stack", string(buf)),
				zap.String("method", c.Request.Method),
				zap.String("path", c.Request.URL.Path),
				zap.String("client_ip", c.ClientIP()),
			)

			c.JSON(http.StatusOK, httputils.Error(http.ErrAbortHandler))
			c.Abort()
		}
	}()

	c.Next()
}
