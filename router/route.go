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

	routerV1.GET("/hello", v1.Hello)

	routerUsers := routerV1.Group("/users")
	routerUsers.POST("/", v1.CreateUser)
	routerUsers.GET("/", v1.ListUsers)
	routerUsers.GET("/:id", v1.GetUser)
	routerUsers.PUT("/:id", v1.UpdateUser)
	routerUsers.DELETE("/:id", v1.DeleteUser)

	return router
}
