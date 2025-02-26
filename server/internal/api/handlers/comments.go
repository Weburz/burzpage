/*
Package handlers provides HTTP handlers for managing comments in the system.

This package includes various handler functions related to comment management,
including:
  - Retrieving all comments (`GetComments`)
  - Adding a new comment (`AddComment`)
  - Removing an existing comment (`RemoveComment`)
  - (Planned) Retrieving comments for a specific article (`GetCommentsFromArticle`)

The `CommentHandler` struct defines methods that handle HTTP requests related to
comments.

The package interacts with the `models` package for defining comment structures and
validation.
*/
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Weburz/burzcontent/server/internal/api/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// CommentHandler handles HTTP requests related to comments.
type CommentHandler struct{}

// NewCommentHandler creates and returns a new instance of CommentHandler.
func NewCommentHandler() *CommentHandler {
	return &CommentHandler{}
}

// GetComments will fetch all comments available on the system.
// It returns a list of comments in JSON format.
func (cr *CommentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	commentID, err := uuid.NewV7()
	if err != nil {
		http.Error(w, "Comment ID Not Found", http.StatusNotFound)
		return
	}

	comments := []models.Comment{
		{
			ID:      commentID,
			Name:    "John Doe",
			Email:   "john.doe@example.com",
			Content: "This is a great article",
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
		{
			ID:      commentID,
			Name:    "Somraj Saha",
			Email:   "somraj.saha@weburz.com",
			Content: "This is a test comment for experimental purposes ONLY!",
		},
	}

	response := map[string][]models.Comment{
		"comments": comments,
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}

// GetCommentsFromArticle will fetch all comments for a specific article.
// This method currently has no implementation.
func (cr *CommentHandler) GetCommentsFromArticle(
	w http.ResponseWriter,
	r *http.Request,
) {
	// TODO: Implement functionality to fetch comments by article ID
}

// AddComment will add a new comment to an article.
// It accepts a comment in the request body, validates it, and adds it to the system.
// Returns the created comment in the response.
func (cr *CommentHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	// Unmarshal the request into a struct for processing
	var newComment models.Comment
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newComment); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the request body
	validate := validator.New()
	if err := validate.Struct(newComment); err != nil {
		http.Error(w, "Request body validation failed", http.StatusUnprocessableEntity)
		return
	}

	// Generate an unique identifier for the comment
	commentID, err := uuid.NewV7()
	if err != nil {
		http.Error(
			w,
			"Unable to generate the Comment ID",
			http.StatusInternalServerError,
		)
		return
	}

	// Create a comment instance
	comment := &models.Comment{
		ID:      commentID,
		Name:    newComment.Name,
		Email:   newComment.Email,
		Content: newComment.Content,
	}

	// Create a response to encode and send to the client
	response := map[string]models.Comment{
		"comment": *comment,
	}

	// Set appropriate HTTP headers
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusCreated)

	// Encode the response appropriately before sending it to the client
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Unable to encoder JSON", http.StatusInternalServerError)
		return
	}
}

// RemoveComment will remove a comment from the system by its ID.
// It expects the comment ID as a URL parameter and returns a successful response if
// the comment is deleted.
func (cr *CommentHandler) RemoveComment(w http.ResponseWriter, r *http.Request) {
	commentID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Comment ID Not Found", http.StatusNotFound)
		return
	}

	comment := models.Comment{ID: commentID}

	// Log the deletion (this would be replaced by actual deletion logic)
	fmt.Printf("%q is deleted!\n", comment)

	// Set HTTP response headers and status code for a successful deletion
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusNoContent)
}
