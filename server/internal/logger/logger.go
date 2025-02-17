// Package logger provides a simple logging utility using the standard library's
// log/slog package. It initializes a global logger that can be accessed globally
// throughout the application.
package logger

import (
	"log/slog"
	"os"
)

var (
	// logger holds the global instance of the logger used throughout the application.
	logger *slog.Logger
)

// init initializes the logger with a text handler that writes logs to standard output.
// The logger instance is created and assigned to the global 'logger' variable.
func init() {
	var handler slog.Handler

	// Get the runtime environment type (could be "development" or "production"). It
	// should be fetched from an environment variable
	mode := os.Getenv("ENV")

	// Choose the appropriate handler based on the mode
	if mode == "production" {
		// Attempt to open a JSON-based log, else panic and exit the execution
		file, err := os.OpenFile("logs.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}

		// Custom JSONHandler options
		opts := &slog.HandlerOptions{
			AddSource: false,          // Disable stack tracing in production logs
			Level:     slog.LevelInfo, // Set the log level to "ERROR" in production
		}
		// Production mode: use JsonHandler
		handler = slog.NewJSONHandler(file, opts)
	} else {
		// Custom TextHandler options
		opts := &slog.HandlerOptions{
			AddSource: true,            // Show the location of the stack in the log
			Level:     slog.LevelDebug, // Set the log level to "DEBUG"
		}
		// Development mode: use TextHandler
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	// Initialize the logger with the chosen handler
	logger = slog.New(handler)
}

// NewLogger returns the global logger instance, which can be used to log messages
// throughout the application. This function provides easy access to the logger
// without needing to create a new instance.
func NewLogger() *slog.Logger {
	return logger
}
