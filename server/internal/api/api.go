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
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Weburz/burzcontent/server/internal/api/handlers"
	"github.com/Weburz/burzcontent/server/internal/api/routes"
)

/*
API represents the configuration of the HTTP server, including its router
and other components like database and configuration that could be added later.

Fields:
  - Router: The router (of type *chi.Mux) used for routing HTTP requests to handlers.
    It is based on the chi router, which provides a fast and lightweight way to handle
    routing for RESTful APIs and other HTTP requests.

Future Enhancements:
  - Additional fields like Db and config can be added to this struct to include a
    database connection and configuration settings required for the server to function.

Example Usage:

	server := &API{
		Router: chi.NewRouter(),
	}
*/
type API struct {
	Router *chi.Mux
}

/*
NewAPI creates a new instance of the API server with the given handlers.

This function performs the following steps:

 1. Initializes a new router using `chi.NewRouter()` for routing HTTP requests.
 2. Adds middleware to the router, such as the `Logger` middleware for logging HTTP
    requests.
 3. Sets up the server's routes by calling `routes.SetupRoutes(router, h)`, where the
    routes are defined based on the provided handlers.
 4. Returns a pointer to an `API` instance, which contains the configured router.

The returned `API` instance is ready to handle incoming HTTP requests, with the routes
and middleware set up according to the provided handlers.

Example:
  - This function can be used to create a new API instance with custom request handlers
    for various routes.
*/
func NewAPI(h *handlers.Handlers) *API {
	// Initialise a new `Router` object
	router := chi.NewRouter()

	// Register the in-built logger
	router.Use(middleware.Logger)

	// Setup the routes (aka the API endpoints) to receive HTTP requests on
	routes.SetupRoutes(router, h)

	// Return an instance of the `API` struct
	return &API{
		Router: router,
	}
}

/*
Run sets up all middleware and routes for the server.

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

func (a *API) Run() error {
	r := a.Router

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

	// Set up the HTTP server
	srv := http.Server{
		// TODO: Remove the hardcoded port number later down the line
		Addr:         ":8000",
		Handler:      a.Router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// TODO: Remove the hardcoded port number later down the line
	log.Printf("Starting the server at [::]:%s", "8000")

	// Start the server
	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Error starting server: ", err)
		return err
	}

	return nil
}
