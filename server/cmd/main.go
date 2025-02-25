/*
Package main is the entry point for the web server application.

This package is responsible for setting up and starting the server. It initializes the
necessary configuration, sets up request handlers, and runs the server to listen for
incoming HTTP requests.

The main function in this package follows these high-level steps:
1. Configures the server by creating a new configuration and initializing the handlers.
2. Creates a new API server and configures the necessary routes and middleware.
3. Starts the server to listen for requests on port 8000.

The application does not exit until the server is stopped.
*/
package main

import (
	"github.com/Weburz/burzcontent/server/internal/api"
	"github.com/Weburz/burzcontent/server/internal/config"
)

/*
Main is the entry point for the web server application.

This function performs the following steps:

 1. Initializes a new configuration instance using `config.NewConfig()` to retrieve
    necessary configurations for the server.
 2. Initializes the request handlers by calling `cfg.InitialiseHandlers()` to set up
    handler functions based on the configuration.
 3. Creates a new API instance using `api.NewAPI(handlers)` and initializes it with
    the handlers.
 4. Starts the server with the `server.Run()` function, which listens for HTTP requests
    and processes them based on the defined handlers.

The server will run continuously, handling incoming requests until it is manually
stopped.

Example:
  - The server will be accessible at http://localhost:8000.
  - Requests will be routed according to the handler configuration.

This function is the main entry point for running the server and will not exit until
the server is stopped.
*/
func main() {
	cfg := config.NewConfig()
	handlers := cfg.InitialiseHandlers()
	server := api.NewAPI(handlers)
	server.Run()
}
