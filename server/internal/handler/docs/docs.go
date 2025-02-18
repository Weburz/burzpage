// Package docs provides HTTP handler functions to serve documentation pages as HTML
// templates. It handles requests to render static HTML templates and logs relevant
// details for each request.
package docs

import (
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Weburz/burzcontent/server/internal/logger"
)

// DocsHandler handles the HTTP request to render the "docs.html" template.
// It parses the template and writes the HTML content to the response writer.
// If any error occurs during parsing or execution of the template, it logs the error
// and responds with an internal server error. On successful template execution,
// it logs the details of the successful response.
//
// Logs:
//   - Error logs when parsing or executing the template.
//   - Info log when the template is successfully rendered, including method, path, and
//     status.
func DocsHandler(w http.ResponseWriter, r *http.Request) {
	// Initialise the log handler to log the stack trace
	logger := logger.NewLogger()

	// Get the path to the template files
	tmplPath := filepath.Join("templates", "docs.html")

	// Define the HTML content as a template
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		// Log error with more context
		logger.Error(
			"Error parsing template",
			slog.String("template", tmplPath),
			slog.String("error", err.Error()),
		)
		http.Error(
			w,
			fmt.Sprintf("Error parsing template: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	// Render the template to the response
	err = tmpl.Execute(w, nil)
	if err != nil {
		// Log error with more context
		logger.Error(
			"Error executing template",
			slog.String("template", tmplPath),
			slog.String("error", err.Error()),
		)
		http.Error(
			w,
			fmt.Sprintf("Error executing template: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	// Log a success message if the handler could respond properly
	logger.Info(
		"Successfully served template",
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.Int("status", http.StatusOK),
	)
}
