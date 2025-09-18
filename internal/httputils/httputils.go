package httputils

import "github.com/gin-gonic/gin"

func SuccessWithData(data interface{}) gin.H {
	return gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	}
}
