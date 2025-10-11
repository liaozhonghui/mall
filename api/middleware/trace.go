package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
)

func Trace(c *gin.Context) {
	traceId := c.GetHeader("X-Trace-Id")
	if len(traceId) == 0 {
		traceId, _ = uuid.GenerateUUID()
	}
	startTime := time.Now().UnixNano()
	c.Set("trace_id", traceId)
	c.Set("start_time", startTime)

	log.Printf("Trace start: %s %s %s", traceId, c.Request.Method, c.Request.URL.Path)
	c.Next()
}
