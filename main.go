package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Define a struct to represent a Game.
type Game struct {
	ID      int    `json:"id"`       // Unique identifier for the game.
	Name    string `json:"name"`     // Name of the game.
	Console string `json:"console"`  // Gaming console/platform.
	Genre   string `json:"genre"`    // Genre or category of the game.
	DatePub string `json:"datePub"`  // Release date of the game.
}

// Create a map to store Game objects, an ID counter, and a mutex for synchronization.
var Games = make(map[int]Game)
var idCounter int
var GamesMutex sync.Mutex

// Define an HTTP handler function to get all games and return them as JSON.
func getAllGamesHandler(w http.ResponseWriter, r *http.Request) {
	GamesMutex.Lock() // Lock the mutex to protect concurrent access to the Games map.
	defer GamesMutex.Unlock() // Ensure the mutex is unlocked when the function exits.

	var GameList []Game
	for _, Game := range Games {
		GameList = append(GameList, Game)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GameList)
}

// Define an HTTP handler function to get a specific game by ID and return it as JSON.
func getGameHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	if Game, exists := Games[id]; exists {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Game)
	} else {
		// Return an HTTP error with a 404 status code if the game with the specified ID doesn't exist.
		http.Error(w, "Game not found", http.StatusNotFound)
	}
}

// Define an HTTP handler function to create a new game and add it to the map.
func createGameHandler(w http.ResponseWriter, r *http.Request) {
	var newGame Game
	if err := json.NewDecoder(r.Body).Decode(&newGame); err != nil {
		// Return an HTTP error with a 400 status code if the request body is not valid JSON.
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	GamesMutex.Lock() // Lock the mutex to protect concurrent access to the Games map.
	idCounter++
	newGame.ID = idCounter
	Games[idCounter] = newGame
	GamesMutex.Unlock() // Unlock the mutex to allow other operations.

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newGame)
}

// Define an HTTP handler function to update an existing game by ID.
func updateGameHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	var updatedGame Game
	if err := json.NewDecoder(r.Body).Decode(&updatedGame); err != nil {
		// Return an HTTP error with a 400 status code if the request body is not valid JSON.
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	GamesMutex.Lock() // Lock the mutex to protect concurrent access to the Games map.
	defer GamesMutex.Unlock() // Ensure the mutex is unlocked when the function exits.

	if _, exists := Games[id]; exists {
		updatedGame.ID = id
		Games[id] = updatedGame
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedGame)
	} else {
		// Return an HTTP error with a 404 status code if the game with the specified ID doesn't exist.
		http.Error(w, "Game not found", http.StatusNotFound)
	}
}

// Define an HTTP handler function to delete a game by ID.
func deleteGameHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	GamesMutex.Lock() // Lock the mutex to protect concurrent access to the Games map.
	defer GamesMutex.Unlock() // Ensure the mutex is unlocked when the function exits.

	if _, exists := Games[id]; exists {
		delete(Games, id) // Delete the game with the specified ID from the map.
		w.WriteHeader(http.StatusNoContent)
	} else {
		// Return an HTTP error with a 404 status code if the game with the specified ID doesn't exist.
		http.Error(w, "Game not found", http.StatusNotFound)
	}
}

func main() {
	// Define HTTP routes and handlers for various CRUD operations.
	http.HandleFunc("/api/Games", getAllGamesHandler)
	http.HandleFunc("/api/Game", getGameHandler)
	http.HandleFunc("/api/Game/create", createGameHandler)
	http.HandleFunc("/api/Game/update", updateGameHandler)
	http.HandleFunc("/api/Game/delete", deleteGameHandler)

	fmt.Print("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
