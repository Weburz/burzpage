/*
Package articles provides HTTP handlers for managing articles in a RESTful API.

Handlers include:
- GetArticles: Handles the request to retrieve a list of articles.
- GetArticle: Handles the request to retrieve a specific article by its ID.
- CreateArticle: Handles the request to create a new article.
- EditArticle: Handles the request to update an existing article by its ID.
- DeleteArticle: Handles the request to delete an article by its ID.
*/
package articles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/Weburz/burzcontent/server/internal/api/models"
)

/*
GetArticles handles the HTTP GET request to retrieve a list of articles.

It simulates fetching a list of articles from an in-memory data source and returns
a JSON response with a 200 OK status code.

Response:
  - 200 OK: A JSON object containing an array of articles.
*/
func GetArticles(w http.ResponseWriter, r *http.Request) {
	articleID, err := uuid.NewV7()
	if err != nil {
		http.Error(
			w,
			"Unable to generate the Article ID",
			http.StatusInternalServerError,
		)
		return
	}
	articles := []models.Article{
		{
			ID:        articleID,
			Title:     "Go Programming Basics",
			Author:    "John Doe",
			Published: true,
		},
		{
			ID:        articleID,
			Title:     "Advanced Go Techniques",
			Author:    "Jane Smith",
			Published: false,
		},
		{
			ID:        articleID,
			Title:     "Understanding Go Concurrency",
			Author:    "Alice Johnson",
			Published: true,
		},
	}

	// Create response object containing articles
	response := map[string][]models.Article{
		"articles": articles,
	}

	// Encode the response object to JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Set HTTP status code to 200 OK
	w.WriteHeader(http.StatusOK)
}

/*
GetArticle handles the HTTP GET request to retrieve a specific article by its ID.

It extracts the article ID from the URL parameters, simulates fetching the article,
and returns it as a JSON response with a 200 OK status code.

Response:
  - 200 OK: A JSON object containing the requested article.
*/
func GetArticle(w http.ResponseWriter, r *http.Request) {
	// Extract article ID from URL parameters
	articleID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Article ID Not Found", http.StatusNotFound)
		return
	}

	// Simulate fetching the article
	article := models.Article{
		ID:        articleID,
		Title:     "Understanding Go Concurrency",
		Author:    "Alice Johnson",
		Published: true,
	}

	// Create response object containing the article
	response := map[string]models.Article{
		"article": article,
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Set HTTP status code to 200 OK
	w.WriteHeader(http.StatusOK)

	// Encode and return the JSON response
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}

/*
CreateArticle handles the HTTP POST request to create a new article.

It simulates creating a new article (in-memory) and returns the created article
as a JSON response with a 201 Created status code.

Response:
  - 201 Created: A JSON object containing the newly created article.
*/
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	// Simulate article creation (in-memory)
	articleID, err := uuid.NewV7()
	if err != nil {
		http.Error(
			w,
			"Unable to generate the Article ID",
			http.StatusInternalServerError,
		)
		return
	}
	article := models.Article{
		ID:        articleID,
		Title:     "Learn to Build REST API in Go",
		Author:    "Somraj Saha",
		Published: false,
	}

	// Create response object containing the article
	response := map[string]models.Article{
		"article": article,
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Set HTTP status code to 201 Created
	w.WriteHeader(http.StatusCreated)

	// Encode and return the JSON response
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}

/*
EditArticle handles the HTTP PUT request to update an existing article by its ID.

It simulates updating the article and returns the updated article as a JSON response
with a 200 OK status code.

Response:
  - 200 OK: A JSON object containing the updated article.
*/
func EditArticle(w http.ResponseWriter, r *http.Request) {
	// Extract article ID from URL parameters
	articleID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Article ID Not Found", http.StatusNotFound)
		return
	}

	// Simulate creating a new article to replace the old one
	newArticle := models.Article{
		ID:        articleID,
		Title:     "Learn to Build REST API in Go",
		Author:    "John Doe",
		Published: false,
	}

	// Create response object containing the updated article
	response := map[string]models.Article{
		"article": newArticle,
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Set HTTP status code to 200 OK
	w.WriteHeader(http.StatusOK)

	// Encode and return the JSON response
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}

/*
DeleteArticle handles the HTTP DELETE request to remove an article by its ID.

It simulates the deletion of an article and returns a 204 No Content status code to
indicate successful deletion with no response body.

Response:
  - 204 No Content: Successful deletion of the article.
*/
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	// Extract article ID from URL parameters
	articleID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Article ID Not Found", http.StatusNotFound)
		return
	}

	// Simulate article deletion
	article := &models.Article{
		ID:        articleID,
		Title:     "Learn to Build REST API in Go",
		Author:    "John Doe",
		Published: false,
	}

	// Log the deletion (for simulation purposes)
	fmt.Printf("Article: %v\n", article)

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Set HTTP status code to 204 No Content
	w.WriteHeader(http.StatusNoContent)
}
