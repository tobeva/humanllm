package main

import (
	"log"
	"net/http"
	"humanllm/backend/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", loggingMiddleware(fs))

	log.Println("Server running at http://localhost:8080")
	
	http.HandleFunc("/session/", handlers.SessionHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// loggingMiddleware logs all incoming requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
