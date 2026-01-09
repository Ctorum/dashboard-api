package src

import "github.com/gin-gonic/gin"

type ResponseEnveloper struct {
	Succes  bool   `json:"succes"`
	Message string `json:"message"`
	Data    gin.H  `json:"data"`
}
