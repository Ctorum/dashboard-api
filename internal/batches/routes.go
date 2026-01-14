package batches

import "github.com/gin-gonic/gin"

func BatchRoutes(r *gin.RouterGroup) {
	r.GET("", GetBatches)
}
