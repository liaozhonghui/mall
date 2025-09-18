package controller

import "github.com/gin-gonic/gin"

func TouchPanic(c *gin.Context) {
	panic("触发了 panic")
}
