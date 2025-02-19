/*
Package articles provides HTTP handlers for managing articles.

- GetArticles: Handles the request to retrieve a list of articles.
- GetArticle: Handles the request to retrieve a single article by its ID.
- CreateArticle: Handles the request to create a new article.
- EditArticle: Handles the request to update an existing article by its ID.
- DeleteArticle: Handles the request to delete an article by its ID.
*/
package articles

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
Article represents an article with a title, author, and publication status.

Fields:
- Title: The title of the article.
- Author: The author of the article.
- Published: A boolean indicating if the article is published.
*/
type Article struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Published bool   `json:"published"`
}

/*
GetArticles handles the HTTP GET request to retrieve a list of articles.

It simulates fetching articles from an in-memory data source and returns the list
as a JSON response with a 200 OK status code.

Response:
  - A 200 OK status code with a JSON array of articles.
*/
func GetArticles(w http.ResponseWriter, r *http.Request) {
	// Simulate a list of in-memory article
	articles := []Article{
		{
			ID:        "2eee23d1-f1d3-4c1e-b22d-f789098c55e7",
			Title:     "Go Programming Basics",
			Author:    "John Doe",
			Published: true,
		},
		{
			ID:        "f83d03bd-cc83-4a19-ae27-099fc8d8fa66",
			Title:     "Advanced Go Techniques",
			Author:    "Jane Smith",
			Published: false,
		},
		{
			ID:        "377be0b6-49ec-43e5-b02e-05e15c95e964",
			Title:     "Understanding Go Concurrency",
			Author:    "Alice Johnson",
			Published: true,
		},
	}

	// Create a response object for the handler to return
	response := map[string][]Article{
		"articles": articles,
	}

	// Encode the struct object to JSON for the handler to return
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type to be of JSON
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status code to be of "200 OK"
	w.WriteHeader(http.StatusOK)
}

/*
GetArticle handles the HTTP GET request to retrieve a specific article by its ID.

It extracts the article ID from the URL parameters, simulates fetching the article,
and returns it as a JSON response with a 200 OK status code.

Response:
  - A 200 OK status code with a JSON object containing the article.
*/
func GetArticle(w http.ResponseWriter, r *http.Request) {
	// Get the article ID from the URL parameters
	articleID := chi.URLParam(r, "id")

	article := Article{
		ID:        articleID,
		Title:     "Understanding Go Concurrency",
		Author:    "Alice Johnson",
		Published: true,
	}

	response := map[string]Article{
		"article": article,
	}

	// Set the "Content-Type" to be of JSON
	w.Header().Set("Content-Type", "application/json")

	// Set an appropriate HTTP status code to return along with the response
	w.WriteHeader(http.StatusOK)

	// Return an JSON response after serialising the struct
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
  - A 201 Created status code with a JSON object containing the created article.
*/
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	// Simulate article creation (in-memory only for now)
	article := Article{
		ID:        "1dd9eb62-5af5-44b8-bc5f-b256f6f8f2ee",
		Title:     "Learn to Build REST API in Go",
		Author:    "John Doe",
		Published: false,
	}

	// Create a response object using the "Article" struct
	response := map[string]Article{
		"article": article,
	}

	// Set the Content-Type to be of JSON
	w.Header().Set("Content-Type", "application/json")

	// Set an appropriate HTTP status code to return along with the response
	w.WriteHeader(http.StatusCreated)

	// Return a JSON response by the handler
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}

func EditArticle(w http.ResponseWriter, r *http.Request) {}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {}
