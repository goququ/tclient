package server

import (
	"github.com/gin-gonic/gin"
)

func createApiError(e error) gin.H {
	return gin.H{
		"error":  e.Error(),
		"status": "falied",
	}
}
