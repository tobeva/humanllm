package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Package-level variable to store sessions
var sessions = []Session{
	{Name: "alpha", Players: 3, Link: "/session/alpha"},
	{Name: "beta", Players: 1, Link: "/session/beta"},
	{Name: "gamma", Players: 5, Link: "/session/gamma"},
}

type Session struct {
	Name    string `json:"name"`
	Players int    `json:"players"`
	Link    string `json:"link"`
}

// Handle /api/sessions
func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(sessions)
	} else if r.Method == http.MethodPost {
		newSession := Session{
			Name:    fmt.Sprintf("session-%d", len(sessions)+1),
			Players: 0,
			Link:    fmt.Sprintf("/session/session-%d", len(sessions)+1),
		}
		sessions = append(sessions, newSession)
		log.Printf("Added new session: %+v", newSession)
		log.Printf("Current sessions: %+v", sessions)
		w.WriteHeader(http.StatusCreated)
	}
}
