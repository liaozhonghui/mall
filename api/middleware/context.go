package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
)

func Context(c *gin.Context) {
	traceId := c.GetHeader("trace-Id")
	if traceId == "" {
		traceId, _ = uuid.GenerateUUID()
	}
	c.Set("traceId", traceId)
	c.Set("startTime", time.Now().UnixNano())

}
