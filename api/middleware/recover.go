package middleware

import (
	"fmt"
	"mall/internal/httputils"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

const size = 64 << 10 // 64KB

func RecoverMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Printf("panic recovered: %s\n%s", r, buf)
			c.JSON(http.StatusOK, httputils.Error(http.ErrAbortHandler))
			c.Abort()
		}
	}()

	c.Next()
}
