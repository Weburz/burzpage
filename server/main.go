package main

import (
	"net/http"

	"github.com/Weburz/burzcontent/server/internal/server"
)

/*
Main is the entry point for the web server application.

This function performs the following steps:

 1. Initializes a new server instance using `server.CreateNewServer()`, which sets up
    the server and its router.
 2. Configures the server's routes and middleware by calling `s.MountHandlers()`.
 3. Starts the HTTP server on port 8000 using `http.ListenAndServe`, passing the
    server's router (`s.Router`) to handle incoming requests.

The server listens for HTTP requests on port 8000, where the routes defined in
`MountHandlers` will be available.

Example:
  - The server will be accessible at http://localhost:8000.
  - Requests to the root path ("/") will trigger the `HelloWorld` handler.

This main function is the entry point for running the server and does not exit until
the server is stopped.
*/
func main() {
	s := server.CreateNewServer()
	s.MountHandlers()
	http.ListenAndServe(":8000", s.Router)
}
