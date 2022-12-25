package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Room   string `json:"room"`
	UserId string `json:"userId"`
	Type   string `json:"type"`
	Data   string `json:"data"`
}

type Response struct {
	Users []User  `json:"users"`
	Round Round   `json:"round"`
	Score []Score `json:"score"`
}

type Team struct {
	name  string
	score int
	clues []Clue
}

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Team     int    `json:"team"`
	Ready    bool   `json:"ready"`
	Step     int    `json:"step"`
}

type Clue struct {
	Id      string `json:"id"`
    UserId  string `json:"userId"`
	Word    string `json:"word"`
	Type    string `json:"type"`
	Guessed bool   `json:"guessed"`
}

type Round struct {
	Game     int  `json:"game"`
	Team     int  `json:"team"`
	Clue     Clue `json:"clue"`
	User     User `json:"user"`
	NextUser User `json:"nextUser"`
	Time     int  `json:"time"`
}

type Score struct {
	Game      int      `json:"game"`
	TeamClues [][]Clue `json:"teamClues"`
}

var usersPerRoom = make(map[string][]User)
var cluesPerRoom = make(map[string][]Clue)
var roundPerRoom = make(map[string]Round)
var scorePerRoom = make(map[string][]Score)
var tickerPerRoom = make(map[string]*time.Ticker)

var playerPerTeamPerRoom = make(map[string][]int)

func runGame(c *Client, msg []byte, roomId string) {

	users := usersPerRoom[roomId]
	clues := cluesPerRoom[roomId]
	round := roundPerRoom[roomId]
	score := scorePerRoom[roomId]
	ticker := tickerPerRoom[roomId]

	playerPerTeam := playerPerTeamPerRoom[roomId]

	var message Message

	err := json.Unmarshal(msg, &message)
	if err != nil {
		log.Println(err)
		return
	}

	if message.Type == "nickname" {
		var user User
		// str to byte
		err := json.Unmarshal([]byte(message.Data), &user)
		if err != nil {
			log.Println(err)
			return
		}

		user.Team = -1
		user.Ready = false
		user.Step = 1

		users = append(users, user)
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
			rand.Shuffle(len(users), func(i, j int) { users[i], users[j] = users[j], users[i] })
			for i := range users {
				if i%2 == 0 {
					users[i].Team = 0
				} else {
					users[i].Team = 1
				}
			}
		}
	}

	if message.Type == "go-to-clues" {
		for i, user := range users {
			if user.Id == message.Data {
				users[i].Step = 2
			}
		}
	}

	// TODO - validate clues
	if message.Type == "send-clues" {
		var newClues []Clue
		err := json.Unmarshal([]byte(message.Data), &newClues)
		if err != nil {
			log.Println(err)
			return
		}
		for _, clue := range newClues {
			clue.Id = uuid.New().String()
			clue.Guessed = false
			clues = append(clues, clue)
		}

		for i, user := range users {
			if user.Id == message.UserId {
				users[i].Step = 3
			}
		}

		/**
		 * Prepare round
		 */
		if len(clues) == len(users)*3 {
			for i := range users {
				users[i].Step = 4
			}
			// randomly select a team
			if rand.Intn(2) == 0 {
				round.Team = 0
			} else {
				round.Team = 1
			}

			round.Clue = clues[rand.Intn(len(clues))]
			round.Game = 0
			round.Time = -1
			score = append(score, Score{Game: 0, TeamClues: [][]Clue{{}, {}}})

			// select a user from a team
			var usersFromTeam []User
			for _, user := range users {
				if user.Team == round.Team {
					usersFromTeam = append(usersFromTeam, user)
				}
			}
			// TODO - what if less then 2
			round.User = usersFromTeam[0]
			round.NextUser = usersFromTeam[1]


			playerPerTeam = []int{0, 0}
		}

	}

	if message.Type == "start-round" {
		round.Time = 60
        for i := range users {
            users[i].Step = 5
        }
		ticker = time.NewTicker(time.Second)
		go tickRound(roomId, ticker, *c)
	}

	if message.Type == "send-guess" {
		if message.Data == "correct" {
			correctGuess(ticker, &round, users, clues, &score, playerPerTeam)
		}
		// TODO - error

	}

	var response Response
	response.Users = users
	response.Round = round
	response.Score = score

	val, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return
	}
	c.hub.broadcast <- []byte(val)

	usersPerRoom[roomId] = users
	roundPerRoom[roomId] = round
	cluesPerRoom[roomId] = clues
	scorePerRoom[roomId] = score
	tickerPerRoom[roomId] = ticker
	playerPerTeamPerRoom[roomId] = playerPerTeam
}

func startTimer(c *Client, roomId string) {
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
