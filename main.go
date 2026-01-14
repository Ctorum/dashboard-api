package main

import (
	"github.com/gin-gonic/gin"
	api "github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/internal"
	"github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/internal/infra"
)

func main() {
	api.LoadSettings()
	infra.InitDatabase()

	r := gin.Default()
	routerGroup := r.Group("")
	api.SetupRoutes(routerGroup)

	r.Run()
}
