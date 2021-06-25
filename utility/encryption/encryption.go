package encryption

import (
	"golangtest/utility/env"
)

// jwt payload data structure
type UserJWT struct {
	Uuid      string
	Type      string
	Resources []string
}

// generate jwt, wrapping the payload into hash string
func GenerateJWT(UserJWT) string {
	// get environment variables
	appSettings := env.GetENV()

	return appSettings.JwtSecret
}

// checks the jswt if still valid,
// if not return error
func ValidateJWT(jwt string) (UserJWT, error) {
	return UserJWT{}, nil
}

// Hash password
func GenerateHash() string {
	return ""
}

// check if the string is equevalint to the hash
func ValidateHash(pass string, hash string) bool {
	return false
}
