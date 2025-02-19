/*
Package comments provides HTTP handlers for managing comments on articles.

  - GetComments: Handles the request to retrieve a list of all comments.
  - GetCommentFromArticle: Handles the request to retrieve comments from a specific
    article by its ID.
  - AddComment: Handles the request to add a new comment to an article.
  - RemoveComment: Handles the request to delete a specific comment by its ID.
*/
package comments

import "net/http"

func GetComments(w http.ResponseWriter, r *http.Request) {}

func GetCommentFromArticle(w http.ResponseWriter, r *http.Request) {}

func AddComment(w http.ResponseWriter, r *http.Request) {}

func RemoveComment(w http.ResponseWriter, r *http.Request) {}
