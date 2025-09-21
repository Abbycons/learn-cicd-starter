package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from the Authorization header.
// Expected format: "Authorization: Bearer <API_KEY>"
func GetAPIKey(r *http.Request) (string, error) {
	const prefix = "Bearer "

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	if !strings.HasPrefix(authHeader, prefix) {
		return "", errors.New("authorization header format must be 'Bearer <token>'")
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, prefix))
	if token == "" {
		return "", errors.New("authorization header missing token")
	}

	return token, nil
}
