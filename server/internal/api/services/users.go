/*
Package services provides user management operations for interacting with the user data
in the system. This package includes functionality to retrieve, create, update, and
delete user information.

The primary interface, `IUserService`, defines the methods for managing users, and the
`UserService` struct provides the concrete implementation of these methods. These
operations are intended to be used for managing user-related data such as user profiles
and account details.

The package contains the following key functionalities:

- GetAllUsers: Retrieves a list of all users in the system.
- GetUserByID: Fetches a user based on their unique ID.
- CreateUser: Creates a new user with a given name and email.
- UpdateUser: Updates the details of an existing user.
- DeleteUser: Removes a user from the system by their ID.

This package is meant to handle typical CRUD operations related to users in the system,
with the methods returning appropriate data or errors as needed.

The package also defines a constructor function, `NewUserService`, to initialize and
return an instance of `UserService`, which implements the `IUserService` interface.
*/
package services

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/Weburz/burzcontent/server/internal/api/models"
)

// UserService defines the methods for user management.
type UserService interface {
	// GetAllUsers retrieves all users and returns a slice of User models and an error.
	GetAllUsers() ([]models.User, error)

	// GetUserByID fetches a user by ID and returns the User model and an error (if
	// any).
	GetUserByID(id uuid.UUID) (models.User, error)

	// CreateUser creates a new user with the given name and email and returns the
	// created User model and an error (if any).
	CreateUser(name, email string) (models.User, error)

	// UpdatedUser updates an existing user's details identified by their unique ID and
	// returns the updated User model and an error (if any).
	UpdateUser(id uuid.UUID, name, email string) (models.User, error)

	// DeleteUser removes a user identified by their unique ID from the system.
	DeleteUser(id uuid.UUID) error
}

// The `UserServiceImpl` struct implements the IUserService interface
type UserServiceImpl struct{}

/*
NewUserService creates and returns a new instance of the UserService struct.

This constructor function initializes a zero-value UserService struct, returning a
pointer to it. It does not perform any additional setup or initialization logic beyond
the default struct creation. This is typically used to instantiate the UserService when
no custom initialization is required.

Returns:
- *UserService: A pointer to the newly created UserService instance.
*/
func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{}
}

/*
GetAllUsers retrieves all users from the system. It generates a new unique user ID and
returns a slice of hardcoded User models along with any error encountered. In case of
error during UUID generation, it returns an empty slice of users and an error message.
*/
func (us *UserServiceImpl) GetAllUsers() ([]models.User, error) {
	userID, err := uuid.NewV7()
	if err != nil {
		return []models.User{}, fmt.Errorf("%w\n", err)
	}

	users := []models.User{
		{
			ID:    userID,
			Name:  "Somraj Saha",
			Email: "somraj.saha@weburz.com",
		},
		{
			ID:    userID,
			Name:  "John Doe",
			Email: "john.doe@example.com",
		},
		{
			ID:    userID,
			Name:  "Sagar Kapoor",
			Email: "sagar.kapoor@weburz.com",
		},
	}

	return users, nil
}

/*
GetUserByID retrieves a user by their unique ID. It returns the corresponding User
model with hardcoded details and an error (if any). If no error occurs, the user details
are returned with a nil error.
*/
func (us *UserServiceImpl) GetUserByID(id uuid.UUID) (models.User, error) {
	userID, err := uuid.Parse(id.String())
	if err != nil {
		return models.User{}, fmt.Errorf("%w\n", err)
	}

	user := models.User{
		ID:    userID,
		Name:  "Somraj Saha",
		Email: "somraj.saha@weburz.com",
	}

	return user, nil
}

/*
CreateUser creates a new user with the provided name and email. It generates a new
unique user ID and returns the newly created User model along with any error encountered
during UUID generation or other issues.
*/
func (us *UserServiceImpl) CreateUser(name, email string) (models.User, error) {
	userID, err := uuid.NewV7()
	if err != nil {
		return models.User{}, fmt.Errorf("%w\n", err)
	}

	user := models.User{
		ID:    userID,
		Name:  name,
		Email: email,
	}

	return user, nil
}

/*
UpdateUser updates an existing user's details using the provided ID, name, and email. It
parses the ID to ensure it's valid and returns the updated User model along with any
error encountered during ID parsing or other operations.
*/
func (us *UserServiceImpl) UpdateUser(
	id uuid.UUID,
	name, email string,
) (models.User, error) {
	userID, err := uuid.Parse(id.String())
	if err != nil {
		return models.User{}, fmt.Errorf("%w\n", err)
	}

	user := models.User{
		ID:    userID,
		Name:  name,
		Email: email,
	}

	return user, nil
}

/*
DeleteUser removes a user from the system using the provided unique user ID. It prints
a message confirming the user has been deleted.
*/
func (us *UserServiceImpl) DeleteUser(id uuid.UUID) error {
	fmt.Printf("User %s deleted!", id)
	return nil
}
