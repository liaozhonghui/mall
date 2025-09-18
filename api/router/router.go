package router

import (
	"mall/api/controller"
	"mall/api/middleware"
	"mall/internal/httputils"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	admin := router.Group("/admin")
	registerAdminRoutes(admin)

	api := router.Group("/api")
	registerAPIRoutes(api)
}

func registerAdminRoutes(rg *gin.RouterGroup) {
	// Define admin routes here
}
func registerAPIRoutes(rg *gin.RouterGroup) {
	// Define API routes here
	rg.Any("/healthCheck", func(c *gin.Context) {
		t := rand.Intn(10000)
		c.JSON(http.StatusOK, httputils.SuccessWithData(t))
	})

	rg.PUT("/users", middleware.CheckLogin, controller.SetUserInfo)
}
