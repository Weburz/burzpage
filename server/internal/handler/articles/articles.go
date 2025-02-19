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
)

/*
Article represents an article with a title, author, and publication status.

Fields:
- Title: The title of the article.
- Author: The author of the article.
- Published: A boolean indicating if the article is published.
*/
type Article struct {
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
			Title:     "Go Programming Basics",
			Author:    "John Doe",
			Published: true,
		},
		{
			Title:     "Advanced Go Techniques",
			Author:    "Jane Smith",
			Published: false,
		},
		{
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

func GetArticle(w http.ResponseWriter, r *http.Request) {}

func CreateArticle(w http.ResponseWriter, r *http.Request) {}

func EditArticle(w http.ResponseWriter, r *http.Request) {}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {}
