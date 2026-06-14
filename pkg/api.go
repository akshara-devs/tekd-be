package pkg

import (
	"github.com/gin-gonic/gin"
)

type HTTPResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
}

func JSON(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, HTTPResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}
