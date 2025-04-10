package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Session struct {
	Name    string `json:"name"`
	Players int    `json:"players"`
	Link    string `json:"link"`
}

// Handle /api/sessions
func SessionsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving /api/sessions")

	// Create a slice of sessions
	sessions := []Session{
		{Name: "alpha", Players: 3, Link: "/session/alpha"},
		{Name: "beta", Players: 1, Link: "/session/beta"},
		{Name: "gamma", Players: 5, Link: "/session/gamma"},
	}

	// Serialize the sessions to JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}
