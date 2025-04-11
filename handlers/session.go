package handlers

import (
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/gorilla/mux"
)

// SessionHandler handles requests to /session/{name}
func SessionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling /session/ request")
	vars := mux.Vars(r)
	sessionName := vars["name"]
	if sessionName == "" {
		log.Println("404 Not Found: Missing session name")
		http.NotFound(w, r)
		return
	}

	log.Printf("Looking up session: %s\n", sessionName)

	// Read the session.html file
	htmlFilePath := "static/session.html"
	htmlContent, err := os.ReadFile(htmlFilePath)
	if err != nil {
		log.Printf("Error reading HTML file: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Replace the placeholder with the session name
	pageContent := strings.ReplaceAll(string(htmlContent), "{{sessionName}}", sessionName)

	// Serve the modified HTML
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(pageContent))
}
