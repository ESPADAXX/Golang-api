package middleware

import (
	"net/http"
)

// RoleMiddleware enforces role-based access control
func RoleMiddleware(requiredRole string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the role from the context
			role, ok := r.Context().Value(ContextKeyRole).(string)
			if !ok || role == "" {
				http.Error(w, "Unauthorized: no role found", http.StatusUnauthorized)
				return
			}

			// Check if the role matches the required role
			if role != requiredRole {
				http.Error(w, "Forbidden: insufficient permissions", http.StatusForbidden)
				return
			}

			// Proceed to the next handler
			next.ServeHTTP(w, r)
		})
	}
}
