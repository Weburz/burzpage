/*
Package models provides data structures related to the entities in the system.

It includes:
  - The `Article` struct that represents an article with fields for its unique ID,
    title, author, and publication status.
*/

package models

import "github.com/google/uuid"

/*
Article represents an article with its associated data.

Fields:
  - ID: The unique identifier for the article (UUID).
  - Title: The title of the article.
  - Author: The author of the article.
  - Published: A boolean indicating if the article is published.
*/
type Article struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Published bool      `json:"published"`
}
