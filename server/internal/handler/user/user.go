/*
Package user provides HTTP handlers for user-related endpoints.

This package includes a handler `UserHandler` that responds with a simple JSON
message. The handler logs both incoming request details and the outcome of processing
(success or error).

Key functionalities:
  - Logs the method, URL, and user-agent for incoming requests.
  - Sends a JSON response with a "Hello Jarmos!" message.
  - Logs error details if JSON encoding fails, and responds with an internal server
    error (500).
  - Logs a success message and returns a 200 OK status upon successful response.
*/
package user

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// message represents the structure of the response message
type message struct {
	Message string `json:"message"` // The message returned in the response
}

/*
GetUserHandler handles HTTP requests to the user endpoint.
It responds with a simple message and logs the request and response details.

The handler performs the following tasks:
 1. Sets the Content-Type header to "application/json".
 2. Logs the incoming request, including the HTTP method, URL, and user-agent.
 3. Attempts to encode a message into JSON and send it in the response body.
 4. Logs any errors if JSON encoding fails, and returns an internal server error
    (500) to the client.
 5. Logs a success message with a 200 OK status after successfully
    sending the response.
*/
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate that the response is JSON
	w.Header().Set("Content-Type", "application/json")

	// Create the response message
	response := message{
		Message: "Hello Jarmos!",
	}

	// Log details about the incoming request (method, URL, and user-agent)
	slog.Info(
		"Handling request",
		slog.String("method", r.Method),
		slog.String("url", r.URL.String()),
		slog.String("user-agent", r.UserAgent()),
	)

	// Attempt to encode the response into JSON and send it to the client
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// If encoding fails, log the error and return an internal server error response
		slog.Error(
			"Unable to encode JSON",
			slog.String("error", err.Error()),
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
		)
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}

	// Log the successful response along with the HTTP status code (200 OK)
	slog.Info(
		"Successfully processed request",
		slog.String("method", r.Method),
		slog.String("url", r.URL.String()),
		slog.Int("status", http.StatusOK),
	)
}
