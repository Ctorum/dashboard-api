package main

import (
	"github.com/gin-gonic/gin"
	"github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/api"
	"github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/api/infra"
)

func main() {
	api.LoadSettings()
	infra.InitDatabase()

	r := gin.Default()
	routerGroup := r.Group("")
	api.SetupRoutes(routerGroup)
	r.Run()
}
