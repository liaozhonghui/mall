package middleware

import (
	"bytes"
	"io"
	"mall/internal/log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	clientIP := c.ClientIP()

	// 使用结构化日志记录访问信息
	log.Info("HTTP Request",
		zap.String("method", method),
		zap.String("path", path),
		zap.Int("status", statusCode),
		zap.Duration("latency", latency),
		zap.String("client_ip", clientIP),
		zap.String("user_agent", c.Request.UserAgent()),
	)

	// 如果是 debug 级别，记录请求体（注意不要记录敏感信息）
	if len(body) > 0 && len(body) < 1024 { // 限制body大小
		log.Debug("Request body", zap.ByteString("body", body))
	}
}
