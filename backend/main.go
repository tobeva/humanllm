package main

import (
	"log"
	"net/http"
	"strings"
)

// loggingMiddleware logs all incoming requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", loggingMiddleware(fs))
	
	log.Println("Server running at http://localhost:8080")

	http.HandleFunc("/api/sessions", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving /api/sessions")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[
			{ "name": "alpha", "players": 3, "link": "/session/alpha" },
			{ "name": "beta", "players": 1, "link": "/session/beta" },
			{ "name": "gamma", "players": 5, "link": "/session/gamma" }
		]`))
	})

	http.HandleFunc("/session/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /session/ request") // Log when the handler is invoked
		sessionName := strings.TrimPrefix(r.URL.Path, "/session/")
		if sessionName == "" || strings.Contains(sessionName, "/") {
			log.Println("404 Not Found: Invalid session name")
			http.NotFound(w, r)
			return
		}
		log.Printf("Looking up session: %s\n", sessionName)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html>
			<head><title>Session: ` + sessionName + `</title></head>
			<body>
				<h1>Session: ` + sessionName + `</h1>
			</body>
		</html>`))
	})
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
