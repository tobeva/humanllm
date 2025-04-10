package handlers

import (
	"log"
	"net/http"
	"os"
	"strings"
)

// SessionHandler handles requests to /session/{sessionName}
func SessionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling /session/ request")
	sessionName := strings.TrimPrefix(r.URL.Path, "/session/")
	if sessionName == "" || strings.Contains(sessionName, "/") {
		log.Println("404 Not Found: Invalid session name")
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
