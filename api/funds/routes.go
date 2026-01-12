package funds

import "github.com/gin-gonic/gin"

func FundRoutes(r *gin.RouterGroup) {
	r.GET("", GetFundsUseCase)
}
