// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	utils "github.com/mpiorowski/golang"
)

var (
	DOMAIN = utils.MustGetenv("DOMAIN")
	PORT   = utils.MustGetenv("PORT")
	ENV    = utils.MustGetenv("ENV")
)

var hubs = make(map[string]*Hub)

func main() {
	log.Println("Starting server...")
	if ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	url := "http://www." + DOMAIN + ":3000"
	urlShort := "http://" + DOMAIN + ":3000"
	if ENV == "production" {
		url = "https://www." + DOMAIN
		urlShort = "https://" + DOMAIN
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{url, urlShort}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(config))

	router.POST("/rooms", joinRoom)
	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		hub, ok := hubs[roomId]
		if !ok {
			hub = newHub()
			go hub.run()
			hubs[roomId] = hub
		}
		serveWs(hub, c.Writer, c.Request, roomId)
	})

	if err := router.Run(fmt.Sprintf("0.0.0.0:%v", PORT)); err != nil {
		panic(err)
	}
}
