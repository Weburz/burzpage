/*
Package services provides the implementation of the CommentService interface for
managing comments in the system.

This package contains functionality related to comment management, including retrieving,
adding, and deleting comments. It defines the CommentService interface and provides the
concrete implementation through the CommentServiceImpl struct.

Key Components:

  - CommentService: An interface defining methods to manage comments.
  - CommentServiceImpl: A struct that implements the CommentService interface.
  - NewCommentService: A constructor function to create a new CommentServiceImpl
    instance.
  - GetAllComments: Retrieves all comments for an article.
  - GetCommentsFromComment: Retrieves comments associated with a specific comment.
  - AddCommentToArticle: Adds a new comment to an article.
  - DeleteCommentFromArticle: Removes a comment from an article (currently
    unimplemented).

The functionality is primarily focused on handling comment-related operations, which
can be extended or modified based on the requirements of the application.
*/
package services

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/Weburz/burzcontent/server/internal/api/models"
)

/*
CommentService defines the methods for managing comments in the system.

This interface outlines the core operations related to comments, including retrieving,
adding, and deleting comments. Any type that implements this interface is expected to
provide the logic for each of the following methods:

Methods:

	GetAllComments(): Retrieves all the comments.
	GetCommentsFromArticle(): Retrieves comments associated with a specific article.
	AddComment(name, email, content string): Adds a new comment.
	DeleteComment(): Deletes a comment.
*/
type CommentService interface {
	GetAllComments() ([]models.Comment, error)
	GetCommentsFromArticle() ([]models.Comment, error)
	AddCommentToArticle(name, email, content string) (*models.Comment, error)
	DeleteCommentFromArticle() error
}

/*
CommentServiceImpl is a struct that implements the CommentService interface.

This struct is used to manage operations related to comments, such as adding, deleting
and retrieving comments. The struct itself does not contain any fields but serves as the
concrete implementation for the methods defined in the CommentService interface.
*/
type CommentServiceImpl struct{}

/*
NewCommentService creates and returns a new instance of CommentServiceImpl.

This function initializes a new CommentServiceImpl object and returns it as a pointer.
It serves as a constructor for the CommentServiceImpl type.

Returns:

	*CommentServiceImpl: A pointer to a newly created CommentServiceImpl instance.
*/
func NewCommentService() *CommentServiceImpl {
	return &CommentServiceImpl{}
}

/*
GetAllComments retrieves all the comments for an article.

This function simulates the retrieval of all comments by generating a new comment ID
using uuid.NewV7(). It creates a list of pre-defined comments with unique names, emails,
and content. If there is an error while generating the comment ID, it returns an empty
slice of comments and the error.

Returns:

	[]models.Comment: A slice of pre-defined comments.
	error: An error if there was an issue generating the comment ID.
*/
func (cs *CommentServiceImpl) GetAllComments() ([]models.Comment, error) {
	commentID, err := uuid.NewV7()
	if err != nil {
		return []models.Comment{}, fmt.Errorf("%w", err)
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

	return comments, nil
}

/*
GetCommentsFromComment retrieves a list of comments for a given article.

This function simulates the retrieval of comments by generating a new comment ID using
uuid.NewV7(). It creates a list of pre-defined comments with unique names, emails, and
content. If there is an error while generating the comment ID, it returns an empty slice
of comments and the error.

Returns:

	[]models.Comment: A slice of pre-defined comments.
	error: An error if there was an issue generating the comment ID.
*/
func (cs *CommentServiceImpl) GetCommentsFromArticle() ([]models.Comment, error) {
	commentID, err := uuid.NewV7()
	if err != nil {
		return []models.Comment{}, fmt.Errorf("%w", err)
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

	return comments, nil
}

/*
AddCommentToArticle adds a new comment to an article.

This function generates a new unique comment ID using uuid.NewV7() and then  creates a
new comment object with the provided name, email, and content. If there is an error
while generating the comment ID, it returns an empty comment object and the error.

Parameters:

	name (string): The name of the commenter.
	email (string): The email of the commenter.
	content (string): The content of the comment.

Returns:

	*models.Comment: The newly created comment with the generated ID.
	error: An error if there was an issue generating the comment ID.
*/
func (cs *CommentServiceImpl) AddCommentToArticle(
	name, email, content string,
) (*models.Comment, error) {
	commentID, err := uuid.NewV7()
	if err != nil {
		return &models.Comment{}, fmt.Errorf("%w", err)
	}

	comment := &models.Comment{
		ID:      commentID,
		Name:    name,
		Email:   email,
		Content: content,
	}

	return comment, nil
}

/*
DeleteCommentFromArticle removes a comment from an article.

This function currently does not implement any logic for deleting a comment. It is
intended to be used in the context of a comment service, where the deletion of a
comment from an article would be handled in the future.

Returns:

	nil: Always returns nil as there is no implementation yet.
*/
func (cs *CommentServiceImpl) DeleteCommentFromArticle() error {
	return nil
}
