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
package users

import (
	"encoding/json"
	"net/http"
)

// message represents the structure of the response message
type message struct {
	Message string `json:"message"` // The message returned in the response
}

/*
GetUsersHandler handles HTTP requests to the user endpoint.
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
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate that the response is JSON
	w.Header().Set("Content-Type", "application/json")

	// Create the response message
	response := message{
		Message: "Hello Jarmos!",
	}

	// Attempt to encode the response into JSON and send it to the client
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}
