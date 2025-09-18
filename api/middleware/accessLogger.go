package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func AccessLogger(c *gin.Context) {
	startTime := time.Now()

	c.Next()
	endTime := time.Now()
	latency := endTime.Sub(startTime)

	method := c.Request.Method
	path := c.Request.RequestURI
	statusCode := c.Writer.Status()

	msg := fmt.Sprintf("| %3d | %13v | %s | %s |", statusCode, latency, method, path)

	fmt.Println(msg)
}
