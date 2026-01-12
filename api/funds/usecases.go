package funds

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/api/infra"
)

func GetFundsUseCase(c *gin.Context) {
	var funds []Funds
	result := infra.DB.Find(&funds)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, infra.BaseResponseEnveloper{
		Success: true,
		Message: "Funds retrieved successfully",
		Data:    funds,
	})
}
