package middleware

import (
	"log"
	"mall/internal/httputils"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

const size = 64 << 10 // 64KB

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("[Recovery] panic recovered:\n%s\n%s\n", r, buf)
			c.JSON(http.StatusOK, httputils.Error(http.ErrAbortHandler))
			c.Abort()
		}
	}()

	c.Next()
}
