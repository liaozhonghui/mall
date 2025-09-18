package middleware

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func AccessLogger(c *gin.Context) {
	startTime := time.Now()

	var body []byte
	if c.Request.Body != nil {
		body, _ = io.ReadAll(c.Request.Body)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	c.Next()
	endTime := time.Now()
	latency := endTime.Sub(startTime)

	method := c.Request.Method
	path := c.Request.RequestURI
	statusCode := c.Writer.Status()

	msg := fmt.Sprintf("| %v | %v | %v | %v |%v", method, path, statusCode, latency, cast.ToString(body))

	fmt.Println(msg)
}
