package router

import (
	"errors"
	"mall/api/controller"
	"mall/api/middleware"

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

	rg.Any("/healthCheck", controller.HealthCheck)
	rg.Any("/healthCheckV1", controller.HealthCheckV1)
	rg.PUT("/users", controller.SetUserInfo)
}
