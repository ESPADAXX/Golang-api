package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key for signing and validating JWTs (keep this secret)
var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

// ContextKey is the type for context keys to avoid collisions
type ContextKey string

const (
	ContextKeyRole ContextKey = "role" // Context key for storing the role
)

// JWTMiddleware verifies the JWT token and adds claims to the context
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Extract the token (expected format: "Bearer <token>")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]
		
		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtSecret, nil
		})


		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Extract role from claims
		role, ok := claims["role"].(string)
		if !ok || role == "" {
			http.Error(w, "Missing role in token claims", http.StatusUnauthorized)
			return
		}

		// Add role to context
		ctx := context.WithValue(r.Context(), ContextKeyRole, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
