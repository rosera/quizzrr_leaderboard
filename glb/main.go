package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"gopkg.in/yaml.v3" // Import the YAML library
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
)

// Config struct to hold the IP and Port from the YAML file
type Config struct {
	Server struct {
		IP   string `yaml:"ip"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
}

// Player represents an entry in the leaderboard
type Player struct {
	Game  string `json:"game"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

var (
	// leaderboards is a map that stores leaderboards for different games
	leaderboards = map[string][]Player{}
	config       Config
)

// loadConfig reads the configuration from the specified YAML file
func loadConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return nil
}

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
	playerNameRaw := vars[len(vars)-1]
	playerNameDecoded, err := url.PathUnescape(playerNameRaw)
	if err != nil {
		http.Error(w, "Failed to decode player name", http.StatusBadRequest)
		return
	}
	playerName := strings.TrimSpace(playerNameDecoded)

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

func exportLeaderboards(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=leaderboards.json")
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(leaderboards)
	if err != nil {
		http.Error(w, "Failed to export leaderboards", http.StatusInternalServerError)
	}
}

func importLeaderboards(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var imported map[string][]Player
	err := json.NewDecoder(r.Body).Decode(&imported)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	leaderboards = imported

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Leaderboards successfully imported.")
}

// listGamesHandler returns a list of all available game IDs.
func listGamesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	games := make([]string, 0, len(leaderboards))
	for gameID := range leaderboards {
		games = append(games, gameID)
	}
	sort.Strings(games) // Sort the game IDs for consistent output

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(games); err != nil {
		http.Error(w, "Failed to encode game list", http.StatusInternalServerError)
	}
}

func main() {
	// Load configuration from YAML file
	configFile := "glb.yaml" // You can change the filename here
	err := loadConfig(configFile)
	if err != nil {
		log.Fatalf("Error loading %s configuration: %v", configFile, err)
	}

	// Configure the CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Replace "*" with specific origins in production
		AllowedMethods:   []string{"DELETE", "GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Wrap the default HTTP server with CORS middleware
	http.Handle("/games/", c.Handler(http.HandlerFunc(getGameLeaderboard)))
	http.Handle("/points/", c.Handler(http.HandlerFunc(setScoreGameLeaderboard)))
	http.Handle("/cancels/", c.Handler(http.HandlerFunc(setCancelGameLeaderboard)))
	http.Handle("/remove/", c.Handler(http.HandlerFunc(removeScoreGameLeaderboard)))
	http.Handle("/export", c.Handler(http.HandlerFunc(exportLeaderboards)))
	http.Handle("/import", c.Handler(http.HandlerFunc(importLeaderboards)))

	// Add the new handler for listing games
	http.Handle("/list", c.Handler(http.HandlerFunc(listGamesHandler)))

	fmt.Println("Leaderboard REST API listening on", fmt.Sprintf("%s:%d", config.Server.IP, config.Server.Port))

	// TCP: IPv6 + IPv4 wildcards
	// Reverse proxy: ensure IP is set to the Host IP not a Wildcard
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.Server.IP, config.Server.Port), nil))
}
