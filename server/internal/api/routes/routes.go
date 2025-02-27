/*
Package routes defines the application's routing configuration.

This package is responsible for setting up the routes and their associated handlers
for various HTTP endpoints. The routes are mapped to functions defined in the
handlers package, which process incoming requests and generate appropriate responses.

The main function in this package, `SetupRoutes`, configures the application's
routes and binds them to specific handlers for resource management, such as
user-related routes.
*/
package routes

import (
	chi "github.com/go-chi/chi/v5"

	"github.com/Weburz/burzcontent/server/internal/api/handlers"
)

/*
SetupRoutes sets up the application's HTTP routes and maps them to their corresponding
handlers.

This function performs the following steps:

 1. Configures the `/users` route using `r.Route("/users", ...)` for handling
    user-related requests.
 2. Binds the `GET` method for the `/users` route to the `GetUsers` method of
    the `UserHandler` defined in the `handlers.Handlers` instance.

The routes are now ready to process incoming requests related to users.
*/
func SetupRoutes(r *chi.Mux, h *handlers.Handlers) {
	// Mount all handlers related to the users
	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.UserHandler.GetAllUsers)
		r.Put("/new", h.UserHandler.CreateUser)
		r.Get("/{id}", h.UserHandler.GetUserByID)
		r.Post("/{id}/edit", h.UserHandler.UpdateUser)
		r.Delete("/{id}/delete", h.UserHandler.DeleteUser)
	})

	// Mount all handlers related to the articles
	r.Route("/articles", func(r chi.Router) {
		r.Get("/", h.ArticleHandler.GetAllArticles)
		r.Put("/new", h.ArticleHandler.CreateArticle)
		r.Get("/{id}", h.ArticleHandler.GetArticleByID)
		r.Post("/{id}/edit", h.ArticleHandler.UpdateArticle)
		r.Delete("/{id}/delete", h.ArticleHandler.DeleteArticle)
	})

	// Mount all handlers related to the comments
	r.Route("/comments", func(r chi.Router) {
		r.Get("/", h.CommentHandler.GetAllComments)
		r.Get("/article/{id}", h.CommentHandler.GetCommentsFromArticle)
		r.Post("/article/{id}/new", h.CommentHandler.AddCommentToArticle)
		r.Delete("/{id}/delete", h.CommentHandler.DeleteCommentFromArticle)
	})
}
