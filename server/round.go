package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

func correctGuess(ticker *time.Ticker, round *Round, users []User, clues []Clue, score *[]Score, playerPerTeam []int) {

	for i, clue := range clues {
		if clue.Id == round.Clue.Id {
			clues[i].Guessed = true
		}
	}

	// update score
	(*score)[round.Game].TeamClues[round.Team] = append((*score)[round.Game].TeamClues[round.Team], round.Clue)

	// change user
	var usersFromTeam []User
	for _, user := range users {
		if user.Team == round.Team {
			usersFromTeam = append(usersFromTeam, user)
		}
	}
	round.User = round.NextUser
	playerPerTeam[round.Team] = (playerPerTeam[round.Team] + 1) % len(usersFromTeam)
	round.NextUser = usersFromTeam[(playerPerTeam[round.Team]+1)%len(usersFromTeam)]

	// assign new clue
	notGuessedClues := []Clue{}
	for _, clue := range clues {
		if !clue.Guessed {
			notGuessedClues = append(notGuessedClues, clue)
		}
	}

	if len(notGuessedClues) > 0 {
		round.Clue = notGuessedClues[rand.Intn(len(notGuessedClues))]
	}

	// last clue, finish game, start new round
	if len(notGuessedClues) == 0 {
		ticker.Stop()
		round.Game = round.Game + 1

		// Finish game
		if round.Game == 3 {
			for i := range users {
				users[i].Step = 6
			}
			return
		} else {
			for i := range users {
				users[i].Step = 4
			}
		}

		round.Time = -1
		round.Team = 1 - round.Team

		newScore := Score{
			Game:      round.Game,
			TeamClues: [][]Clue{{}, {}},
		}

		*score = append(*score, newScore)

		// make all clues available again
		for i := range clues {
			clues[i].Guessed = false
		}

		// change user
		var usersFromTeam []User
		for _, user := range users {
			if user.Team == round.Team {
				usersFromTeam = append(usersFromTeam, user)
			}
		}
		round.User = usersFromTeam[playerPerTeam[round.Team]]
		round.NextUser = usersFromTeam[(playerPerTeam[round.Team]+1)%len(usersFromTeam)]
	}

}

func endRound(ticker *time.Ticker, round *Round, users []User, clues []Clue, playerPerTeam []int) {
	ticker.Stop()

	// change team
	round.Time = -1
	round.Team = 1 - round.Team

	// change user
	var usersFromTeam []User
	for _, user := range users {
		if user.Team == round.Team {
			usersFromTeam = append(usersFromTeam, user)
		}
	}
	round.User = usersFromTeam[playerPerTeam[round.Team]]
	round.NextUser = usersFromTeam[(playerPerTeam[round.Team]+1)%len(usersFromTeam)]

	// assign new clue
	notGuessedClues := []Clue{}
	for _, clue := range clues {
		if !clue.Guessed {
			notGuessedClues = append(notGuessedClues, clue)
		}
	}
	round.Clue = notGuessedClues[rand.Intn(len(notGuessedClues))]
}

func tickRound(roomId string, ticker *time.Ticker, c Client) {
	for range ticker.C {
		users := usersPerRoom[roomId]
		clues := cluesPerRoom[roomId]
		round := roundPerRoom[roomId]
		score := scorePerRoom[roomId]
		ticker := tickerPerRoom[roomId]

		playerPerTeam := playerPerTeamPerRoom[roomId]

		round.Time -= 1
		/**
		 * End round
		 */
		if round.Time == 0 {
            endRound(ticker, &round, users, clues, playerPerTeam)
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
}
