/*
Package config defines the configuration settings for the server.

This package is responsible for managing and providing configuration settings
such as the server port and environment type. It also handles the initialization
of the request handlers used by the server. The `Config` struct holds the configuration
data, while functions like `NewConfig` and `InitialiseHandlers` help initialize
and set up the necessary components for the server.
*/
package config

import "github.com/Weburz/burzcontent/server/internal/api/handlers"

// Config holds the server configuration settings, such as the port and environment
// type.
type Config struct {
	Port string // The port on which the server will listen
	Env  string // The environment type (e.g., "development", "production")
}

/*
NewConfig creates and initializes a new instance of Config with default values.

This function returns a new `Config` instance with default values:
  - Port: "8000"
  - Env: "development"

These default values can be overridden by setting the respective fields after
creating the `Config` instance.

Example:
  - This function is used to create a configuration object before initializing
    the handlers or setting up the server.
*/
func NewConfig() *Config {
	return &Config{
		Port: "8000",        // Default port
		Env:  "development", // Default environment
	}
}

/*
InitialiseHandlers initializes and returns a new instance of Handlers.

This function calls the `handlers.NewHandlers()` function to create a new
`Handlers` instance, which contains the necessary request handlers for the server.

Example:
  - This function can be used to set up the handlers needed by the server,
    including those for user-related HTTP requests.
*/
func (c *Config) InitialiseHandlers() *handlers.Handlers {
	return handlers.NewHandlers()
}
