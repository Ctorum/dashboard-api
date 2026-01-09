package main

import (
	"github.com/gin-gonic/gin"
	"github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/src"
)

func main() {
	r := gin.Default()
	routerGroup := r.Group("/api")
	src.SetupRoutes(routerGroup)
	r.Run(":18000")
}
