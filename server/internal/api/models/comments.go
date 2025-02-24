/*
Package models provides data structures related to entities in the system.

It includes:
  - The `Comment` struct that represents a comment made by a user on an article,
    including fields for the unique ID, name, email, and content of the comment.
*/

package models

import "github.com/google/uuid"

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
