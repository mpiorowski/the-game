package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
    ENV = os.Getenv("ENV")
    PORT = os.Getenv("PORT")
    DOMAIN = os.Getenv("DOMAIN")
)

func main() {
	if ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	url := "http://" + DOMAIN + ":3000"
	if ENV == "production" {
		url = "https://www." + DOMAIN
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{url}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(config))

	if err := router.Run(fmt.Sprintf("0.0.0.0:%v", PORT)); err != nil {
		panic(err)
	}
}
