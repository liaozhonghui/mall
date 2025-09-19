package router

import (
	"errors"
	"mall/api/controller"
	"mall/api/httputils"
	"mall/api/middleware"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	router.Use(middleware.Context, middleware.AccessLogger)

	// 管理后台接口
	admin := router.Group("/admin")
	registerAdminRoutes(admin)

	// 对外API接口
	api := router.Group("/api")
	registerAPIRoutes(api)
}

func registerAdminRoutes(rg *gin.RouterGroup) {
	// Define admin routes here
}
func registerAPIRoutes(rg *gin.RouterGroup) {
	// Define API routes here
	rg.Any("/panic", middleware.Recover, func(c *gin.Context) {
		panic(errors.New("this is a panic test"))
	})

	rg.Any("/healthCheck", func(c *gin.Context) {
		t := rand.Intn(10000)
		c.JSON(http.StatusOK, httputils.SuccessWithData(t))
	})
	rg.PUT("/users", controller.SetUserInfo)
}
