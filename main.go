package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akshara-devs/tekd-be/pkg"
	"github.com/akshara-devs/tekd-be/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config := pkg.LoadConfig()
	pkg.ConnectDatabase(config.DatabaseURL)
	gin.SetMode(config.RunMode)

	initRouter := router.InitRouter()
	endPoint := fmt.Sprintf(":%d", config.HttpPort)

	server := &http.Server{
		Addr:         endPoint,
		Handler:      initRouter,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
