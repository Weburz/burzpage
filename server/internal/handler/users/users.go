/*
Package user provides HTTP handlers for user-related endpoints in a RESTful API.

It includes:
  - Logging of request details (method, URL, user-agent).
  - Sending a list of users in JSON format.
  - Error handling and appropriate responses (500).
  - Logging success messages with 200 OK responses.
  - User validation and detailed error messages on failure.
*/
package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/Weburz/burzcontent/server/internal/api/models"
)

// Global instance of a validator
var validate = validator.New()

// ErrorResponse represents the structure of a validation error response.
type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}

// ErrorDetail represents the details of a single validation error.
type ErrorDetail struct {
	Status int         `json:"status"` // The HTTP status code associated with the error
	Source ErrorSource `json:"source"` // A pointer to the error source
	Title  string      `json:"title"`  // A brief description of the error
	Detail string      `json:"detail"` // A detailed explanation of the error
}

// ErrorSource represents the source of the validation error, usually a field.
type ErrorSource struct {
	Pointer string `json:"pointer"` // The field pointer
}

/*
GetUsers handles HTTP GET requests to retrieve a list of users.

The handler:
 1. Sets the Content-Type header to "application/json".
 2. Logs request details including method, URL, and user-agent.
 3. Creates and sends a list of users in JSON format.
 4. Returns a 500 status if JSON encoding fails.
 5. Logs a success message and returns a 200 OK status.
*/
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Simulate generating a unique user ID for the response
	userID, err := uuid.NewV7()
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		http.Error(w, "Unable to Generate User ID", http.StatusInternalServerError)
		return
	}

	// Create a list of users to simulate a user database
	users := []models.User{
		{
			ID:    userID,
			Name:  "Somraj Saha",
			Email: "somraj.saha@weburz.com",
		},
		{
			ID:    userID,
			Name:  "John Doe",
			Email: "john.doe@example.com",
		},
	}

	// Prepare the response object
	response := map[string][]models.User{
		"users": users,
	}

	// Set an appropriate header for the response
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)

	// Encode the user data as JSON and send it in the response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}

/*
GetUser handles HTTP GET requests to retrieve a user by their unique ID.

URL Parameter:
  - id (string): The unique identifier of the user.

Response:
  - A JSON object containing the user's ID, name, and email.
  - Returns a 404 Not Found status if the user ID is invalid.
*/
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Fetch the user ID from the URL parameter
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "User ID Not Found", http.StatusNotFound)
		return
	}

	// Simulate fetching the user from a database
	user := models.User{
		ID:    userID,
		Name:  "Somraj Saha",
		Email: "somraj.saha@weburz.com",
	}

	// Prepare the response object
	response := map[string]models.User{
		"user": user,
	}

	// Set appropriate headers for the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the user data and return it as a JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}

/*
UpdateUser handles HTTP PUT requests to update a user's details based on their ID.

URL Parameter:
  - id (string): The unique identifier of the user to update.

Request Body:
  - A JSON object containing the user's updated name and email.

Response:
  - A JSON object containing the updated user's ID, name, and email.
  - Returns a 400 Bad Request status if the request body is invalid.
  - Returns a 404 Not Found status if the user ID is invalid.
  - Returns a 201 Created status after successfully updating the user.
*/
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the URL parameter
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "User ID Not Found", http.StatusNotFound)
		return
	}

	// Decode the incoming request body into the User struct
	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the user struct
	if err := validate.Struct(updatedUser); err != nil {
		// Create a custom error response
		var errorResponse ErrorResponse
		for _, err := range err.(validator.ValidationErrors) {
			// Create an error detail based on the validation failure
			errorDetail := ErrorDetail{
				Status: http.StatusUnprocessableEntity,
				Source: ErrorSource{
					Pointer: "/data/attributes/" + err.Field(),
				},
				Title: "Invalid Attribute",
				Detail: fmt.Sprintf(
					"'%s' validation failed for field: '%s'",
					err.Tag(),
					err.Field(),
				),
			}
			errorResponse.Errors = append(errorResponse.Errors, errorDetail)
		}

		// Set Content-Type header to JSON
		w.Header().Set("Content-Type", "application/vnd.api+json")

		// Set HTTP status code to 422 Unprocessable Entity
		w.WriteHeader(http.StatusUnprocessableEntity)

		// Return the custom error response in JSON format
		if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
			http.Error(
				w,
				"Unable to encode error response",
				http.StatusInternalServerError,
			)
		}
		return
	}

	// Simulate updating the user (e.g., in a database)
	// In this example, we just replace the user data for simplicity
	user := map[string]models.User{
		"user": {
			ID:    userID,
			Name:  updatedUser.Name,
			Email: updatedUser.Email,
		},
	}

	// Set an appropriate header for the response
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusCreated)

	// Encode the updated user and return it in the response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}

/*
DeleteUser handles HTTP DELETE requests to remove a user based on their ID.

URL Parameter:
  - id (string): The unique identifier of the user to delete.

Response:
  - A 204 No Content status code indicating successful deletion of the user.
  - No response body is returned.
*/
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Fetch the user ID from the URL parameter
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "User ID Not Found", http.StatusNotFound)
		return
	}

	// Simulate deleting the user (e.g., from a database)
	user := models.User{
		ID:    userID,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	// Log the deletion
	fmt.Printf("%q is deleted\n", user)

	// Set an appropriate header for the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
