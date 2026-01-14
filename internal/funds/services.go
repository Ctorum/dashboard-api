package funds

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/internal/infra"
)

func GetFundsService(c *gin.Context) {
	var funds []BasicFundSchema
	result := infra.DB.Model(&Funds{}).Find(&funds)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, infra.BaseResponseEnveloper{
			Success: false,
			Message: "Failed to retrieve funds",
			Detail:  result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, infra.BaseResponseEnveloper{
		Success: true,
		Message: "Funds retrieved successfully",
		Data:    funds,
	})
}

func GetFundByAliasService(c *gin.Context) {
	var fund Funds
	result := infra.DB.Where("alias = ?", c.Param("alias")).First(&fund)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, infra.BaseResponseEnveloper{
			Success: false,
			Message: "Fund not found",
			Detail:  result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, infra.BaseResponseEnveloper{
		Success: true,
		Message: "Fund retrieved successfully",
		Data:    fund,
	})
}
