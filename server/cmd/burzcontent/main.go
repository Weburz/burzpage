package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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

	// Get the runtime environment type
	mode := os.Getenv("ENV")

	// Get the port to access the server at
	port := os.Getenv("PORT")

	// Conditionally load the HTTP server according to the runtime environment it is
	// invoked from
	if mode == "production" {
		// Start the production HTTP server
		msg := fmt.Sprintf(
			"Server started in PRODUCTION mode at http://:::%s",
			port,
		)
		logger.Info(msg)
		http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	} else {
		// Start the development HTTP server
		msg := fmt.Sprintf(
			"Server started in DEVELOPMENT mode at http://127.0.0.1:%s",
			port,
		)
		logger.Info(msg)
		http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	}
}
