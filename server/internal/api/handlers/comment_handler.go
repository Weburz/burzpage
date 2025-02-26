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

type CommentHandler struct{}

func NewCommentHandler() *CommentHandler {
	return &CommentHandler{}
}

/* GetComments will fetch all comments available on the system. */
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

/* GetCommentsFromArticle will fetch all comments from an article (by ID). */
func (cr *CommentHandler) GetCommentsFromArticle(
	w http.ResponseWriter,
	r *http.Request,
) {
}

/* AddComment will add a new comment to an article (by ID). */
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

/* RemoveComment will remove a comment from an article (by ID). */
func (cr *CommentHandler) RemoveComment(w http.ResponseWriter, r *http.Request) {
	commentID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Comment ID Not Found", http.StatusNotFound)
		return
	}

	comment := models.Comment{ID: commentID}

	fmt.Printf("%q is deleted!\n", comment)

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusNoContent)
}
