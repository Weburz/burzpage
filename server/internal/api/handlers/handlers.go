/*
Package handlers defines the handler processes for handling incoming HTTP requests.

This package provides a structure for managing different request handlers.
It includes a main handler, `Handlers`, that contains specific handlers for
various resources, such as `UserHandler` for user-related requests.

The package is designed to be used within the broader application to handle
and process requests routed to different endpoints.
*/
package handlers

// Handlers holds the handler instances for the various resources in the application.
type Handlers struct {
	UserHandler    *UserHandler
	ArticleHandler *ArticleHandler
}

/*
NewHandlers creates and initializes a new Handlers instance.

This function performs the following steps:

 1. Creates a new `UserHandler` instance by calling `NewUserHandler()`.
 2. Returns a new `Handlers` instance that contains the `UserHandler`.

This function provides an easy way to initialize all the handlers needed
for the application, including user-related handlers.
*/
func NewHandlers() *Handlers {
	return &Handlers{
		UserHandler:    NewUserHandler(),
		ArticleHandler: NewArticleHandler(),
	}
}
