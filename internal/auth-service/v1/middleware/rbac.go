package middleware

import (
	"net/http"

	"idp_mvp/internal/auth-service/v1/repository"
)

func RequireRole(role string, repo repository.UserRepository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}
