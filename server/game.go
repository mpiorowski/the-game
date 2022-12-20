package main

import (
	"encoding/json"
	"log"
	"time"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func readMessage(c *Client, msg []byte) {

	var message Message
	err := json.Unmarshal(msg, &message)
	if err != nil {
		log.Println(err)
		return
	}

	switch message.Type {
	case "message":
		c.hub.broadcast <- msg
	case "nickname":
		c.hub.broadcast <- []byte(message.Data)
	case "start-timer":
		startTimer(c)

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
