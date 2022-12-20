package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Room struct {
	Id       string  `json:"id"`
	Created  string  `json:"created"`
	Updated  string  `json:"updated"`
	Deleted  *string `json:"deleted"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
}

func joinRoom(c *gin.Context) {
	var request Room
	var response Room

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	row := db.QueryRow("SELECT id, created, updated, deleted, name, password FROM rooms WHERE name = $1", request.Name)
	err = row.Scan(&response.Id, &response.Created, &response.Updated, &response.Deleted, &response.Name, &response.Password)

	if err != nil && err != sql.ErrNoRows {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err == nil {
		compareErr := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(request.Password))
		if compareErr != nil {
			c.JSON(401, gin.H{"error": "Invalid password"})
			return
		} else {
			c.JSON(200, response)
			return
		}
	}

	if err == sql.ErrNoRows {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		row := db.QueryRow("INSERT INTO rooms (name, password) VALUES ($1, $2) RETURNING *", request.Name, hashedPassword)
		err = row.Scan(
			&response.Id,
			&response.Created,
			&response.Updated,
			&response.Deleted,
			&response.Name,
			&response.Password,
		)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, response)
}
