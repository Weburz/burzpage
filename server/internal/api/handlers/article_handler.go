/*
Package handler provides the implementation for managing article-related
operations, such as creating, retrieving, and updating articles. This package
handles HTTP requests and responses for articles, including validation and
JSON encoding/decoding.

It includes the following key functionalities:
  - GetArticles: Retrieves a list of all articles.
  - GetArticle: Retrieves a specific article by its ID.
  - CreateArticle: Creates a new article with a given title and author.

Each handler ensures that proper HTTP status codes are returned along with
appropriate JSON responses. The package also handles error scenarios, such as
invalid requests, failed UUID generation, and JSON encoding failures.
*/
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Weburz/burzcontent/server/internal/api/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

/*
ArticleHandler is responsible for handling HTTP requests related to articles.

This struct provides methods for creating, retrieving, updating, and deleting
articles. It serves as a handler for HTTP routes that manage article-related
operations and returns appropriate responses to the client.

Methods:
  - GetArticles: Retrieves a list of all articles.
  - GetArticle: Retrieves a single article by its ID.
  - CreateArticle: Creates a new article.
  - EditArticle: Updates an existing article.
  - DeleteArticle: Deletes an article.

The `ArticleHandler` struct does not store any state itself but relies on
external services, such as models and validators, to handle article data
and validation.
*/
type ArticleHandler struct{}

/*
NewArticleHandler creates and returns a new instance of ArticleHandler.

This function initializes and returns a new `ArticleHandler` struct, which can be
used to manage article-related operations such as creating, reading, updating,
and deleting articles.

Returns:
  - *ArticleHandler: A new instance of `ArticleHandler`.

Example:
  - Call `NewArticleHandler()` to create a new `ArticleHandler` instance.
*/
func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{}
}

