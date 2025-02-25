/*
Package handlers defines various request handlers, including user-related operations.

This package provides handlers for processing HTTP requests. The `UserHandler`
in this file handles user-related operations, such as retrieving a list of users.
The handlers are used by the application to define specific routes for user requests
and send responses with appropriate data or error messages.
*/
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/Weburz/burzcontent/server/internal/api/models"
)

// UserHandler handles HTTP requests related to users, including retrieving user data.
type UserHandler struct{}

/*
NewUserHandler creates and initializes a new instance of UserHandler.

This function returns a new `UserHandler` instance, which is ready to handle
user-related HTTP requests.
*/
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

/*
GetUsers handles HTTP requests to retrieve a list of users.

This function performs the following steps:

 1. Generates a new user ID using `uuid.NewV7()`. If it fails, it logs the error
    and returns an HTTP error response.
 2. Creates a hardcoded list of two users, with each user having a unique ID,
    name, and email address.
 3. Responds with the user data in a JSON format under the key "users".
 4. Sets the `Content-Type` header to `application/vnd.api+json` and returns an
    HTTP 200 status code if successful. If encoding the JSON fails, it returns
    an HTTP error response.

The response will contain a JSON array of users, each represented by their ID,
name, and email address.
*/
func (ur *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.NewV7()
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		http.Error(w, "Unable to Generate User ID", http.StatusInternalServerError)
		return
	}

	users := []models.User{
		{
			ID:    userID,
			Name:  "Somraj Saha",
			Email: "somraj.saha@weburz.com",
		},
		{
			ID:    userID,
			Name:  "John Doe",
			Email: "john.doe@weburz.com",
		},
	}

	response := map[string][]models.User{
		"users": users,
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}
