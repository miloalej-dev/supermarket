package middlewares

import (
	"encoding/json"
	"github.com/miloalej-dev/supermarket/internal/infrastructure/adapters/http/responses"
	"net/http"
	"os"
)

var authToken = os.Getenv("AUTH_TOKEN")

// AuthMiddleware is a middleware that checks for the presence of an authorization token in the request headers.
// If the token is missing or invalid, it returns a 401 Unauthorized response.
// If the token is valid, it allows the request to proceed to the next handler.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Here you can add logic to validate the authToken, e.g., check against a database or a list of valid tokens.
		// For simplicity, let's assume any non-empty token is valid.
		if r.Header.Get("Authorization") != authToken {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			unauthorizeError, _ := json.Marshal(responses.ErrorResponse{
				Status:  401,
				Message: "Unauthorized access, token is missing or invalid",
			})

			_, _ = w.Write(unauthorizeError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
