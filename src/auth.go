package src

import (
	"net/http"
	"strings"
)

// TokenValidator checks if the request has a valid API token
func TokenValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		// Strip "Bearer " prefix
		token = strings.TrimPrefix(token, "Bearer ")
		
		if !validateToken(token) {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func validateToken(token string) bool {
	// In production, verify against your auth service
	return len(token) > 0
}

// CheckCredentials verifies username and password
func CheckCredentials(username, password string) bool {
	// Placeholder - connect to your user database
	return username != "" && password != ""
}
