// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	utils "github.com/mpiorowski/golang"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DOMAIN  = utils.MustGetenv("DOMAIN")
	PORT    = utils.MustGetenv("PORT")
	ENV     = utils.MustGetenv("ENV")
	DB_USER = utils.MustGetenv("DB_USER")
	DB_PASS = utils.MustGetenv("DB_PASS")
	DB_HOST = utils.MustGetenv("DB_HOST")
	DB_NAME = utils.MustGetenv("DB_NAME")
)

var hubs = make(map[string]*Hub)
var db *sql.DB

func init() {
	// Db connection
	var err error
	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s",
		DB_USER, DB_PASS, DB_NAME, DB_HOST)
	if db, err = sql.Open("pgx", dbURI); err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected to database")

	// Migrations
	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Migrations failed: %v", err)
	}
	log.Printf("Applied %d migrations", n)
}

func main() {
	log.Println("Starting server...")
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
