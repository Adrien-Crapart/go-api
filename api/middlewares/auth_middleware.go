package middlewares

import (
	"go-api/auth"
	"net/http"
)

// AuthMiddleware verifies the JWT token in the request header
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extract and verify JWT token from the request header
        token, err := auth.ExtractToken(r)
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Validate token
        if !auth.VerifyToken(token) {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Call the next handler
        next.ServeHTTP(w, r)
    })
}