package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/api/funds"
	"github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/api/infra"
)

func SetupRoutes(r *gin.RouterGroup) {
	fundGroup := r.Group("/funds")
	funds.FundRoutes(fundGroup)
	r.GET("/health", getHealth)
}

func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, infra.BaseResponseEnveloper{
		Success: true,
		Message: "Service healthy",
	})
}
