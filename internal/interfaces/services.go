package interfaces

import (
	"github.com/gin-gonic/gin"
)

type Servicer interface {
	Execute(c *gin.Context)
}
