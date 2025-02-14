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
	handler := slog.NewTextHandler(
		os.Stdout,
		nil,
	) // Create a text handler for standard output
	logger = slog.New(
		handler,
	) // Initialize the logger with the text handler
}

// NewLogger returns the global logger instance, which can be used to log messages
// throughout the application. This function provides easy access to the logger
// without needing to create a new instance.
func NewLogger() *slog.Logger {
	return logger
}
