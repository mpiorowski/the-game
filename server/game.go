package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

type Message struct {
	Room   string `json:"room"`
	UserId string `json:"userId"`
	Type   string `json:"type"`
	Data   string `json:"data"`
}

type Respone struct {
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
	Team     string `json:"team"`
	Ready    bool   `json:"ready"`
	Step     int    `json:"step"`
}

type Clue struct {
	Id      int    `json:"id"`
	Word    string `json:"word"`
	Type    string `json:"type"`
	Guessed bool   `json:"guessed"`
}

type Round struct {
	Step     int    `json:"step"`
	Team     string `json:"team"`
	Clue     Clue   `json:"clue"`
	User     User   `json:"user"`
	NextUser User   `json:"nextUser"`
	Time     int    `json:"time"`
}

type Score struct {
	Round      int    `json:"round"`
	Team1Clues []Clue `json:"team1Clues"`
	Team2Clues []Clue `json:"team2Clues"`
}

var usersPerRoom = make(map[string][]User)
var cluesPerRoom = make(map[string][]Clue)
var roundPerRoom = make(map[string]Round)
var scorePerRoom = make(map[string][]Score)
var tickerPerRoom = make(map[string]*time.Ticker)

func runGame(c *Client, msg []byte, roomId string) {

	users := usersPerRoom[roomId]
	clues := cluesPerRoom[roomId]
	round := roundPerRoom[roomId]
	score := scorePerRoom[roomId]
	ticker := tickerPerRoom[roomId]

	var message Message
	var respone Respone

	err := json.Unmarshal(msg, &message)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(message)

	if message.Type == "nickname" {
		var user User
		// str to byte
		err := json.Unmarshal([]byte(message.Data), &user)
		if err != nil {
			log.Println(err)
			return
		}

		user.Team = ""
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
					users[i].Team = "red"
				} else {
					users[i].Team = "blue"
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
		clues = append(clues, newClues...)

		for i, user := range users {
			if user.Id == message.UserId {
				users[i].Step = 3
			}
		}

		// check if clues lenght equal users lenght * 3, if so, prepare the round
		if len(clues) == len(users)*3 {
			for i := range users {
				users[i].Step = 4
			}
			// randomly select a team
			if rand.Intn(2) == 0 {
				round.Team = "red"
			} else {
				round.Team = "blue"
			}

			round.Clue = clues[rand.Intn(len(clues))]
			round.Step = 1
			round.Time = -1

			// select a user from a team
			var usersFromTeam []User
			for _, user := range users {
				if user.Team == round.Team {
					usersFromTeam = append(usersFromTeam, user)
				}
			}
			round.User = usersFromTeam[0]
			round.NextUser = usersFromTeam[1]

			score = append(score, Score{Round: 1, Team1Clues: []Clue{}, Team2Clues: []Clue{}})
		}

	}

	if message.Type == "start-round" {
		round.Time = 10
		ticker = time.NewTicker(time.Second)
		go func() {
			for range ticker.C {
				round.Time -= 1

				val, err := json.Marshal(respone)
				if err != nil {
					log.Println(err)
					return
				}
				c.hub.broadcast <- []byte(val)
				if round.Time == 0 {
					ticker.Stop()

                    round.Time = -1
                    if round.Team == "red" {
                        round.Team = "blue"
                    } else {
                        round.Team = "red"
                    }

                    notGuessedClues := []Clue{}
                    for _, clue := range clues {
                        if !clue.Guessed {
                            notGuessedClues = append(notGuessedClues, clue)
                        }
                    }
                    round.Clue = notGuessedClues[rand.Intn(len(notGuessedClues))]

                    round.User
				}
			}
		}()
	}

	if message.Type == "send-guess" {
		if message.Data == "correct" {

            for i, clue := range clues {
                if clue.Id == round.Clue.Id {
                    clues[i].Guessed = true
                }
            }

            // update score
            if round.Team == "red" {
                score[len(score)-1].Team1Clues = append(score[len(score)-1].Team1Clues, round.Clue)
            } else {
                score[len(score)-1].Team2Clues = append(score[len(score)-1].Team2Clues, round.Clue)
            }

			round.User = round.NextUser
			for i, user := range users {
				if user.Id == round.User.Id {
					if i == len(users)-1 {
						round.NextUser = users[0]
					} else {
						round.NextUser = users[i+1]
					}
				}
			}
			round.Clue = clues[rand.Intn(len(clues))]

		}

	}

	respone.Users = users
	respone.Round = round
	respone.Score = score

	val, err := json.Marshal(respone)
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
