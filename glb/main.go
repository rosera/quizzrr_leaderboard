package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

// Player represents an entry in the leaderboard
type Player struct {
	Game  string `json:"game"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

var (
	// leaderboards is a map that stores leaderboards for different games
	leaderboards = map[string][]Player{}
	port         = 8080
)

func getGameLeaderboardEntries(game string, w http.ResponseWriter) {

  // Add Cors Headers
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if game == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: Game parameter is required")
		return
	}
	if players, ok := leaderboards[game]; ok {
		json.NewEncoder(w).Encode(players)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Leaderboard for %s not found", game)
	}

	return
}

func setGameLeaderboardScore(newPlayer Player, w http.ResponseWriter) {
	leaderboard, ok := leaderboards[newPlayer.Game]

	if !ok {
		leaderboard = []Player{}
	}

	leaderboard = append(leaderboard, newPlayer)
	// Sort based on score in descending order
	sort.Slice(leaderboard, func(i, j int) bool {
		return leaderboard[i].Score > leaderboard[j].Score
	})

	leaderboards[newPlayer.Game] = leaderboard
	json.NewEncoder(w).Encode(leaderboard)
}

func setGameLeaderBoardDelete(game string, w http.ResponseWriter) {
	if game == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: Game parameter is required")
		return
	}

	delete(leaderboards, game)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, fmt.Sprintf("Leaderboard for %s reset successfully", game))
}

func main() {
	http.HandleFunc("/leaderboard", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			game := r.URL.Query().Get("game")
			getGameLeaderboardEntries(game, w)

		case http.MethodPost:
			var newPlayer Player
			err := json.NewDecoder(r.Body).Decode(&newPlayer)
			if err != nil || newPlayer.Name == "" || newPlayer.Score == 0 || newPlayer.Game == "" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Error: Game, Name and Score are required")
				return
			}
			setGameLeaderboardScore(newPlayer, w)

		case http.MethodDelete:
			game := r.URL.Query().Get("game")
			setGameLeaderBoardDelete(game, w)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method %s not allowed", r.Method)
		}
	})

	fmt.Println("Leaderboard API listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
