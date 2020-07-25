package auth

import (
	"Project/auth-cookie/model"
	"errors"
)

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

// ValidCredentials isn't a production example.
// Check credentials is out of scope of this project
// Here I simulate a map to store user data, but in production
// you shouldn't to save user password in database.
func ValidCredentials(creds Credentials) (*model.User, error) {

	var mockUserCredentials = map[string]string{
		"samu": "123456",
	}

	expectedPassword, ok := mockUserCredentials[creds.Email]
	if !ok || expectedPassword != creds.Password {
		return nil, errors.New("User unauthorized")
	}

	var mockUserData = &model.User{
		ID:    22,
		Name:  "Samuel",
		Admin: true,
	}

	return mockUserData, nil
}
