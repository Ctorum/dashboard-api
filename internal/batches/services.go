package batches

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/internal/infra"
)

type BatchesService struct{}

func GetBatches(c *gin.Context) {
	var batches []Batches
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	if page <= 0 {
		page = 1
	}

	switch {
	case size <= 0:
		size = 10
	case size > 100:
		size = 100
	}

	offset := (page - 1) * size
	infra.DB.Offset(offset).Limit(size).Order("created_at desc").Find(&batches)
	c.JSON(http.StatusOK, infra.BaseResponseEnveloper{
		Success: true,
		Message: "Batches retrieved successfully",
		Data:    batches,
	})
}