/*
GetArticles handles the retrieval of all articles.

This function performs the following actions:

 1. Generates a new article ID using `uuid.NewV7()`. In this implementation, the
    same ID is used for all the sample articles (this could be a mistake as each
    article should likely have a unique ID).
 2. Creates a list of sample articles with predefined titles, authors, and
    publication statuses.
 3. Encodes the list of articles into a JSON response and sends it back to the
    client with a status of `200 OK`.

The response JSON object contains an array of articles, each with the following
structure:
  - `ID`: The unique identifier of the article.
  - `Title`: The title of the article.
  - `Author`: The author of the article.
  - `Published`: A boolean indicating whether the article is published or not.

Example Response:

	{
	  "articles": [
	    {
	      "id": "some-uuid",
	      "title": "Go Programming Basics",
	      "author": "John Doe",
	      "published": true
	    },
	    {
	      "id": "some-uuid",
	      "title": "Advanced Go Techniques",
	      "author": "Jane Smith",
	      "published": false
	    },
	    {
	      "id": "some-uuid",
	      "title": "Understanding Go Concurrency",
	      "author": "Alice Johnson",
	      "published": true
	    }
	  ]
	}

If the UUID generation or JSON encoding fails, the function returns a
corresponding error message with an appropriate HTTP status.

Possible Errors:
  - If the UUID generation fails, a `500 Internal Server Error` is returned with
    the message "Unable to generate the Article ID".
  - If JSON encoding fails, a `500 Internal Server Error` is returned with the
    message "Unable to encode JSON".

Example:
  - Request: GET /articles
  - Response: HTTP 200 OK with a JSON body containing a list of articles.
*/
func (ar *ArticleHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
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

	response := map[string][]models.Article{
		"articles": articles,
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
GetArticle handles the retrieval of a single article by its ID.

This function performs the following actions:

 1. Retrieves the article ID from the URL path parameter using
    `chi.URLParam(r, "id")`.
 2. Attempts to parse the ID into a UUID using `uuid.Parse()`. If the parsing
    fails, it returns a `404 Not Found` error with the message "Article ID Not
    Found".
 3. Creates a sample article with predefined title, author, and publication
    status.
 4. Encodes the article into a JSON response and sends it back to the client with
    a status of `200 OK`.

The response JSON object contains the article with the following structure:
  - `ID`: The unique identifier of the article.
  - `Title`: The title of the article.
  - `Author`: The author of the article.
  - `Published`: A boolean indicating whether the article is published or not.

Example Response:

	{
	  "article": {
	    "id": "some-uuid",
	    "title": "Go Programming Basics",
	    "author": "John Doe",
	    "published": true
	  }
	}

Possible Errors:
  - If the article ID is not found or cannot be parsed, a `404 Not Found` error is
    returned with the message "Article ID Not Found".
  - If JSON encoding fails, a `500 Internal Server Error` is returned with the
    message "Unable to encode JSON".

Example:
  - Request: GET /articles/{id}
  - Response: HTTP 200 OK with a JSON body containing the requested article.
*/
func (ar *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	articleID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Article ID Not Found", http.StatusNotFound)
		return
	}

	article := models.Article{
		ID:        articleID,
		Title:     "Go Programmign Basics",
		Author:    "John Doe",
		Published: true,
	}

	response := map[string]models.Article{
		"article": article,
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
CreateArticle handles the creation of a new article.

This function performs the following actions:

 1. Decodes the incoming request body into a new `Article` object.
 2. Validates the new article using the `validator` package. If validation fails,
    it returns a `422 Unprocessable Entity` error with the message "Request validation
    failed".
 3. Generates a new UUID for the article ID using `uuid.NewV7()`. If the UUID
    generation fails, it returns a `500 Internal Server Error` with the message
    "Failed to generate the Article ID".
 4. Creates a new article with the given title and author, and sets the article's
    publication status to `false`.
 5. Encodes the newly created article into a JSON response and returns it to the
    client with a status of `201 Created`.

The response JSON object contains the created article with the following structure:
- `ID`: The unique identifier of the article.
- `Title`: The title of the article.
- `Author`: The author of the article.
- `Published`: A boolean indicating whether the article is published or not.

Example Response:

	{
	  "article": {
	    "id": "some-uuid",
	    "title": "Go Programming for Beginners",
	    "author": "John Doe",
	    "published": false
	  }
	}

Possible Errors:
  - If the request body is invalid or cannot be decoded, a `400 Bad Request` error
    is returned with the message "Invalid request body".
  - If the request validation fails, a `422 Unprocessable Entity` error is returned
    with the message "Request validation failed".
  - If UUID generation fails, a `500 Internal Server Error` is returned with the
    message "Failed to generate the Article ID".
  - If JSON encoding fails, a `500 Internal Server Error` is returned with the
    message "Unable to encode JSON".

Example:
  - Request: POST /articles
  - Request Body: JSON object with title and author fields.
  - Response: HTTP 201 Created with a JSON body containing the created article.
*/
func (ar *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()

	var newArticle models.Article
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newArticle); err != nil {
		http.Error(
			w,
			"Invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	if err := validate.Struct(newArticle); err != nil {
		http.Error(w, "Request validation failed", http.StatusUnprocessableEntity)
		return
	}

	articleID, err := uuid.NewV7()
	if err != nil {
		http.Error(
			w,
			"Failed to generate the Article ID",
			http.StatusInternalServerError,
		)
		return
	}

	article := &models.Article{
		ID:        articleID,
		Title:     newArticle.Title,
		Author:    newArticle.Author,
		Published: false,
	}

	response := map[string]models.Article{
		"article": *article,
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
EditArticle handles the updating of an existing article.

This function performs the following actions:

 1. Initializes a new `Article` object and validates the incoming request body
    using the `validator` package. If validation fails, it returns a `422
    Unprocessable Entity` error with the message "Request body validation failed".
 2. Decodes the incoming request body into the `updatedArticle` object. If
    decoding fails, it returns a `400 Bad Request` error with the message "Invalid
    Request Body".
 3. Retrieves and parses the article ID from the URL path parameter using
    `chi.URLParam(r, "id")`. If the ID is invalid or missing, it returns a `404
    Not Found` error with the message "Article ID Not Found".
 4. Updates the article with the parsed ID, and the updated title, author, and
    publication status.
 5. Encodes the updated article into a JSON response and sends it back to the
    client with a status of `201 Created`.

The response JSON object contains the updated article with the following structure:
  - `ID`: The unique identifier of the article.
  - `Title`: The updated title of the article.
  - `Author`: The updated author of the article.
  - `Published`: The updated publication status of the article.

Possible Errors:
  - If the request body validation fails, a `422 Unprocessable Entity` error is
    returned with the message "Request body validation failed".
  - If the request body is invalid or cannot be decoded, a `400 Bad Request` error
    is returned with the message "Invalid Request Body".
  - If the article ID is not found or cannot be parsed, a `404 Not Found` error is
    returned with the message "Article ID Not Found".
  - If JSON encoding fails, a `500 Internal Server Error` is returned with the
    message "Unable to encode JSON".

Example:
  - Request: PUT /articles/{id}
  - Request Body: JSON object with updated title, author, and publication status.
  - Response: HTTP 201 Created with a JSON body containing the updated article.
*/
func (ar *ArticleHandler) EditArticle(w http.ResponseWriter, r *http.Request) {
	var updatedArticle models.Article
	validate := validator.New()
	if err := validate.Struct(updatedArticle); err != nil {
		http.Error(w, "Request body validation failed", http.StatusUnprocessableEntity)
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedArticle); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	articleID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Article ID Not Found", http.StatusNotFound)
		return
	}

	article := models.Article{
		ID:        articleID,
		Title:     updatedArticle.Title,
		Author:    updatedArticle.Author,
		Published: updatedArticle.Published,
	}

	response := map[string]models.Article{
		"article": article,
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
DeleteArticle handles the deletion of an article.

This function performs the following actions:

 1. Retrieves and parses the article ID from the URL path parameter using
    `chi.URLParam(r, "id")`. If the ID is invalid or missing, it returns a
    `404 Not Found` error with the message "Article ID Not Found".
 2. If the article ID is valid, it returns an empty response with a `204 No
    Content` status indicating the article was successfully deleted.

Possible Errors:
  - If the article ID is not found or cannot be parsed, a `404 Not Found` error is
    returned with the message "Article ID Not Found".

Example:
  - Request: DELETE /articles/{id}
  - Response: HTTP 204 No Content, indicating successful deletion.
*/
func (ar *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	_, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Article ID Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusNoContent)
}
