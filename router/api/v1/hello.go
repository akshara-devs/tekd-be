package v1

import (
	"net/http"

	"github.com/akshara-devs/tekd-be/pkg"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	data := gin.H{
		"name": "akshara",
	}

	pkg.JSON(c, http.StatusOK, "Hello, World!", data)
}
