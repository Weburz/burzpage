/*
Package testutils provides utility functions for testing HTTP handlers.

These functions help in simulating HTTP requests and verifying responses in unit tests.
It includes utilities for executing HTTP requests using a handler and checking the
response code.

Functions:
  - ExecuteRequest: Executes an HTTP request using a handler and returns
    the recorded response.
  - CheckResponseCode: Compares the expected and actual HTTP response
    codes, reporting errors if they don't match.
*/
package testutils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
ExecuteRequest creates a new ResponseRecorder, executes the provided HTTP request using
the given handler, and returns the recorded response.

This utility function is useful for testing HTTP handlers by capturing their responses
to requests for further inspection.

Parameters:

	req: The HTTP request to be executed.
	h: The HTTP handler that processes the request.

Returns:

	A pointer to an httptest.ResponseRecorder which holds the response data.
*/
func ExecuteRequest(req *http.Request, h http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	return rr
}

/*
CheckResponseCode compares the expected and actual HTTP response codes. If they do not
match, it reports an error in the test output.

This utility function simplifies the process of verifying that the handler returned the
correct response code during testing.

Parameters:

	t: The testing.T instance to report errors.
	expected: The expected HTTP status code.
	actual: The actual HTTP status code returned by the handler.

Returns:

	Nothing. If the codes do not match, an error will be logged to
	the test output.
*/
func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
