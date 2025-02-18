package users

import (
	"net/http"
	"testing"

	"github.com/Weburz/burzcontent/server/internal/testutils"
)

func TestGetUserHandler(t *testing.T) {
	// Create a new HTTP request for the /users route
	req, _ := http.NewRequest("GET", "/users", nil)

	// Create the handler and execute the request directly
	handler := http.HandlerFunc(GetUsersHandler)
	response := testutils.ExecuteRequest(req, handler)

	// Check the response code
	testutils.CheckResponseCode(t, http.StatusOK, response.Code)
}
