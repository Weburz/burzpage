/*
Package comments provides HTTP handlers for managing comments on articles in a RESTful
API.

Handlers include:
  - GetComments: Handles the request to retrieve a list of all comments.
  - GetCommentFromArticle: Handles the request to retrieve comments from a specific
    article by its ID.
  - AddComment: Handles the request to add a new comment to an article.
  - RemoveComment: Handles the request to delete a specific comment by its ID.
*/
package comments

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

/*
Comment represents a user comment on an article.

Fields:
  - ID: The unique identifier for the comment (UUID).
  - Name: The name of the person who made the comment.
  - Email: The email address of the person who made the comment.
  - Content: The text content of the comment.
*/
type Comment struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Content string    `json:"content"`
}

/*
GetComments handles the HTTP GET request to retrieve a list of all comments.

It simulates fetching a list of comments from an in-memory data source and returns
a JSON response with a 200 OK status code.

Response:
  - 200 OK: A JSON object containing an array of comments.
*/
func GetComments(w http.ResponseWriter, r *http.Request) {
	// Simulate a list of comments to return
	commentID, err := uuid.NewV7()
	if err != nil {
		http.Error(w, "Comment ID Not Found", http.StatusNotFound)
		return
	}

	comments := []Comment{
		{
			ID:      commentID,
			Name:    "John Doe",
			Email:   "john.doe@example.com",
			Content: "This is a great article!",
		},
		{
			ID:      commentID,
			Name:    "Jane Smith",
			Email:   "jane.smith@example.com",
			Content: "I found this article really helpful, thanks!",
		},
		{
			ID:      commentID,
			Name:    "Alice Johnson",
			Email:   "alice.johnson@example.com",
			Content: "Interesting perspective, I learned a lot!",
		},
	}

	// Create a response object containing the list of comments
	response := map[string][]Comment{
		"comments": comments,
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Set HTTP status code to 200 OK
	w.WriteHeader(http.StatusOK)

	// Encode and return the JSON response
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
	}
}

/*
GetCommentFromArticle handles the HTTP GET request to retrieve comments from a specific
article by its ID.

Currently, this function does not contain any functionality.

Response:
  - 200 OK: A JSON object containing the comments related to the specific article.
*/
func GetCommentFromArticle(w http.ResponseWriter, r *http.Request) {}

/*
AddComment handles the HTTP POST request to add a new comment to an article.

It simulates adding a new comment (in-memory) and returns the newly added comment
as a JSON response with a 201 Created status code.

Response:
  - 201 Created: A JSON object containing the newly added comment.
*/
func AddComment(w http.ResponseWriter, r *http.Request) {
	// Simulate creating a new comment (in-memory)
	commentID, err := uuid.NewV7()
	if err != nil {
		http.Error(
			w,
			"Unable to generate the Comment ID",
			http.StatusInternalServerError,
		)
		return
	}
	comment := &Comment{
		ID:      commentID,
		Name:    "Somraj Saha",
		Content: `This is an experimental comment & it doesn't serve any other purpose`,
	}

	// Create a response object containing the newly added comment
	response := map[string]Comment{
		"comment": *comment,
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Set HTTP status code to 201 Created
	w.WriteHeader(http.StatusCreated)

	// Encode and return the JSON response
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
	}
}

/*
RemoveComment handles the HTTP DELETE request to remove a specific comment by its ID.

It simulates the removal of a comment (in-memory) and returns a 204 No Content
status code indicating successful deletion with no response body.

Response:
  - 204 No Content: Successful deletion of the comment.
*/
func RemoveComment(w http.ResponseWriter, r *http.Request) {
	// Extract comment ID from URL parameters
	commentID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Comment ID Not Found", http.StatusNotFound)
		return
	}
	comment := &Comment{
		ID: commentID,
	}

	// Simulate comment removal
	fmt.Printf("Deleted comment %v\n", comment)
	comment = nil

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Set HTTP status code to 204 No Content
	w.WriteHeader(http.StatusNoContent)
}
