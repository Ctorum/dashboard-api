package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/internal/batches"
	"github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/internal/funds"
	"github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/internal/infra"
)

func SetupRoutes(r *gin.RouterGroup) {
	fundGroup := r.Group("/funds")
	batchGroup := r.Group("/batches")

	batches.BatchRoutes(batchGroup)
	funds.FundRoutes(fundGroup)

	r.GET("/health", getHealth)
}

func getHealth(c *gin.Context) {

	c.JSON(http.StatusOK, infra.BaseResponseEnveloper{
		Success: true,
		Message: "Service healthy",
		Data: gin.H{
			"version":         "0.0.1",
			"database_status": infra.DBInterface.Stats(),
		},
	})
}
