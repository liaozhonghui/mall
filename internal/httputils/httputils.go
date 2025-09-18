package httputils

import "github.com/gin-gonic/gin"

func SuccessWithData(data interface{}) gin.H {
	return gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	}
}
func Error(err error) gin.H {
	return gin.H{
		"code": 50000,
		"msg":  err.Error(),
		"data": nil,
	}
}
