package api

import (
	"fmt"
	"os"

	"github.com/adridevelopsthings/openapi-change-notification/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var apiEngine *gin.Engine

func createRouter() {
	apiEngine = gin.Default()
	apiEngine.Use(cors.Default())
	apiEngine.Use(static.Serve("/", static.LocalFile(config.CurrentConfig.FrontendStaticServe, false)))

	apiEngine.POST("/api/subscribe", subscribePath)
	apiEngine.POST("/api/unsubscribe", unsubscribe)
	apiEngine.POST("/api/email/verify/:code", finishEmailVerification)
	apiEngine.POST("/api/unsubscribe/verify/:code", finishUnsubscribeVerification)
}

func StartServer() {
	createRouter()
	host := os.Getenv("LISTEN_HOST")
	port := os.Getenv("LISTEN_PORT")

	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "8080"
	}

	listen := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("Server is running on http://%s/\n", listen)
	apiEngine.Run(listen)
}
