package middleware

import (
	"bytes"
	"fmt"
	"io"
	"mall/internal/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func AccessLogger(c *gin.Context) {
	// 开始时间
	startTime := time.Now()
	var body []byte
	if c.Request.Body != nil {
		body, _ = io.ReadAll(c.Request.Body)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	blw := &bodyLogWriter{body: bytes.NewBuffer([]byte{}), ResponseWriter: c.Writer}
	c.Writer = blw

	c.Next()

	// 结束时间
	endTime := time.Now()
	//执行时间
	latencyTime := endTime.Sub(startTime)
	// 请求方式
	reqMethod := c.Request.Method
	// 请求路由
	reqUri := c.Request.RequestURI
	// 状态码
	statusCode := c.Writer.Status()

	msg := fmt.Sprintf("method:%v uri:%v req_body:%v status_code:%v latency:%v",
		reqMethod, reqUri, cast.ToString(body), statusCode, latencyTime)

	logger.WithContext(c).Info(msg)
	// logger.WithGoID().Info(msg)
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 包装后的对象：write写，写自己的buffer，再调用原始的write写响应
func (w bodyLogWriter) Write(b []byte) (int, error) {
	if _, err := w.body.Write(b); err != nil {
		fmt.Println("bodyLogWriter err:", err)
	}
	return w.ResponseWriter.Write(b)
}
