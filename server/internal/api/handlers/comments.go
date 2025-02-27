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
	"net/http"

	validator "github.com/go-playground/validator/v10"

	"github.com/Weburz/burzcontent/server/internal/api/models"
	"github.com/Weburz/burzcontent/server/internal/api/services"
)

/*
CommentHandler is a struct that handles HTTP requests related to comments.

This struct holds a reference to the CommentService, which is used to perform operations
such as adding, retrieving, and deleting comments. It acts as a controller in an MVC
architecture, handling requests from the client and delegating comment-related logic to
the CommentService.

Fields:

	CommentService (services.CommentService): A service for managing comments.
*/
type CommentHandler struct {
	CommentService services.CommentService
}

/*
NewCommentHandler creates and returns a new instance of CommentHandler.

This function initializes a new CommentHandler object, providing it with the given
CommentService to handle comment-related operations. It serves as a constructor for the
CommentHandler type.

Parameters:

	commentService (services.CommentService): The service to be used for comment
	    operations.

Returns:

	*CommentHandler: A pointer to a newly created CommentHandler instance.
*/
func NewCommentHandler(commentService services.CommentService) *CommentHandler {
	return &CommentHandler{
		CommentService: commentService,
	}
}

/*
GetAllComments handles HTTP requests to retrieve all comments for an article.

This method interacts with the CommentService to fetch all comments. If successful,
it returns the comments in a JSON format with a "comments" key. If any error occurs
while retrieving the comments or encoding the response, it returns an appropriate
error message with an HTTP status code of 500 (Internal Server Error).

Parameters:

	w (http.ResponseWriter): The HTTP response writer used to send the response.
	r (*http.Request): The HTTP request containing information about the request.

Returns:

	None: Writes the response directly to the HTTP client.

HTTP Status Codes:
  - 200 (OK): If the comments are successfully retrieved and returned.
  - 500 (Internal Server Error): If there is an error while retrieving comments
    or encoding the response.
*/
func (cr *CommentHandler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := cr.CommentService.GetAllComments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string][]models.Comment{"comments": comments}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
GetCommentsFromArticle handles HTTP requests to retrieve comments from a specific
article.

This method interacts with the CommentService to fetch the comments associated with
an article. If successful, it returns the comments in a JSON format with a "comments"
key. If any error occurs while retrieving the comments or encoding the response,
it returns an appropriate error message with an HTTP status code of 500 (Internal
Server Error).

Parameters:

	w (http.ResponseWriter): The HTTP response writer used to send the response.
	r (*http.Request): The HTTP request containing information about the request.

Returns:

	None: Writes the response directly to the HTTP client.

HTTP Status Codes:
  - 200 (OK): If the comments are successfully retrieved and returned.
  - 500 (Internal Server Error): If there is an error while retrieving comments
    or encoding the response.
*/
func (cr *CommentHandler) GetCommentsFromArticle(
	w http.ResponseWriter,
	r *http.Request,
) {
	comments, err := cr.CommentService.GetCommentsFromArticle()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string][]models.Comment{"comments": comments}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
AddCommentToArticle handles HTTP requests to add a new comment to an article.

This method receives a new comment in JSON format, validates it, and then uses
the CommentService to add the comment. If the comment is successfully added,
it returns the newly created comment in a JSON format with a "comment" key.
If any error occurs during the process, it returns an appropriate error message
with the corresponding HTTP status code.

Parameters:

	w (http.ResponseWriter): The HTTP response writer used to send the response.
	r (*http.Request): The HTTP request containing the comment data.

Returns:

	None: Writes the response directly to the HTTP client.

HTTP Status Codes:
  - 201 (Created): If the comment is successfully added.
  - 400 (Bad Request): If there is an error decoding the request body.
  - 422 (Unprocessable Entity): If the comment fails validation.
  - 500 (Internal Server Error): If there is an error while adding the comment
    or encoding the response.
*/
func (cr *CommentHandler) AddCommentToArticle(w http.ResponseWriter, r *http.Request) {
	var newComment models.Comment
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newComment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the request body
	validate := validator.New()
	if err := validate.Struct(newComment); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// Create a comment instance
	comment, err := cr.CommentService.AddCommentToArticle(
		newComment.Name,
		newComment.Email,
		newComment.Content,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

/*
DeleteCommentFromArticle handles HTTP requests to delete a comment from an article.

This method interacts with the CommentService to delete a comment. If the deletion
is successful, it returns an HTTP status code of 204 (No Content). If there is any
error while deleting the comment, it returns an appropriate error message with a
status code of 500 (Internal Server Error).

Parameters:

	w (http.ResponseWriter): The HTTP response writer used to send the response.
	r (*http.Request): The HTTP request containing information about the request.

Returns:

	None: Writes the response directly to the HTTP client.

HTTP Status Codes:
  - 204 (No Content): If the comment is successfully deleted.
  - 500 (Internal Server Error): If there is an error while deleting the comment.
*/
func (cr *CommentHandler) DeleteCommentFromArticle(
	w http.ResponseWriter,
	r *http.Request,
) {
	if err := cr.CommentService.DeleteCommentFromArticle(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusNoContent)
}
