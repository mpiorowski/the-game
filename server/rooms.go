package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Room struct {
	Id       string `json:"id"`
	Created  string `json:"created"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var rooms = make(map[string]*Room)

func joinRoom(c *gin.Context) {
	var request Room

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if room exists
    var room *Room
    for _, r := range rooms {
        if r.Name == request.Name {
            room = r
        }
    }
	if room != nil {
		compareErr := bcrypt.CompareHashAndPassword([]byte(room.Password), []byte(request.Password))
		if compareErr != nil {
			c.JSON(404, gin.H{"error": "Invalid password"})
			return
		} else {
			c.JSON(200, room)
		}
	}

	if room == nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		room = &Room{
			Id:       uuid.New().String(),
			Created:  time.Now().Format("2006-01-02 15:04:05"),
			Name:     request.Name,
			Password: string(hashedPassword),
		}
		rooms[room.Id] = room
		c.JSON(200, room)
	}


}
