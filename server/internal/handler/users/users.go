/*
Package user provides HTTP handlers for user-related endpoints.

This package includes a handler `UserHandler` that responds with a list of users
in JSON format. The handler logs both incoming request details and the outcome
of processing (success or error).

Key functionalities:
  - Logs the method, URL, and user-agent for incoming requests.
  - Sends a JSON response with a list of users.
  - Logs error details if JSON encoding fails, and responds with an internal
    server error (500).
  - Logs a success message and returns a 200 OK status upon successful response.
*/
package users

import (
	"encoding/json"
	"net/http"
)

// User represents the structure of a User
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

/*
GetUsersHandler handles HTTP requests to the user endpoint.
It responds with a list of users in JSON format.

The handler performs the following tasks:
 1. Sets the Content-Type header to "application/json".
 2. Logs the incoming request, including the HTTP method, URL, and user-agent.
 3. Attempts to encode the list of users into JSON and send it in the response body.
 4. Logs any errors if JSON encoding fails, and returns an internal server error
    (500) to the client.
 5. Logs a success message with a 200 OK status after successfully sending the response.
*/
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate that the response is JSON
	w.Header().Set("Content-Type", "application/json")

	// Create a list of users
	users := []User{
		{
			ID:    "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
			Name:  "Somraj Saha",
			Email: "somraj.saha@weburz.com",
		},
		{
			ID:    "4f3e23f4-d5d9-4886-90de-f07a93d3c7f5",
			Name:  "John Doe",
			Email: "john.doe@example.com",
		},
	}

	// The response to return with when calling this handler
	response := map[string][]User{
		"users": users,
	}

	// Attempt to encode the list of users into JSON and send it to the client
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}
