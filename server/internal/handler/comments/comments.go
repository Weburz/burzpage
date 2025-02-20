/*
Package comments provides HTTP handlers for managing comments on articles.

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
)

/*
Comment represents a user comment on an article.

Fields:
  - Name: The name of the person who made the comment.
  - Email: The email address of the person who made the comment.
  - Content: The text content of the comment.
*/
type Comment struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Content string `json:"content"`
}

/*
GetComments handles the HTTP GET request to retrieve a list of all comments.

It simulates fetching a list of comments and returns the comments as a JSON response
with a 200 OK status code.

Response:
  - A 200 OK status code with a JSON array containing the comments.
*/
func GetComments(w http.ResponseWriter, r *http.Request) {
	// Simulate a list of comments to return
	comments := []Comment{
		{
			Name:    "John Doe",
			Email:   "john.doe@example.com",
			Content: "This is a great article!",
		},
		{
			Name:    "Jane Smith",
			Email:   "jane.smith@example.com",
			Content: "I found this article really helpful, thanks!",
		},
		{
			Name:    "Alice Johnson",
			Email:   "alice.johnson@example.com",
			Content: "Interesting perspective, I learned a lot!",
		},
	}

	// Create an object to encode as a JSON response
	response := map[string][]Comment{
		"comments": comments,
	}

	// Set the Content-Type to be of JSON
	w.Header().Set("Content-Type", "application/json")

	// Set an appropriate HTTP status code for the response
	w.WriteHeader(http.StatusOK)

	// Return an appropriate JSON response from the handler
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
	}
}

func GetCommentFromArticle(w http.ResponseWriter, r *http.Request) {}

func AddComment(w http.ResponseWriter, r *http.Request) {
	// Simulate an example comment for experimentation
	comment := &Comment{
		Name:    "Somraj Saha",
		Content: `This is an experimental comment & it doesn't server anyother purpose`,
	}

	// Create an object to encode as a JSON response
	response := map[string]Comment{
		"comment": *comment,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
	}
}

func RemoveComment(w http.ResponseWriter, r *http.Request) {
	commentID := chi.URLParam(r, "id")
	comment := &Comment{
		ID: commentID,
	}
	fmt.Printf("Deleted comment %v\n", comment)
	comment = nil

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
