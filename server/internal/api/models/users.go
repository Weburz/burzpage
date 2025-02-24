/*
Package models provides the data structures and definitions related to the user entity.

It includes:
  - The `User` struct that represents a user in the system with fields for the unique
    ID, name, and email.
  - Validation tags for ensuring that the `Name` field is at least 5 characters long and
    that the `Email` field is a valid email address.
*/

package models

import "github.com/google/uuid"

/*
User represents the structure of a User entity.

Fields:
  - ID: A unique identifier for the user (UUID).
  - Name: The user's name, which must be at least 5 characters long.
  - Email: The user's email address, which must be in a valid email format.
*/
type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"  validate:"required,min=5"`
	Email string    `json:"email" validate:"required,email"`
}
