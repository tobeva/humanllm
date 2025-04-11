package main

import (
	"log"
	"net/http"
	"github.com/tobeva/humanllm/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(loggingMiddleware(fs))

	log.Println("Server running at http://localhost:8080")
	
	router.HandleFunc("/api/homepage", handlers.HomepageHandler)
	router.HandleFunc("/session/{name}", handlers.SessionHandler).Methods("GET")

	err := http.ListenAndServe(":8080", router)
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
