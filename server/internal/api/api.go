/*
Package server provides the HTTP server setup, including route configuration,
middleware, and handler initialization.

This package is responsible for:
- Creating and configuring the server (`Server` struct).
- Defining and mounting routes and middleware on the server.
- Providing utility functions to create a new server and set up its routes.

It uses the `github.com/go-chi/chi` package for routing and middleware management,
allowing for flexible and efficient HTTP request handling.

Functions:
  - CreateNewServer: Initializes a new `Server` instance with a configured router.
  - MountHandlers: Sets up middleware and request handlers for the server, including
    routes and their corresponding logic.
*/
package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Weburz/burzcontent/server/internal/handler/articles"
	"github.com/Weburz/burzcontent/server/internal/handler/comments"
	"github.com/Weburz/burzcontent/server/internal/handler/users"
)

/*
Server represents the configuration of the HTTP server, including its router
and other components like database and configuration that could be added later.

Fields:
  - Router: The router (of type *chi.Mux) used for routing HTTP requests to handlers.
    It is based on the chi router, which provides a fast and lightweight way to handle
    routing for RESTful APIs and other HTTP requests.

Future Enhancements:
  - Additional fields like Db and config can be added to this struct to include a
    database connection and configuration settings required for the server to function.

Example Usage:

	server := &Server{
		Router: chi.NewRouter(),
	}
*/
type Server struct {
	Router *chi.Mux
	// Db, config can be added here
}

/*
CreateNewServer initializes a new instance of the Server struct, setting up its router.

The function does the following:
  - Creates a new `Server` instance.
  - Initializes the `Router` field with a new chi.Mux router, which is used for routing
    HTTP requests.

This function provides a convenient way to create a fully configured server instance,
ready to have routes defined on its router.

Returns:
- A pointer to a `Server` struct with an initialized `Router` (of type *chi.Mux).

Example Usage:

	server := CreateNewServer()
	// You can now define routes on the server's router using `server.Router`
*/
func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

/*
MountHandlers sets up all middleware and routes for the server.

This function is responsible for:
  - Mounting the middleware to be used in the server, such as logging middleware.
  - Mounting the HTTP request handlers for various routes.

Currently, it mounts the following:
  - `middleware.Logger`: A middleware that logs all incoming HTTP requests.
  - `"/"` route: A handler function (`HelloWorld`) that responds to requests made to the
    root path.

This function can be expanded in the future to add more middleware or routes to the
server.
*/
func (s *Server) MountHandlers() {
	r := s.Router

	// Mount all Middleware here
	r.Use(middleware.Logger)

	// Mount all handlers related to users
	r.Route("/users", func(r chi.Router) {
		r.Get("/", users.GetUsers)
		r.Get("/{id}", users.GetUser)
		r.Post("/{id}/edit", users.UpdateUser)
		r.Delete("/{id}/delete", users.DeleteUser)
	})

	// Mount all handlers related to the articles
	r.Route("/articles", func(r chi.Router) {
		r.Get("/", articles.GetArticles)
		r.Get("/{id}", articles.GetArticle)
		r.Post("/new", articles.CreateArticle)
		r.Put("/{id}/edit", articles.EditArticle)
		r.Delete("/{id}/delete", articles.DeleteArticle)
	})

	// Mount all handlers related to the comments
	r.Route("/comments", func(r chi.Router) {
		r.Get("/", comments.GetComments)
		r.Get("/article/{id}", comments.GetCommentFromArticle)
		r.Post("/article/{id}/new", comments.AddComment)
		r.Delete("/{id}/delete", comments.RemoveComment)
	})
}
