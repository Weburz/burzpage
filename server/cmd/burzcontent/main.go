package main

import (
	"encoding/json"
	"net/http"

	"github.com/Weburz/burzcontent/server/internal/logger"
)

// Define a struct to represent the JSON response
type Message struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Set the response header to indicate we're sending JSON
	w.Header().Set("Content-Type", "application/json")

	// Create the response object
	response := Message{
		Message: "Hello World!",
	}

	// Encode the response struct as JSON and write it to the response body
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
	}
}

func main() {
	// Handle the root path and link it to the handler function
	http.HandleFunc("/", handler)

	// Create an instance of a logger
	logger := logger.NewLogger()

	// Start the HTTP server on port 8080
	logger.Info("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
