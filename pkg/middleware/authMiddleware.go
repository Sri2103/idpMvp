package middleware

import (
	"context"
	"net/http"
	"strings"
)

// AuthMiddleware validates the presence and structure of an Authorization token (JWT, API key, etc.)
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// TODO: replace with real JWT validation
		if token != "valid-token" {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add user info into context if valid
		ctx := context.WithValue(r.Context(), "user", "sample-user")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
