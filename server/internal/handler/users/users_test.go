package users

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Weburz/burzcontent/server/internal/testutils"
)

// Define the structure of the expected JSON response
type UserResponse struct {
	Users []User `json:"users"`
}

func TestGetUserHandler(t *testing.T) {
	// Create a new HTTP request for the /users route
	req, _ := http.NewRequest("GET", "/users", nil)

	// Create the handler and execute the request directly
	handler := http.HandlerFunc(GetUsers)
	response := testutils.ExecuteRequest(req, handler)

	// Initialize the test table with appropriate configurations to test with
	tests := []struct {
		name         string
		method       string
		url          string
		statusCode   int
		expectedBody UserResponse // Change expectedBody to use unmarshaled JSON struct
	}{
		{
			name:       "ValidRequest",
			method:     "GET",
			url:        "/users",
			statusCode: http.StatusOK,
			expectedBody: UserResponse{
				Users: []User{
					{
						ID:    "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
						Name:  "Somraj Saha",
						Email: "somraj.saha@weburz.com",
					},
					{
						ID:    "4f3e23f4-d5d9-4886-90de-f07a93d3c7f5",
						Name:  "John Doe",
						Email: "john.doe@example.com",
					},
				},
			},
		},
	}

	// Run the group of tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check the response code
			testutils.CheckResponseCode(t, tt.statusCode, response.Code)

			// Unmarshal the response body into a UserResponse struct
			var actualBody UserResponse
			if err := json.NewDecoder(response.Body).Decode(&actualBody); err != nil {
				t.Errorf("Failed to unmarshal response body: %v", err)
			}

			// Compare the actual and expected response bodies
			if !compareUserResponses(tt.expectedBody, actualBody) {
				t.Errorf("Expected body %v, but got %v", tt.expectedBody, actualBody)
			}
		})
	}
}

// Helper function to compare UserResponse structs
func compareUserResponses(expected, actual UserResponse) bool {
	if len(expected.Users) != len(actual.Users) {
		return false
	}
	for i := range expected.Users {
		if expected.Users[i].ID != actual.Users[i].ID ||
			expected.Users[i].Name != actual.Users[i].Name ||
			expected.Users[i].Email != actual.Users[i].Email {
			return false
		}
	}
	return true
}
