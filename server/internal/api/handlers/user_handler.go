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
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
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

/*
GetUser handles HTTP requests to retrieve a user's information by their ID.

This function performs the following steps:

 1. Retrieves the user ID from the URL parameter `id` using `chi.URLParam(r, "id")`.
 2. Attempts to parse the user ID using `uuid.Parse()`. If parsing fails, an error
    response is returned with an HTTP 404 (Not Found) status, indicating that the user
    ID could not be found or is invalid.
 3. Creates a mock `User` object with the parsed user ID, name, and email.
 4. Responds with the user data in a JSON format and a HTTP 200 (OK) status code,
    indicating that the user data has been successfully retrieved.

Example:
  - When a GET request is made to `/users/{id}`, this function will retrieve the user
    associated with the specified ID (mocked data for now) and return the user details
    in the response body.

Error Handling:
  - If the user ID cannot be parsed from the URL, the function responds with a 404
    status and an error message indicating that the user ID was not found.
  - If the JSON encoding for the response fails, the function responds with a 500 status
    (Internal Server Error).

Note: The user retrieval process in this function is mocked; no actual user data is
fetched from a database or persistent storage.
*/
func (ur *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "User ID Not Found", http.StatusNotFound)
		return
	}

	user := models.User{
		ID:    userID,
		Name:  "Somraj Saha",
		Email: "somraj.saha@weburz.com",
	}

	response := map[string]models.User{
		"user": user,
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}

/*
EditUser handles HTTP requests to update an existing user's information.

This function performs the following steps:

 1. Retrieves the user ID from the URL parameter `id` using `chi.URLParam(r, "id")`.
 2. Attempts to parse the user ID using `uuid.Parse()`. If parsing fails, an error
    response is returned with an HTTP 404 (Not Found) status, indicating that the user
    ID could not be found or is invalid.
 3. Decodes the incoming request body into an updated `models.User` object, which
    contains the new user details (name, email).
 4. Validates the decoded user data using the `validator` package. If validation fails,
    an HTTP 422 (Unprocessable Entity) status is returned with an error message.
 5. Creates a new `User` object with the updated user ID, name, and email.
 6. Responds with a JSON representation of the updated user, along with an HTTP 201
    (Created) status code, indicating the update was successful.

Example:
  - When a PUT request is made to `/users/{id}/edit` with a valid user ID in the URL
    and updated user data in the request body
    (e.g., `{"name": "Jane Doe", "email": "jane.doe@example.com"}`), this function will
    update the user information and return the updated user data in the response.

Error Handling:
  - If the user ID cannot be parsed from the URL or does not exist, the function
    responds with a 404 status and an error message.
  - If the request body is invalid or cannot be parsed, the function responds with a
    400 status (Bad Request) and an error message.
  - If the request validation fails (e.g., missing or invalid fields), the function
    responds with a 422 status and an error message indicating validation failure.
  - If the JSON encoding for the response fails, the function responds with a 500
    status (Internal Server Error).

Note: The user update process in this function is mocked; no actual user data is stored
in a database or persistent storage.
*/
func (ur *UserHandler) EditUser(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()

	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "User ID Not Found", http.StatusNotFound)
		return
	}

	var updatedUser models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(updatedUser); err != nil {
		http.Error(w, "Request validation failed", http.StatusUnprocessableEntity)
		return
	}

	user := &models.User{
		ID:    userID,
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
	}

	response := map[string]models.User{
		"user": *user,
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}

/*
CreateUser handles HTTP requests to create a new user.

This function performs the following steps:

 1. Decodes the incoming request body into a `models.User` object.
 2. Validates the decoded user data using the `validator` package. If validation fails,
    an HTTP 422 (Unprocessable Entity) status is returned along with an error message.
 3. Attempts to generate a new user ID using `uuid.NewV7()`. If ID generation fails,
    an HTTP 500 (Internal Server Error) is returned.
 4. Creates a new `User` object with the generated user ID, name, and email.
 5. Returns a JSON response with the newly created user, including their ID, along with
    a HTTP 201 (Created) status code.

Example:
  - When a POST request is made to `/users/new` with a valid user JSON payload (e.g.,
    `{"name": "Jane Doe", "email": "jane.doe@example.com"}`), this function will create
    a new user, assign them a unique ID, and respond with a 201 status along with the
    user data in the response body.

Error Handling:
  - If the request body is invalid or cannot be parsed, the function responds with a
    500 status and an error message.
  - If the request validation fails, the function responds with a 422 status and an
    error message indicating validation failure.
  - If the user ID generation fails, the function responds with a 500 status.

Note: The user creation process in this function is mocked; no actual user is stored
in a database or persistent storage.
*/
func (ur *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()

	var newUser models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusInternalServerError)
		return
	}

	if err := validate.Struct(newUser); err != nil {
		http.Error(w, "Request validation failed", http.StatusUnprocessableEntity)
		return
	}

	userID, err := uuid.NewV7()
	if err != nil {
		http.Error(w, "Failed to generate a User ID", http.StatusInternalServerError)
	}
	user := &models.User{
		ID:    userID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	response := map[string]models.User{
		"user": *user,
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
	}
}

/*
DeleteUser handles HTTP requests to delete a user by their ID.

This function performs the following steps:

 1. Retrieves the user ID from the URL parameter `id` using `chi.URLParam(r, "id")`.
 2. Attempts to parse the user ID using `uuid.Parse()`. If parsing fails, an error
    response is returned with an HTTP 404 (Not Found) status, indicating that the
    user ID could not be found or is invalid.
 3. Creates a mock `User` object with the parsed user ID, name, and email.
 4. Prints a message to the server log indicating that the user has been deleted.
 5. Responds with an HTTP 204 (No Content) status code, indicating successful
    deletion, though no content is returned in the response body.

Note: In this implementation, the user deletion is mocked, and no actual user is deleted
from a database or persistent storage.
*/
func (ur *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "User ID Not Found", http.StatusNotFound)
	}

	user := models.User{
		ID:    userID,
		Name:  "John Doe",
		Email: "john.doe@weburz.com",
	}

	fmt.Printf("%q is deleted!\n", user)

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusNoContent)
}
