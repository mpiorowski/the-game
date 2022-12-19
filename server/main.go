// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var addr = flag.String("addr", ":8080", "http service address")

var hubs = make(map[string]*Hub)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	log.Println("Starting server...")
	flag.Parse()

	router := gin.New()

	router.LoadHTMLFiles("home.html")

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		hub, ok := hubs[roomId]
		if !ok {
			hub = newHub()
			go hub.run()
			hubs[roomId] = hub
		}
		serveWs(hub, c.Writer, c.Request)
	})

    router.Run("0.0.0.0:8080")
}
