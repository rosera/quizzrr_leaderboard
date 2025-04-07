package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"net/url" // Add this import
	"sort"
	"strings"
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
	ipAddress    = "0.0.0.0"
)

// getGameLeaderboard view the leaderboard for a specific game
func getGameLeaderboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := strings.Split(r.URL.Path, "/")
	gameID := vars[len(vars)-1]

	// Check if gameID is empty before passing to the function
	if gameID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: Game ID is required")
		return
	}

	if players, ok := leaderboards[gameID]; ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(players)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Leaderboard for %s not found", gameID)
	}

	return
}

// setScoreGameLeaderboard updates the leaderboard for a specific game
func setScoreGameLeaderboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the incoming player data
	var newPlayer Player
	err := json.NewDecoder(r.Body).Decode(&newPlayer)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Retrieve the leaderboard for the game
	leaderboard, ok := leaderboards[newPlayer.Game]
	if !ok {
		leaderboard = []Player{}
	}

	// Append the new player and sort the leaderboard by score (descending)
	leaderboard = append(leaderboard, newPlayer)
	sort.Slice(leaderboard, func(i, j int) bool {
		return leaderboard[i].Score > leaderboard[j].Score
	})

	// Update the leaderboard
	leaderboards[newPlayer.Game] = leaderboard

	// Respond with the updated leaderboard
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leaderboard)
}

// setCancelGameLeaderboard cancels the leaderboard for a specific game
func setCancelGameLeaderboard(w http.ResponseWriter, r *http.Request) {
	vars := strings.Split(r.URL.Path, "/")
	gameID := vars[len(vars)-1]

	// Check if gameID is empty before passing to the function
	if gameID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: Game ID is required")
		return
	}

	delete(leaderboards, gameID)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, fmt.Sprintf("Leaderboard for %s reset successfully", gameID))

}

// removeScoreGameLeaderboard removes a player's score from the leaderboard
func removeScoreGameLeaderboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := strings.Split(r.URL.Path, "/")
	if len(vars) < 4 {
		http.Error(w, "Invalid request format: /remove/{gameID}/{playerName}", http.StatusBadRequest)
		return
	}

	gameID := vars[len(vars)-2]
	playerName, err := url.PathUnescape(vars[len(vars)-1])
	if err != nil {
		http.Error(w, "Failed to decode player name", http.StatusBadRequest)
		return
	}

	if gameID == "" || playerName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: Game ID and Player Name are required")
		return
	}

	leaderboard, ok := leaderboards[gameID]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Leaderboard for %s not found", gameID)
		return
	}

	found := false
	newLeaderboard := []Player{}
	for _, player := range leaderboard {
		if player.Name != playerName {
			newLeaderboard = append(newLeaderboard, player)
		} else {
			found = true
		}
	}

	if found {
		leaderboards[gameID] = newLeaderboard
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Score for player %s in game %s removed successfully", playerName, gameID)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Player %s not found in the leaderboard for %s", playerName, gameID)
	}
}


func main() {
	// Configure the CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Replace "*" with specific origins in production
		AllowedMethods:   []string{"DELETE", "GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Wrap the default HTTP server with CORS middleware
	http.Handle("/games/", c.Handler(http.HandlerFunc(getGameLeaderboard)))
	// Register the handler function
	// http.Handle("/getGameLeaderboard", handlerLeaderboard)

	// Wrap the default HTTP server with CORS middleware
	http.Handle("/points/", c.Handler(http.HandlerFunc(setScoreGameLeaderboard)))
	// Register the handler function
	// http.Handle("/setScoreGameLeaderboard", handlerScores)

	// Wrap the default HTTP server with CORS middleware
	http.Handle("/cancels/", c.Handler(http.HandlerFunc(setCancelGameLeaderboard)))
	// Register the handler function
	// http.Handle("/setCancelGameLeaderboard", handlerCancel)

	// Wrap the default HTTP server with CORS middleware for remove score
	http.Handle("/remove/", c.Handler(http.HandlerFunc(removeScoreGameLeaderboard)))

	fmt.Println("Leaderboard REST API listening on port", port)
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", ipAddress, port), nil))
}
