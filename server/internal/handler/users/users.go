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

	"github.com/go-chi/chi/v5"
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

/*
GetUser retrieves a user by their ID from the URL, encodes the user data as JSON,
and sends it as a response. If encoding fails, an internal server error is returned.

URL Parameter:
  - id (string): The user's unique ID.

Response:
  - JSON object containing the user's ID, name, and email.
*/
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Fetch the user ID from the request URL parameter
	userID := chi.URLParam(r, "id")

	// Get the user from the database and store it for further encoding
	user := User{
		ID:    userID,
		Name:  "Somraj Saha",
		Email: "somraj.saha@weburz.com",
	}

	// Return the response as a JSON object
	response := map[string]User{
		"user": user,
	}

	// Set the Content-Type header to indicate that the response is JSON
	w.Header().Set("Content-Type", "application/json")

	// Attempt to encode the list of users into JSON and send it to the client
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}

/*
UpdateUser updates the user details based on the provided ID in the URL and the request
body. It reads the user data from the request body, updates the user information, and
returns the updated user as a JSON response. If the user is not found or there is an
error, an appropriate error response is returned.

URL Parameter:
  - id (string): The unique identifier of the user to update.

Request Body:
  - A JSON object containing the user's new name and email.

Response:
  - A JSON object containing the updated user's ID, name, and email.
  - A 201 Created status code on successful update.
*/
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL parameter
	userID := chi.URLParam(r, "id")

	// Decode the incoming request body into the User struct
	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Simulate updating the user (e.g., in a database)
	// Here we just replace the data for the sake of example
	user := map[string]User{
		"user": {
			ID:    userID,
			Name:  updatedUser.Name,
			Email: updatedUser.Email,
		},
	}

	// Set the Content-Type header to indicate the response is JSON
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status code to be "Created 201"
	w.WriteHeader(http.StatusCreated)

	// Return the updated user as JSON
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}
