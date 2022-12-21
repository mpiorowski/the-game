package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/gofrs/uuid"
)

type Message struct {
	Room string `json:"room"`
	Type string `json:"type"`
	Data string `json:"data"`
}

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Team     string `json:"team"`
	Ready    bool   `json:"ready"`
    Step     int    `json:"step"`
}

type Respone struct {
    Users []User `json:"users"`
}

var users []User

func runGame(c *Client, msg []byte) {

	var message Message
	err := json.Unmarshal(msg, &message)
	if err != nil {
		log.Println(err)
		return
	}

    var respone Respone

	if message.Type == "join" {
        respone.Users = users
		val, err := json.Marshal(respone)
		if err != nil {
			log.Println(err)
			return
		}
		c.send <- []byte(val)
	}

	if message.Type == "nickname" {
		uuid, err := uuid.NewV4()
		if err != nil {
			log.Println(err)
			return
		}
		user := &User{
			Id:       uuid.String(),
			Nickname: message.Data,
			Team:     "",
			Ready:    false,
            Step:     1,
		}
		users = append(users, *user)
        respone.Users = users
        val, err := json.Marshal(respone)
		if err != nil {
			log.Println(err)
			return
		}
		c.hub.broadcast <- []byte(val)

		jsonUser, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
			return
		}
		c.send <- []byte("{\"user\":" + string(jsonUser) + "}")
	}

	if message.Type == "ready" {
		for i, user := range users {
			if user.Id == message.Data {
				users[i].Ready = true
			}
		}

		// if all users are ready, randomly assign teams
		allReady := true
		for _, user := range users {
			if !user.Ready {
				allReady = false
			}
		}
		if allReady {
			for i := range users {
				j := rand.Intn(i + 1)
				users[i], users[j] = users[j], users[i]
			}
			for i := range users {
				if i%2 == 0 {
					users[i].Team = "red"
				} else {
					users[i].Team = "blue"
				}
			}
		}

        respone.Users = users
        val, err := json.Marshal(respone)
		if err != nil {
			log.Println(err)
			return
		}
		c.hub.broadcast <- []byte(val)
	}

    if message.Type == "clues" {
        for i, user := range users {
            if user.Id == message.Data {
                users[i].Step = 2
            }
        }

        respone.Users = users
        val, err := json.Marshal(respone)
        if err != nil {
            log.Println(err)
            return
        }
        c.hub.broadcast <- []byte(val)
    }


}

func startTimer(c *Client) {
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		// send message every sec for one minute
		for i := 0; i < 10; i++ {
			select {
			case <-ticker.C:
				c.hub.broadcast <- []byte("tick")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

}
