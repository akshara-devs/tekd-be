package router

import (
	"github.com/akshara-devs/tekd-be/router/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routerV1 := router.Group("/api/v1")

	{
		routerV1.GET("/hello", v1.Hello)
	}

	return router
}
