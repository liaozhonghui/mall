package middleware

import (
	"log"
	"mall/api/httputils"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

const size = 64 << 10 // 64kb

/**
 * 防止异常崩溃
 */
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("http: panic serving error: %v\n%s", r, buf)
			c.JSON(http.StatusOK, httputils.Error(httputils.InternalError))
			c.Abort()
		}
	}()
	c.Next()
}
