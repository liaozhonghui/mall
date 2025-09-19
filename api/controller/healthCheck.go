package controller

import (
	"mall/api/httputils"
	"mall/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	r, _ := service.HealthCheck(c)
	c.JSON(http.StatusOK, httputils.SuccessWithData(r))
}

func HealthCheckV1(c *gin.Context) {
	r, _ := service.HealthCheckV1(c)
	c.JSON(http.StatusOK, httputils.SuccessWithData(r))
}
